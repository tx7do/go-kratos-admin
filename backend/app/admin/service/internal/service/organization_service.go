package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	userV1 "go-wind-admin/api/gen/go/user/service/v1"

	"go-wind-admin/pkg/middleware/auth"
	"go-wind-admin/pkg/utils/name_set"
)

type OrganizationService struct {
	adminV1.OrganizationServiceHTTPServer

	log *log.Helper

	organizationRepo *data.OrganizationRepo
	userRepo         *data.UserRepo
}

func NewOrganizationService(
	ctx *bootstrap.Context,
	organizationRepo *data.OrganizationRepo,
	userRepo *data.UserRepo,
) *OrganizationService {
	return &OrganizationService{
		log:              ctx.NewLoggerHelper("organization/service/admin-service"),
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
	}
}

func (s *OrganizationService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListOrganizationResponse, error) {
	resp, err := s.organizationRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var userSet = make(name_set.UserNameSetMap)

	InitOrganizationNameSetMap(resp.Items, &userSet)

	QueryUserInfoFromRepo(ctx, s.userRepo, &userSet)

	FileOrganizationInfo(resp.Items, &userSet)

	return resp, nil
}

func (s *OrganizationService) Get(ctx context.Context, req *userV1.GetOrganizationRequest) (*userV1.Organization, error) {
	resp, err := s.organizationRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.ManagerId != nil {
		manager, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
			QueryBy: &userV1.GetUserRequest_Id{
				Id: resp.GetManagerId(),
			},
		})
		if err == nil && manager != nil {
			resp.ManagerName = manager.Nickname
		} else {
			s.log.Warnf("Get organization manager user failed: %v", err)
		}
	}

	return resp, nil
}

func (s *OrganizationService) Create(ctx context.Context, req *userV1.CreateOrganizationRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.organizationRepo.Create(ctx, req); err != nil {
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
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.organizationRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OrganizationService) Delete(ctx context.Context, req *userV1.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	if err := s.organizationRepo.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
