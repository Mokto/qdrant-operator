package statushandler

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
	"qdrantoperator.io/operator/internal/qdrant"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type StatusHandler struct {
	ctx             context.Context
	log             logr.Logger
	manager         manager.Manager
	grpcConnections map[string]*grpc.ClientConn
}

func NewStatusHandler(mngr manager.Manager) *StatusHandler {

	return &StatusHandler{
		ctx:             context.Background(),
		log:             ctrl.Log.WithName("statusHandler"),
		manager:         mngr,
		grpcConnections: map[string]*grpc.ClientConn{},
	}
}

func (s *StatusHandler) Run() {

	go s.watchAndRemoveDeletedPeers()
	go s.checkForConsensusThreadStatus()

	for {
		time.Sleep(2 * time.Second)
		// start := time.Now()
		clusters := &qdrantv1alpha1.QdrantClusterList{}
		err := s.manager.GetClient().List(s.ctx, clusters)
		if err != nil {
			s.log.Error(err, "unable to list QdrantClusters")
		}

		for _, cluster := range clusters.Items {
			// Run each cluster independently
			s.runCluster(cluster)
		}

		// elapsed := time.Since(start)
		// s.log.Info(fmt.Sprintf("Fetching cluster data took %s", elapsed))
	}
}

func (s *StatusHandler) runCluster(cluster qdrantv1alpha1.QdrantCluster) {
	defer func() {
		falseValue := false
		trueValue := true

		ctx := context.Background()
		if cluster.Spec.ApiKey != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, "api-key", cluster.Spec.ApiKey)
		}
		patch := client.MergeFrom(cluster.DeepCopy())

		cluster.Status.UnknownStatus = &falseValue
		serviceName := cluster.GetServiceName()
		bodyString, err := s.getClusterInfo(ctx, serviceName+"."+cluster.Namespace, cluster.Spec.ApiKey)
		if err != nil {
			s.log.Error(err, "unable to get cluster info")
			return
		}

		// Getting peers from main service endpoint
		peers := qdrantv1alpha1.Peers{}

		currentPeerId := gjson.Get(bodyString, "result.peer_id").String()
		currentPeerName := ""
		leaderId := gjson.Get(bodyString, "result.raft_info.leader").String()
		for peerId, result := range gjson.Get(bodyString, "result.peers").Map() {
			serviceName := strings.Replace(strings.Replace(result.Get("uri").String(), "http://", "", 1), ":6335/", "", 1)
			podName := strings.Replace(strings.Replace(serviceName, "."+cluster.GetHeadlessServiceName(), "", 1), "."+cluster.GetNamespace(), "", 1)
			dns := podName + "." + cluster.GetHeadlessServiceName() + "." + cluster.Namespace
			statefulsetName := strings.Join(strings.Split(podName, "-")[:len(strings.Split(podName, "-"))-1], "-")
			if peerId == currentPeerId {
				currentPeerName = podName
			}

			resp, err := http.Get("http://" + dns + ":6333/readyz")
			isReady := err == nil && resp.StatusCode == 200

			foundStatefulset := qdrantv1alpha1.StatefulSet{}
			for _, statefulsetConfig := range cluster.Spec.Statefulsets {
				if cluster.Name+"-"+statefulsetConfig.Name == statefulsetName {
					foundStatefulset = statefulsetConfig
					break
				}
			}

			peers[peerId] = &qdrantv1alpha1.Peer{
				IsLeader:         leaderId == peerId,
				StatefulSetName:  statefulsetName,
				PodName:          podName,
				DNS:              dns,
				IsReady:          isReady,
				EphemeralStorage: foundStatefulset.EphemeralStorage,
			}
		}

		if peers.GetLeader() == nil {
			s.log.Error(errors.New("leader not found on peer "+currentPeerId+" / "+currentPeerName), fmt.Sprintf("Leader not found amongst %d peers on cluster %s: %s / %s", len(peers), cluster.Namespace, currentPeerId, currentPeerName))
			return
		}
		cluster.Status.Peers = peers
		cluster.Status.HasBeenInited = &trueValue

		hasDeletedPeers, err := s.clearDuplicatePeers(&cluster)
		if err != nil {
			s.log.Error(err, "unable to read response body")
			return
		}
		if hasDeletedPeers {
			s.log.Info("Deleted duplicate peers")
			return
		}

		conn := s.getGrpcConnection(cluster.GetServiceName() + "." + cluster.Namespace + ":6334")
		if conn == nil {
			// error is already logged
			return
		}
		collectionsList, err := s.getCollections(ctx, conn)
		if err != nil {
			s.log.Error(err, "unable to get collections")
			return
		}

		collections := map[string]*qdrantv1alpha1.Collection{}
		for _, collection := range collectionsList {
			collections[collection] = &qdrantv1alpha1.Collection{}

			shards, shardInProgress, err := s.getShardsInfo(ctx, conn, collection)
			if err != nil {
				cluster.Status.UnknownStatus = &trueValue
				continue
			}
			collections[collection].Shards = shards

			hasInProgressShards := false
			for _, shardInProgress := range shardInProgress {
				fmt.Println(shardInProgress)
				peerIdTo := strconv.FormatUint(shardInProgress.To, 10)
				peerIdFrom := strconv.FormatUint(shardInProgress.From, 10)
				if cluster.Status.Peers[peerIdTo] == nil {
					s.abortShardTransfer(ctx, conn, collection, shardInProgress.From, shardInProgress.To, shardInProgress.ShardId)
				} else if cluster.Status.Peers[peerIdFrom] == nil {
					s.abortShardTransfer(ctx, conn, collection, shardInProgress.From, shardInProgress.To, shardInProgress.ShardId)
				} else {
					hasInProgressShards = true
				}
			}
			collections[collection].ShardsInProgress = hasInProgressShards

			collectionInfo, err := s.getCollectionInfo(ctx, conn, collection)
			if err != nil {
				cluster.Status.UnknownStatus = &trueValue
				continue
			}
			collections[collection].Status = collectionInfo.Status.String()
			if collectionInfo.Config.Params.ReplicationFactor != nil {
				collections[collection].ReplicationFactor = *collectionInfo.Config.Params.ReplicationFactor
			} else {
				collections[collection].ReplicationFactor = 1
			}
			collections[collection].ShardNumber = collectionInfo.Config.Params.ShardNumber

		}

		cluster.Status.Collections = collections

		err = s.manager.GetClient().Status().Patch(s.ctx, &cluster, patch)
		if err != nil {
			s.log.Error(err, "unable to update QdrantCluster status")
		}
	}()
}

