package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type ApiResourceService struct {
	adminV1.ApiResourceServiceHTTPServer

	RestServer *http.Server

	log *log.Helper

	uc *data.ApiResourceRepo
}

func NewApiResourceService(uc *data.ApiResourceRepo, logger log.Logger) *ApiResourceService {
	l := log.NewHelper(log.With(logger, "module", "api-resource/service/admin-service"))
	return &ApiResourceService{
		log: l,
		uc:  uc,
	}
}

func (s *ApiResourceService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListApiResourceResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *ApiResourceService) Get(ctx context.Context, req *adminV1.GetApiResourceRequest) (*adminV1.ApiResource, error) {
	return s.uc.Get(ctx, req)
}

func (s *ApiResourceService) Create(ctx context.Context, req *adminV1.CreateApiResourceRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) Update(ctx context.Context, req *adminV1.UpdateApiResourceRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.uc.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) Delete(ctx context.Context, req *adminV1.DeleteApiResourceRequest) (*emptypb.Empty, error) {
	if err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) SyncApiResources(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	var routes []http.RouteInfo
	if s.RestServer != nil {
		_ = s.RestServer.WalkRoute(func(info http.RouteInfo) error {
			//log.Infof("Path[%s] Method[%s]", info.Path, info.Method)
			routes = append(routes, info)
			return nil
		})
	}

	if len(routes) == 0 {
		return &emptypb.Empty{}, nil
	}

	return &emptypb.Empty{}, nil
}
