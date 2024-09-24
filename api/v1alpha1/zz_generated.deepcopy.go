//go:build !ignore_autogenerated

/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collection) DeepCopyInto(out *Collection) {
	*out = *in
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make(ShardsPerPeer, len(*in))
		for key, val := range *in {
			var outVal []*ShardInfo
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(ShardsList, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(ShardInfo)
						(*in).DeepCopyInto(*out)
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ShardsInProgress != nil {
		in, out := &in.ShardsInProgress, &out.ShardsInProgress
		*out = make(ShardsList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ShardInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collection.
func (in *Collection) DeepCopy() *Collection {
	if in == nil {
		return nil
	}
	out := new(Collection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Collections) DeepCopyInto(out *Collections) {
	{
		in := &in
		*out = make(Collections, len(*in))
		for key, val := range *in {
			var outVal *Collection
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(Collection)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collections.
func (in Collections) DeepCopy() Collections {
	if in == nil {
		return nil
	}
	out := new(Collections)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Peer) DeepCopyInto(out *Peer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Peer.
func (in *Peer) DeepCopy() *Peer {
	if in == nil {
		return nil
	}
	out := new(Peer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Peers) DeepCopyInto(out *Peers) {
	{
		in := &in
		*out = make(Peers, len(*in))
		for key, val := range *in {
			var outVal *Peer
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(Peer)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Peers.
func (in Peers) DeepCopy() Peers {
	if in == nil {
		return nil
	}
	out := new(Peers)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QdrantCluster) DeepCopyInto(out *QdrantCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QdrantCluster.
func (in *QdrantCluster) DeepCopy() *QdrantCluster {
	if in == nil {
		return nil
	}
	out := new(QdrantCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QdrantCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QdrantClusterList) DeepCopyInto(out *QdrantClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QdrantCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QdrantClusterList.
func (in *QdrantClusterList) DeepCopy() *QdrantClusterList {
	if in == nil {
		return nil
	}
	out := new(QdrantClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QdrantClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QdrantClusterSpec) DeepCopyInto(out *QdrantClusterSpec) {
	*out = *in
	if in.Statefulsets != nil {
		in, out := &in.Statefulsets, &out.Statefulsets
		*out = make([]StatefulSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QdrantClusterSpec.
func (in *QdrantClusterSpec) DeepCopy() *QdrantClusterSpec {
	if in == nil {
		return nil
	}
	out := new(QdrantClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QdrantClusterStatus) DeepCopyInto(out *QdrantClusterStatus) {
	*out = *in
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make(Peers, len(*in))
		for key, val := range *in {
			var outVal *Peer
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(Peer)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Collections != nil {
		in, out := &in.Collections, &out.Collections
		*out = make(Collections, len(*in))
		for key, val := range *in {
			var outVal *Collection
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(Collection)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.CordonedPeerIds != nil {
		in, out := &in.CordonedPeerIds, &out.CordonedPeerIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QdrantClusterStatus.
func (in *QdrantClusterStatus) DeepCopy() *QdrantClusterStatus {
	if in == nil {
		return nil
	}
	out := new(QdrantClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardInfo) DeepCopyInto(out *ShardInfo) {
	*out = *in
	if in.ShardId != nil {
		in, out := &in.ShardId, &out.ShardId
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardInfo.
func (in *ShardInfo) DeepCopy() *ShardInfo {
	if in == nil {
		return nil
	}
	out := new(ShardInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ShardsList) DeepCopyInto(out *ShardsList) {
	{
		in := &in
		*out = make(ShardsList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ShardInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardsList.
func (in ShardsList) DeepCopy() ShardsList {
	if in == nil {
		return nil
	}
	out := new(ShardsList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ShardsPerPeer) DeepCopyInto(out *ShardsPerPeer) {
	{
		in := &in
		*out = make(ShardsPerPeer, len(*in))
		for key, val := range *in {
			var outVal []*ShardInfo
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(ShardsList, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(ShardInfo)
						(*in).DeepCopyInto(*out)
					}
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardsPerPeer.
func (in ShardsPerPeer) DeepCopy() ShardsPerPeer {
	if in == nil {
		return nil
	}
	out := new(ShardsPerPeer)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSet) DeepCopyInto(out *StatefulSet) {
	*out = *in
	if in.VolumeClaim != nil {
		in, out := &in.VolumeClaim, &out.VolumeClaim
		*out = new(VolumeClaim)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSet.
func (in *StatefulSet) DeepCopy() *StatefulSet {
	if in == nil {
		return nil
	}
	out := new(StatefulSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeClaim) DeepCopyInto(out *VolumeClaim) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeClaim.
func (in *VolumeClaim) DeepCopy() *VolumeClaim {
	if in == nil {
		return nil
	}
	out := new(VolumeClaim)
	in.DeepCopyInto(out)
	return out
}
