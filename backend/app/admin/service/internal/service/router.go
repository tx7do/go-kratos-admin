package service

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type RouterService struct {
	adminV1.RouterServiceHTTPServer

	log *log.Helper
	uc  *data.MenuRepo
}

func NewRouterService(logger log.Logger, uc *data.MenuRepo) *RouterService {
	l := log.NewHelper(log.With(logger, "module", "router/service/admin-service"))
	return &RouterService{
		log: l,
		uc:  uc,
	}
}

func (s *RouterService) ListPermissionCode(_ context.Context, _ *emptypb.Empty) (*adminV1.ListPermissionCodeResponse, error) {
	return &adminV1.ListPermissionCodeResponse{}, nil
}

func (s *RouterService) fillRouteItem(menus []*systemV1.Menu) []*systemV1.RouteItem {
	if len(menus) == 0 {
		return nil
	}

	var routers []*systemV1.RouteItem

	for _, v := range menus {
		if v.GetStatus() != systemV1.MenuStatus_ON {
			continue
		}
		if v.GetType() == systemV1.MenuType_BUTTON {
			continue
		}

		item := &systemV1.RouteItem{
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
	menuList, err := s.uc.TravelList(ctx, &pagination.PagingRequest{
		NoPaging: trans.Ptr(true),
	})
	if err != nil {
		s.log.Errorf("查询列表发生错误[%s]", err.Error())
		return nil, errors.New("读取列表发生错误")
	}

	resp := &adminV1.ListRouteResponse{Items: s.fillRouteItem(menuList.Items)}

	return resp, nil
}
