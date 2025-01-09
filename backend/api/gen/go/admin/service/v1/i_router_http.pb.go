// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             (unknown)
// source: admin/service/v1/i_router.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRouterServiceListPermissionCode = "/admin.service.v1.RouterService/ListPermissionCode"
const OperationRouterServiceListRoute = "/admin.service.v1.RouterService/ListRoute"

type RouterServiceHTTPServer interface {
	// ListPermissionCode 查询权限码列表
	ListPermissionCode(context.Context, *emptypb.Empty) (*ListPermissionCodeResponse, error)
	// ListRoute 查询路由列表
	ListRoute(context.Context, *emptypb.Empty) (*ListRouteResponse, error)
}

func RegisterRouterServiceHTTPServer(s *http.Server, srv RouterServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/routes", _RouterService_ListRoute0_HTTP_Handler(srv))
	r.GET("/admin/v1/perm-codes", _RouterService_ListPermissionCode0_HTTP_Handler(srv))
}

func _RouterService_ListRoute0_HTTP_Handler(srv RouterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRouterServiceListRoute)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRoute(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRouteResponse)
		return ctx.Result(200, reply)
	}
}

func _RouterService_ListPermissionCode0_HTTP_Handler(srv RouterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRouterServiceListPermissionCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPermissionCode(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPermissionCodeResponse)
		return ctx.Result(200, reply)
	}
}

type RouterServiceHTTPClient interface {
	ListPermissionCode(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListPermissionCodeResponse, err error)
	ListRoute(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListRouteResponse, err error)
}

type RouterServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewRouterServiceHTTPClient(client *http.Client) RouterServiceHTTPClient {
	return &RouterServiceHTTPClientImpl{client}
}

func (c *RouterServiceHTTPClientImpl) ListPermissionCode(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListPermissionCodeResponse, error) {
	var out ListPermissionCodeResponse
	pattern := "/admin/v1/perm-codes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRouterServiceListPermissionCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RouterServiceHTTPClientImpl) ListRoute(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListRouteResponse, error) {
	var out ListRouteResponse
	pattern := "/admin/v1/routes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRouterServiceListRoute))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
