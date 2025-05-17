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

func (s *DepartmentService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListDepartmentResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *DepartmentService) Get(ctx context.Context, req *userV1.GetDepartmentRequest) (*userV1.Department, error) {
	return s.uc.Get(ctx, req)
}

func (s *DepartmentService) Create(ctx context.Context, req *userV1.CreateDepartmentRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DepartmentService) Update(ctx context.Context, req *userV1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.uc.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DepartmentService) Delete(ctx context.Context, req *userV1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
