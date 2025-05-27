package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	authnEngine "github.com/tx7do/kratos-authn/engine"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/jwt"
	"kratos-admin/pkg/middleware/auth"
)

type AuthenticationService struct {
	adminV1.AuthenticationServiceHTTPServer

	userRepo           *data.UserRepo
	userCredentialRepo *data.UserCredentialRepo
	roleRepo           *data.RoleRepo
	tenantRepo         *data.TenantRepo

	userToken *data.UserToken

	authenticator authnEngine.Authenticator

	log *log.Helper
}

func NewAuthenticationService(
	logger log.Logger,
	userRepo *data.UserRepo,
	userCredentialRepo *data.UserCredentialRepo,
	tenantRepo *data.TenantRepo,
	roleRepo *data.RoleRepo,
	userToken *data.UserToken,
	authenticator authnEngine.Authenticator,
) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authn/service/admin-service"))
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
	case authenticationV1.GrantType_password.String():
		return s.doGrantTypePassword(ctx, req)

	case authenticationV1.GrantType_refresh_token.String():
		return s.doGrantTypeRefreshToken(ctx, req)

	case authenticationV1.GrantType_client_credentials.String():
		return s.doGrantTypeClientCredentials(ctx, req)

	default:
		return nil, authenticationV1.ErrorInvalidGrantType("invalid grant type")
	}
}

// doGrantTypePassword 处理授权类型 - 密码
func (s *AuthenticationService) doGrantTypePassword(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	var err error
	if _, err = s.userCredentialRepo.VerifyCredential(ctx, &authenticationV1.VerifyCredentialRequest{
		IdentityType: authenticationV1.IdentityType_PASSWORD,
		Identifier:   req.GetUsername(),
		Credential:   req.GetPassword(),
	}); err != nil {
		return nil, err
	}

	// 获取用户信息
	var user *userV1.User
	user, err = s.userRepo.GetUserByUserName(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	// 验证权限
	if user.GetAuthority() != userV1.UserAuthority_SYS_ADMIN {
		return &authenticationV1.LoginResponse{}, authenticationV1.ErrorForbidden("权限不够")
	}

	// 生成令牌
	accessToken, refreshToken, err := s.userToken.GenerateToken(ctx, user, req.GetClientId())
	if err != nil {
		return nil, err
	}

	return &authenticationV1.LoginResponse{
		TokenType:    authenticationV1.TokenType_bearer.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// doGrantTypeAuthorizationCode 处理授权类型 -
func (s *AuthenticationService) doGrantTypeRefreshToken(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := s.userRepo.Get(ctx, operator.UserId)
	if err != nil {
		return &authenticationV1.LoginResponse{}, err
	}

	// 校验刷新令牌
	if !s.userToken.IsExistRefreshToken(ctx, operator.UserId, req.GetRefreshToken()) {
		return nil, authenticationV1.ErrorIncorrectRefreshToken("invalid refresh token")
	}

	if err = s.userToken.RemoveRefreshToken(ctx, operator.UserId, req.GetRefreshToken()); err != nil {
		s.log.Errorf("remove refresh token failed [%s]", err.Error())
	}

	// 生成令牌
	accessToken, refreshToken, err := s.userToken.GenerateToken(ctx, user, req.GetClientId())
	if err != nil {
		return nil, authenticationV1.ErrorServiceUnavailable("generate token failed")
	}

	return &authenticationV1.LoginResponse{
		TokenType:    authenticationV1.TokenType_bearer.String(),
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
	if req.GetGrantType() != authenticationV1.GrantType_refresh_token.String() {
		return nil, authenticationV1.ErrorInvalidGrantType("invalid grant type")
	}

	return s.doGrantTypeRefreshToken(ctx, req)
}

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

func (s *AuthenticationService) RegisterUser(ctx context.Context, req *authenticationV1.RegisterUserRequest) (*authenticationV1.RegisterUserResponse, error) {
	var err error

	var tenantId *uint32
	tenant, err := s.tenantRepo.GetTenantByTenantCode(ctx, req.GetTenantCode())
	if tenant != nil {
		tenantId = tenant.Id
	}

	user, err := s.userRepo.Create(ctx, &userV1.CreateUserRequest{
		Data: &userV1.User{
			TenantId:  tenantId,
			Username:  trans.Ptr(req.Username),
			Email:     req.Email,
			Authority: trans.Ptr(userV1.UserAuthority_CUSTOMER_USER),
			Status:    trans.Ptr(userV1.UserStatus_ON),
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

			IdentityType: authenticationV1.IdentityType_PASSWORD.Enum(),
			Identifier:   trans.Ptr(req.GetUsername()),

			CredentialType: authenticationV1.CredentialType_PASSWORD_HASH.Enum(),
			Credential:     trans.Ptr(req.GetPassword()),

			IsPrimary: trans.Ptr(true),
			Status:    authenticationV1.UserCredentialStatus_ENABLED.Enum(),
		},
	}); err != nil {
		s.log.Errorf("create user credentials error: %v", err)
		return nil, err
	}

	return &authenticationV1.RegisterUserResponse{
		UserId: user.GetId(),
	}, nil
}

func (s *AuthenticationService) ChangePassword(ctx context.Context, req *authenticationV1.ChangePasswordRequest) (*emptypb.Empty, error) {
	err := s.userCredentialRepo.ChangeCredential(ctx, &authenticationV1.ChangeCredentialRequest{
		IdentityType:  authenticationV1.IdentityType_PASSWORD,
		Identifier:    req.GetUsername(),
		OldCredential: req.GetOldPassword(),
		NewCredential: req.GetNewPassword(),
	})
	return &emptypb.Empty{}, err
}

func (s *AuthenticationService) WhoAmI(ctx context.Context, _ *emptypb.Empty) (*authenticationV1.WhoAmIResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	return &authenticationV1.WhoAmIResponse{
		UserId:    operator.UserId,
		Username:  operator.Username,
		Authority: operator.Authority,
	}, nil
}
