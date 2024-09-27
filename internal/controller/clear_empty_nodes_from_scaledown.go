package controller

import (
	"context"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/go-logr/logr"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (r *QdrantClusterReconciler) clearEmptyNodesFromScaleDown(_ context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) error {
	statefulsetsNumberOfReplicas := map[string]int32{}
	for _, statefulset := range obj.Spec.Statefulsets {
		statefulsetsNumberOfReplicas[obj.Name+"-"+statefulset.Name] = statefulset.Replicas
	}

	for peerId, peer := range obj.Status.Peers {

		if slices.Contains(obj.Status.CordonedPeerIds, peerId) {
			continue
		}
		statefulsetsNumberOfReplicas := statefulsetsNumberOfReplicas[peer.StatefulSetName]
		currentReplicaNumberParsed, err := strconv.ParseInt(strings.Replace(peer.PodName, peer.StatefulSetName+"-", "", 1), 10, 64)
		if err != nil {
			panic(err)
		}
		currentReplicaNumber := int32(currentReplicaNumberParsed)
		if currentReplicaNumber > statefulsetsNumberOfReplicas-1 {
			log.Info("Deleting peer " + peerId + " from the cluster. DNS was " + peer.DNS)

			client := &http.Client{}
			req, err := http.NewRequest("DELETE", "http://"+obj.GetServiceName()+"."+obj.Namespace+":6333/cluster/peer/"+peerId+"?force=true", nil)
			if obj.Spec.ApiKey != "" {
				req.Header.Add("api-key", obj.Spec.ApiKey)
			}
			if err != nil {
				return err
			}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
		}

	}

	return nil
}
