package v1alpha1

import (
	"strconv"

	"qdrantoperator.io/operator/internal/qdrant"
)

type Collections map[string]*Collection

type ShardsList []*ShardInfo

type ShardsPerPeer map[string]ShardsList

type Collection struct {
	Name              string        `json:"name,omitempty"`
	Status            string        `json:"status,omitempty"`
	ReplicationFactor uint32        `json:"replicationFactor,omitempty"`
	ShardNumber       uint32        `json:"shardNumber,omitempty"`
	Shards            ShardsPerPeer `json:"shards,omitempty"`
	ShardsInProgress  bool          `json:"shardsInProgress,omitempty"`
}

func (collection *Collection) IsIdle() bool {
	return collection.ShardsInProgress || collection.Status != qdrant.CollectionStatus_Green.String()
}

type ShardInfo struct {
	PeerId  uint64  `json:"peerId,omitempty"`
	ShardId *uint32 `json:"shardId,omitempty"`
	State   string  `json:"state,omitempty"`
}

func (shardsPerPeer *ShardsPerPeer) GetShardsPerId(shardId uint32) ShardsList {
	shards := ShardsList{}
	for peerId, shardInfos := range *shardsPerPeer {
		peerIdUint, err := strconv.ParseUint(peerId, 10, 64)
		if err != nil {
			panic(err)
		}
		for _, shardInfo := range shardInfos {
			if shardInfo.ShardId != nil && *shardInfo.ShardId == shardId {
				shards = append(shards, &ShardInfo{
					PeerId:  peerIdUint,
					ShardId: shardInfo.ShardId,
					State:   shardInfo.State,
				})
			}
		}
	}
	return shards
}

func (shardsList *ShardsList) HasShardFromPeer(peerId string) bool {
	for _, shardInfo := range *shardsList {
		if strconv.FormatUint(shardInfo.PeerId, 10) == peerId {
			return true
		}
	}
	return false
}

func (shardsList *ShardsList) AllActive() bool {
	for _, shardInfo := range *shardsList {
		if shardInfo.State != qdrant.ReplicaState_Active.String() {
			return false
		}
	}
	return true
}
