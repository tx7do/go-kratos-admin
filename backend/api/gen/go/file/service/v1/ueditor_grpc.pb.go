// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: file/service/v1/ueditor.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UEditorService_UEditorAPI_FullMethodName = "/file.service.v1.UEditorService/UEditorAPI"
	UEditorService_UploadFile_FullMethodName = "/file.service.v1.UEditorService/UploadFile"
)

// UEditorServiceClient is the client API for UEditorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// UEditor后端服务
type UEditorServiceClient interface {
	// UEditor API
	UEditorAPI(ctx context.Context, in *UEditorRequest, opts ...grpc.CallOption) (*UEditorResponse, error)
	// 上传文件
	UploadFile(ctx context.Context, in *UEditorUploadRequest, opts ...grpc.CallOption) (*UEditorUploadResponse, error)
}

type uEditorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUEditorServiceClient(cc grpc.ClientConnInterface) UEditorServiceClient {
	return &uEditorServiceClient{cc}
}

func (c *uEditorServiceClient) UEditorAPI(ctx context.Context, in *UEditorRequest, opts ...grpc.CallOption) (*UEditorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UEditorResponse)
	err := c.cc.Invoke(ctx, UEditorService_UEditorAPI_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uEditorServiceClient) UploadFile(ctx context.Context, in *UEditorUploadRequest, opts ...grpc.CallOption) (*UEditorUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UEditorUploadResponse)
	err := c.cc.Invoke(ctx, UEditorService_UploadFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UEditorServiceServer is the server API for UEditorService service.
// All implementations must embed UnimplementedUEditorServiceServer
// for forward compatibility.
//
// UEditor后端服务
type UEditorServiceServer interface {
	// UEditor API
	UEditorAPI(context.Context, *UEditorRequest) (*UEditorResponse, error)
	// 上传文件
	UploadFile(context.Context, *UEditorUploadRequest) (*UEditorUploadResponse, error)
	mustEmbedUnimplementedUEditorServiceServer()
}

// UnimplementedUEditorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUEditorServiceServer struct{}

func (UnimplementedUEditorServiceServer) UEditorAPI(context.Context, *UEditorRequest) (*UEditorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UEditorAPI not implemented")
}
func (UnimplementedUEditorServiceServer) UploadFile(context.Context, *UEditorUploadRequest) (*UEditorUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedUEditorServiceServer) mustEmbedUnimplementedUEditorServiceServer() {}
func (UnimplementedUEditorServiceServer) testEmbeddedByValue()                        {}

// UnsafeUEditorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UEditorServiceServer will
// result in compilation errors.
type UnsafeUEditorServiceServer interface {
	mustEmbedUnimplementedUEditorServiceServer()
}

func RegisterUEditorServiceServer(s grpc.ServiceRegistrar, srv UEditorServiceServer) {
	// If the following call pancis, it indicates UnimplementedUEditorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UEditorService_ServiceDesc, srv)
}

func _UEditorService_UEditorAPI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UEditorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UEditorServiceServer).UEditorAPI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UEditorService_UEditorAPI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UEditorServiceServer).UEditorAPI(ctx, req.(*UEditorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UEditorService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UEditorUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UEditorServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UEditorService_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UEditorServiceServer).UploadFile(ctx, req.(*UEditorUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UEditorService_ServiceDesc is the grpc.ServiceDesc for UEditorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UEditorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.service.v1.UEditorService",
	HandlerType: (*UEditorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UEditorAPI",
			Handler:    _UEditorService_UEditorAPI_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _UEditorService_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file/service/v1/ueditor.proto",
}
