package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-logr/logr"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (r *QdrantClusterReconciler) ensureShardsSafe(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) (hasDoneAnything bool, err error) {
	if !obj.Status.Peers.HasEphemeralStorage() {
		return
	}

	for collectionName, collection := range obj.Status.Collections {
		for shardNumber := range collection.ShardNumber {
			shardsFromId := collection.Shards.GetShardsPerId(shardNumber)
			isShardSafe := false
			for _, shard := range shardsFromId {
				peer := obj.Status.Peers[strconv.FormatUint(shard.PeerId, 10)]
				if peer != nil && !peer.EphemeralStorage {
					isShardSafe = true
					break
				}
			}

			if !isShardSafe {
				log.Info(fmt.Sprintf("Shard is not safe %d", shardNumber))
				bestCandidateForShardFrom := uint64(0)
				numberOfShards := 0
				for _, shard := range shardsFromId {
					currentNumberOfShards := len(collection.Shards[strconv.FormatUint(shard.PeerId, 10)])
					if bestCandidateForShardFrom == 0 || currentNumberOfShards > numberOfShards {
						bestCandidateForShardFrom = shard.PeerId
						numberOfShards = currentNumberOfShards
					}
				}
				bestCandidateForShardTo := ""
				numberOfShards = 0
				for peerId, peer := range obj.Status.Peers {
					if peer.EphemeralStorage {
						continue
					}
					currentNumberOfShards := len(collection.Shards[peerId])
					if bestCandidateForShardTo == "" || currentNumberOfShards < numberOfShards {
						bestCandidateForShardTo = peerId
						numberOfShards = currentNumberOfShards
					}
				}

				hasDoneAnything, err = r.moveShardSafely(ctx, log, collectionName, shardNumber, strconv.FormatUint(bestCandidateForShardFrom, 10), bestCandidateForShardTo, obj.Status.Peers[bestCandidateForShardTo].DNS)
				if err != nil {
					return
				}
				if hasDoneAnything {
					return
				}
			}

		}
	}
	return
}
