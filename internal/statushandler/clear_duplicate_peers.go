package statushandler

import (
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (s *StatusHandler) clearDuplicatePeers(cluster *qdrantv1alpha1.QdrantCluster) (hasDoneAnything bool, err error) {
	/**
	* This is addressing localstorage nodes that could be join the cluster with the same DNS
	* The watcher (watch remove peers) is supposed to handle that earlier but this would be used
	* in case the controller is down when the pod restarts
	 */
	peerDnsToId := map[string]string{}
	peersDnsToIds := map[string][]string{}
	for peerId, peer := range cluster.Status.Peers {
		if peerDnsToId[peer.DNS] != "" {
			peersDnsToIds[peer.DNS] = append(peersDnsToIds[peer.DNS], peerId)
			peersDnsToIds[peer.DNS] = append(peersDnsToIds[peer.DNS], peerDnsToId[peer.DNS])
		}
		peerDnsToId[peer.DNS] = peerId
	}

	if len(peersDnsToIds) == 0 {
		return hasDoneAnything, nil
	}

	for dns, ids := range peersDnsToIds {
		// Get unique elements from the slice
		ids = uniqueStrings(ids)
		fmt.Println(ids)
		bodyString, err := s.getClusterInfo(s.ctx, dns, cluster.Spec.ApiKey)
		if err != nil {
			s.log.Info("unable to get cluster info. Deleting all peers with the same DNS.")
			for _, id := range ids {
				err := s.deletePeer(cluster, id)
				if err != nil {
					s.log.Error(err, "unable to delete peer")
					return false, err
				}
			}
			return true, nil
		}

		foundPeerId := gjson.Get(bodyString, "result.peer_id").String()

		peersToDelete := []string{}
		for _, id := range ids {
			if foundPeerId != id {
				peersToDelete = append(peersToDelete, id)
			}
		}
		for _, id := range peersToDelete {
			err := s.deletePeer(cluster, id)
			if err != nil {
				s.log.Error(err, "unable to delete peer")
				return false, err
			}
		}
		return true, nil

	}

	return
}

func (s *StatusHandler) deletePeer(obj *qdrantv1alpha1.QdrantCluster, peerToDelete string) error {
	s.log.Info("Deleting peer " + peerToDelete + " from the cluster.")

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://"+obj.GetServiceName()+"."+obj.Namespace+":6333/cluster/peer/"+peerToDelete+"?force=true", nil)
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

	s.log.Info("Done deleting peer " + peerToDelete + " from the cluster.")
	return nil
}

func uniqueStrings(slice []string) []string {
	// Create a map to track seen elements
	seen := make(map[string]bool)
	var result []string

	// Iterate over the slice
	for _, str := range slice {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}
