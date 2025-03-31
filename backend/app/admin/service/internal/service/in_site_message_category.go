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
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type InSiteMessageCategoryService struct {
	adminV1.InSiteMessageCategoryServiceHTTPServer

	log *log.Helper

	uc *data.InSiteMessageCategoryRepo
}

func NewInSiteMessageCategoryService(uc *data.InSiteMessageCategoryRepo, logger log.Logger) *InSiteMessageCategoryService {
	l := log.NewHelper(log.With(logger, "module", "in-site-message-category/service/admin-service"))
	return &InSiteMessageCategoryService{
		log: l,
		uc:  uc,
	}
}

func (s *InSiteMessageCategoryService) ListInSiteMessageCategory(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListInSiteMessageCategoryResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *InSiteMessageCategoryService) GetInSiteMessageCategory(ctx context.Context, req *systemV1.GetInSiteMessageCategoryRequest) (*systemV1.InSiteMessageCategory, error) {
	return s.uc.Get(ctx, req)
}

func (s *InSiteMessageCategoryService) CreateInSiteMessageCategory(ctx context.Context, req *systemV1.CreateInSiteMessageCategoryRequest) (*emptypb.Empty, error) {
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

func (s *InSiteMessageCategoryService) UpdateInSiteMessageCategory(ctx context.Context, req *systemV1.UpdateInSiteMessageCategoryRequest) (*emptypb.Empty, error) {
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

func (s *InSiteMessageCategoryService) DeleteInSiteMessageCategory(ctx context.Context, req *systemV1.DeleteInSiteMessageCategoryRequest) (*emptypb.Empty, error) {
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
