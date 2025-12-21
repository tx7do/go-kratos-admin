package service

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	authenticationV1 "go-wind-admin/api/gen/go/authentication/service/v1"
	userV1 "go-wind-admin/api/gen/go/user/service/v1"

	"go-wind-admin/pkg/middleware/auth"
	"go-wind-admin/pkg/utils/name_set"
)

type UserService struct {
	adminV1.UserServiceHTTPServer

	log *log.Helper

	userRepo           *data.UserRepo
	roleRepo           *data.RoleRepo
	userCredentialRepo *data.UserCredentialRepo
	positionRepo       *data.PositionRepo
	departmentRepo     *data.DepartmentRepo
	organizationRepo   *data.OrganizationRepo
	tenantRepo         *data.TenantRepo

	userRoleRepo     *data.UserRoleRepo
	userPositionRepo *data.UserPositionRepo
}

func NewUserService(
	ctx *bootstrap.Context,
	userRepo *data.UserRepo,
	roleRepo *data.RoleRepo,
	userCredentialRepo *data.UserCredentialRepo,
	positionRepo *data.PositionRepo,
	departmentRepo *data.DepartmentRepo,
	organizationRepo *data.OrganizationRepo,
	tenantRepo *data.TenantRepo,
	userRoleRepo *data.UserRoleRepo,
	userPositionRepo *data.UserPositionRepo,
) *UserService {
	l := log.NewHelper(log.With(ctx.Logger, "module", "user/service/admin-service"))
	svc := &UserService{
		log:                l,
		userRepo:           userRepo,
		roleRepo:           roleRepo,
		userCredentialRepo: userCredentialRepo,
		positionRepo:       positionRepo,
		departmentRepo:     departmentRepo,
		organizationRepo:   organizationRepo,
		tenantRepo:         tenantRepo,
		userRoleRepo:       userRoleRepo,
		userPositionRepo:   userPositionRepo,
	}

	svc.init()

	return svc
}

func (s *UserService) init() {
	ctx := context.Background()
	if count, _ := s.userRepo.Count(ctx, []func(s *sql.Selector){}); count == 0 {
		_ = s.CreateDefaultUser(ctx)
	}
}

func (s *UserService) initUserNameSetMap(
	users []*userV1.User,
	tenantSet *name_set.UserNameSetMap,
	orgSet *name_set.UserNameSetMap,
	deptSet *name_set.UserNameSetMap,
	posSet *name_set.UserNameSetMap,
	roleSet *name_set.UserNameSetMap,
) {
	for _, v := range users {
		if v.TenantId != nil {
			(*tenantSet)[v.GetTenantId()] = nil
		}
		if v.OrgId != nil {
			(*orgSet)[v.GetOrgId()] = nil
		}
		if v.DepartmentId != nil {
			(*deptSet)[v.GetDepartmentId()] = nil
		}
		if v.PositionId != nil {
			(*posSet)[v.GetPositionId()] = nil
		}
		for _, roleId := range v.RoleIds {
			(*roleSet)[roleId] = nil
		}
	}
}

func (s *UserService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	resp, err := s.userRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var roleSet = make(name_set.UserNameSetMap)
	var tenantSet = make(name_set.UserNameSetMap)
	var orgSet = make(name_set.UserNameSetMap)
	var deptSet = make(name_set.UserNameSetMap)
	var posSet = make(name_set.UserNameSetMap)

	s.initUserNameSetMap(resp.Items, &tenantSet, &orgSet, &deptSet, &posSet, &roleSet)

	QueryTenantInfoFromRepo(ctx, s.tenantRepo, &tenantSet)
	QueryOrganizationInfoFromRepo(ctx, s.organizationRepo, &orgSet)
	QueryDepartmentInfoFromRepo(ctx, s.departmentRepo, &deptSet)
	QueryPositionInfoFromRepo(ctx, s.positionRepo, &posSet)
	QueryRoleInfoFromRepo(ctx, s.roleRepo, &roleSet)

	for k, v := range tenantSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(resp.Items); i++ {
			if resp.Items[i].TenantId != nil && resp.Items[i].GetTenantId() == k {
				resp.Items[i].TenantName = &v.UserName
			}
		}
	}

	for k, v := range orgSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(resp.Items); i++ {
			if resp.Items[i].OrgId != nil && resp.Items[i].GetOrgId() == k {
				resp.Items[i].OrgName = &v.UserName
			}
		}
	}

	for k, v := range deptSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(resp.Items); i++ {
			if resp.Items[i].DepartmentId != nil && resp.Items[i].GetDepartmentId() == k {
				resp.Items[i].DepartmentName = &v.UserName
			}
		}
	}

	for k, v := range posSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(resp.Items); i++ {
			if resp.Items[i].PositionId != nil && resp.Items[i].GetPositionId() == k {
				resp.Items[i].PositionName = &v.UserName
			}
		}
	}

	for k, v := range roleSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(resp.Items); i++ {
			for _, roleId := range resp.Items[i].RoleIds {
				if roleId == k {
					resp.Items[i].RoleNames = append(resp.Items[i].RoleNames, v.UserName)
					resp.Items[i].Roles = append(resp.Items[i].Roles, v.Code)
				}
			}
		}
	}

	return resp, nil
}

