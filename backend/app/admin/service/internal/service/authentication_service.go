package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	authenticationV1 "go-wind-admin/api/gen/go/authentication/service/v1"
	userV1 "go-wind-admin/api/gen/go/user/service/v1"

	"go-wind-admin/pkg/jwt"
	"go-wind-admin/pkg/middleware/auth"
)

type AuthenticationService struct {
	adminV1.AuthenticationServiceHTTPServer

	userRepo           *data.UserRepo
	userCredentialRepo *data.UserCredentialRepo
	roleRepo           *data.RoleRepo
	tenantRepo         *data.TenantRepo

	userToken *data.UserTokenCacheRepo

	authenticator authnEngine.Authenticator

	log *log.Helper
}

func NewAuthenticationService(
	ctx *bootstrap.Context,
	userRepo *data.UserRepo,
	userCredentialRepo *data.UserCredentialRepo,
	tenantRepo *data.TenantRepo,
	roleRepo *data.RoleRepo,
	userToken *data.UserTokenCacheRepo,
	authenticator authnEngine.Authenticator,
) *AuthenticationService {
	l := log.NewHelper(log.With(ctx.Logger, "module", "authn/service/admin-service"))
	return &AuthenticationService{
		log:                l,
		userRepo:           userRepo,
		userCredentialRepo: userCredentialRepo,
		tenantRepo:         tenantRepo,
		roleRepo:           roleRepo,
		userToken:          userToken,
		authenticator:      authenticator,
	}
}

// Login 登录
func (s *AuthenticationService) Login(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	switch req.GetGrantType() {
	case authenticationV1.GrantType_password:
		return s.doGrantTypePassword(ctx, req)

	case authenticationV1.GrantType_refresh_token:
		return s.doGrantTypeRefreshToken(ctx, req)

	case authenticationV1.GrantType_client_credentials:
		return s.doGrantTypeClientCredentials(ctx, req)

	default:
		return nil, authenticationV1.ErrorInvalidGrantType("invalid grant type")
	}
}

// checkAuthority 检查用户权限
func (s *AuthenticationService) checkAuthority(user *userV1.User) error {
	if user == nil {
		return authenticationV1.ErrorUnauthorized("用户不存在")
	}

	// 仅允许系统管理员和租户管理员登录后台管理系统
	if user.GetAuthority() != userV1.User_SYS_ADMIN && user.GetAuthority() != userV1.User_TENANT_ADMIN {
		s.log.Errorf("user [%d] authority [%s] is not allowed to login admin system", user.GetId(), user.GetAuthority().String())
		return authenticationV1.ErrorForbidden("权限不够")
	}

	return nil
}

