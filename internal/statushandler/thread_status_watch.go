package statushandler

import (
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (s *StatusHandler) checkForConsensusThreadStatus() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(10 * time.Second)

		clusters := &qdrantv1alpha1.QdrantClusterList{}
		err := s.manager.GetClient().List(s.ctx, clusters)
		if err != nil {
			s.log.Error(err, "unable to list QdrantClusters")
		}

		for _, cluster := range clusters.Items {

			for _, peer := range cluster.Status.Peers {
				threadStatus, err := s.getConsensusThreadStatus(peer.DNS)
				if err != nil {
					s.log.Error(err, "unable to get consensus thread status")
					continue
				}
				if threadStatus == "stopped" {
					err := clientset.CoreV1().Pods(cluster.Namespace).Delete(s.ctx, peer.PodName, v1.DeleteOptions{})
					if err != nil {
						s.log.Error(err, "unable to delete pod")
						continue
					}
					s.log.Info("Consensus thread stopped. Restarting the pod " + peer.PodName + ".")
				}
			}
		}

	}
}

func (s *StatusHandler) getConsensusThreadStatus(serviceName string) (string, error) {

	resp, err := http.Get("http://" + serviceName + ":6333/cluster")
	if err != nil {
		s.log.Error(err, "unable to get cluster info")
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.Error(err, "unable to read response body")
		return "", err
	}
	resp.Body.Close()
	bodyString := string(body)

	return gjson.Get(bodyString, "result.consensus_thread_status.consensus_thread_status").String(), nil
}
