// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: authentication/service/v1/user_credential.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserCredentialService_List_FullMethodName             = "/authentication.service.v1.UserCredentialService/List"
	UserCredentialService_Get_FullMethodName              = "/authentication.service.v1.UserCredentialService/Get"
	UserCredentialService_GetByIdentifier_FullMethodName  = "/authentication.service.v1.UserCredentialService/GetByIdentifier"
	UserCredentialService_Create_FullMethodName           = "/authentication.service.v1.UserCredentialService/Create"
	UserCredentialService_Update_FullMethodName           = "/authentication.service.v1.UserCredentialService/Update"
	UserCredentialService_Delete_FullMethodName           = "/authentication.service.v1.UserCredentialService/Delete"
	UserCredentialService_VerifyCredential_FullMethodName = "/authentication.service.v1.UserCredentialService/VerifyCredential"
	UserCredentialService_ChangeCredential_FullMethodName = "/authentication.service.v1.UserCredentialService/ChangeCredential"
	UserCredentialService_ResetCredential_FullMethodName  = "/authentication.service.v1.UserCredentialService/ResetCredential"
)

// UserCredentialServiceClient is the client API for UserCredentialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 用户认证服务
type UserCredentialServiceClient interface {
	// 查询列表
	List(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListUserCredentialResponse, error)
	// 查询
	Get(ctx context.Context, in *GetUserCredentialRequest, opts ...grpc.CallOption) (*UserCredential, error)
	GetByIdentifier(ctx context.Context, in *GetUserCredentialByIdentifierRequest, opts ...grpc.CallOption) (*UserCredential, error)
	// 创建
	Create(ctx context.Context, in *CreateUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新
	Update(ctx context.Context, in *UpdateUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除
	Delete(ctx context.Context, in *DeleteUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 验证凭证
	VerifyCredential(ctx context.Context, in *VerifyCredentialRequest, opts ...grpc.CallOption) (*VerifyCredentialResponse, error)
	// 修改凭证
	ChangeCredential(ctx context.Context, in *ChangeCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 重设凭证
	ResetCredential(ctx context.Context, in *ResetCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userCredentialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCredentialServiceClient(cc grpc.ClientConnInterface) UserCredentialServiceClient {
	return &userCredentialServiceClient{cc}
}

func (c *userCredentialServiceClient) List(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*ListUserCredentialResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserCredentialResponse)
	err := c.cc.Invoke(ctx, UserCredentialService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) Get(ctx context.Context, in *GetUserCredentialRequest, opts ...grpc.CallOption) (*UserCredential, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserCredential)
	err := c.cc.Invoke(ctx, UserCredentialService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) GetByIdentifier(ctx context.Context, in *GetUserCredentialByIdentifierRequest, opts ...grpc.CallOption) (*UserCredential, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserCredential)
	err := c.cc.Invoke(ctx, UserCredentialService_GetByIdentifier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) Create(ctx context.Context, in *CreateUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserCredentialService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) Update(ctx context.Context, in *UpdateUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserCredentialService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) Delete(ctx context.Context, in *DeleteUserCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserCredentialService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) VerifyCredential(ctx context.Context, in *VerifyCredentialRequest, opts ...grpc.CallOption) (*VerifyCredentialResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyCredentialResponse)
	err := c.cc.Invoke(ctx, UserCredentialService_VerifyCredential_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) ChangeCredential(ctx context.Context, in *ChangeCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserCredentialService_ChangeCredential_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialServiceClient) ResetCredential(ctx context.Context, in *ResetCredentialRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserCredentialService_ResetCredential_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCredentialServiceServer is the server API for UserCredentialService service.
// All implementations must embed UnimplementedUserCredentialServiceServer
// for forward compatibility.
//
// 用户认证服务
type UserCredentialServiceServer interface {
	// 查询列表
	List(context.Context, *v1.PagingRequest) (*ListUserCredentialResponse, error)
	// 查询
	Get(context.Context, *GetUserCredentialRequest) (*UserCredential, error)
	GetByIdentifier(context.Context, *GetUserCredentialByIdentifierRequest) (*UserCredential, error)
	// 创建
	Create(context.Context, *CreateUserCredentialRequest) (*emptypb.Empty, error)
	// 更新
	Update(context.Context, *UpdateUserCredentialRequest) (*emptypb.Empty, error)
	// 删除
	Delete(context.Context, *DeleteUserCredentialRequest) (*emptypb.Empty, error)
	// 验证凭证
	VerifyCredential(context.Context, *VerifyCredentialRequest) (*VerifyCredentialResponse, error)
	// 修改凭证
	ChangeCredential(context.Context, *ChangeCredentialRequest) (*emptypb.Empty, error)
	// 重设凭证
	ResetCredential(context.Context, *ResetCredentialRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserCredentialServiceServer()
}

// UnimplementedUserCredentialServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserCredentialServiceServer struct{}

func (UnimplementedUserCredentialServiceServer) List(context.Context, *v1.PagingRequest) (*ListUserCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserCredentialServiceServer) Get(context.Context, *GetUserCredentialRequest) (*UserCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserCredentialServiceServer) GetByIdentifier(context.Context, *GetUserCredentialByIdentifierRequest) (*UserCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdentifier not implemented")
}
func (UnimplementedUserCredentialServiceServer) Create(context.Context, *CreateUserCredentialRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserCredentialServiceServer) Update(context.Context, *UpdateUserCredentialRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserCredentialServiceServer) Delete(context.Context, *DeleteUserCredentialRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserCredentialServiceServer) VerifyCredential(context.Context, *VerifyCredentialRequest) (*VerifyCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCredential not implemented")
}
func (UnimplementedUserCredentialServiceServer) ChangeCredential(context.Context, *ChangeCredentialRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCredential not implemented")
}
func (UnimplementedUserCredentialServiceServer) ResetCredential(context.Context, *ResetCredentialRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCredential not implemented")
}
func (UnimplementedUserCredentialServiceServer) mustEmbedUnimplementedUserCredentialServiceServer() {}
func (UnimplementedUserCredentialServiceServer) testEmbeddedByValue()                               {}

// UnsafeUserCredentialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCredentialServiceServer will
// result in compilation errors.
type UnsafeUserCredentialServiceServer interface {
	mustEmbedUnimplementedUserCredentialServiceServer()
}

func RegisterUserCredentialServiceServer(s grpc.ServiceRegistrar, srv UserCredentialServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserCredentialServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserCredentialService_ServiceDesc, srv)
}

func _UserCredentialService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).List(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).Get(ctx, req.(*GetUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_GetByIdentifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCredentialByIdentifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).GetByIdentifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_GetByIdentifier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).GetByIdentifier(ctx, req.(*GetUserCredentialByIdentifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).Create(ctx, req.(*CreateUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).Update(ctx, req.(*UpdateUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).Delete(ctx, req.(*DeleteUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_VerifyCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).VerifyCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_VerifyCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).VerifyCredential(ctx, req.(*VerifyCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_ChangeCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).ChangeCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_ChangeCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).ChangeCredential(ctx, req.(*ChangeCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialService_ResetCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialServiceServer).ResetCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCredentialService_ResetCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialServiceServer).ResetCredential(ctx, req.(*ResetCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCredentialService_ServiceDesc is the grpc.ServiceDesc for UserCredentialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCredentialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.service.v1.UserCredentialService",
	HandlerType: (*UserCredentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserCredentialService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserCredentialService_Get_Handler,
		},
		{
			MethodName: "GetByIdentifier",
			Handler:    _UserCredentialService_GetByIdentifier_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserCredentialService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserCredentialService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserCredentialService_Delete_Handler,
		},
		{
			MethodName: "VerifyCredential",
			Handler:    _UserCredentialService_VerifyCredential_Handler,
		},
		{
			MethodName: "ChangeCredential",
			Handler:    _UserCredentialService_ChangeCredential_Handler,
		},
		{
			MethodName: "ResetCredential",
			Handler:    _UserCredentialService_ResetCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication/service/v1/user_credential.proto",
}
