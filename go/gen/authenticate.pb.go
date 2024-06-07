// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: protobuf/authenticate.proto

package gen

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

type Authenticate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Authenticate) Reset() {
	*x = Authenticate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_authenticate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authenticate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authenticate) ProtoMessage() {}

func (x *Authenticate) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_authenticate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authenticate.ProtoReflect.Descriptor instead.
func (*Authenticate) Descriptor() ([]byte, []int) {
	return file_protobuf_authenticate_proto_rawDescGZIP(), []int{0}
}

func (x *Authenticate) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_protobuf_authenticate_proto protoreflect.FileDescriptor

var file_protobuf_authenticate_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobuf_authenticate_proto_rawDescOnce sync.Once
	file_protobuf_authenticate_proto_rawDescData = file_protobuf_authenticate_proto_rawDesc
)

func file_protobuf_authenticate_proto_rawDescGZIP() []byte {
	file_protobuf_authenticate_proto_rawDescOnce.Do(func() {
		file_protobuf_authenticate_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_authenticate_proto_rawDescData)
	})
	return file_protobuf_authenticate_proto_rawDescData
}

var file_protobuf_authenticate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_authenticate_proto_goTypes = []interface{}{
	(*Authenticate)(nil), // 0: Authenticate
}
var file_protobuf_authenticate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_authenticate_proto_init() }
func file_protobuf_authenticate_proto_init() {
	if File_protobuf_authenticate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_authenticate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authenticate); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_authenticate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_authenticate_proto_goTypes,
		DependencyIndexes: file_protobuf_authenticate_proto_depIdxs,
		MessageInfos:      file_protobuf_authenticate_proto_msgTypes,
	}.Build()
	File_protobuf_authenticate_proto = out.File
	file_protobuf_authenticate_proto_rawDesc = nil
	file_protobuf_authenticate_proto_goTypes = nil
	file_protobuf_authenticate_proto_depIdxs = nil
}
