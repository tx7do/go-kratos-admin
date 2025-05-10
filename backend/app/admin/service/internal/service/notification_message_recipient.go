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

type NotificationMessageRecipientService struct {
	adminV1.NotificationMessageRecipientServiceHTTPServer

	log *log.Helper

	uc *data.NotificationMessageRecipientRepo
}

func NewNotificationMessageRecipientService(uc *data.NotificationMessageRecipientRepo, logger log.Logger) *NotificationMessageRecipientService {
	l := log.NewHelper(log.With(logger, "module", "notification-message-recipient/service/admin-service"))
	return &NotificationMessageRecipientService{
		log: l,
		uc:  uc,
	}
}

func (s *NotificationMessageRecipientService) ListNotificationMessageRecipient(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageRecipientResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *NotificationMessageRecipientService) GetNotificationMessageRecipient(ctx context.Context, req *internalMessageV1.GetNotificationMessageRecipientRequest) (*internalMessageV1.NotificationMessageRecipient, error) {
	return s.uc.Get(ctx, req)
}

func (s *NotificationMessageRecipientService) CreateNotificationMessageRecipient(ctx context.Context, req *internalMessageV1.CreateNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	if err := s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageRecipientService) UpdateNotificationMessageRecipient(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	if err := s.uc.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageRecipientService) DeleteNotificationMessageRecipient(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
