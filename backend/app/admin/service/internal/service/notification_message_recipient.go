package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"
	"kratos-admin/app/admin/service/internal/middleware/auth"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
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
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	err = s.uc.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageRecipientService) UpdateNotificationMessageRecipient(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	err = s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageRecipientService) DeleteNotificationMessageRecipient(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageRecipientRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	_, err = s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
