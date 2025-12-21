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

type PositionService struct {
	adminV1.PositionServiceHTTPServer

	log *log.Helper

	positionRepo     *data.PositionRepo
	departmentRepo   *data.DepartmentRepo
	organizationRepo *data.OrganizationRepo
}

func NewPositionService(
	ctx *bootstrap.Context,
	positionRepo *data.PositionRepo,
	departmentRepo *data.DepartmentRepo,
	organizationRepo *data.OrganizationRepo,
) *PositionService {
	l := log.NewHelper(log.With(ctx.Logger, "module", "position/service/admin-service"))
	return &PositionService{
		log:              l,
		positionRepo:     positionRepo,
		departmentRepo:   departmentRepo,
		organizationRepo: organizationRepo,
	}
}

func (s *PositionService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListPositionResponse, error) {
	resp, err := s.positionRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var deptSet = make(name_set.UserNameSetMap)
	var orgSet = make(name_set.UserNameSetMap)

	InitPositionNameSetMap(resp.Items, &orgSet, &deptSet)

	QueryOrganizationInfoFromRepo(ctx, s.organizationRepo, &orgSet)
	QueryDepartmentInfoFromRepo(ctx, s.departmentRepo, &deptSet)

	FillPositionOrganizationInfo(resp.Items, &orgSet)
	FillPositionDepartmentInfo(resp.Items, &deptSet)

	return resp, nil
}

func (s *PositionService) Get(ctx context.Context, req *userV1.GetPositionRequest) (*userV1.Position, error) {
	resp, err := s.positionRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.OrganizationId != nil {
		organization, err := s.organizationRepo.Get(ctx, &userV1.GetOrganizationRequest{QueryBy: &userV1.GetOrganizationRequest_Id{Id: resp.GetOrganizationId()}})
		if err == nil && organization != nil {
			resp.OrganizationName = organization.Name
		} else {
			s.log.Warnf("Get position organization failed: %v", err)
		}
	}

	if resp.DepartmentId != nil {
		department, err := s.departmentRepo.Get(ctx, &userV1.GetDepartmentRequest{QueryBy: &userV1.GetDepartmentRequest_Id{Id: resp.GetDepartmentId()}})
		if err == nil && department != nil {
			resp.DepartmentName = department.Name
		} else {
			s.log.Warnf("Get position department failed: %v", err)
		}
	}

	return resp, nil
}

func (s *PositionService) Create(ctx context.Context, req *userV1.CreatePositionRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.positionRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PositionService) Update(ctx context.Context, req *userV1.UpdatePositionRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.positionRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PositionService) Delete(ctx context.Context, req *userV1.DeletePositionRequest) (*emptypb.Empty, error) {
	if err := s.positionRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
