package v1alpha1

type Peers map[string]*Peer

type Peer struct {
	IsLeader         bool   `json:"isLeader,omitempty"`
	PodName          string `json:"podName,omitempty"`
	StatefulSetName  string `json:"statefulSetName,omitempty"`
	DNS              string `json:"dns,omitempty"`
	IsReady          bool   `json:"isReady,omitempty"`
	EphemeralStorage bool   `json:"ephemeralStorage,omitempty"`
}

func (peers *Peers) GetLeader() *Peer {
	for _, peer := range *peers {
		if peer.IsLeader {
			return peer
		}
	}
	return nil
}
func (peers *Peers) AllReady() bool {
	for _, peer := range *peers {
		if !peer.IsReady {
			return false
		}
	}
	return true
}

func (peers *Peers) FindPeerId(podName string) string {
	for peerId, peer := range *peers {
		if peer.PodName == podName {
			return peerId
		}
	}
	return ""
}
