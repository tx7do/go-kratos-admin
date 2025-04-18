// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_oss.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "kratos-admin/api/gen/go/file/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OssService_OssUploadUrl_FullMethodName   = "/admin.service.v1.OssService/OssUploadUrl"
	OssService_PostUploadFile_FullMethodName = "/admin.service.v1.OssService/PostUploadFile"
	OssService_PutUploadFile_FullMethodName  = "/admin.service.v1.OssService/PutUploadFile"
)

// OssServiceClient is the client API for OssService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// OSS服务
type OssServiceClient interface {
	// 获取对象存储（OSS）上传用的预签名链接
	OssUploadUrl(ctx context.Context, in *v1.OssUploadUrlRequest, opts ...grpc.CallOption) (*v1.OssUploadUrlResponse, error)
	// POST方法上传文件
	PostUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadOssFileRequest, v1.UploadOssFileResponse], error)
	// PUT方法上传文件
	PutUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadOssFileRequest, v1.UploadOssFileResponse], error)
}

type ossServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOssServiceClient(cc grpc.ClientConnInterface) OssServiceClient {
	return &ossServiceClient{cc}
}

func (c *ossServiceClient) OssUploadUrl(ctx context.Context, in *v1.OssUploadUrlRequest, opts ...grpc.CallOption) (*v1.OssUploadUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.OssUploadUrlResponse)
	err := c.cc.Invoke(ctx, OssService_OssUploadUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossServiceClient) PostUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadOssFileRequest, v1.UploadOssFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OssService_ServiceDesc.Streams[0], OssService_PostUploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[v1.UploadOssFileRequest, v1.UploadOssFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OssService_PostUploadFileClient = grpc.ClientStreamingClient[v1.UploadOssFileRequest, v1.UploadOssFileResponse]

func (c *ossServiceClient) PutUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadOssFileRequest, v1.UploadOssFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OssService_ServiceDesc.Streams[1], OssService_PutUploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[v1.UploadOssFileRequest, v1.UploadOssFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OssService_PutUploadFileClient = grpc.ClientStreamingClient[v1.UploadOssFileRequest, v1.UploadOssFileResponse]

// OssServiceServer is the server API for OssService service.
// All implementations must embed UnimplementedOssServiceServer
// for forward compatibility.
//
// OSS服务
type OssServiceServer interface {
	// 获取对象存储（OSS）上传用的预签名链接
	OssUploadUrl(context.Context, *v1.OssUploadUrlRequest) (*v1.OssUploadUrlResponse, error)
	// POST方法上传文件
	PostUploadFile(grpc.ClientStreamingServer[v1.UploadOssFileRequest, v1.UploadOssFileResponse]) error
	// PUT方法上传文件
	PutUploadFile(grpc.ClientStreamingServer[v1.UploadOssFileRequest, v1.UploadOssFileResponse]) error
	mustEmbedUnimplementedOssServiceServer()
}

// UnimplementedOssServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOssServiceServer struct{}

func (UnimplementedOssServiceServer) OssUploadUrl(context.Context, *v1.OssUploadUrlRequest) (*v1.OssUploadUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssUploadUrl not implemented")
}
func (UnimplementedOssServiceServer) PostUploadFile(grpc.ClientStreamingServer[v1.UploadOssFileRequest, v1.UploadOssFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PostUploadFile not implemented")
}
func (UnimplementedOssServiceServer) PutUploadFile(grpc.ClientStreamingServer[v1.UploadOssFileRequest, v1.UploadOssFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PutUploadFile not implemented")
}
func (UnimplementedOssServiceServer) mustEmbedUnimplementedOssServiceServer() {}
func (UnimplementedOssServiceServer) testEmbeddedByValue()                    {}

// UnsafeOssServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OssServiceServer will
// result in compilation errors.
type UnsafeOssServiceServer interface {
	mustEmbedUnimplementedOssServiceServer()
}

func RegisterOssServiceServer(s grpc.ServiceRegistrar, srv OssServiceServer) {
	// If the following call pancis, it indicates UnimplementedOssServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OssService_ServiceDesc, srv)
}

func _OssService_OssUploadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.OssUploadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServiceServer).OssUploadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OssService_OssUploadUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServiceServer).OssUploadUrl(ctx, req.(*v1.OssUploadUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OssService_PostUploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OssServiceServer).PostUploadFile(&grpc.GenericServerStream[v1.UploadOssFileRequest, v1.UploadOssFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OssService_PostUploadFileServer = grpc.ClientStreamingServer[v1.UploadOssFileRequest, v1.UploadOssFileResponse]

func _OssService_PutUploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OssServiceServer).PutUploadFile(&grpc.GenericServerStream[v1.UploadOssFileRequest, v1.UploadOssFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OssService_PutUploadFileServer = grpc.ClientStreamingServer[v1.UploadOssFileRequest, v1.UploadOssFileResponse]

// OssService_ServiceDesc is the grpc.ServiceDesc for OssService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OssService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.OssService",
	HandlerType: (*OssServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OssUploadUrl",
			Handler:    _OssService_OssUploadUrl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostUploadFile",
			Handler:       _OssService_PostUploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PutUploadFile",
			Handler:       _OssService_PutUploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "admin/service/v1/i_oss.proto",
}
