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

	repo *data.RoleRepo
}

func NewRoleService(logger log.Logger, repo *data.RoleRepo) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service/admin-service"))
	svc := &RoleService{
		log:  l,
		repo: repo,
	}

	svc.init()

	return svc
}

func (s *RoleService) init() {
	ctx := context.Background()
	if count, _ := s.repo.Count(ctx, []func(s *sql.Selector){}); count == 0 {
		_ = s.createDefaultRoles(ctx)
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

// createDefaultRoles 创建默认角色(包括超级管理员)
func (s *RoleService) createDefaultRoles(ctx context.Context) error {
	var err error

	defaultRoles := []*userV1.CreateRoleRequest{
		{
			Data: &userV1.Role{
				Id:     trans.Ptr(uint32(1)),
				Name:   trans.Ptr("超级管理员"),
				Code:   trans.Ptr("super"),
				Status: trans.Ptr("ON"),
				Remark: trans.Ptr("超级管理员拥有对系统的最高权限"),
				Menus:  []uint32{1, 2, 10, 11, 12, 13, 14, 20, 21, 22, 15, 16, 17, 18, 23, 24, 25, 26, 27, 30, 31, 32},
				Apis:   []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102},
			},
		},
		{
			Data: &userV1.Role{
				Id:     trans.Ptr(uint32(2)),
				Name:   trans.Ptr("管理员"),
				Code:   trans.Ptr("admin"),
				Status: trans.Ptr("ON"),
				Remark: trans.Ptr("系统管理员拥有对整个系统的管理权限"),
				Menus:  []uint32{1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14},
				Apis:   []uint32{},
			},
		},
		{
			Data: &userV1.Role{
				Id:     trans.Ptr(uint32(3)),
				Name:   trans.Ptr("普通用户"),
				Code:   trans.Ptr("user"),
				Status: trans.Ptr("ON"),
				Remark: trans.Ptr("普通用户没有管理权限，只有设备和APP的使用权限"),
				Menus:  []uint32{},
				Apis:   []uint32{},
			},
		},
		{
			Data: &userV1.Role{
				Id:     trans.Ptr(uint32(4)),
				Name:   trans.Ptr("游客"),
				Code:   trans.Ptr("guest"),
				Status: trans.Ptr("ON"),
				Remark: trans.Ptr("游客只有非常有限的数据读取权限"),
				Menus:  []uint32{},
				Apis:   []uint32{},
			},
		},
	}

	for _, roleReq := range defaultRoles {
		err = s.repo.Create(ctx, roleReq)
		if err != nil {
			return err
		}
	}

	return nil
}
