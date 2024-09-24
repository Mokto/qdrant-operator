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
	hasDuplicatedShards = false
	for collectionName, collection := range obj.Status.Collections {
		for shardNumber := range collection.ShardNumber {
			shardsFromId := collection.Shards.GetShardsPerId(shardNumber)
			if len(shardsFromId) < int(collection.ReplicationFactor) {
				log.Info(fmt.Sprintf("Shard %d is missing on some peers", shardNumber))
				peerIdsWithShard := []string{}
				for _, shard := range shardsFromId {
					peerIdsWithShard = append(peerIdsWithShard, strconv.FormatUint(shard.PeerId, 10))
				}

				for peerId := range obj.Status.Collections[collectionName].Shards {
					if !slices.Contains(peerIdsWithShard, peerId) {
						log.Info(fmt.Sprintf("Replicating shard %d to %s", shardNumber, peerId))
						hasDuplicatedShards = true
						conn, err := grpc.NewClient(obj.Status.Peers[peerId].DNS+":6334", grpc.WithTransportCredentials(insecure.NewCredentials()))
						if err != nil {
							log.Error(err, "grpc.NewClient")
							return false, err
						}
						defer conn.Close()

						fromPeerId, err := strconv.ParseUint(peerIdsWithShard[0], 10, 64)
						if err != nil {
							panic(err)
						}
						toPeerId, err := strconv.ParseUint(peerId, 10, 64)
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
								},
							},
						})
						if err != nil {
							log.Error(err, "unable to replicate shard")
							return false, err
						}
						break
					}
				}
			}
		}
		// conn, err := grpc.NewClient(obj.Status.Peers[*to].DNS+":6334", grpc.WithTransportCredentials(insecure.NewCredentials()))
		// if err != nil {
		// 	log.Error(err, "grpc.NewClient")
		// 	return nil
		// }
		// defer conn.Close()

		// client := qdrant.NewCollectionsClient(conn)
		// _, err = client.UpdateCollectionClusterSetup(ctx, &qdrant.UpdateCollectionClusterSetupRequest{
		// 	CollectionName: collectionName,
		// 	Operation: &qdrant.UpdateCollectionClusterSetupRequest_MoveShard{
		// 		MoveShard: &qdrant.MoveShard{
		// 			ShardId:    *foundShardNumber,
		// 			FromPeerId: fromPeerId,
		// 			ToPeerId:   toPeerId,
		// 		},
		// 	},
		// })
		// if err != nil {
		// 	log.Error(err, "unable to move shards")
		// 	return nil
		// }
		// log.Info(fmt.Sprintf("Shard %d moved from %s to %s", *foundShardNumber, *from, *to))

	}
	return hasDuplicatedShards, nil
}
