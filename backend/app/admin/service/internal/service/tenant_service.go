package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
	"kratos-admin/pkg/utils/name_set"
)

type TenantService struct {
	adminV1.TenantServiceHTTPServer

	log *log.Helper

	tenantRepo          *data.TenantRepo
	userRepo            *data.UserRepo
	userCredentialsRepo *data.UserCredentialRepo
}

func NewTenantService(
	logger log.Logger,
	tenantRepo *data.TenantRepo,
	userRepo *data.UserRepo,
	userCredentialsRepo *data.UserCredentialRepo,
) *TenantService {
	l := log.NewHelper(log.With(logger, "module", "tenant/service/admin-service"))
	return &TenantService{
		log:                 l,
		tenantRepo:          tenantRepo,
		userRepo:            userRepo,
		userCredentialsRepo: userCredentialsRepo,
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
		userResp, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
			QueryBy: &userV1.GetUserRequest_Id{
				Id: resp.GetAdminUserId(),
			},
		})
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
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if _, err = s.tenantRepo.Create(ctx, req.Data); err != nil {
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
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

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

func (s *TenantService) TenantExists(ctx context.Context, req *userV1.TenantExistsRequest) (*userV1.TenantExistsResponse, error) {
	return s.tenantRepo.TenantExists(ctx, req)
}

func (s *TenantService) CreateTenantWithAdminUser(ctx context.Context, req *adminV1.CreateTenantWithAdminUserRequest) (*emptypb.Empty, error) {
	if req.Tenant == nil || req.User == nil {
		s.log.Error("invalid parameter: tenant or user is nil", req)
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Tenant.CreatedBy = trans.Ptr(operator.UserId)
	req.User.CreatedBy = trans.Ptr(operator.UserId)

	// Check if tenant code or admin username already exists
	if _, err = s.tenantRepo.TenantExists(ctx, &userV1.TenantExistsRequest{
		Code: req.GetTenant().GetCode(),
	}); err != nil {
		s.log.Errorf("check tenant code exists err: %v", err)
		return nil, err
	}

	// Check if admin user exists
	if _, err = s.userRepo.UserExists(ctx, &userV1.UserExistsRequest{
		Username: req.GetUser().GetUsername(),
	}); err != nil {
		s.log.Errorf("check admin user exists err: %v", err)
		return nil, err
	}

	// Create tenant
	var tenant *userV1.Tenant
	if tenant, err = s.tenantRepo.Create(ctx, req.Tenant); err != nil {
		s.log.Errorf("create tenant err: %v", err)
		return nil, err
	}

	req.User.Authority = userV1.User_TENANT_ADMIN.Enum()
	req.User.TenantId = tenant.Id

	// Create tenant admin user
	var adminUser *userV1.User
	if adminUser, err = s.userRepo.Create(ctx, &userV1.CreateUserRequest{
		Data: req.User,
	}); err != nil {
		s.log.Errorf("create tenant admin user err: %v", err)
		return nil, err
	}

	// Create user credential
	if err = s.userCredentialsRepo.Create(ctx, &authenticationV1.CreateUserCredentialRequest{
		Data: &authenticationV1.UserCredential{
			UserId:         adminUser.Id,
			IdentityType:   authenticationV1.UserCredential_USERNAME.Enum(),
			Identifier:     adminUser.Username,
			CredentialType: authenticationV1.UserCredential_PASSWORD_HASH.Enum(),
			Credential:     trans.Ptr(req.GetPassword()),
			IsPrimary:      trans.Ptr(true),
			Status:         authenticationV1.UserCredential_ENABLED.Enum(),
		},
	}); err != nil {
		s.log.Errorf("create tenant admin user credential err: %v", err)
		return nil, err
	}

	// Update tenant with admin user id
	if err = s.tenantRepo.Update(ctx, &userV1.UpdateTenantRequest{
		Data: &userV1.Tenant{
			Id:          tenant.Id,
			AdminUserId: adminUser.Id,
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"id", "admin_user_id"},
		},
	}); err != nil {
		s.log.Errorf("update tenant with admin user id err: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
