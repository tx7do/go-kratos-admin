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

type RoleService struct {
	adminV1.RoleServiceHTTPServer

	log *log.Helper

	uc *data.RoleRepo
}

func NewRoleService(uc *data.RoleRepo, logger log.Logger) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service/admin-service"))
	return &RoleService{
		log: l,
		uc:  uc,
	}
}

func (s *RoleService) ListRole(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListRoleResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *RoleService) GetRole(ctx context.Context, req *userV1.GetRoleRequest) (*userV1.Role, error) {
	return s.uc.GetRole(ctx, req.GetId())
}

func (s *RoleService) CreateRole(ctx context.Context, req *userV1.CreateRoleRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	err = s.uc.CreateRole(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *RoleService) UpdateRole(ctx context.Context, req *userV1.UpdateRoleRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	err = s.uc.UpdateRole(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *RoleService) DeleteRole(ctx context.Context, req *userV1.DeleteRoleRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	_, err = s.uc.DeleteRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