func (s *StatusHandler) getCollections(ctx context.Context, conn *grpc.ClientConn) ([]string, error) {
	collections := []string{}
	client := qdrant.NewCollectionsClient(conn)
	collectionsResponse, err := client.List(ctx, &qdrant.ListCollectionsRequest{})

	if err != nil {
		s.log.Error(err, "unable to list collections")
		return nil, err
	}
	for _, collection := range collectionsResponse.Collections {
		collections = append(collections, collection.Name)
	}
	return collections, nil
}

func (s *StatusHandler) getShardsInfo(ctx context.Context, conn *grpc.ClientConn, collectionName string) (shards map[string]qdrantv1alpha1.ShardsList, shardsInProgress []*qdrant.ShardTransferInfo, err error) {
	client := qdrant.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	clusterInfoResponse, err := client.CollectionClusterInfo(ctx, &qdrant.CollectionClusterInfoRequest{
		CollectionName: collectionName,
	})
	if err != nil {
		s.log.Error(err, "unable to list shards")
		return nil, nil, err
	}

	shards = map[string]qdrantv1alpha1.ShardsList{}
	localPeerId := strconv.FormatUint(clusterInfoResponse.PeerId, 10)
	for _, shard := range clusterInfoResponse.LocalShards {
		shards[localPeerId] = append(shards[localPeerId], &qdrantv1alpha1.ShardInfo{
			ShardId: &shard.ShardId,
			State:   shard.State.String(),
		})
	}
	for _, shard := range clusterInfoResponse.RemoteShards {
		peerId := strconv.FormatUint(shard.PeerId, 10)
		shards[peerId] = append(shards[peerId], &qdrantv1alpha1.ShardInfo{
			ShardId: &shard.ShardId,
			State:   shard.State.String(),
		})
	}
	for _, shards := range shards {
		slices.SortFunc(shards, func(a, b *qdrantv1alpha1.ShardInfo) int {
			return cmp.Compare(*a.ShardId, *b.ShardId)
		})
	}
	return shards, clusterInfoResponse.ShardTransfers, nil
}

func (s *StatusHandler) getCollectionInfo(ctx context.Context, conn *grpc.ClientConn, collectionName string) (*qdrant.CollectionInfo, error) {
	client := qdrant.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	clusterInfoResponse, err := client.Get(ctx, &qdrant.GetCollectionInfoRequest{
		CollectionName: collectionName,
	})
	if err != nil {
		s.log.Error(err, "unable to get collection status")
		return nil, err
	}
	return clusterInfoResponse.Result, nil
}

func (s *StatusHandler) getGrpcConnection(url string) *grpc.ClientConn {
	if s.grpcConnections[url] == nil {
		conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			s.log.Error(err, "grpc.NewClient")
			return nil
		}
		s.grpcConnections[url] = conn
	}
	return s.grpcConnections[url]
}

func (s *StatusHandler) abortShardTransfer(ctx context.Context, conn *grpc.ClientConn, collectionName string, fromPeerId uint64, toPeerId uint64, shardId uint32) {
	client := qdrant.NewCollectionsClient(conn)
	_, err := client.UpdateCollectionClusterSetup(ctx, &qdrant.UpdateCollectionClusterSetupRequest{
		CollectionName: collectionName,
		Operation: &qdrant.UpdateCollectionClusterSetupRequest_AbortTransfer{
			AbortTransfer: &qdrant.AbortShardTransfer{
				FromPeerId: fromPeerId,
				ToPeerId:   toPeerId,
				ShardId:    shardId,
			},
		},
	})
	if err != nil {
		s.log.Error(err, "unable to move shards")
		return
	}
	s.log.Info("Shard transfer aborted.")
}

func (s *StatusHandler) getClusterInfo(_ context.Context, dns string, apiKey string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://"+dns+":6333/cluster", nil)
	if apiKey != "" {
		req.Header.Set("api-key", apiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		s.log.Error(err, "unable to get cluster info")
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.Error(err, "unable to read response body")
		return "", err
	}
	return string(body), nil
}
