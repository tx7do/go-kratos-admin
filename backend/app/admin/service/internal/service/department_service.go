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
	"kratos-admin/pkg/utils/name_set"
)

type DepartmentService struct {
	adminV1.DepartmentServiceHTTPServer

	log *log.Helper

	departmentRepo   *data.DepartmentRepo
	organizationRepo *data.OrganizationRepo
	userRepo         *data.UserRepo
}

func NewDepartmentService(
	logger log.Logger,
	departmentRepo *data.DepartmentRepo,
	organizationRepo *data.OrganizationRepo,
	userRepo *data.UserRepo,
) *DepartmentService {
	l := log.NewHelper(log.With(logger, "module", "department/service/admin-service"))
	return &DepartmentService{
		log:              l,
		departmentRepo:   departmentRepo,
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
	}
}

func (s *DepartmentService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListDepartmentResponse, error) {
	resp, err := s.departmentRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var userSet = make(name_set.UserNameSetMap)
	var orgSet = make(name_set.UserNameSetMap)

	InitDepartmentNameSetMap(resp.Items, &userSet, &orgSet)

	QueryUserInfoFromRepo(ctx, s.userRepo, &userSet)
	QueryOrganizationInfoFromRepo(ctx, s.organizationRepo, &orgSet)

	FillDepartmentUserInfo(resp.Items, &userSet)
	FillDepartmentOrganizationInfo(resp.Items, &orgSet)

	return resp, nil
}

func (s *DepartmentService) Get(ctx context.Context, req *userV1.GetDepartmentRequest) (*userV1.Department, error) {
	resp, err := s.departmentRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.ManagerId != nil {
		manager, err := s.userRepo.Get(ctx, resp.GetManagerId())
		if err == nil && manager != nil {
			resp.ManagerName = manager.Username
		} else {
			s.log.Warnf("Get organization manager user failed: %v", err)
		}
	}

	if resp.OrganizationId != nil {
		organization, err := s.organizationRepo.Get(ctx, &userV1.GetOrganizationRequest{Id: resp.GetOrganizationId()})
		if err == nil && organization != nil {
			resp.OrganizationName = organization.Name
		} else {
			s.log.Warnf("Get department organization failed: %v", err)
		}
	}

	return resp, nil
}

func (s *DepartmentService) Create(ctx context.Context, req *userV1.CreateDepartmentRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.departmentRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DepartmentService) Update(ctx context.Context, req *userV1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.departmentRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DepartmentService) Delete(ctx context.Context, req *userV1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	if err := s.departmentRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
