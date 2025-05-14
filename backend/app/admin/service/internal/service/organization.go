package service

import (
	"context"
	"kratos-admin/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type OrganizationService struct {
	adminV1.OrganizationServiceHTTPServer

	log *log.Helper

	uc *data.OrganizationRepo
}

func NewOrganizationService(uc *data.OrganizationRepo, logger log.Logger) *OrganizationService {
	l := log.NewHelper(log.With(logger, "module", "organization/service/admin-service"))
	return &OrganizationService{
		log: l,
		uc:  uc,
	}
}

func (s *OrganizationService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListOrganizationResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *OrganizationService) Get(ctx context.Context, req *userV1.GetOrganizationRequest) (*userV1.Organization, error) {
	return s.uc.Get(ctx, req)
}

func (s *OrganizationService) Create(ctx context.Context, req *userV1.CreateOrganizationRequest) (*emptypb.Empty, error) {
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

func (s *OrganizationService) Update(ctx context.Context, req *userV1.UpdateOrganizationRequest) (*emptypb.Empty, error) {
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

func (s *OrganizationService) Delete(ctx context.Context, req *userV1.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
