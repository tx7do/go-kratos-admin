package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
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

func (s *PositionService) ListPosition(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListPositionResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *PositionService) GetPosition(ctx context.Context, req *userV1.GetPositionRequest) (*userV1.Position, error) {
	return s.uc.Get(ctx, req)
}

func (s *PositionService) CreatePosition(ctx context.Context, req *userV1.CreatePositionRequest) (*emptypb.Empty, error) {
	err := s.uc.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PositionService) UpdatePosition(ctx context.Context, req *userV1.UpdatePositionRequest) (*emptypb.Empty, error) {
	err := s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PositionService) DeletePosition(ctx context.Context, req *userV1.DeletePositionRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
