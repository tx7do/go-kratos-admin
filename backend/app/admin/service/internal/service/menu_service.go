package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type MenuService struct {
	adminV1.MenuServiceHTTPServer

	log *log.Helper

	repo *data.MenuRepo
}

func NewMenuService(logger log.Logger, repo *data.MenuRepo) *MenuService {
	l := log.NewHelper(log.With(logger, "module", "menu/service/admin-service"))
	return &MenuService{
		log:  l,
		repo: repo,
	}
}

func (s *MenuService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListMenuResponse, error) {
	ret, err := s.repo.List(ctx, req, false)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *MenuService) Get(ctx context.Context, req *adminV1.GetMenuRequest) (*adminV1.Menu, error) {
	ret, err := s.repo.Get(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *MenuService) Create(ctx context.Context, req *adminV1.CreateMenuRequest) (*emptypb.Empty, error) {
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

func (s *MenuService) Update(ctx context.Context, req *adminV1.UpdateMenuRequest) (*emptypb.Empty, error) {
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

func (s *MenuService) Delete(ctx context.Context, req *adminV1.DeleteMenuRequest) (*emptypb.Empty, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.OperatorId = trans.Ptr(operator.UserId)

	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
