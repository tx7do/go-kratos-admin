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
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type TenantService struct {
	adminV1.TenantServiceHTTPServer

	log *log.Helper

	uc *data.TenantRepo
}

func NewTenantService(uc *data.TenantRepo, logger log.Logger) *TenantService {
	l := log.NewHelper(log.With(logger, "module", "tenant/service/admin-service"))
	return &TenantService{
		log: l,
		uc:  uc,
	}
}

func (s *TenantService) ListTenant(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListTenantResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *TenantService) GetTenant(ctx context.Context, req *userV1.GetTenantRequest) (*userV1.Tenant, error) {
	return s.uc.Get(ctx, req)
}

func (s *TenantService) CreateTenant(ctx context.Context, req *userV1.CreateTenantRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	err = s.uc.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TenantService) UpdateTenant(ctx context.Context, req *userV1.UpdateTenantRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	err = s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TenantService) DeleteTenant(ctx context.Context, req *userV1.DeleteTenantRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	_, err = s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
