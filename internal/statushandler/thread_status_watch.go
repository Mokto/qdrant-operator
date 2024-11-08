package statushandler

import (
	"time"

	"github.com/tidwall/gjson"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (s *StatusHandler) getKubernetesClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func (s *StatusHandler) DeletePod(namespace string, name string) {
	clientset := s.getKubernetesClient()
	err := clientset.CoreV1().Pods(namespace).Delete(s.ctx, name, v1.DeleteOptions{})
	if err != nil {
		s.log.Error(err, "unable to delete pod")
	}
}

func (s *StatusHandler) checkForConsensusThreadStatus() {
	for {
		time.Sleep(30 * time.Second)

		clusters := &qdrantv1alpha1.QdrantClusterList{}
		err := s.manager.GetClient().List(s.ctx, clusters)
		if err != nil {
			s.log.Error(err, "unable to list QdrantClusters")
		}

		for _, cluster := range clusters.Items {

			for _, peer := range cluster.Status.Peers {
				if !peer.EphemeralStorage {
					continue
				}
				threadStatus, raftTerm, raftCommit, err := s.getConsensusThreadStatus(peer.DNS, cluster.Spec.ApiKey)
				if err != nil {
					s.log.Error(err, "unable to get consensus thread status")
					continue
				}
				if threadStatus == "stopped" {
					s.DeletePod(cluster.Namespace, peer.PodName)
					s.log.Info("Consensus thread stopped. Restarting the pod " + peer.PodName + ".")
				}
				if raftTerm == 0 && raftCommit == 0 {
					s.log.Info("raftTerm and raftCommit are 0 for peer " + peer.DNS + ". Waiting 30s and killing the pod if that still the case")
					time.Sleep(15 * time.Second)
					_, raftTerm, raftCommit, err = s.getConsensusThreadStatus(peer.DNS, cluster.Spec.ApiKey)
					if err != nil {
						s.log.Error(err, "unable to get consensus thread status")
						continue
					}
					if raftTerm == 0 && raftCommit == 0 {
						s.DeletePod(cluster.Namespace, peer.PodName)
						s.log.Info("raftTerm and raftCommit are 0 for peer " + peer.DNS + ". Restarting the pod " + peer.PodName + ".")
					}
				}
			}
		}

	}
}

func (s *StatusHandler) getConsensusThreadStatus(serviceName string, apiKey string) (status string, raftTerm uint64, raftCommit uint64, err error) {
	bodyString, err := s.getClusterInfo(s.ctx, serviceName, apiKey)
	if err != nil {
		s.log.Error(err, "unable to get cluster info")
		return "", 0, 0, err
	}

	return gjson.Get(bodyString, "result.consensus_thread_status.consensus_thread_status").String(), gjson.Get(bodyString, "result.raft_info.term").Uint(), gjson.Get(bodyString, "result.raft_info.commit").Uint(), nil
}
