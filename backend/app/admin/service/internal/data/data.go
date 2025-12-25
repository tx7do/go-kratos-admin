package data

import (
	"github.com/redis/go-redis/v9"
	"github.com/tx7do/go-utils/password"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	"github.com/tx7do/kratos-bootstrap/bootstrap"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"

	"go-wind-admin/pkg/oss"
)

// NewRedisClient 创建Redis客户端
func NewRedisClient(ctx *bootstrap.Context) (*redis.Client, func(), error) {
	cfg := ctx.GetConfig()
	if cfg == nil {
		return nil, func() {}, nil
	}

	l := ctx.NewLoggerHelper("redis/data/admin-service")

	cli := redisClient.NewClient(cfg.Data, l)

	return cli, func() {
		if err := cli.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewAuthenticator 创建认证器
func NewAuthenticator(ctx *bootstrap.Context) authnEngine.Authenticator {
	cfg := ctx.GetConfig()
	if cfg == nil || cfg.Authn == nil {
		return nil
	}

	switch cfg.GetAuthn().GetType() {
	default:
		return nil

	case "jwt":
		authenticator, err := jwt.NewAuthenticator(
			jwt.WithKey([]byte(cfg.Authn.GetJwt().GetKey())),
			jwt.WithSigningMethod(cfg.Authn.GetJwt().GetMethod()),
		)
		if err != nil {
			return nil
		}
		return authenticator

	case "oidc":
		return nil

	case "preshared_key":
		return nil
	}
}

func NewUserTokenRepo(ctx *bootstrap.Context, rdb *redis.Client, authenticator authnEngine.Authenticator) *UserTokenCacheRepo {
	cfg := ctx.GetConfig()
	if cfg == nil || cfg.Server == nil || cfg.Server.Rest == nil {
		return nil
	}

	return NewUserTokenCacheRepo(
		ctx,
		rdb,
		authenticator,
		cfg.GetServer().GetRest().GetMiddleware().GetAuth().GetAccessTokenKeyPrefix(),
		cfg.GetServer().GetRest().GetMiddleware().GetAuth().GetRefreshTokenKeyPrefix(),
		cfg.GetServer().GetRest().GetMiddleware().GetAuth().GetAccessTokenExpires().AsDuration(),
		cfg.GetServer().GetRest().GetMiddleware().GetAuth().GetRefreshTokenExpires().AsDuration(),
	)
}

func NewMinIoClient(ctx *bootstrap.Context) *oss.MinIOClient {
	return oss.NewMinIoClient(ctx.GetConfig(), ctx.GetLogger())
}

func NewPasswordCrypto() password.Crypto {
	crypto, err := password.CreateCrypto("bcrypt")
	if err != nil {
		panic(err)
	}
	return crypto
}
