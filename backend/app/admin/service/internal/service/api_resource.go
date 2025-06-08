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

	repo *data.ApiResourceRepo
}

func NewApiResourceService(repo *data.ApiResourceRepo, logger log.Logger) *ApiResourceService {
	l := log.NewHelper(log.With(logger, "module", "api-resource/service/admin-service"))
	return &ApiResourceService{
		log:  l,
		repo: repo,
	}
}

func (s *ApiResourceService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListApiResourceResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *ApiResourceService) Get(ctx context.Context, req *adminV1.GetApiResourceRequest) (*adminV1.ApiResource, error) {
	return s.repo.Get(ctx, req)
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

	if err = s.repo.Create(ctx, req); err != nil {
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

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) Delete(ctx context.Context, req *adminV1.DeleteApiResourceRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) SyncApiResources(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	if s.RestServer != nil {
		var count uint32 = 0
		_ = s.RestServer.WalkRoute(func(info http.RouteInfo) error {
			//log.Infof("Path[%s] Method[%s]", info.Path, info.Method)
			count++
			_ = s.repo.Update(ctx, &adminV1.UpdateApiResourceRequest{
				AllowMissing: trans.Ptr(true),
				Data: &adminV1.ApiResource{
					Id:     trans.Ptr(count),
					Path:   trans.Ptr(info.Path),
					Method: trans.Ptr(info.Method),
				},
			})

			return nil
		})
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) GetWalkRouteData(_ context.Context, _ *emptypb.Empty) (*adminV1.ListApiResourceResponse, error) {
	resp := &adminV1.ListApiResourceResponse{
		Items: []*adminV1.ApiResource{},
	}
	var count uint32 = 0
	if s.RestServer != nil {
		_ = s.RestServer.WalkRoute(func(info http.RouteInfo) error {
			//log.Infof("Path[%s] Method[%s]", info.Path, info.Method)
			count++
			resp.Items = append(resp.Items, &adminV1.ApiResource{
				Id:     trans.Ptr(count),
				Path:   trans.Ptr(info.Path),
				Method: trans.Ptr(info.Method),
			})
			return nil
		})
	}
	resp.Total = count

	return resp, nil
}
