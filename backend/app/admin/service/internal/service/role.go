package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type RoleService struct {
	userV1.UnimplementedRoleServiceServer

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
	return s.uc.Get(ctx, req)
}

func (s *RoleService) CreateRole(ctx context.Context, req *userV1.CreateRoleRequest) (*emptypb.Empty, error) {
	err := s.uc.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *RoleService) UpdateRole(ctx context.Context, req *userV1.UpdateRoleRequest) (*emptypb.Empty, error) {
	err := s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *RoleService) DeleteRole(ctx context.Context, req *userV1.DeleteRoleRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
