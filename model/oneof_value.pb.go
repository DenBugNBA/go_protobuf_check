// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/oneof_value.proto

package model

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

type SomeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SomeField:
	//
	//	*SomeValue_Str
	//	*SomeValue_I32
	SomeField isSomeValue_SomeField `protobuf_oneof:"someField"`
}

func (x *SomeValue) Reset() {
	*x = SomeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oneof_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeValue) ProtoMessage() {}

func (x *SomeValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oneof_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeValue.ProtoReflect.Descriptor instead.
func (*SomeValue) Descriptor() ([]byte, []int) {
	return file_proto_oneof_value_proto_rawDescGZIP(), []int{0}
}

func (m *SomeValue) GetSomeField() isSomeValue_SomeField {
	if m != nil {
		return m.SomeField
	}
	return nil
}

func (x *SomeValue) GetStr() string {
	if x, ok := x.GetSomeField().(*SomeValue_Str); ok {
		return x.Str
	}
	return ""
}

func (x *SomeValue) GetI32() int32 {
	if x, ok := x.GetSomeField().(*SomeValue_I32); ok {
		return x.I32
	}
	return 0
}

type isSomeValue_SomeField interface {
	isSomeValue_SomeField()
}

type SomeValue_Str struct {
	Str string `protobuf:"bytes,10,opt,name=str,proto3,oneof"`
}

type SomeValue_I32 struct {
	I32 int32 `protobuf:"varint,12,opt,name=i32,proto3,oneof"`
}

func (*SomeValue_Str) isSomeValue_SomeField() {}

func (*SomeValue_I32) isSomeValue_SomeField() {}

var File_proto_oneof_value_proto protoreflect.FileDescriptor

var file_proto_oneof_value_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x22, 0x40, 0x0a, 0x09, 0x53, 0x6f, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x73, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x03, 0x69, 0x33, 0x32, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_oneof_value_proto_rawDescOnce sync.Once
	file_proto_oneof_value_proto_rawDescData = file_proto_oneof_value_proto_rawDesc
)

func file_proto_oneof_value_proto_rawDescGZIP() []byte {
	file_proto_oneof_value_proto_rawDescOnce.Do(func() {
		file_proto_oneof_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_oneof_value_proto_rawDescData)
	})
	return file_proto_oneof_value_proto_rawDescData
}

var file_proto_oneof_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_oneof_value_proto_goTypes = []interface{}{
	(*SomeValue)(nil), // 0: protobuf.SomeValue
}
var file_proto_oneof_value_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_oneof_value_proto_init() }
func file_proto_oneof_value_proto_init() {
	if File_proto_oneof_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_oneof_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeValue); i {
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
	file_proto_oneof_value_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SomeValue_Str)(nil),
		(*SomeValue_I32)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_oneof_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_oneof_value_proto_goTypes,
		DependencyIndexes: file_proto_oneof_value_proto_depIdxs,
		MessageInfos:      file_proto_oneof_value_proto_msgTypes,
	}.Build()
	File_proto_oneof_value_proto = out.File
	file_proto_oneof_value_proto_rawDesc = nil
	file_proto_oneof_value_proto_goTypes = nil
	file_proto_oneof_value_proto_depIdxs = nil
}