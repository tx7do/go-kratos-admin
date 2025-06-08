package service

import (
	"context"
	"encoding/json"
	"fmt"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
	"kratos-admin/pkg/utils/slice"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"

	"kratos-admin/pkg/middleware/auth"
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
		query["type"] = adminV1.MenuType_BUTTON.String()
	} else {
		query["type__not"] = adminV1.MenuType_BUTTON.String()
	}

	query["status"] = "ON"

	queryStr, err := json.Marshal(query)
	if err != nil {
		return ""
	}

	return string(queryStr)
}

// queryRoleMenus 使用RoleID查询菜单，即单个角色的菜单
func (s *RouterService) queryRoleMenus(ctx context.Context, roleId uint32) ([]uint32, error) {
	role, err := s.roleRepo.Get(ctx, roleId)
	if err != nil {
		s.log.Errorf("query role failed[%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("query role failed")
	}
	return role.GetMenus(), nil
}

// queryRolesMenus 使用Roles查询菜单，即多个角色的菜单
func (s *RouterService) queryRolesMenus(ctx context.Context, roles []string) ([]uint32, error) {
	var err error
	var menus []uint32
	var role *userV1.Role
	for _, code := range roles {
		if role, err = s.roleRepo.GetRoleByCode(ctx, code); err != nil {
			s.log.Errorf("query role failed [%s] [%s]", code, err.Error())
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

	user, err := s.userRepo.Get(ctx, operator.UserId)
	if err != nil {
		s.log.Errorf("query user failed[%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("query user failed")
	}

	// 单角色的菜单
	//roleMenus, err := s.queryRoleMenus(ctx, user.GetRoleId())
	//if err != nil {
	//	return nil, err
	//}

	// 多角色的菜单
	roleMenus, err := s.queryRolesMenus(ctx, user.GetRoles())
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
		if v.GetStatus() != adminV1.MenuStatus_ON {
			continue
		}
		if v.GetType() == adminV1.MenuType_BUTTON {
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

	user, err := s.userRepo.Get(ctx, operator.UserId)
	if err != nil {
		s.log.Errorf("query user failed[%s]", err.Error())
		return nil, adminV1.ErrorInternalServerError("query user failed")
	}

	// 单角色的菜单
	//roleMenus, err := s.queryRoleMenus(ctx, user.GetRoleId())
	//if err != nil {
	//	return nil, err
	//}

	// 多角色的菜单
	roleMenus, err := s.queryRolesMenus(ctx, user.GetRoles())
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
