package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type MenuService struct {
	adminV1.MenuServiceHTTPServer

	log *log.Helper

	uc *data.MenuRepo
}

func NewMenuService(uc *data.MenuRepo, logger log.Logger) *MenuService {
	l := log.NewHelper(log.With(logger, "module", "menu/service/admin-service"))
	return &MenuService{
		log: l,
		uc:  uc,
	}
}

func (s *MenuService) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListMenuResponse, error) {
	ret, err := s.uc.List(ctx, req, false)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *MenuService) Get(ctx context.Context, req *systemV1.GetMenuRequest) (*systemV1.Menu, error) {
	ret, err := s.uc.Get(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *MenuService) Create(ctx context.Context, req *systemV1.CreateMenuRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.uc.Create(ctx, req, operator); err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MenuService) Update(ctx context.Context, req *systemV1.UpdateMenuRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.uc.Update(ctx, req, operator); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MenuService) Delete(ctx context.Context, req *systemV1.DeleteMenuRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
