package controller

import (
	"cmp"
	"context"
	"fmt"
	"math"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
	"qdrantoperator.io/operator/internal/qdrant"
)

func (r *QdrantClusterReconciler) moveShards(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) error {

	for collectionName, collection := range obj.Status.Collections {
		if collection.Status != qdrant.CollectionStatus_Green.String() {
			log.Info(collectionName + " is not ready to be optimized (status " + collection.Status + "). Skipping.")
			continue
		}
		if collection.ShardsInProgress {
			log.Info(collectionName + " has shards in progress. Skipping.")
			continue
		}

		isCordoning := len(obj.Status.CordonedPeerIds) > 0

		totalShardCount := collection.ShardNumber * collection.ReplicationFactor
		optimalShardsPerHost := float64(totalShardCount) / float64(len(obj.Status.Peers))

		// Getting a list of peers that don't have an optimal number of shards
		abovePeerIds := []string{}
		belowPeerIds := []string{}
		shardsPerPeer := map[string]int{}
		for peerId := range obj.Status.Peers {
			shards := collection.Shards[peerId]
			shardsPerPeer[peerId] = len(shards)
			if len(shards) > int(math.Ceil(optimalShardsPerHost)) {
				abovePeerIds = append(abovePeerIds, peerId)
			} else if len(shards) < int(math.Floor(optimalShardsPerHost)) {
				belowPeerIds = append(belowPeerIds, peerId)
			}
		}

		// If we are cordoning, we need to move shards from the cordoned peer to the best candidate
		if isCordoning {
			abovePeerIds = []string{obj.Status.CordonedPeerIds[0]}
			for peerId := range obj.Status.Peers {
				if !slices.Contains(obj.Status.CordonedPeerIds, peerId) {
					belowPeerIds = append(belowPeerIds, peerId)
				}
			}
		}

		if len(abovePeerIds) == 0 || len(belowPeerIds) == 0 {
			continue
		}
		slices.SortFunc(belowPeerIds, func(a, b string) int {
			return cmp.Compare(shardsPerPeer[b], shardsPerPeer[a])
		})
		slices.SortFunc(abovePeerIds, func(a, b string) int {
			return cmp.Compare(shardsPerPeer[a], shardsPerPeer[b])
		})
		// Findind the best pair of peers to move shards between
		var foundShardNumber *uint32
		var from *string
		var to *string
	out:
		for _, abovePeerId := range abovePeerIds {
			for _, belowPeerId := range belowPeerIds {
				for shardNumber := range collection.ShardNumber {
					shardsFromId := collection.Shards.GetShardsPerId(shardNumber)
					// If the shard is on the above peer but not on the below peer, we can move it!
					if shardsFromId.AllActive() && shardsFromId.HasShardFromPeer(abovePeerId) && !shardsFromId.HasShardFromPeer(belowPeerId) && obj.Status.Peers[belowPeerId].IsReady && obj.Status.Peers[abovePeerId].IsReady && !slices.Contains(obj.Status.CordonedPeerIds, belowPeerId) {
						foundShardNumber = &shardNumber
						from = &abovePeerId
						to = &belowPeerId
						break
					}
				}
				if foundShardNumber != nil {
					break out
				}
			}
		}

		if foundShardNumber != nil {
			log.Info(fmt.Sprintf("Moving shard number: %d from %s to %s", *foundShardNumber, *from, *to))

			conn, err := grpc.NewClient(obj.Status.Peers[*to].DNS+":6334", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Error(err, "grpc.NewClient")
				return err
			}
			defer conn.Close()

			fromPeerId, err := strconv.ParseUint(*from, 10, 64)
			if err != nil {
				panic(err)
			}
			toPeerId, err := strconv.ParseUint(*to, 10, 64)
			if err != nil {
				panic(err)
			}

			resp, err := http.Get("http://" + obj.Status.Peers[*to].DNS + ":6333/readyz")
			if err != nil {
				log.Error(err, "unable to get pod status")
				continue
			}

			if resp.StatusCode == 200 {
				client := qdrant.NewCollectionsClient(conn)
				connectionExists, err := client.CollectionExists(ctx, &qdrant.CollectionExistsRequest{CollectionName: collectionName})
				if err != nil {
					log.Error(err, "unable to check if the collection exists")
					return err
				}
				if connectionExists.Result.Exists {
					_, err = client.UpdateCollectionClusterSetup(ctx, &qdrant.UpdateCollectionClusterSetupRequest{
						CollectionName: collectionName,
						Operation: &qdrant.UpdateCollectionClusterSetupRequest_MoveShard{
							MoveShard: &qdrant.MoveShard{
								ShardId:    *foundShardNumber,
								FromPeerId: fromPeerId,
								ToPeerId:   toPeerId,
								Method:     qdrant.ShardTransferMethod_StreamRecords.Enum(),
							},
						},
					})
					if err != nil {
						log.Error(err, "unable to move shards")
						return err
					}
					log.Info(fmt.Sprintf("Shard %d moved from %s to %s", *foundShardNumber, *from, *to))
				} else {
					log.Info("Collection doesn't exist on receiving node. Skipping for now.")
				}
			} else {
				log.Info("Receiving node is not ready. Skipping for now.")
			}
		}

	}
	return nil
}
