// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             (unknown)
// source: admin/service/v1/i_admin_login_log.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAdminLoginLogServiceGet = "/admin.service.v1.AdminLoginLogService/Get"
const OperationAdminLoginLogServiceList = "/admin.service.v1.AdminLoginLogService/List"

type AdminLoginLogServiceHTTPServer interface {
	// Get 查询后台登录日志详情
	Get(context.Context, *GetAdminLoginLogRequest) (*AdminLoginLog, error)
	// List 查询后台登录日志列表
	List(context.Context, *v1.PagingRequest) (*ListAdminLoginLogResponse, error)
}

func RegisterAdminLoginLogServiceHTTPServer(s *http.Server, srv AdminLoginLogServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/admin_login_logs", _AdminLoginLogService_List0_HTTP_Handler(srv))
	r.GET("/admin/v1/admin_login_logs/{id}", _AdminLoginLogService_Get0_HTTP_Handler(srv))
}

func _AdminLoginLogService_List0_HTTP_Handler(srv AdminLoginLogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminLoginLogServiceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAdminLoginLogResponse)
		return ctx.Result(200, reply)
	}
}

func _AdminLoginLogService_Get0_HTTP_Handler(srv AdminLoginLogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAdminLoginLogRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminLoginLogServiceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*GetAdminLoginLogRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminLoginLog)
		return ctx.Result(200, reply)
	}
}

type AdminLoginLogServiceHTTPClient interface {
	Get(ctx context.Context, req *GetAdminLoginLogRequest, opts ...http.CallOption) (rsp *AdminLoginLog, err error)
	List(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *ListAdminLoginLogResponse, err error)
}

type AdminLoginLogServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAdminLoginLogServiceHTTPClient(client *http.Client) AdminLoginLogServiceHTTPClient {
	return &AdminLoginLogServiceHTTPClientImpl{client}
}

func (c *AdminLoginLogServiceHTTPClientImpl) Get(ctx context.Context, in *GetAdminLoginLogRequest, opts ...http.CallOption) (*AdminLoginLog, error) {
	var out AdminLoginLog
	pattern := "/admin/v1/admin_login_logs/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminLoginLogServiceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AdminLoginLogServiceHTTPClientImpl) List(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*ListAdminLoginLogResponse, error) {
	var out ListAdminLoginLogResponse
	pattern := "/admin/v1/admin_login_logs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminLoginLogServiceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
