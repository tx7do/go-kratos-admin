// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_notification_message_recipient.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-admin/api/gen/go/internal_message/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NotificationMessageRecipientService_List_FullMethodName   = "/admin.service.v1.NotificationMessageRecipientService/List"
	NotificationMessageRecipientService_Get_FullMethodName    = "/admin.service.v1.NotificationMessageRecipientService/Get"
	NotificationMessageRecipientService_Create_FullMethodName = "/admin.service.v1.NotificationMessageRecipientService/Create"
	NotificationMessageRecipientService_Update_FullMethodName = "/admin.service.v1.NotificationMessageRecipientService/Update"
	NotificationMessageRecipientService_Delete_FullMethodName = "/admin.service.v1.NotificationMessageRecipientService/Delete"
)

// NotificationMessageRecipientServiceClient is the client API for NotificationMessageRecipientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 通知消息接收者管理服务
type NotificationMessageRecipientServiceClient interface {
	// 查询通知消息接收者列表
	List(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListNotificationMessageRecipientResponse, error)
	// 查询通知消息接收者详情
	Get(ctx context.Context, in *v11.GetNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*v11.NotificationMessageRecipient, error)
	// 创建通知消息接收者
	Create(ctx context.Context, in *v11.CreateNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新通知消息接收者
	Update(ctx context.Context, in *v11.UpdateNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除通知消息接收者
	Delete(ctx context.Context, in *v11.DeleteNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type notificationMessageRecipientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationMessageRecipientServiceClient(cc grpc.ClientConnInterface) NotificationMessageRecipientServiceClient {
	return &notificationMessageRecipientServiceClient{cc}
}

func (c *notificationMessageRecipientServiceClient) List(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListNotificationMessageRecipientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.ListNotificationMessageRecipientResponse)
	err := c.cc.Invoke(ctx, NotificationMessageRecipientService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationMessageRecipientServiceClient) Get(ctx context.Context, in *v11.GetNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*v11.NotificationMessageRecipient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.NotificationMessageRecipient)
	err := c.cc.Invoke(ctx, NotificationMessageRecipientService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationMessageRecipientServiceClient) Create(ctx context.Context, in *v11.CreateNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NotificationMessageRecipientService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationMessageRecipientServiceClient) Update(ctx context.Context, in *v11.UpdateNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NotificationMessageRecipientService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationMessageRecipientServiceClient) Delete(ctx context.Context, in *v11.DeleteNotificationMessageRecipientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NotificationMessageRecipientService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationMessageRecipientServiceServer is the server API for NotificationMessageRecipientService service.
// All implementations must embed UnimplementedNotificationMessageRecipientServiceServer
// for forward compatibility.
//
// 通知消息接收者管理服务
type NotificationMessageRecipientServiceServer interface {
	// 查询通知消息接收者列表
	List(context.Context, *v1.PagingRequest) (*v11.ListNotificationMessageRecipientResponse, error)
	// 查询通知消息接收者详情
	Get(context.Context, *v11.GetNotificationMessageRecipientRequest) (*v11.NotificationMessageRecipient, error)
	// 创建通知消息接收者
	Create(context.Context, *v11.CreateNotificationMessageRecipientRequest) (*emptypb.Empty, error)
	// 更新通知消息接收者
	Update(context.Context, *v11.UpdateNotificationMessageRecipientRequest) (*emptypb.Empty, error)
	// 删除通知消息接收者
	Delete(context.Context, *v11.DeleteNotificationMessageRecipientRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNotificationMessageRecipientServiceServer()
}

// UnimplementedNotificationMessageRecipientServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotificationMessageRecipientServiceServer struct{}

func (UnimplementedNotificationMessageRecipientServiceServer) List(context.Context, *v1.PagingRequest) (*v11.ListNotificationMessageRecipientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNotificationMessageRecipientServiceServer) Get(context.Context, *v11.GetNotificationMessageRecipientRequest) (*v11.NotificationMessageRecipient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNotificationMessageRecipientServiceServer) Create(context.Context, *v11.CreateNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNotificationMessageRecipientServiceServer) Update(context.Context, *v11.UpdateNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNotificationMessageRecipientServiceServer) Delete(context.Context, *v11.DeleteNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNotificationMessageRecipientServiceServer) mustEmbedUnimplementedNotificationMessageRecipientServiceServer() {
}
func (UnimplementedNotificationMessageRecipientServiceServer) testEmbeddedByValue() {}

// UnsafeNotificationMessageRecipientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationMessageRecipientServiceServer will
// result in compilation errors.
type UnsafeNotificationMessageRecipientServiceServer interface {
	mustEmbedUnimplementedNotificationMessageRecipientServiceServer()
}

func RegisterNotificationMessageRecipientServiceServer(s grpc.ServiceRegistrar, srv NotificationMessageRecipientServiceServer) {
	// If the following call pancis, it indicates UnimplementedNotificationMessageRecipientServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotificationMessageRecipientService_ServiceDesc, srv)
}

func _NotificationMessageRecipientService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationMessageRecipientServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationMessageRecipientService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationMessageRecipientServiceServer).List(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationMessageRecipientService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetNotificationMessageRecipientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationMessageRecipientServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationMessageRecipientService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationMessageRecipientServiceServer).Get(ctx, req.(*v11.GetNotificationMessageRecipientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationMessageRecipientService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreateNotificationMessageRecipientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationMessageRecipientServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationMessageRecipientService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationMessageRecipientServiceServer).Create(ctx, req.(*v11.CreateNotificationMessageRecipientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationMessageRecipientService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateNotificationMessageRecipientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationMessageRecipientServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationMessageRecipientService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationMessageRecipientServiceServer).Update(ctx, req.(*v11.UpdateNotificationMessageRecipientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationMessageRecipientService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteNotificationMessageRecipientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationMessageRecipientServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationMessageRecipientService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationMessageRecipientServiceServer).Delete(ctx, req.(*v11.DeleteNotificationMessageRecipientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationMessageRecipientService_ServiceDesc is the grpc.ServiceDesc for NotificationMessageRecipientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationMessageRecipientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.NotificationMessageRecipientService",
	HandlerType: (*NotificationMessageRecipientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _NotificationMessageRecipientService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NotificationMessageRecipientService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _NotificationMessageRecipientService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NotificationMessageRecipientService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NotificationMessageRecipientService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_notification_message_recipient.proto",
}
