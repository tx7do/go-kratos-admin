package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"
)

type PrivateMessageService struct {
	adminV1.PrivateMessageServiceHTTPServer

	log *log.Helper

	repo *data.PrivateMessageRepo
}

func NewPrivateMessageService(logger log.Logger, repo *data.PrivateMessageRepo) *PrivateMessageService {
	l := log.NewHelper(log.With(logger, "module", "private-message/service/admin-service"))
	return &PrivateMessageService{
		log:  l,
		repo: repo,
	}
}

func (s *PrivateMessageService) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListPrivateMessageResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *PrivateMessageService) Get(ctx context.Context, req *internalMessageV1.GetPrivateMessageRequest) (*internalMessageV1.PrivateMessage, error) {
	return s.repo.Get(ctx, req)
}

func (s *PrivateMessageService) Create(ctx context.Context, req *internalMessageV1.CreatePrivateMessageRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	var err error

	if err = s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PrivateMessageService) Update(ctx context.Context, req *internalMessageV1.UpdatePrivateMessageRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	var err error

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PrivateMessageService) Delete(ctx context.Context, req *internalMessageV1.DeletePrivateMessageRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
