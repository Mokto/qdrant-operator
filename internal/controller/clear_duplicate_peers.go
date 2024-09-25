package controller

import (
	"context"
	"io"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/tidwall/gjson"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (r *QdrantClusterReconciler) clearDuplicatePeers(_ context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) (hasDoneAnything bool, err error) {

	hasDoneAnything = false

	peerDnsToId := map[string]string{}
	for peerId, peer := range obj.Status.Peers {
		if peerDnsToId[peer.DNS] != "" {
			log.Info("Peer " + peerId + " has the same DNS as " + peerDnsToId[peer.DNS] + ". Deleting one of them.")
			resp, err := http.Get("http://" + peer.DNS + ":6333/cluster")
			if err != nil {
				log.Error(err, "unable to get cluster info")
				break
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Error(err, "unable to read response body")
				break
			}
			resp.Body.Close()
			bodyString := string(body)

			foundPeerId := gjson.Get(bodyString, "result.peer_id").String()

			peerToDelete := ""
			if foundPeerId == peerId {
				peerToDelete = peerDnsToId[peer.DNS]
			} else if foundPeerId == peerDnsToId[peer.DNS] {
				peerToDelete = peerId
			}
			if peerToDelete != "" {
				log.Info("Deleting peer " + peerToDelete + " from the cluster.")

				client := &http.Client{}
				req, err := http.NewRequest("DELETE", "http://"+obj.GetServiceName()+":6333/cluster/peer/"+peerToDelete, nil)
				if err != nil {
					return false, err
				}
				resp, err := client.Do(req)
				if err != nil {
					return false, err
				}
				defer resp.Body.Close()
			}

			hasDoneAnything = true
			break
		}
		peerDnsToId[peer.DNS] = peerId
	}

	return hasDoneAnything, nil
}
