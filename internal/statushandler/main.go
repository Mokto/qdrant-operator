package statushandler

import (
	"cmp"
	"context"
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
	for {
		time.Sleep(5 * time.Second)

		// start := time.Now()
		clusters := &qdrantv1alpha1.QdrantClusterList{}
		err := s.manager.GetClient().List(s.ctx, clusters)
		if err != nil {
			s.log.Error(err, "unable to list QdrantClusters")
		}

		for _, cluster := range clusters.Items {
			// cluster.Status

			patch := client.MergeFrom(cluster.DeepCopy())

			// Getting peers from main service endpoint
			peers := qdrantv1alpha1.Peers{}
			resp, err := http.Get("http://" + cluster.GetServiceName() + ":6333/cluster")
			if err != nil {
				s.log.Error(err, "unable to get cluster info")
				continue
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				s.log.Error(err, "unable to read response body")
				continue
			}
			resp.Body.Close()
			bodyString := string(body)

			leaderId := gjson.Get(bodyString, "result.raft_info.leader").String()
			for peerId, result := range gjson.Get(bodyString, "result.peers").Map() {
				serviceName := strings.Replace(strings.Replace(result.Get("uri").String(), "http://", "", 1), ":6335/", "", 1)
				podName := strings.Replace(strings.Replace(serviceName, "."+cluster.GetHeadlessServiceName(), "", 1), "."+cluster.GetNamespace(), "", 1)
				dns := podName + "." + cluster.GetHeadlessServiceName() + "." + cluster.Namespace

				_, err := http.Get("http://" + dns + ":6333/readyz")
				isReady := err == nil

				peers[peerId] = &qdrantv1alpha1.Peer{
					IsLeader:        leaderId == peerId,
					StatefulSetName: strings.Join(strings.Split(podName, "-")[:len(strings.Split(podName, "-"))-1], "-"),
					PodName:         podName,
					DNS:             dns,
					IsReady:         isReady,
				}
			}

			cluster.Status.Peers = peers

			conn := s.getGrpcConnection(cluster.GetServiceName() + ":6334")
			if conn == nil {
				// error is already logged
				continue
			}
			collectionsList, err := s.getCollections(conn)
			if err != nil {
				s.log.Error(err, "unable to get collections")
				continue
			}

			collections := map[string]*qdrantv1alpha1.Collection{}
			for _, collection := range collectionsList {
				collections[collection] = &qdrantv1alpha1.Collection{}

				collectionInfo, err := s.getCollectionInfo(conn, collection)
				if err != nil {
					continue
				}
				collections[collection].Status = collectionInfo.Status.String()
				if collectionInfo.Config.Params.ReplicationFactor != nil {
					collections[collection].ReplicationFactor = *collectionInfo.Config.Params.ReplicationFactor
				} else {
					collections[collection].ReplicationFactor = 1
				}
				collections[collection].ShardNumber = collectionInfo.Config.Params.ShardNumber

				shards, _, err := s.getShardsInfo(conn, collection)
				if err != nil {
					continue
				}
				collections[collection].Shards = shards
			}

			cluster.Status.Collections = collections

			err = s.manager.GetClient().Status().Patch(s.ctx, &cluster, patch)
			if err != nil {
				s.log.Error(err, "unable to update QdrantCluster status")
			}

		}

		// elapsed := time.Since(start)
		// s.log.Info(fmt.Sprintf("Fetching cluster data took %s", elapsed))
	}
}

func (s *StatusHandler) getCollections(conn *grpc.ClientConn) ([]string, error) {
	collections := []string{}
	client := qdrant.NewCollectionsClient(conn)
	collectionsResponse, err := client.List(s.ctx, &qdrant.ListCollectionsRequest{})

	if err != nil {
		s.log.Error(err, "unable to list collections")
		return nil, err
	}
	for _, collection := range collectionsResponse.Collections {
		collections = append(collections, collection.Name)
	}
	return collections, nil
}

func (s *StatusHandler) getShardsInfo(conn *grpc.ClientConn, collectionName string) (shards map[string]qdrantv1alpha1.ShardsList, shardsInProgress []*qdrantv1alpha1.ShardInfo, err error) {
	client := qdrant.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(s.ctx, 500*time.Millisecond)
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
	for _, shard := range clusterInfoResponse.ShardTransfers {
		fmt.Println("SHARD TRANSFER", shard)
	}
	return shards, nil, nil
}

func (s *StatusHandler) getCollectionInfo(conn *grpc.ClientConn, collectionName string) (*qdrant.CollectionInfo, error) {
	client := qdrant.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(s.ctx, 500*time.Millisecond)
	defer cancel()
	clusterInfoResponse, err := client.Get(ctx, &qdrant.GetCollectionInfoRequest{
		CollectionName: collectionName,
	})
	if err != nil {
		s.log.Error(err, "unable to get collection statu")
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
