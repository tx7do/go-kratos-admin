package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/trans"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	userV1 "go-wind-admin/api/gen/go/user/service/v1"

	"go-wind-admin/pkg/middleware/auth"
	"go-wind-admin/pkg/utils/slice"
)

type RouterService struct {
	adminV1.RouterServiceHTTPServer

	log *log.Helper

	menuRepo *data.MenuRepo
	roleRepo *data.RoleRepo
	userRepo *data.UserRepo
}

func NewRouterService(
	logger log.Logger,
	menuRepo *data.MenuRepo,
	roleRepo *data.RoleRepo,
	userRepo *data.UserRepo,
) *RouterService {
	l := log.NewHelper(log.With(logger, "module", "router/service/admin-service"))
	return &RouterService{
		log:      l,
		menuRepo: menuRepo,
		roleRepo: roleRepo,
		userRepo: userRepo,
	}
}

func (s *RouterService) menuListToQueryString(menus []uint32, onlyButton bool) string {
	var ids []string
	for _, menu := range menus {
		ids = append(ids, fmt.Sprintf("\"%d\"", menu))
	}
	idsStr := fmt.Sprintf("[%s]", strings.Join(ids, ", "))
	query := map[string]string{"id__in": idsStr}

	if onlyButton {
		query["type"] = adminV1.Menu_BUTTON.String()
	} else {
		query["type__not"] = adminV1.Menu_BUTTON.String()
	}

	query["status"] = "ON"

	queryStr, err := json.Marshal(query)
	if err != nil {
		return ""
	}

	return string(queryStr)
}

// queryOneRoleMenus 使用RoleID查询菜单，即单个角色的菜单
func (s *RouterService) queryOneRoleMenus(ctx context.Context, roleId uint32) ([]uint32, error) {
	role, err := s.roleRepo.Get(ctx, &userV1.GetRoleRequest{QueryBy: &userV1.GetRoleRequest_Id{Id: roleId}})
	if err != nil {
		s.log.Errorf("query role by role id failed[%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("query role failed")
	}
	return role.GetMenus(), nil
}

// queryMultipleRolesMenusByRoleCodes 使用RoleCodes查询菜单，即多个角色的菜单
func (s *RouterService) queryMultipleRolesMenusByRoleCodes(ctx context.Context, roleCodes []string) ([]uint32, error) {
	roles, err := s.roleRepo.ListRolesByRoleCodes(ctx, roleCodes)
	if err != nil {
		return nil, adminV1.ErrorInternalServerError("query roles failed")
	}

	var menus []uint32
	for _, role := range roles {
		if role == nil {
			continue
		}
		menus = slice.MergeAndDeduplicate(menus, role.GetMenus())
	}

	return menus, nil
}

// queryMultipleRolesMenusByRoleIds 使用RoleIDs查询菜单，即多个角色的菜单
func (s *RouterService) queryMultipleRolesMenusByRoleIds(ctx context.Context, roleIds []uint32) ([]uint32, error) {
	roles, err := s.roleRepo.ListRolesByRoleIds(ctx, roleIds)
	if err != nil {
		return nil, adminV1.ErrorInternalServerError("query roles failed")
	}

	var menus []uint32
	for _, role := range roles {
		if role == nil {
			continue
		}
		menus = slice.MergeAndDeduplicate(menus, role.GetMenus())
	}

	return menus, nil
}

func (s *RouterService) ListPermissionCode(ctx context.Context, _ *emptypb.Empty) (*adminV1.ListPermissionCodeResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		s.log.Errorf("query user failed[%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("query user failed")
	}

	// 单角色的菜单
	//roleMenus, err := s.queryOneRoleMenus(ctx, user.GetRoleId())
	//if err != nil {
	//	return nil, err
	//}

	// 多角色的菜单
	roleMenus, err := s.queryMultipleRolesMenusByRoleIds(ctx, user.GetRoleIds())
	if err != nil {
		return nil, err
	}

	menus, err := s.menuRepo.List(ctx, &pagination.PagingRequest{
		NoPaging: trans.Ptr(true),
		Query:    trans.Ptr(s.menuListToQueryString(roleMenus, true)),
		FieldMask: &fieldmaskpb.FieldMask{
			Paths: []string{"id", "meta"},
		},
	}, false)
	if err != nil {
		s.log.Errorf("list permission code failed [%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("list permission code failed")
	}

	var codes []string
	for menu := range menus.Items {
		if menus.Items[menu].GetMeta() == nil {
			continue
		}
		if len(menus.Items[menu].GetMeta().GetAuthority()) == 0 {
			continue
		}

		codes = append(codes, menus.Items[menu].GetMeta().GetAuthority()...)
	}

	return &adminV1.ListPermissionCodeResponse{
		Codes: codes,
	}, nil
}

func (s *RouterService) fillRouteItem(menus []*adminV1.Menu) []*adminV1.RouteItem {
	if len(menus) == 0 {
		return nil
	}

	var routers []*adminV1.RouteItem

	for _, v := range menus {
		if v.GetStatus() != adminV1.Menu_ON {
			continue
		}
		if v.GetType() == adminV1.Menu_BUTTON {
			continue
		}

		item := &adminV1.RouteItem{
			Path:      v.Path,
			Component: v.Component,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Alias:     v.Alias,
			Meta:      v.Meta,
		}

		if len(v.Children) > 0 {
			item.Children = s.fillRouteItem(v.Children)
		}

		routers = append(routers, item)
	}

	return routers
}

func (s *RouterService) ListRoute(ctx context.Context, _ *emptypb.Empty) (*adminV1.ListRouteResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		s.log.Errorf("query user failed[%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("query user failed")
	}

	// 单角色的菜单
	//roleMenus, err := s.queryOneRoleMenus(ctx, user.GetRoleId())
	//if err != nil {
	//	return nil, err
	//}

	// 多角色的菜单
	roleMenus, err := s.queryMultipleRolesMenusByRoleCodes(ctx, user.GetRoles())
	if err != nil {
		return nil, err
	}

	menuList, err := s.menuRepo.List(ctx, &pagination.PagingRequest{
		NoPaging: trans.Ptr(true),
		Query:    trans.Ptr(s.menuListToQueryString(roleMenus, false)),
	}, true)
	if err != nil {
		s.log.Errorf("list route failed [%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("list route failed")
	}

	resp := &adminV1.ListRouteResponse{Items: s.fillRouteItem(menuList.Items)}

	return resp, nil
}
