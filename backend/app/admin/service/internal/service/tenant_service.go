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

type TenantService struct {
	adminV1.TenantServiceHTTPServer

	log *log.Helper

	tenantRepo *data.TenantRepo
	userRepo   *data.UserRepo
}

func NewTenantService(
	logger log.Logger,
	tenantRepo *data.TenantRepo,
	userRepo *data.UserRepo,
) *TenantService {
	l := log.NewHelper(log.With(logger, "module", "tenant/service/admin-service"))
	return &TenantService{
		log:        l,
		tenantRepo: tenantRepo,
		userRepo:   userRepo,
	}
}

func (s *TenantService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListTenantResponse, error) {
	resp, err := s.tenantRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var userSet = make(name_set.UserNameSetMap)

	for _, v := range resp.Items {
		if v.AdminUserId != nil {
			userSet[v.GetAdminUserId()] = nil
		}
	}

	QueryUserInfoFromRepo(ctx, s.userRepo, &userSet)

	for _, v := range resp.Items {
		if v.AdminUserId != nil {
			if userInfo, ok := userSet[v.GetAdminUserId()]; ok && userInfo != nil {
				v.AdminUserName = &userInfo.UserName
			}
		}
	}

	return resp, nil
}

func (s *TenantService) Get(ctx context.Context, req *userV1.GetTenantRequest) (*userV1.Tenant, error) {
	resp, err := s.tenantRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.AdminUserId != nil {
		userResp, err := s.userRepo.Get(ctx, resp.GetAdminUserId())
		if err != nil {
			s.log.Errorf("failed to get admin user info: %v", err)
		} else {
			resp.AdminUserName = userResp.Username
		}
	}

	return resp, nil
}

func (s *TenantService) Create(ctx context.Context, req *userV1.CreateTenantRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.tenantRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TenantService) Update(ctx context.Context, req *userV1.UpdateTenantRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.tenantRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TenantService) Delete(ctx context.Context, req *userV1.DeleteTenantRequest) (*emptypb.Empty, error) {
	if err := s.tenantRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
