package statushandler

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (s *StatusHandler) watchAndRemoveDeletedPeers() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	watcher, err := clientset.CoreV1().Pods("").Watch(context.Background(), v1.ListOptions{
		LabelSelector: "qdrant-ephemeral-storage",
	})
	if err != nil {
		s.log.Error(err, "unable to watch pods")
		return
	}

	for event := range watcher.ResultChan() {
		item := event.Object.(*corev1.Pod)

		if item.Labels["qdrant-ephemeral-storage"] != "true" {
			continue
		}

		switch event.Type {
		case watch.Modified:
			if item.ObjectMeta.DeletionTimestamp != nil {
				// Peer is being deleted
			}
		case watch.Deleted:
			clusters := &qdrantv1alpha1.QdrantClusterList{}
			err := s.manager.GetClient().List(s.ctx, clusters)
			if err != nil {
				s.log.Error(err, "unable to list QdrantClusters")
			}
			for _, cluster := range clusters.Items {
				peerId := cluster.Status.Peers.FindPeerId(item.Name)
				if peerId != "" {
					err := s.deletePeer(&cluster, peerId)
					if err != nil {
						s.log.Error(err, "unable to delete peer")
					}
				}
			}
		}

	}

}
