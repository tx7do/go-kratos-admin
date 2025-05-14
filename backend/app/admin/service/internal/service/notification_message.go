package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type NotificationMessageService struct {
	adminV1.NotificationMessageServiceHTTPServer

	log *log.Helper

	uc *data.NotificationMessageRepo
}

func NewNotificationMessageService(uc *data.NotificationMessageRepo, logger log.Logger) *NotificationMessageService {
	l := log.NewHelper(log.With(logger, "module", "notification-message/service/admin-service"))
	return &NotificationMessageService{
		log: l,
		uc:  uc,
	}
}

func (s *NotificationMessageService) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *NotificationMessageService) Get(ctx context.Context, req *internalMessageV1.GetNotificationMessageRequest) (*internalMessageV1.NotificationMessage, error) {
	return s.uc.Get(ctx, req)
}

func (s *NotificationMessageService) Create(ctx context.Context, req *internalMessageV1.CreateNotificationMessageRequest) (*emptypb.Empty, error) {
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

func (s *NotificationMessageService) Update(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageRequest) (*emptypb.Empty, error) {
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

func (s *NotificationMessageService) Delete(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