func (s *UserService) fillUserInfo(ctx context.Context, user *userV1.User) error {
	if user.TenantId != nil {
		tenant, err := s.tenantRepo.Get(ctx, &userV1.GetTenantRequest{QueryBy: &userV1.GetTenantRequest_Id{Id: user.GetTenantId()}})
		if err == nil && tenant != nil {
			user.TenantName = tenant.Name
		} else {
			s.log.Warnf("Get user tenant failed: %v", err)
		}
	}

	if user.OrgId != nil {
		organization, err := s.organizationRepo.Get(ctx, &userV1.GetOrganizationRequest{QueryBy: &userV1.GetOrganizationRequest_Id{Id: user.GetOrgId()}})
		if err == nil && organization != nil {
			user.OrgName = organization.Name
		} else {
			s.log.Warnf("Get user organization failed: %v", err)
		}
	}

	if user.DepartmentId != nil {
		department, err := s.departmentRepo.Get(ctx, &userV1.GetDepartmentRequest{QueryBy: &userV1.GetDepartmentRequest_Id{Id: user.GetDepartmentId()}})
		if err == nil && department != nil {
			user.DepartmentName = department.Name
		} else {
			s.log.Warnf("Get user department failed: %v", err)
		}
	}

	if user.PositionId != nil {
		position, err := s.positionRepo.Get(ctx, &userV1.GetPositionRequest{QueryBy: &userV1.GetPositionRequest_Id{Id: user.GetPositionId()}})
		if err == nil && position != nil {
			user.PositionName = position.Name
		} else {
			s.log.Warnf("Get user position failed: %v", err)
		}
	}

	if len(user.RoleIds) > 0 {
		roles, err := s.roleRepo.ListRolesByRoleIds(ctx, user.RoleIds)
		if err == nil && roles != nil {
			var roleNames []string
			var roleCodes []string
			for _, role := range roles {
				roleNames = append(roleNames, role.GetName())
				roleCodes = append(roleCodes, role.GetCode())
			}
			user.RoleNames = roleNames
			user.Roles = roleCodes
		} else {
			s.log.Warnf("Get user roles failed: %v", err)
		}
	}

	return nil
}

func (s *UserService) Get(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	resp, err := s.userRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	_ = s.fillUserInfo(ctx, resp)

	return resp, nil
}

