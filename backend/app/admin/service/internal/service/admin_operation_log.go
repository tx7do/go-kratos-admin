package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"
	"kratos-admin/app/admin/service/internal/middleware/auth"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
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

func (s *AdminOperationLogService) ListAdminOperationLog(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListAdminOperationLogResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *AdminOperationLogService) GetAdminOperationLog(ctx context.Context, req *systemV1.GetAdminOperationLogRequest) (*systemV1.AdminOperationLog, error) {
	return s.uc.Get(ctx, req)
}

func (s *AdminOperationLogService) CreateAdminOperationLog(ctx context.Context, req *systemV1.CreateAdminOperationLogRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	if err = s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AdminOperationLogService) UpdateAdminOperationLog(ctx context.Context, req *systemV1.UpdateAdminOperationLogRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	if err = s.uc.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AdminOperationLogService) DeleteAdminOperationLog(ctx context.Context, req *systemV1.DeleteAdminOperationLogRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	if _, err = s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
