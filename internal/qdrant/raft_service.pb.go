// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: raft_service.proto

package qdrant

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RaftMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RaftMessage) Reset() {
	*x = RaftMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftMessage) ProtoMessage() {}

func (x *RaftMessage) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftMessage.ProtoReflect.Descriptor instead.
func (*RaftMessage) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{0}
}

func (x *RaftMessage) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type AllPeers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllPeers    []*Peer `protobuf:"bytes,1,rep,name=all_peers,json=allPeers,proto3" json:"all_peers,omitempty"`
	FirstPeerId uint64  `protobuf:"varint,2,opt,name=first_peer_id,json=firstPeerId,proto3" json:"first_peer_id,omitempty"`
}

func (x *AllPeers) Reset() {
	*x = AllPeers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPeers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPeers) ProtoMessage() {}

func (x *AllPeers) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllPeers.ProtoReflect.Descriptor instead.
func (*AllPeers) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{1}
}

func (x *AllPeers) GetAllPeers() []*Peer {
	if x != nil {
		return x.AllPeers
	}
	return nil
}

func (x *AllPeers) GetFirstPeerId() uint64 {
	if x != nil {
		return x.FirstPeerId
	}
	return 0
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Id  uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{2}
}

func (x *Peer) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Peer) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddPeerToKnownMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri  *string `protobuf:"bytes,1,opt,name=uri,proto3,oneof" json:"uri,omitempty"`
	Port *uint32 `protobuf:"varint,2,opt,name=port,proto3,oneof" json:"port,omitempty"`
	Id   uint64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddPeerToKnownMessage) Reset() {
	*x = AddPeerToKnownMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPeerToKnownMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPeerToKnownMessage) ProtoMessage() {}

func (x *AddPeerToKnownMessage) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPeerToKnownMessage.ProtoReflect.Descriptor instead.
func (*AddPeerToKnownMessage) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddPeerToKnownMessage) GetUri() string {
	if x != nil && x.Uri != nil {
		return *x.Uri
	}
	return ""
}

func (x *AddPeerToKnownMessage) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *AddPeerToKnownMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PeerId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PeerId) Reset() {
	*x = PeerId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerId) ProtoMessage() {}

func (x *PeerId) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerId.ProtoReflect.Descriptor instead.
func (*PeerId) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{4}
}

func (x *PeerId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Uri struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *Uri) Reset() {
	*x = Uri{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uri) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uri) ProtoMessage() {}

func (x *Uri) ProtoReflect() protoreflect.Message {
	mi := &file_raft_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uri.ProtoReflect.Descriptor instead.
func (*Uri) Descriptor() ([]byte, []int) {
	return file_raft_service_proto_rawDescGZIP(), []int{5}
}

func (x *Uri) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

var File_raft_service_proto protoreflect.FileDescriptor

var file_raft_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x52, 0x61, 0x66,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x59, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x29,
	0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x08, 0x61, 0x6c, 0x6c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x28, 0x0a,
	0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x50, 0x65,
	0x65, 0x72, 0x54, 0x6f, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x69, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x18, 0x0a, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x03, 0x55,
	0x72, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x32, 0xe4, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x66, 0x74, 0x12, 0x33, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52,
	0x61, 0x66, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x57, 0x68, 0x6f, 0x49, 0x73, 0x12, 0x0e, 0x2e, 0x71, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x71, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x72, 0x69, 0x12, 0x41, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50,
	0x65, 0x65, 0x72, 0x54, 0x6f, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x1d, 0x2e, 0x71, 0x64, 0x72,
	0x61, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x65, 0x72, 0x54, 0x6f, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x71, 0x64, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x14, 0x41,
	0x64, 0x64, 0x50, 0x65, 0x65, 0x72, 0x41, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x67, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x10, 0x52, 0x61, 0x66, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0xa2, 0x02,
	0x03, 0x51, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x51, 0x64, 0x72, 0x61, 0x6e, 0x74, 0xca, 0x02, 0x06,
	0x51, 0x64, 0x72, 0x61, 0x6e, 0x74, 0xe2, 0x02, 0x12, 0x51, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x51, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raft_service_proto_rawDescOnce sync.Once
	file_raft_service_proto_rawDescData = file_raft_service_proto_rawDesc
)

func file_raft_service_proto_rawDescGZIP() []byte {
	file_raft_service_proto_rawDescOnce.Do(func() {
		file_raft_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_service_proto_rawDescData)
	})
	return file_raft_service_proto_rawDescData
}

var file_raft_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_raft_service_proto_goTypes = []any{
	(*RaftMessage)(nil),           // 0: qdrant.RaftMessage
	(*AllPeers)(nil),              // 1: qdrant.AllPeers
	(*Peer)(nil),                  // 2: qdrant.Peer
	(*AddPeerToKnownMessage)(nil), // 3: qdrant.AddPeerToKnownMessage
	(*PeerId)(nil),                // 4: qdrant.PeerId
	(*Uri)(nil),                   // 5: qdrant.Uri
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_raft_service_proto_depIdxs = []int32{
	2, // 0: qdrant.AllPeers.all_peers:type_name -> qdrant.Peer
	0, // 1: qdrant.Raft.Send:input_type -> qdrant.RaftMessage
	4, // 2: qdrant.Raft.WhoIs:input_type -> qdrant.PeerId
	3, // 3: qdrant.Raft.AddPeerToKnown:input_type -> qdrant.AddPeerToKnownMessage
	4, // 4: qdrant.Raft.AddPeerAsParticipant:input_type -> qdrant.PeerId
	6, // 5: qdrant.Raft.Send:output_type -> google.protobuf.Empty
	5, // 6: qdrant.Raft.WhoIs:output_type -> qdrant.Uri
	1, // 7: qdrant.Raft.AddPeerToKnown:output_type -> qdrant.AllPeers
	6, // 8: qdrant.Raft.AddPeerAsParticipant:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_raft_service_proto_init() }
func file_raft_service_proto_init() {
	if File_raft_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raft_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RaftMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AllPeers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Peer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddPeerToKnownMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PeerId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Uri); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_raft_service_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raft_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raft_service_proto_goTypes,
		DependencyIndexes: file_raft_service_proto_depIdxs,
		MessageInfos:      file_raft_service_proto_msgTypes,
	}.Build()
	File_raft_service_proto = out.File
	file_raft_service_proto_rawDesc = nil
	file_raft_service_proto_goTypes = nil
	file_raft_service_proto_depIdxs = nil
}
