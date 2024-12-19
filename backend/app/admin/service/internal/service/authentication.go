package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/cache"
	"kratos-admin/pkg/middleware/auth"
)

type AuthenticationService struct {
	adminV1.AuthenticationServiceHTTPServer

	uc   *data.UserRepo
	utuc *cache.UserToken

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, uc *data.UserRepo, utuc *cache.UserToken) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authn/service/admin-service"))
	return &AuthenticationService{
		log:  l,
		uc:   uc,
		utuc: utuc,
	}
}

// Login 登录
func (s *AuthenticationService) Login(ctx context.Context, req *adminV1.LoginRequest) (*adminV1.LoginResponse, error) {
	if req.GetGrantType() != adminV1.GrantType_password.String() {
		return nil, adminV1.ErrorInvalidGrantType("invalid grant type")
	}

	var err error
	if _, err = s.uc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return nil, err
	}

	// 获取用户信息
	var user *userV1.User
	user, err = s.uc.GetUserByUserName(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	// 验证权限
	if user.GetAuthority() != userV1.UserAuthority_SYS_ADMIN && user.GetAuthority() != userV1.UserAuthority_SYS_MANAGER {
		return &adminV1.LoginResponse{}, adminV1.ErrorAccessForbidden("权限不够")
	}

	// 生成访问令牌
	accessToken, refreshToken, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return &adminV1.LoginResponse{
		TokenType:    adminV1.TokenType_bearer.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Logout 登出
func (s *AuthenticationService) Logout(ctx context.Context, req *adminV1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.utuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthenticationService) GetMe(ctx context.Context, req *adminV1.GetMeRequest) (*userV1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.Id = authInfo.UserId
	ret, err := s.uc.GetUser(ctx, req.GetId())
	return ret, err
}

// RefreshToken 刷新令牌
func (s *AuthenticationService) RefreshToken(ctx context.Context, req *adminV1.RefreshTokenRequest) (*adminV1.LoginResponse, error) {
	// 校验授权类型
	if req.GetGrantType() != adminV1.GrantType_refresh_token.String() {
		return nil, adminV1.ErrorInvalidGrantType("invalid grant type")
	}

	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	// 校验刷新令牌
	refreshToken := s.utuc.GetRefreshToken(ctx, authInfo.UserId)
	if refreshToken != req.GetRefreshToken() {
		return nil, adminV1.ErrorIncorrectRefreshToken("invalid refresh token")
	}

	// 生成新的访问令牌
	accessToken, err := s.utuc.GenerateAccessToken(ctx, authInfo.UserId, authInfo.UserName)
	if err != nil {
		return nil, err
	}

	return &adminV1.LoginResponse{
		TokenType:    adminV1.TokenType_bearer.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
