package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type PositionService struct {
	adminV1.PositionServiceHTTPServer

	log *log.Helper

	uc *data.PositionRepo
}

func NewPositionService(uc *data.PositionRepo, logger log.Logger) *PositionService {
	l := log.NewHelper(log.With(logger, "module", "position/service/admin-service"))
	return &PositionService{
		log: l,
		uc:  uc,
	}
}

func (s *PositionService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListPositionResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *PositionService) Get(ctx context.Context, req *userV1.GetPositionRequest) (*userV1.Position, error) {
	return s.uc.Get(ctx, req)
}

func (s *PositionService) Create(ctx context.Context, req *userV1.CreatePositionRequest) (*emptypb.Empty, error) {
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

func (s *PositionService) Update(ctx context.Context, req *userV1.UpdatePositionRequest) (*emptypb.Empty, error) {
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

func (s *PositionService) Delete(ctx context.Context, req *userV1.DeletePositionRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
