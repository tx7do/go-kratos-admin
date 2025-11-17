package service

import (
	"context"

	"entgo.io/ent/dialect/sql"
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

	authorizer *data.Authorizer

	roleRepo *data.RoleRepo

	roleApiRepo      *data.RoleApiRepo
	roleMenuRepo     *data.RoleMenuRepo
	roleOrgRepo      *data.RoleOrgRepo
	roleDeptRepo     *data.RoleDeptRepo
	rolePositionRepo *data.RolePositionRepo
}

func NewRoleService(
	logger log.Logger,
	authorizer *data.Authorizer,
	roleRepo *data.RoleRepo,
	roleApiRepo *data.RoleApiRepo,
	roleMenuRepo *data.RoleMenuRepo,
	roleOrgRepo *data.RoleOrgRepo,
	roleDeptRepo *data.RoleDeptRepo,
	rolePositionRepo *data.RolePositionRepo,
) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service/admin-service"))
	svc := &RoleService{
		log:              l,
		authorizer:       authorizer,
		roleRepo:         roleRepo,
		roleApiRepo:      roleApiRepo,
		roleMenuRepo:     roleMenuRepo,
		roleOrgRepo:      roleOrgRepo,
		roleDeptRepo:     roleDeptRepo,
		rolePositionRepo: rolePositionRepo,
	}

	svc.init()

	return svc
}

func (s *RoleService) init() {
	ctx := context.Background()
	if count, _ := s.roleRepo.Count(ctx, []func(s *sql.Selector){}); count == 0 {
		_ = s.createDefaultRoles(ctx)
	}
}

func (s *RoleService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListRoleResponse, error) {
	return s.roleRepo.List(ctx, req)
}

func (s *RoleService) Get(ctx context.Context, req *userV1.GetRoleRequest) (*userV1.Role, error) {
	return s.roleRepo.Get(ctx, req)
}

func (s *RoleService) Create(ctx context.Context, req *userV1.CreateRoleRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.roleRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	if err = s.authorizer.ResetPolicies(ctx); err != nil {
		s.log.Errorf("reset policies error: %v", err)
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
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.roleRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	if err = s.authorizer.ResetPolicies(ctx); err != nil {
		s.log.Errorf("reset policies error: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *RoleService) Delete(ctx context.Context, req *userV1.DeleteRoleRequest) (*emptypb.Empty, error) {
	var err error

	if err = s.roleRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	if err = s.authorizer.ResetPolicies(ctx); err != nil {
		s.log.Errorf("reset policies error: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *RoleService) GetRoleCodesByRoleIds(ctx context.Context, req *userV1.GetRoleCodesByRoleIdsRequest) (*userV1.GetRoleCodesByRoleIdsResponse, error) {
	ids, err := s.roleRepo.GetRoleCodesByRoleIds(ctx, req.GetRoleIds())
	if err != nil {
		return nil, err
	}

	return &userV1.GetRoleCodesByRoleIdsResponse{
		RoleCodes: ids,
	}, nil
}

func (s *RoleService) GetRolesByRoleCodes(ctx context.Context, req *userV1.GetRolesByRoleCodesRequest) (*userV1.ListRoleResponse, error) {
	roles, err := s.roleRepo.GetRolesByRoleCodes(ctx, req.GetRoleCodes())
	if err != nil {
		return nil, err
	}

	return &userV1.ListRoleResponse{
		Items: roles,
		Total: uint32(len(roles)),
	}, nil
}

func (s *RoleService) GetRolesByRoleIds(ctx context.Context, req *userV1.GetRolesByRoleIdsRequest) (*userV1.ListRoleResponse, error) {
	roles, err := s.roleRepo.GetRolesByRoleIds(ctx, req.GetRoleIds())
	if err != nil {
		return nil, err
	}

	return &userV1.ListRoleResponse{
		Items: roles,
		Total: uint32(len(roles)),
	}, nil
}

// createDefaultRoles 创建默认角色(包括超级管理员)
func (s *RoleService) createDefaultRoles(ctx context.Context) error {
	var err error

	defaultRoles := []*userV1.CreateRoleRequest{
		{
			Data: &userV1.Role{
				Id:     trans.Ptr(uint32(1)),
				Name:   trans.Ptr("超级管理员"),
				Code:   trans.Ptr("super"),
				Status: trans.Ptr(userV1.Role_ON),
				Remark: trans.Ptr("超级管理员拥有对系统的最高权限"),
				Menus:  []uint32{1, 2, 10, 11, 20, 21, 22, 23, 24, 25, 30, 31, 32, 40, 41, 42, 50, 51, 52, 60, 61, 62, 63, 64, 65},
				Apis:   []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107},
			},
		},
	}

	for _, roleReq := range defaultRoles {
		err = s.roleRepo.Create(ctx, roleReq)
		if err != nil {
			return err
		}
	}

	return nil
}