func (s *UserService) Create(ctx context.Context, req *userV1.CreateUserRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 获取操作者的用户信息
	operatorUser, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	// 校验操作者的权限
	if operatorUser.GetAuthority() != userV1.User_SYS_ADMIN && operatorUser.GetAuthority() != userV1.User_TENANT_ADMIN {
		s.log.Infof("operator authority: %v", operatorUser.GetAuthority())
		return nil, adminV1.ErrorForbidden("权限不够")
	}

	if req.Data.Authority == nil {
		req.Data.Authority = userV1.User_CUSTOMER_USER.Enum()
	}

	if req.Data.Authority != nil {
		if operatorUser.GetAuthority() < req.Data.GetAuthority() {
			s.log.Infof("operator authority: %v, create authority: %v", operatorUser.GetAuthority(), req.Data.GetAuthority())
			return nil, adminV1.ErrorForbidden("不能够创建同级用户或者比自己权限高的用户")
		}
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)
	req.Data.TenantId = operator.TenantId

	// 创建用户
	var user *userV1.User
	if user, err = s.userRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	if len(req.GetPassword()) == 0 {
		// 如果没有设置密码，则默认设置为 666666
		req.Password = trans.Ptr("666666")
	}

	if len(req.GetPassword()) > 0 {
		if err = s.userCredentialRepo.Create(ctx, &authenticationV1.CreateUserCredentialRequest{
			Data: &authenticationV1.UserCredential{
				UserId:   user.Id,
				TenantId: user.TenantId,

				IdentityType: authenticationV1.UserCredential_USERNAME.Enum(),
				Identifier:   req.Data.Username,

				CredentialType: authenticationV1.UserCredential_PASSWORD_HASH.Enum(),
				Credential:     req.Password,

				IsPrimary: trans.Ptr(true),
				Status:    authenticationV1.UserCredential_ENABLED.Enum(),
			},
		}); err != nil {
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *UserService) Update(ctx context.Context, req *userV1.UpdateUserRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 获取操作者的用户信息
	operatorUser, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	// 校验操作者的权限
	if operatorUser.GetAuthority() != userV1.User_SYS_ADMIN {
		return nil, adminV1.ErrorForbidden("权限不够")
	}

	if req.Data.Authority != nil {
		if operatorUser.GetAuthority() < req.Data.GetAuthority() {
			return nil, adminV1.ErrorForbidden("不能够赋权同级用户或者比自己权限高的用户")
		}
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	// 更新用户
	if err = s.userRepo.Update(ctx, req); err != nil {
		s.log.Error(err)
		return nil, err
	}

	if len(req.GetPassword()) > 0 {
		if err = s.userCredentialRepo.ResetCredential(ctx, &authenticationV1.ResetCredentialRequest{
			IdentityType:  authenticationV1.UserCredential_USERNAME,
			Identifier:    req.Data.GetUsername(),
			NewCredential: req.GetPassword(),
		}); err != nil {
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *UserService) Delete(ctx context.Context, req *userV1.DeleteUserRequest) (*emptypb.Empty, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 获取操作者的用户信息
	operatorUser, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	// 校验操作者的权限
	if operatorUser.GetAuthority() != userV1.User_SYS_ADMIN {
		return nil, adminV1.ErrorForbidden("权限不够")
	}

	// 获取将被删除的用户信息
	user, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: req.GetId(),
		},
	})
	if err != nil {
		return nil, err
	}

	// 不能删除超级管理员
	if user.GetAuthority() == userV1.User_SYS_ADMIN {
		return nil, adminV1.ErrorForbidden("闹哪样？不能删除超级管理员！")
	}

	if operatorUser.GetAuthority() == user.GetAuthority() {
		return nil, adminV1.ErrorForbidden("不能删除同级用户！")
	}

	// 删除用户
	err = s.userRepo.Delete(ctx, req.GetId())

	return &emptypb.Empty{}, err
}

func (s *UserService) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	return s.userRepo.UserExists(ctx, req)
}

// EditUserPassword 修改用户密码
func (s *UserService) EditUserPassword(ctx context.Context, req *userV1.EditUserPasswordRequest) (*emptypb.Empty, error) {
	// 获取操作者的用户信息
	u, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: req.GetUserId(),
		},
	})
	if err != nil {
		return nil, err
	}

	if err = s.userCredentialRepo.ResetCredential(ctx, &authenticationV1.ResetCredentialRequest{
		IdentityType:  authenticationV1.UserCredential_USERNAME,
		Identifier:    u.GetUsername(),
		NewCredential: req.GetNewPassword(),
		NeedDecrypt:   false,
	}); err != nil {
		s.log.Errorf("reset user password err: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// CreateDefaultUser 创建默认用户，即超级用户
func (s *UserService) CreateDefaultUser(ctx context.Context) error {
	const (
		defaultUsername = "admin"
		defaultPassword = "admin"
	)

	var err error

	if _, err = s.userRepo.Create(ctx, &userV1.CreateUserRequest{
		Data: &userV1.User{
			Id:        trans.Ptr(uint32(1)),
			Username:  trans.Ptr(defaultUsername),
			Realname:  trans.Ptr("喵个咪"),
			Nickname:  trans.Ptr("鹳狸猿"),
			Region:    trans.Ptr("中国"),
			Email:     trans.Ptr("admin@gmail.com"),
			Authority: userV1.User_SYS_ADMIN.Enum(),
			RoleIds:   []uint32{1},
		},
	}); err != nil {
		s.log.Errorf("create default user err: %v", err)
		return err
	}

	if err = s.userCredentialRepo.Create(ctx, &authenticationV1.CreateUserCredentialRequest{
		Data: &authenticationV1.UserCredential{
			UserId:         trans.Ptr(uint32(1)),
			IdentityType:   authenticationV1.UserCredential_USERNAME.Enum(),
			Identifier:     trans.Ptr(defaultUsername),
			CredentialType: authenticationV1.UserCredential_PASSWORD_HASH.Enum(),
			Credential:     trans.Ptr(defaultPassword),
			IsPrimary:      trans.Ptr(true),
			Status:         authenticationV1.UserCredential_ENABLED.Enum(),
		},
	}); err != nil {
		s.log.Errorf("create default user credential err: %v", err)
		return err
	}

	return err
}
