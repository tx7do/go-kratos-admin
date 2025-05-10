package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
	"kratos-admin/app/admin/service/internal/data"
)

type DepartmentService struct {
	adminV1.DepartmentServiceHTTPServer

	log *log.Helper

	uc *data.DepartmentRepo
}

func NewDepartmentService(uc *data.DepartmentRepo, logger log.Logger) *DepartmentService {
	l := log.NewHelper(log.With(logger, "module", "department/service/admin-service"))
	return &DepartmentService{
		log: l,
		uc:  uc,
	}
}

func (s *DepartmentService) ListDepartment(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListDepartmentResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *DepartmentService) GetDepartment(ctx context.Context, req *userV1.GetDepartmentRequest) (*userV1.Department, error) {
	return s.uc.Get(ctx, req)
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, req *userV1.CreateDepartmentRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	if err := s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DepartmentService) UpdateDepartment(ctx context.Context, req *userV1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	if err := s.uc.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DepartmentService) DeleteDepartment(ctx context.Context, req *userV1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
