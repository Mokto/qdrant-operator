package v1alpha1

import (
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
	Replicas          int32        `json:"replicas,omitempty"`
	PriorityClassName string       `json:"priorityClassName,omitempty"`
	VolumeClaim       *VolumeClaim `json:"volumeClaim,omitempty"`
	EphemeralStorage  bool         `json:"ephemeralStorage,omitempty"`
}

// QdrantClusterSpec defines the desired state of QdrantCluster
type QdrantClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of QdrantCluster. Edit qdrantcluster_types.go to remove/update
	Statefulsets []StatefulSet `json:"statefulsets,omitempty"`
}

// QdrantClusterStatus defines the observed state of QdrantCluster
type QdrantClusterStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	Peers                         Peers             `json:"peers,omitempty"`
	Collections                   Collections       `json:"collections,omitempty"`
	DesiredReplicasPerStatefulSet map[string]*int32 `json:"desiredReplicasPerStatefulset,omitempty"`
}

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
