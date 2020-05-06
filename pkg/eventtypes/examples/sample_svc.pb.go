// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: schemas/examples/sample_svc.proto

package examples

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yolocs/ce-proto/pkg/eventtypes/options"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FileUploadedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Ext  string `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *FileUploadedEvent) Reset() {
	*x = FileUploadedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_examples_sample_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadedEvent) ProtoMessage() {}

func (x *FileUploadedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_examples_sample_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadedEvent.ProtoReflect.Descriptor instead.
func (*FileUploadedEvent) Descriptor() ([]byte, []int) {
	return file_schemas_examples_sample_svc_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadedEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileUploadedEvent) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileUploadedEvent) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type PaymentReceivedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	UserGroup string `protobuf:"bytes,2,opt,name=user_group,json=userGroup,proto3" json:"user_group,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PaymentReceivedEvent) Reset() {
	*x = PaymentReceivedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_examples_sample_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentReceivedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentReceivedEvent) ProtoMessage() {}

func (x *PaymentReceivedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_examples_sample_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentReceivedEvent.ProtoReflect.Descriptor instead.
func (*PaymentReceivedEvent) Descriptor() ([]byte, []int) {
	return file_schemas_examples_sample_svc_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentReceivedEvent) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PaymentReceivedEvent) GetUserGroup() string {
	if x != nil {
		return x.UserGroup
	}
	return ""
}

func (x *PaymentReceivedEvent) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_schemas_examples_sample_svc_proto protoreflect.FileDescriptor

var file_schemas_examples_sample_svc_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x76, 0x63, 0x1a, 0x20, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x11,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x03, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0x88, 0x27, 0x01, 0x52, 0x03, 0x65, 0x78,
	0x74, 0x3a, 0x08, 0x8a, 0x88, 0x27, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x77, 0x0a, 0x14, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x90, 0x88, 0x27, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x90, 0x88, 0x27, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x08, 0x8a, 0x88, 0x27, 0x04,
	0x4a, 0x53, 0x4f, 0x4e, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x6c, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_schemas_examples_sample_svc_proto_rawDescOnce sync.Once
	file_schemas_examples_sample_svc_proto_rawDescData = file_schemas_examples_sample_svc_proto_rawDesc
)

func file_schemas_examples_sample_svc_proto_rawDescGZIP() []byte {
	file_schemas_examples_sample_svc_proto_rawDescOnce.Do(func() {
		file_schemas_examples_sample_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_schemas_examples_sample_svc_proto_rawDescData)
	})
	return file_schemas_examples_sample_svc_proto_rawDescData
}

var file_schemas_examples_sample_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_schemas_examples_sample_svc_proto_goTypes = []interface{}{
	(*FileUploadedEvent)(nil),    // 0: com.example.samplesvc.FileUploadedEvent
	(*PaymentReceivedEvent)(nil), // 1: com.example.samplesvc.PaymentReceivedEvent
}
var file_schemas_examples_sample_svc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schemas_examples_sample_svc_proto_init() }
func file_schemas_examples_sample_svc_proto_init() {
	if File_schemas_examples_sample_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schemas_examples_sample_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadedEvent); i {
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
		file_schemas_examples_sample_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentReceivedEvent); i {
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
			RawDescriptor: file_schemas_examples_sample_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schemas_examples_sample_svc_proto_goTypes,
		DependencyIndexes: file_schemas_examples_sample_svc_proto_depIdxs,
		MessageInfos:      file_schemas_examples_sample_svc_proto_msgTypes,
	}.Build()
	File_schemas_examples_sample_svc_proto = out.File
	file_schemas_examples_sample_svc_proto_rawDesc = nil
	file_schemas_examples_sample_svc_proto_goTypes = nil
	file_schemas_examples_sample_svc_proto_depIdxs = nil
}
