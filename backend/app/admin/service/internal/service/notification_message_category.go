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

type NotificationMessageCategoryService struct {
	adminV1.NotificationMessageCategoryServiceHTTPServer

	log *log.Helper

	uc *data.NotificationMessageCategoryRepo
}

func NewNotificationMessageCategoryService(uc *data.NotificationMessageCategoryRepo, logger log.Logger) *NotificationMessageCategoryService {
	l := log.NewHelper(log.With(logger, "module", "notification-message-category/service/admin-service"))
	return &NotificationMessageCategoryService{
		log: l,
		uc:  uc,
	}
}

func (s *NotificationMessageCategoryService) ListNotificationMessageCategory(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageCategoryResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *NotificationMessageCategoryService) GetNotificationMessageCategory(ctx context.Context, req *internalMessageV1.GetNotificationMessageCategoryRequest) (*internalMessageV1.NotificationMessageCategory, error) {
	return s.uc.Get(ctx, req)
}

func (s *NotificationMessageCategoryService) CreateNotificationMessageCategory(ctx context.Context, req *internalMessageV1.CreateNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
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

func (s *NotificationMessageCategoryService) UpdateNotificationMessageCategory(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
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

func (s *NotificationMessageCategoryService) DeleteNotificationMessageCategory(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
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
