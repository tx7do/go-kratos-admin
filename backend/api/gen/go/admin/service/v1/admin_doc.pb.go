// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: admin/service/v1/admin_doc.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_admin_service_v1_admin_doc_proto protoreflect.FileDescriptor

const file_admin_service_v1_admin_doc_proto_rawDesc = "" +
	"\n" +
	" admin/service/v1/admin_doc.proto\x12\x10admin.service.v1\x1a$gnostic/openapi/v3/annotations.protoB\xaf\x06\xbaG\xf0\x04\x12\xb8\x01\n" +
	"\x10Kratos Admin API\x12\x10Kratos Admin API\"C\n" +
	"\x05tx7do\x12%https://github.com/tx7do/kratos-admin\x1a\x13yanglinbo@gmail.com*H\n" +
	"\vMIT License\x129https://github.com/tx7do/kratos-admin/blob/master/LICENSE2\x031.0*\x96\x03\n" +
	"\xd4\x01\n" +
	"\xd1\x01\n" +
	"\fKratosStatus\x12\xc0\x01\n" +
	"\xbd\x01\xca\x01\x06object\xfa\x01\x9b\x01\n" +
	"'\n" +
	"\x04code\x12\x1f\n" +
	"\x1d\xca\x01\x06number\x92\x02\t错误码\x9a\x02\x05int32\n" +
	"%\n" +
	"\amessage\x12\x1a\n" +
	"\x18\xca\x01\x06string\x92\x02\f错误消息\n" +
	"$\n" +
	"\x06reason\x12\x1a\n" +
	"\x18\xca\x01\x06string\x92\x02\f错误原因\n" +
	"#\n" +
	"\bmetadata\x12\x17\n" +
	"\x15\xca\x01\x06object\x92\x02\t元数据\x92\x02\x12Kratos错误返回\x12g\n" +
	"e\n" +
	"\adefault\x12Z\n" +
	"X\n" +
	"\x17default kratos response\x1a=\n" +
	";\n" +
	"\x10application/json\x12'\n" +
	"%\x12#\n" +
	"!#/components/schemas/KratosStatus:T\n" +
	"R\n" +
	"\x14OAuth2PasswordBearer\x12:\n" +
	"8\n" +
	"\x06oauth2:.\x12,\x12\x0f/admin/v1/login\x1a\x17/admin/v1/refresh_token\"\x002\x1a\n" +
	"\x18\n" +
	"\x14OAuth2PasswordBearer\x12\x00\n" +
	"\x14com.admin.service.v1B\rAdminDocProtoP\x01Z2kratos-admin/api/gen/go/admin/service/v1;servicev1\xa2\x02\x03ASX\xaa\x02\x10Admin.Service.V1\xca\x02\x10Admin\\Service\\V1\xe2\x02\x1cAdmin\\Service\\V1\\GPBMetadata\xea\x02\x12Admin::Service::V1b\x06proto3"

var file_admin_service_v1_admin_doc_proto_goTypes = []any{}
var file_admin_service_v1_admin_doc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_service_v1_admin_doc_proto_init() }
func file_admin_service_v1_admin_doc_proto_init() {
	if File_admin_service_v1_admin_doc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_admin_service_v1_admin_doc_proto_rawDesc), len(file_admin_service_v1_admin_doc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_service_v1_admin_doc_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_admin_doc_proto_depIdxs,
	}.Build()
	File_admin_service_v1_admin_doc_proto = out.File
	file_admin_service_v1_admin_doc_proto_goTypes = nil
	file_admin_service_v1_admin_doc_proto_depIdxs = nil
}
