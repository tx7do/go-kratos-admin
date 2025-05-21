package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminOperationLogService struct {
	adminV1.AdminOperationLogServiceHTTPServer

	log *log.Helper

	uc *data.AdminOperationLogRepo
}

func NewAdminOperationLogService(uc *data.AdminOperationLogRepo, logger log.Logger) *AdminOperationLogService {
	l := log.NewHelper(log.With(logger, "module", "admin-operation-log/service/admin-service"))
	return &AdminOperationLogService{
		log: l,
		uc:  uc,
	}
}

func (s *AdminOperationLogService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminOperationLogResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *AdminOperationLogService) Get(ctx context.Context, req *adminV1.GetAdminOperationLogRequest) (*adminV1.AdminOperationLog, error) {
	return s.uc.Get(ctx, req)
}

func (s *AdminOperationLogService) Create(ctx context.Context, req *adminV1.CreateAdminOperationLogRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	if err := s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
