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

func (s *NotificationMessageService) ListNotificationMessage(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *NotificationMessageService) GetNotificationMessage(ctx context.Context, req *internalMessageV1.GetNotificationMessageRequest) (*internalMessageV1.NotificationMessage, error) {
	return s.uc.Get(ctx, req)
}

func (s *NotificationMessageService) CreateNotificationMessage(ctx context.Context, req *internalMessageV1.CreateNotificationMessageRequest) (*emptypb.Empty, error) {
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

func (s *NotificationMessageService) UpdateNotificationMessage(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageRequest) (*emptypb.Empty, error) {
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

func (s *NotificationMessageService) DeleteNotificationMessage(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageRequest) (*emptypb.Empty, error) {
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
