package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type VolumeClaim struct {
	Storage          string `json:"storage,omitempty"`
	StorageClassName string `json:"storageClassName,omitempty"`
}

type StatefulSet struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +required
	Name string `json:"name,omitempty"`
	// +required
	Replicas int32 `json:"replicas,omitempty"`

	VolumeClaim       *VolumeClaim                `json:"volumeClaim,omitempty"`
	EphemeralStorage  bool                        `json:"ephemeralStorage,omitempty"`
	NodeSelector      map[string]string           `json:"nodeSelector,omitempty"`
	Affinity          *corev1.Affinity            `json:"affinity,omitempty"`
	Tolerations       []corev1.Toleration         `json:"tolerations,omitempty"`
	PriorityClassName string                      `json:"priorityClassName,omitempty"`
	Resources         corev1.ResourceRequirements `json:"resources,omitempty"`
}

// QdrantClusterSpec defines the desired state of QdrantCluster
type QdrantClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +required
	Image            string                        `json:"image,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
	Statefulsets     []StatefulSet                 `json:"statefulsets,omitempty"`
}

// QdrantClusterStatus defines the observed state of QdrantCluster
type QdrantClusterStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	Peers       Peers       `json:"peers,omitempty"`
	Collections Collections `json:"collections,omitempty"`
	// When the cluster is draining, no new writes are allowed and shards are moved out of the node.
	CordonedPeerIds []string `json:"cordonedPeerIds,omitempty"`
	// DesiredReplicasPerStatefulSet map[string]*int32 `json:"desiredReplicasPerStatefulset,omitempty"`
}

// func (status *QdrantClusterStatus) SetDesiredReplicasPerStatefulSet(statefulsetName string, replicas int32) *Peer {
// 	if status.DesiredReplicasPerStatefulSet == nil {
// 		status.DesiredReplicasPerStatefulSet = map[string]*int32{}
// 	}
// 	status.DesiredReplicasPerStatefulSet[statefulsetName] = &replicas
// 	return nil
// }

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// QdrantCluster is the Schema for the qdrantclusters API
type QdrantCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QdrantClusterSpec   `json:"spec,omitempty"`
	Status QdrantClusterStatus `json:"status,omitempty"`
}

func (q *QdrantCluster) GetServiceName() string {
	return q.Name
}

func (q *QdrantCluster) GetHeadlessServiceName() string {
	return q.Name + "-headless"
}

// +kubebuilder:object:root=true

// QdrantClusterList contains a list of QdrantCluster
type QdrantClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QdrantCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&QdrantCluster{}, &QdrantClusterList{})
}
