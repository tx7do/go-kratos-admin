package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	entCrud "github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/go-utils/password"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	"github.com/tx7do/kratos-bootstrap/bootstrap"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"

	"go-wind-admin/app/admin/service/internal/data/ent"

	"go-wind-admin/pkg/oss"
)

// Data .
type Data struct {
	log *log.Helper

	rdb *redis.Client
	db  *entCrud.EntClient[*ent.Client]

	authenticator authnEngine.Authenticator
	authorizer    *Authorizer
}

// NewData .
func NewData(
	ctx *bootstrap.Context,
	db *entCrud.EntClient[*ent.Client],
	rdb *redis.Client,
) (*Data, func(), error) {
	d := &Data{
		log: ctx.NewLoggerHelper("data/admin-service"),

		db:  db,
		rdb: rdb,
	}

	return d, func() {
		d.log.Info("closing the data resources")

		if d.db != nil {
			if err := d.db.Close(); err != nil {
				d.log.Error(err)
			}
		}

		if d.rdb != nil {
			if err := d.rdb.Close(); err != nil {
				d.log.Error(err)
			}
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(ctx *bootstrap.Context) *redis.Client {
	cfg := ctx.GetConfig()
	if cfg == nil {
		return nil
	}
	return redisClient.NewClient(cfg.Data, ctx.NewLoggerHelper("redis/data/admin-service"))
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
