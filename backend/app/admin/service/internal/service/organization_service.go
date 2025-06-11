package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type OrganizationService struct {
	adminV1.OrganizationServiceHTTPServer

	log *log.Helper

	repo *data.OrganizationRepo
}

func NewOrganizationService(logger log.Logger, repo *data.OrganizationRepo) *OrganizationService {
	l := log.NewHelper(log.With(logger, "module", "organization/service/admin-service"))
	return &OrganizationService{
		log:  l,
		repo: repo,
	}
}

func (s *OrganizationService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListOrganizationResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *OrganizationService) Get(ctx context.Context, req *userV1.GetOrganizationRequest) (*userV1.Organization, error) {
	return s.repo.Get(ctx, req)
}

func (s *OrganizationService) Create(ctx context.Context, req *userV1.CreateOrganizationRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OrganizationService) Update(ctx context.Context, req *userV1.UpdateOrganizationRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OrganizationService) Delete(ctx context.Context, req *userV1.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
