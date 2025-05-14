package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type AdminLoginLogService struct {
	adminV1.AdminLoginLogServiceHTTPServer

	log *log.Helper

	uc *data.AdminLoginLogRepo
}

func NewAdminLoginLogService(uc *data.AdminLoginLogRepo, logger log.Logger) *AdminLoginLogService {
	l := log.NewHelper(log.With(logger, "module", "admin-login-log/service/admin-service"))
	return &AdminLoginLogService{
		log: l,
		uc:  uc,
	}
}

func (s *AdminLoginLogService) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListAdminLoginLogResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *AdminLoginLogService) Get(ctx context.Context, req *systemV1.GetAdminLoginLogRequest) (*systemV1.AdminLoginLog, error) {
	return s.uc.Get(ctx, req)
}

func (s *AdminLoginLogService) Create(ctx context.Context, req *systemV1.CreateAdminLoginLogRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	if err := s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AdminLoginLogService) Update(ctx context.Context, req *systemV1.UpdateAdminLoginLogRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	if err := s.uc.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AdminLoginLogService) Delete(ctx context.Context, req *systemV1.DeleteAdminLoginLogRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
