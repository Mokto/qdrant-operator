// Fork of the google.protobuf.Value with explicit support for integer values

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: json_with_int.proto

package qdrant

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//	The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

// Enum value maps for NullValue.
var (
	NullValue_name = map[int32]string{
		0: "NULL_VALUE",
	}
	NullValue_value = map[string]int32{
		"NULL_VALUE": 0,
	}
)

func (x NullValue) Enum() *NullValue {
	p := new(NullValue)
	*p = x
	return p
}

func (x NullValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NullValue) Descriptor() protoreflect.EnumDescriptor {
	return file_json_with_int_proto_enumTypes[0].Descriptor()
}

func (NullValue) Type() protoreflect.EnumType {
	return &file_json_with_int_proto_enumTypes[0]
}

func (x NullValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NullValue.Descriptor instead.
func (NullValue) EnumDescriptor() ([]byte, []int) {
	return file_json_with_int_proto_rawDescGZIP(), []int{0}
}

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is a JSON object.
type Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unordered map of dynamically typed values.
	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Struct) Reset() {
	*x = Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_json_with_int_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Struct) ProtoMessage() {}

func (x *Struct) ProtoReflect() protoreflect.Message {
	mi := &file_json_with_int_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Struct.ProtoReflect.Descriptor instead.
func (*Struct) Descriptor() ([]byte, []int) {
	return file_json_with_int_proto_rawDescGZIP(), []int{0}
}

func (x *Struct) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of those
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is a JSON value.
type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of value.
	//
	// Types that are assignable to Kind:
	//
	//	*Value_NullValue
	//	*Value_DoubleValue
	//	*Value_IntegerValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_ListValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_json_with_int_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_json_with_int_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_json_with_int_proto_rawDescGZIP(), []int{1}
}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Value) GetNullValue() NullValue {
	if x, ok := x.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

func (x *Value) GetDoubleValue() float64 {
	if x, ok := x.GetKind().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Value) GetIntegerValue() int64 {
	if x, ok := x.GetKind().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (x *Value) GetStringValue() string {
	if x, ok := x.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Value) GetBoolValue() bool {
	if x, ok := x.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Value) GetStructValue() *Struct {
	if x, ok := x.GetKind().(*Value_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (x *Value) GetListValue() *ListValue {
	if x, ok := x.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	// Represents a null value.
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=qdrant.NullValue,oneof"`
}

type Value_DoubleValue struct {
	// Represents a double value.
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	// Represents an integer value
	IntegerValue int64 `protobuf:"varint,3,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_StringValue struct {
	// Represents a string value.
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BoolValue struct {
	// Represents a boolean value.
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_StructValue struct {
	// Represents a structured value.
	StructValue *Struct `protobuf:"bytes,6,opt,name=struct_value,json=structValue,proto3,oneof"`
}

type Value_ListValue struct {
	// Represents a repeated `Value`.
	ListValue *ListValue `protobuf:"bytes,7,opt,name=list_value,json=listValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_IntegerValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_StructValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is a JSON array.
type ListValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Repeated field of dynamically typed values.
	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_json_with_int_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_json_with_int_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_json_with_int_proto_rawDescGZIP(), []int{2}
}

func (x *ListValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_json_with_int_proto protoreflect.FileDescriptor

var file_json_with_int_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x86, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x48, 0x0a, 0x0b,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x71,
	0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbe, 0x02, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x32, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x4e, 0x75,
	0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x71,
	0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x32, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0x1b, 0x0a, 0x09, 0x4e,
	0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x55, 0x4c, 0x4c,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x42, 0x67, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e,
	0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x10, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x51, 0x58,
	0x58, 0xaa, 0x02, 0x06, 0x51, 0x64, 0x72, 0x61, 0x6e, 0x74, 0xca, 0x02, 0x06, 0x51, 0x64, 0x72,
	0x61, 0x6e, 0x74, 0xe2, 0x02, 0x12, 0x51, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x51, 0x64, 0x72, 0x61, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_json_with_int_proto_rawDescOnce sync.Once
	file_json_with_int_proto_rawDescData = file_json_with_int_proto_rawDesc
)

func file_json_with_int_proto_rawDescGZIP() []byte {
	file_json_with_int_proto_rawDescOnce.Do(func() {
		file_json_with_int_proto_rawDescData = protoimpl.X.CompressGZIP(file_json_with_int_proto_rawDescData)
	})
	return file_json_with_int_proto_rawDescData
}

var file_json_with_int_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_json_with_int_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_json_with_int_proto_goTypes = []any{
	(NullValue)(0),    // 0: qdrant.NullValue
	(*Struct)(nil),    // 1: qdrant.Struct
	(*Value)(nil),     // 2: qdrant.Value
	(*ListValue)(nil), // 3: qdrant.ListValue
	nil,               // 4: qdrant.Struct.FieldsEntry
}
var file_json_with_int_proto_depIdxs = []int32{
	4, // 0: qdrant.Struct.fields:type_name -> qdrant.Struct.FieldsEntry
	0, // 1: qdrant.Value.null_value:type_name -> qdrant.NullValue
	1, // 2: qdrant.Value.struct_value:type_name -> qdrant.Struct
	3, // 3: qdrant.Value.list_value:type_name -> qdrant.ListValue
	2, // 4: qdrant.ListValue.values:type_name -> qdrant.Value
	2, // 5: qdrant.Struct.FieldsEntry.value:type_name -> qdrant.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_json_with_int_proto_init() }
func file_json_with_int_proto_init() {
	if File_json_with_int_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_json_with_int_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Struct); i {
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
		file_json_with_int_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Value); i {
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
		file_json_with_int_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListValue); i {
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
	file_json_with_int_proto_msgTypes[1].OneofWrappers = []any{
		(*Value_NullValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_json_with_int_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_json_with_int_proto_goTypes,
		DependencyIndexes: file_json_with_int_proto_depIdxs,
		EnumInfos:         file_json_with_int_proto_enumTypes,
		MessageInfos:      file_json_with_int_proto_msgTypes,
	}.Build()
	File_json_with_int_proto = out.File
	file_json_with_int_proto_rawDesc = nil
	file_json_with_int_proto_goTypes = nil
	file_json_with_int_proto_depIdxs = nil
}
