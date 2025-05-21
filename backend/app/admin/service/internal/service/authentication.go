package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type AuthenticationService struct {
	adminV1.AuthenticationServiceHTTPServer

	userRepo  *data.UserRepo
	userToken *data.UserToken
	roleRepo  *data.RoleRepo

	log *log.Helper
}

func NewAuthenticationService(
	logger log.Logger,
	userRepo *data.UserRepo,
	userToken *data.UserToken,
	roleRepo *data.RoleRepo,
) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authn/service/admin-service"))
	return &AuthenticationService{
		log:       l,
		userRepo:  userRepo,
		userToken: userToken,
		roleRepo:  roleRepo,
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
	if _, err = s.userRepo.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
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
	accessToken, refreshToken, err := s.userToken.GenerateToken(ctx, user)
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
	// 获取用户信息
	user, err := s.userRepo.Get(ctx, req.GetOperatorId())
	if err != nil {
		return &authenticationV1.LoginResponse{}, err
	}

	// 校验刷新令牌
	if !s.userToken.IsExistRefreshToken(ctx, req.GetOperatorId(), req.GetRefreshToken()) {
		return nil, authenticationV1.ErrorIncorrectRefreshToken("invalid refresh token")
	}

	if err = s.userToken.RemoveRefreshToken(ctx, req.GetOperatorId(), req.GetRefreshToken()); err != nil {
		s.log.Errorf("remove refresh token failed [%s]", err.Error())
	}

	// 生成令牌
	accessToken, refreshToken, err := s.userToken.GenerateToken(ctx, user)
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

func (s *AuthenticationService) GetMe(ctx context.Context, _ *emptypb.Empty) (*userV1.User, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.Get(ctx, operator.UserId)
	if err != nil {
		s.log.Errorf("查询用户失败[%s]", err.Error())
		return nil, authenticationV1.ErrorNotFound("user not found")
	}

	role, err := s.roleRepo.Get(ctx, user.GetRoleId())
	if err == nil && role != nil {
		user.Roles = append(user.Roles, role.GetCode())
	}

	return user, err
}

// RefreshToken 刷新令牌
func (s *AuthenticationService) RefreshToken(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	// 校验授权类型
	if req.GetGrantType() != authenticationV1.GrantType_refresh_token.String() {
		return nil, authenticationV1.ErrorInvalidGrantType("invalid grant type")
	}

	return s.doGrantTypeRefreshToken(ctx, req)
}