// doGrantTypePassword 处理授权类型 - 密码
func (s *AuthenticationService) doGrantTypePassword(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	var err error
	if _, err = s.userCredentialRepo.VerifyCredential(ctx, &authenticationV1.VerifyCredentialRequest{
		IdentityType: authenticationV1.UserCredential_USERNAME,
		Identifier:   req.GetUsername(),
		Credential:   req.GetPassword(),
		NeedDecrypt:  true,
	}); err != nil {
		return nil, err
	}

	// 获取用户信息
	var user *userV1.User
	user, err = s.userRepo.Get(ctx, &userV1.GetUserRequest{QueryBy: &userV1.GetUserRequest_UserName{UserName: req.GetUsername()}})
	if err != nil {
		return nil, err
	}

	// 验证权限
	if err = s.checkAuthority(user); err != nil {
		return nil, err
	}

	roleCodes, err := s.roleRepo.ListRoleCodesByRoleIds(ctx, user.GetRoleIds())
	if err != nil {
		s.log.Errorf("get user role codes failed [%s]", err.Error())
	}
	if roleCodes != nil {
		user.Roles = roleCodes
	}

	// 生成令牌
	accessToken, refreshToken, err := s.userToken.GenerateToken(ctx, user, req.GetClientId())
	if err != nil {
		return nil, err
	}

	return &authenticationV1.LoginResponse{
		TokenType:    authenticationV1.TokenType_bearer,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// doGrantTypeAuthorizationCode 处理授权类型 - 刷新令牌
func (s *AuthenticationService) doGrantTypeRefreshToken(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		return &authenticationV1.LoginResponse{}, err
	}

	// 验证权限
	if err = s.checkAuthority(user); err != nil {
		return nil, err
	}

	// 校验刷新令牌
	if !s.userToken.IsExistRefreshToken(ctx, operator.UserId, req.GetRefreshToken()) {
		return nil, authenticationV1.ErrorIncorrectRefreshToken("invalid refresh token")
	}

	if err = s.userToken.RemoveRefreshToken(ctx, operator.UserId, req.GetRefreshToken()); err != nil {
		s.log.Errorf("remove refresh token failed [%s]", err.Error())
	}

	roleCodes, err := s.roleRepo.ListRoleCodesByRoleIds(ctx, user.GetRoleIds())
	if err != nil {
		s.log.Errorf("get user role codes failed [%s]", err.Error())
	}
	if roleCodes != nil {
		user.Roles = roleCodes
	}

	// 生成令牌
	accessToken, refreshToken, err := s.userToken.GenerateToken(ctx, user, req.GetClientId())
	if err != nil {
		return nil, authenticationV1.ErrorServiceUnavailable("generate token failed")
	}

	return &authenticationV1.LoginResponse{
		TokenType:    authenticationV1.TokenType_bearer,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// doGrantTypeClientCredentials 处理授权类型 - 客户端凭据
func (s *AuthenticationService) doGrantTypeClientCredentials(_ context.Context, _ *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	return nil, authenticationV1.ErrorInvalidGrantType("invalid grant type")
}

// Logout 登出
func (s *AuthenticationService) Logout(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	if err = s.userToken.RemoveToken(ctx, operator.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// RefreshToken 刷新令牌
func (s *AuthenticationService) RefreshToken(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	// 校验授权类型
	if req.GetGrantType() != authenticationV1.GrantType_refresh_token {
		return nil, authenticationV1.ErrorInvalidGrantType("invalid grant type")
	}

	return s.doGrantTypeRefreshToken(ctx, req)
}

// ValidateToken 验证令牌
func (s *AuthenticationService) ValidateToken(_ context.Context, req *authenticationV1.ValidateTokenRequest) (*authenticationV1.ValidateTokenResponse, error) {
	ret, err := s.authenticator.AuthenticateToken(req.GetToken())
	if err != nil {
		return &authenticationV1.ValidateTokenResponse{
			IsValid: false,
		}, err
	}

	claims, err := jwt.NewUserTokenPayloadWithClaims(ret)
	if err != nil {
		return &authenticationV1.ValidateTokenResponse{
			IsValid: false,
		}, err
	}

	return &authenticationV1.ValidateTokenResponse{
		IsValid: true,
		Claim:   claims,
	}, nil
}

// RegisterUser 注册前台用户
func (s *AuthenticationService) RegisterUser(ctx context.Context, req *authenticationV1.RegisterUserRequest) (*authenticationV1.RegisterUserResponse, error) {
	var err error

	var tenantId *uint32
	tenant, err := s.tenantRepo.Get(ctx, &userV1.GetTenantRequest{QueryBy: &userV1.GetTenantRequest_Code{Code: req.GetTenantCode()}})
	if tenant != nil {
		tenantId = tenant.Id
	}

	user, err := s.userRepo.Create(ctx, &userV1.CreateUserRequest{
		Data: &userV1.User{
			TenantId:  tenantId,
			Username:  trans.Ptr(req.Username),
			Email:     req.Email,
			Authority: trans.Ptr(userV1.User_CUSTOMER_USER),
			Status:    trans.Ptr(userV1.User_ON),
		},
	})
	if err != nil {
		s.log.Errorf("create user error: %v", err)
		return nil, err
	}

	if err = s.userCredentialRepo.Create(ctx, &authenticationV1.CreateUserCredentialRequest{
		Data: &authenticationV1.UserCredential{
			UserId:   user.Id,
			TenantId: user.TenantId,

			IdentityType: authenticationV1.UserCredential_USERNAME.Enum(),
			Identifier:   trans.Ptr(req.GetUsername()),

			CredentialType: authenticationV1.UserCredential_PASSWORD_HASH.Enum(),
			Credential:     trans.Ptr(req.GetPassword()),

			IsPrimary: trans.Ptr(true),
			Status:    authenticationV1.UserCredential_ENABLED.Enum(),
		},
	}); err != nil {
		s.log.Errorf("create user credentials error: %v", err)
		return nil, err
	}

	return &authenticationV1.RegisterUserResponse{
		UserId: user.GetId(),
	}, nil
}

func (s *AuthenticationService) WhoAmI(ctx context.Context, _ *emptypb.Empty) (*authenticationV1.WhoAmIResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	return &authenticationV1.WhoAmIResponse{
		UserId:    operator.GetUserId(),
		Username:  operator.GetUsername(),
		Authority: operator.GetAuthority(),
	}, nil
}
