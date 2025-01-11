package auth

import (
	"context"
	"kratos-admin/app/admin/service/internal/data"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"
)

var action = authzEngine.Action("ANY")

// Server 衔接认证和权鉴
func Server(userToken *data.UserToken) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			// 校验访问令牌是否存在
			if err := verifyAccessToken(ctx, userToken, authnClaims); err != nil {
				return nil, err
			}

			sub, _ := authnClaims.GetSubject()
			path := authzEngine.Resource(tr.Operation())

			authzClaims := authzEngine.AuthClaims{
				Subject:  (*authzEngine.Subject)(&sub),
				Action:   &action,
				Resource: &path,
			}

			ctx = authz.NewContext(ctx, &authzClaims)

			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) (*data.UserTokenPayload, error) {
	claims, ok := authnEngine.AuthClaimsFromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}

	return data.NewUserTokenPayloadWithClaims(claims)
}

// verifyAccessToken 校验访问令牌
func verifyAccessToken(ctx context.Context, userToken *data.UserToken, authnClaims *authnEngine.AuthClaims) error {
	ut, err := data.NewUserTokenPayloadWithClaims(authnClaims)
	if err != nil {
		return ErrExtractUserInfoFailed
	}

	// 校验访问令牌是否存在
	if !userToken.IsExistAccessToken(ctx, ut.UserId) {
		return ErrAccessTokenExpired
	}

	return nil
}
