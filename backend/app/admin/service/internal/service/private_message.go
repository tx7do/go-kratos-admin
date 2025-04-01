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

type PrivateMessageService struct {
	adminV1.PrivateMessageServiceHTTPServer

	log *log.Helper

	uc *data.PrivateMessageRepo
}

func NewPrivateMessageService(uc *data.PrivateMessageRepo, logger log.Logger) *PrivateMessageService {
	l := log.NewHelper(log.With(logger, "module", "private-message/service/admin-service"))
	return &PrivateMessageService{
		log: l,
		uc:  uc,
	}
}

func (s *PrivateMessageService) ListPrivateMessage(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListPrivateMessageResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *PrivateMessageService) GetPrivateMessage(ctx context.Context, req *internalMessageV1.GetPrivateMessageRequest) (*internalMessageV1.PrivateMessage, error) {
	return s.uc.Get(ctx, req)
}

func (s *PrivateMessageService) CreatePrivateMessage(ctx context.Context, req *internalMessageV1.CreatePrivateMessageRequest) (*emptypb.Empty, error) {
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

func (s *PrivateMessageService) UpdatePrivateMessage(ctx context.Context, req *internalMessageV1.UpdatePrivateMessageRequest) (*emptypb.Empty, error) {
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

func (s *PrivateMessageService) DeletePrivateMessage(ctx context.Context, req *internalMessageV1.DeletePrivateMessageRequest) (*emptypb.Empty, error) {
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
