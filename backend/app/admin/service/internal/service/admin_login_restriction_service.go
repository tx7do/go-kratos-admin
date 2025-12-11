package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"

	"go-wind-admin/pkg/middleware/auth"
)

type AdminLoginRestrictionService struct {
	adminV1.AdminLoginRestrictionServiceHTTPServer

	log *log.Helper

	repo *data.AdminLoginRestrictionRepo
}

func NewAdminLoginRestrictionService(logger log.Logger, repo *data.AdminLoginRestrictionRepo) *AdminLoginRestrictionService {
	l := log.NewHelper(log.With(logger, "module", "admin-login-restriction/service/admin-service"))
	return &AdminLoginRestrictionService{
		log:  l,
		repo: repo,
	}
}

func (s *AdminLoginRestrictionService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminLoginRestrictionResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *AdminLoginRestrictionService) Get(ctx context.Context, req *adminV1.GetAdminLoginRestrictionRequest) (*adminV1.AdminLoginRestriction, error) {
	return s.repo.Get(ctx, req)
}

func (s *AdminLoginRestrictionService) Create(ctx context.Context, req *adminV1.CreateAdminLoginRestrictionRequest) (*emptypb.Empty, error) {
	if req == nil || req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid request")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AdminLoginRestrictionService) Update(ctx context.Context, req *adminV1.UpdateAdminLoginRestrictionRequest) (*emptypb.Empty, error) {
	if req == nil || req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid request")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AdminLoginRestrictionService) Delete(ctx context.Context, req *adminV1.DeleteAdminLoginRestrictionRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid request")
	}

	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
