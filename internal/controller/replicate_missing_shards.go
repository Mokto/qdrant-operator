package controller

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
	"qdrantoperator.io/operator/internal/qdrant"
)

func (r *QdrantClusterReconciler) replicateMissingShards(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) (hasDuplicatedShards bool, err error) {
	for collectionName, collection := range obj.Status.Collections {
		for shardNumber := range collection.ShardNumber {
			shardsFromId := collection.Shards.GetNonDeadShardsPerId(shardNumber)
			allShardsFromId := collection.Shards.GetShardsPerId(shardNumber)
			if len(shardsFromId) < int(collection.ReplicationFactor) {
				peerIdsWithShard := []string{}
				for _, shard := range shardsFromId {
					peerIdsWithShard = append(peerIdsWithShard, strconv.FormatUint(shard.PeerId, 10))
				}

				log.Info("Looking for the best peer to replicate the shard to...")

				currentMinShardsCount := 100000
				bestPeerId := ""
				for peerId := range obj.Status.Peers {
					if !slices.Contains(peerIdsWithShard, peerId) && obj.Status.Peers[peerId].IsReady {
						shardsOnThatPeer := len(collection.Shards[peerId])
						if shardsOnThatPeer < currentMinShardsCount {
							currentMinShardsCount = shardsOnThatPeer
							bestPeerId = peerId
						}
					}
				}

				log.Info(fmt.Sprintf("Replicating shard %d to %s", shardNumber, bestPeerId))
				hasDuplicatedShards = true
				conn, err := grpc.NewClient(obj.Status.Peers[bestPeerId].DNS+":6334", grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					log.Error(err, "grpc.NewClient")
					return false, err
				}
				defer conn.Close()

				fromPeerId, err := strconv.ParseUint(peerIdsWithShard[0], 10, 64)
				if err != nil {
					panic(err)
				}
				toPeerId, err := strconv.ParseUint(bestPeerId, 10, 64)
				if err != nil {
					panic(err)
				}

				client := qdrant.NewCollectionsClient(conn)
				_, err = client.UpdateCollectionClusterSetup(ctx, &qdrant.UpdateCollectionClusterSetupRequest{
					CollectionName: collectionName,
					Operation: &qdrant.UpdateCollectionClusterSetupRequest_ReplicateShard{
						ReplicateShard: &qdrant.ReplicateShard{
							ShardId:    shardNumber,
							FromPeerId: fromPeerId,
							ToPeerId:   toPeerId,
							Method:     qdrant.ShardTransferMethod_Snapshot.Enum(),
						},
					},
				})
				if err != nil {
					log.Error(err, "unable to replicate shard")
					return false, err
				}
			} else {
				// if we have enough shards, let's kill the Dead ones
				for _, shard := range allShardsFromId {
					if shard.State == "Dead" {
						log.Info("Shard is dead, deleting...")

						peerId := strconv.FormatUint(shard.PeerId, 10)
						if err != nil {
							panic(err)
						}
						conn, err := grpc.NewClient(obj.Status.Peers[peerId].DNS+":6334", grpc.WithTransportCredentials(insecure.NewCredentials()))
						if err != nil {
							log.Error(err, "grpc.NewClient")
							return false, err
						}
						defer conn.Close()

						client := qdrant.NewCollectionsClient(conn)
						_, err = client.UpdateCollectionClusterSetup(ctx, &qdrant.UpdateCollectionClusterSetupRequest{
							CollectionName: collectionName,
							Operation: &qdrant.UpdateCollectionClusterSetupRequest_DropReplica{
								DropReplica: &qdrant.Replica{
									ShardId: *shard.ShardId,
									PeerId:  shard.PeerId,
								},
							},
						})
						if err != nil {
							log.Error(err, "unable to drop dead shard")
							return false, err
						}
						log.Info("Deleted dead shard...")
					}
				}

			}
		}
	}
	return hasDuplicatedShards, nil
}
