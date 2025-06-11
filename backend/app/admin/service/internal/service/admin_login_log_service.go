package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminLoginLogService struct {
	adminV1.AdminLoginLogServiceHTTPServer

	log *log.Helper

	repo *data.AdminLoginLogRepo
}

func NewAdminLoginLogService(logger log.Logger, repo *data.AdminLoginLogRepo) *AdminLoginLogService {
	l := log.NewHelper(log.With(logger, "module", "admin-login-log/service/admin-service"))
	return &AdminLoginLogService{
		log:  l,
		repo: repo,
	}
}

func (s *AdminLoginLogService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminLoginLogResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *AdminLoginLogService) Get(ctx context.Context, req *adminV1.GetAdminLoginLogRequest) (*adminV1.AdminLoginLog, error) {
	return s.repo.Get(ctx, req)
}

func (s *AdminLoginLogService) Create(ctx context.Context, req *adminV1.CreateAdminLoginLogRequest) (*emptypb.Empty, error) {
	if req == nil || req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	if err := s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
