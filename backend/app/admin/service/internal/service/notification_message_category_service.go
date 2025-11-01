package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type NotificationMessageCategoryService struct {
	adminV1.NotificationMessageCategoryServiceHTTPServer

	log *log.Helper

	repo *data.NotificationMessageCategoryRepo
}

func NewNotificationMessageCategoryService(logger log.Logger, repo *data.NotificationMessageCategoryRepo) *NotificationMessageCategoryService {
	l := log.NewHelper(log.With(logger, "module", "notification-message-category/service/admin-service"))
	return &NotificationMessageCategoryService{
		log:  l,
		repo: repo,
	}
}

func (s *NotificationMessageCategoryService) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageCategoryResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *NotificationMessageCategoryService) Get(ctx context.Context, req *internalMessageV1.GetNotificationMessageCategoryRequest) (*internalMessageV1.NotificationMessageCategory, error) {
	return s.repo.Get(ctx, req)
}

func (s *NotificationMessageCategoryService) Create(ctx context.Context, req *internalMessageV1.CreateNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageCategoryService) Update(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NotificationMessageCategoryService) Delete(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageCategoryRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
