// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: internal_message/service/v1/message.proto

package servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 消息状态
type MessageStatus int32

const (
	MessageStatus_MessageStatus_Unknown MessageStatus = 0 // 未知状态
	MessageStatus_DRAFT                 MessageStatus = 1 // 草稿
	MessageStatus_PUBLISHED             MessageStatus = 2 // 已发布
	MessageStatus_SCHEDULED             MessageStatus = 3 // 定时发布
	MessageStatus_REVOKED               MessageStatus = 4 // 已撤销
	MessageStatus_ARCHIVED              MessageStatus = 5 // 已归档
	MessageStatus_DELETED               MessageStatus = 6 // 已删除
	MessageStatus_SENT                  MessageStatus = 7 // 已发送
	MessageStatus_RECEIVED              MessageStatus = 8 // 已接收
	MessageStatus_READ                  MessageStatus = 9 // 已接收
)

// Enum value maps for MessageStatus.
var (
	MessageStatus_name = map[int32]string{
		0: "MessageStatus_Unknown",
		1: "DRAFT",
		2: "PUBLISHED",
		3: "SCHEDULED",
		4: "REVOKED",
		5: "ARCHIVED",
		6: "DELETED",
		7: "SENT",
		8: "RECEIVED",
		9: "READ",
	}
	MessageStatus_value = map[string]int32{
		"MessageStatus_Unknown": 0,
		"DRAFT":                 1,
		"PUBLISHED":             2,
		"SCHEDULED":             3,
		"REVOKED":               4,
		"ARCHIVED":              5,
		"DELETED":               6,
		"SENT":                  7,
		"RECEIVED":              8,
		"READ":                  9,
	}
)

func (x MessageStatus) Enum() *MessageStatus {
	p := new(MessageStatus)
	*p = x
	return p
}

func (x MessageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_message_service_v1_message_proto_enumTypes[0].Descriptor()
}

func (MessageStatus) Type() protoreflect.EnumType {
	return &file_internal_message_service_v1_message_proto_enumTypes[0]
}

func (x MessageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageStatus.Descriptor instead.
func (MessageStatus) EnumDescriptor() ([]byte, []int) {
	return file_internal_message_service_v1_message_proto_rawDescGZIP(), []int{0}
}

var File_internal_message_service_v1_message_proto protoreflect.FileDescriptor

var file_internal_message_service_v1_message_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x9d, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x07,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x08, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x09, 0x42, 0xf8, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x53,
	0x58, 0xaa, 0x02, 0x1a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x26, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internal_message_service_v1_message_proto_rawDescOnce sync.Once
	file_internal_message_service_v1_message_proto_rawDescData []byte
)

func file_internal_message_service_v1_message_proto_rawDescGZIP() []byte {
	file_internal_message_service_v1_message_proto_rawDescOnce.Do(func() {
		file_internal_message_service_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_message_service_v1_message_proto_rawDesc), len(file_internal_message_service_v1_message_proto_rawDesc)))
	})
	return file_internal_message_service_v1_message_proto_rawDescData
}

var file_internal_message_service_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_message_service_v1_message_proto_goTypes = []any{
	(MessageStatus)(0), // 0: internal_message.service.v1.MessageStatus
}
var file_internal_message_service_v1_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_message_service_v1_message_proto_init() }
func file_internal_message_service_v1_message_proto_init() {
	if File_internal_message_service_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_message_service_v1_message_proto_rawDesc), len(file_internal_message_service_v1_message_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_message_service_v1_message_proto_goTypes,
		DependencyIndexes: file_internal_message_service_v1_message_proto_depIdxs,
		EnumInfos:         file_internal_message_service_v1_message_proto_enumTypes,
	}.Build()
	File_internal_message_service_v1_message_proto = out.File
	file_internal_message_service_v1_message_proto_goTypes = nil
	file_internal_message_service_v1_message_proto_depIdxs = nil
}
