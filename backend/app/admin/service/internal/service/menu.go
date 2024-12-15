package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
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

func (s *MenuService) ListMenu(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListMenuResponse, error) {
	ret, err := s.uc.List(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *MenuService) GetMenu(ctx context.Context, req *systemV1.GetMenuRequest) (*systemV1.Menu, error) {
	ret, err := s.uc.Get(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *MenuService) CreateMenu(ctx context.Context, req *systemV1.CreateMenuRequest) (*emptypb.Empty, error) {
	err := s.uc.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MenuService) UpdateMenu(ctx context.Context, req *systemV1.UpdateMenuRequest) (*emptypb.Empty, error) {
	err := s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MenuService) DeleteMenu(ctx context.Context, req *systemV1.DeleteMenuRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
