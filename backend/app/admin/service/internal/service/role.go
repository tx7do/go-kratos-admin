package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type RoleService struct {
	adminV1.RoleServiceHTTPServer

	log *log.Helper

	repo *data.RoleRepo
}

func NewRoleService(logger log.Logger, repo *data.RoleRepo) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service/admin-service"))
	return &RoleService{
		log:  l,
		repo: repo,
	}
}

func (s *RoleService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListRoleResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *RoleService) Get(ctx context.Context, req *userV1.GetRoleRequest) (*userV1.Role, error) {
	return s.repo.Get(ctx, req.GetId())
}

func (s *RoleService) Create(ctx context.Context, req *userV1.CreateRoleRequest) (*emptypb.Empty, error) {
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

func (s *RoleService) Update(ctx context.Context, req *userV1.UpdateRoleRequest) (*emptypb.Empty, error) {
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

func (s *RoleService) Delete(ctx context.Context, req *userV1.DeleteRoleRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
