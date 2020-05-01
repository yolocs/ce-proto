// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: schemas/extensions/options.proto

package options

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var file_schemas_extensions_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         80001,
		Name:          "cloudevents.options.transport_encoding",
		Tag:           "bytes,80001,opt,name=transport_encoding",
		Filename:      "schemas/extensions/options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         80002,
		Name:          "cloudevents.options.event_source",
		Tag:           "bytes,80002,opt,name=event_source",
		Filename:      "schemas/extensions/options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         80003,
		Name:          "cloudevents.options.event_subject",
		Tag:           "bytes,80003,opt,name=event_subject",
		Filename:      "schemas/extensions/options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         80004,
		Name:          "cloudevents.options.event_type",
		Tag:           "bytes,80004,opt,name=event_type",
		Filename:      "schemas/extensions/options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         80002,
		Name:          "cloudevents.options.ce_extension",
		Tag:           "varint,80002,opt,name=ce_extension",
		Filename:      "schemas/extensions/options.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional string transport_encoding = 80001;
	E_TransportEncoding = &file_schemas_extensions_options_proto_extTypes[0]
	// optional string event_source = 80002;
	E_EventSource = &file_schemas_extensions_options_proto_extTypes[1]
	// optional string event_subject = 80003;
	E_EventSubject = &file_schemas_extensions_options_proto_extTypes[2]
	// optional string event_type = 80004;
	E_EventType = &file_schemas_extensions_options_proto_extTypes[3]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional bool ce_extension = 80002;
	E_CeExtension = &file_schemas_extensions_options_proto_extTypes[4]
)

var File_schemas_extensions_options_proto protoreflect.FileDescriptor

var file_schemas_extensions_options_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x50, 0x0a, 0x12, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x81, 0xf1, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x44, 0x0a, 0x0c, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x82, 0xf1, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x3a, 0x46, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x83, 0xf1, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x40, 0x0a, 0x0a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0xf1, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x42, 0x0a, 0x0c, 0x63,
	0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x82, 0xf1, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f,
	0x6c, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_schemas_extensions_options_proto_goTypes = []interface{}{
	(*descriptor.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_schemas_extensions_options_proto_depIdxs = []int32{
	0, // 0: cloudevents.options.transport_encoding:extendee -> google.protobuf.MessageOptions
	0, // 1: cloudevents.options.event_source:extendee -> google.protobuf.MessageOptions
	0, // 2: cloudevents.options.event_subject:extendee -> google.protobuf.MessageOptions
	0, // 3: cloudevents.options.event_type:extendee -> google.protobuf.MessageOptions
	1, // 4: cloudevents.options.ce_extension:extendee -> google.protobuf.FieldOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schemas_extensions_options_proto_init() }
func file_schemas_extensions_options_proto_init() {
	if File_schemas_extensions_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schemas_extensions_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_schemas_extensions_options_proto_goTypes,
		DependencyIndexes: file_schemas_extensions_options_proto_depIdxs,
		ExtensionInfos:    file_schemas_extensions_options_proto_extTypes,
	}.Build()
	File_schemas_extensions_options_proto = out.File
	file_schemas_extensions_options_proto_rawDesc = nil
	file_schemas_extensions_options_proto_goTypes = nil
	file_schemas_extensions_options_proto_depIdxs = nil
}
