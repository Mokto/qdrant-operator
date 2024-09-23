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
	pb "qdrantoperator.io/operator/internal/qdrant"
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
		start := time.Now()
		clusters := &qdrantv1alpha1.QdrantClusterList{}
		err := s.manager.GetClient().List(s.ctx, clusters)
		if err != nil {
			s.log.Error(err, "unable to list QdrantClusters")
		}

		for _, cluster := range clusters.Items {
			patch := client.MergeFrom(cluster.DeepCopy())

			peerIdsToNames := map[string]string{}
			leaderId := uint64(0)
		out:
			for _, statefulset := range cluster.Spec.Statefulsets {
				for i := 0; i < int(statefulset.Replicas); i++ {
					uniqueName := fmt.Sprintf("%s-%d", statefulset.Name, i)
					endpoint := fmt.Sprintf("%s.%s-headless.%s", uniqueName, cluster.Name, cluster.Namespace)
					resp, err := http.Get("http://" + endpoint + ":6333/cluster")
					if err != nil {
						s.log.Error(err, "unable to get cluster info")
						continue
					}
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						s.log.Error(err, "unable to read response body")
						continue
					}
					bodyString := string(body)

					leaderId = gjson.Get(bodyString, "result.raft_info.leader").Uint()
					for peerId, result := range gjson.Get(bodyString, "result.peers").Map() {
						peerIdsToNames[peerId] = strings.Replace(strings.Replace(result.Get("uri").String(), "http://", "", 1), ":6335/", "", 1)
					}
					break out
				}
			}

			cluster.Status.PeerIdsToNames = peerIdsToNames
			cluster.Status.RaftLeaderPeerId = leaderId

			conn := s.getGrpcConnection(cluster.Status.PeerIdsToNames[strconv.FormatUint(cluster.Status.RaftLeaderPeerId, 10)] + ":6334")
			if conn != nil {
				collections, err := s.getCollections(conn)
				if err == nil {
					cluster.Status.Collections = collections
				}
			}

			shardsPerCollection := map[string][]*qdrantv1alpha1.ShardInfo{}
			hasError := false
			for _, collection := range cluster.Status.Collections {
				shards := []*qdrantv1alpha1.ShardInfo{}

				for _, name := range cluster.Status.PeerIdsToNames {
					conn := s.getGrpcConnection(name + ":6334")
					if err != nil {
						continue
					}
					shards_found, err := s.getShardsInfo(conn, collection)
					if err != nil {
						continue
					}
					shards = append(shards, shards_found...)

					slices.SortFunc(shards, func(a, b *qdrantv1alpha1.ShardInfo) int {
						return cmp.Or(
							cmp.Compare(*a.ShardId, *b.ShardId),
							cmp.Compare(a.PeerId, b.PeerId),
						)
					})

					if len(shards) > 0 {
						shardsPerCollection[collection] = shards
						break
					}
				}

				if len(shards) == 0 {
					hasError = true
					s.log.Error(fmt.Errorf("unable to get shards for collection %s", collection), "")
				}
			}
			if !hasError {
				cluster.Status.ShardsPerCollection = shardsPerCollection
			}

			err = s.manager.GetClient().Status().Patch(s.ctx, &cluster, patch)
			if err != nil {
				s.log.Error(err, "unable to update QdrantCluster status")
			}

		}

		elapsed := time.Since(start)
		s.log.Info(fmt.Sprintf("Fetching cluster data took %s", elapsed))
	}
}

func (s *StatusHandler) getCollections(conn *grpc.ClientConn) ([]string, error) {
	collections := []string{}
	client := pb.NewCollectionsClient(conn)
	collectionsResponse, err := client.List(s.ctx, &pb.ListCollectionsRequest{})

	if err != nil {
		s.log.Error(err, "unable to list collections")
		return nil, err
	}
	for _, collection := range collectionsResponse.Collections {
		collections = append(collections, collection.Name)
	}
	return collections, nil
}

func (s *StatusHandler) getShardsInfo(conn *grpc.ClientConn, collectionName string) ([]*qdrantv1alpha1.ShardInfo, error) {
	shards := []*qdrantv1alpha1.ShardInfo{}
	client := pb.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(s.ctx, 500*time.Millisecond)
	defer cancel()
	clusterInfoResponse, err := client.CollectionClusterInfo(ctx, &pb.CollectionClusterInfoRequest{
		CollectionName: collectionName,
	})
	if err != nil {
		s.log.Error(err, "unable to list shards")
		return nil, err
	}
	for _, shard := range clusterInfoResponse.LocalShards {
		shards = append(shards, &qdrantv1alpha1.ShardInfo{
			PeerId:  clusterInfoResponse.PeerId,
			ShardId: &shard.ShardId,
			State:   &shard.State,
		})
	}
	for _, shard := range clusterInfoResponse.RemoteShards {
		shards = append(shards, &qdrantv1alpha1.ShardInfo{
			PeerId:  shard.PeerId,
			ShardId: &shard.ShardId,
			State:   &shard.State,
		})
	}
	return shards, nil
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
