package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	entCrud "github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/go-utils/password"

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
	l := log.NewHelper(log.With(ctx.Logger, "module", "data/admin-service"))

	d := &Data{
		log: l,

		db:  db,
		rdb: rdb,
	}

	return d, func() {
		l.Info("closing the data resources")

		if d.db != nil {
			if err := d.db.Close(); err != nil {
				l.Error(err)
			}
		}

		if d.rdb != nil {
			if err := d.rdb.Close(); err != nil {
				l.Error(err)
			}
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(ctx *bootstrap.Context) *redis.Client {
	l := log.NewHelper(log.With(ctx.Logger, "module", "redis/data/admin-service"))
	return redisClient.NewClient(ctx.Config.Data, l)
}

// NewAuthenticator 创建认证器
func NewAuthenticator(ctx *bootstrap.Context) authnEngine.Authenticator {
	if ctx.Config.Authn == nil {
		return nil
	}

	switch ctx.Config.GetAuthn().GetType() {
	default:
		return nil

	case "jwt":
		authenticator, err := jwt.NewAuthenticator(
			jwt.WithKey([]byte(ctx.Config.Authn.GetJwt().GetKey())),
			jwt.WithSigningMethod(ctx.Config.Authn.GetJwt().GetMethod()),
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
	return NewUserTokenCacheRepo(
		ctx,
		rdb,
		authenticator,
		ctx.Config.GetServer().GetRest().GetMiddleware().GetAuth().GetAccessTokenKeyPrefix(),
		ctx.Config.GetServer().GetRest().GetMiddleware().GetAuth().GetRefreshTokenKeyPrefix(),
		ctx.Config.GetServer().GetRest().GetMiddleware().GetAuth().GetAccessTokenExpires().AsDuration(),
		ctx.Config.GetServer().GetRest().GetMiddleware().GetAuth().GetRefreshTokenExpires().AsDuration(),
	)
}

func NewMinIoClient(ctx *bootstrap.Context) *oss.MinIOClient {
	return oss.NewMinIoClient(ctx.Config, ctx.Logger)
}

func NewPasswordCrypto() password.Crypto {
	crypto, err := password.CreateCrypto("bcrypt")
	if err != nil {
		panic(err)
	}
	return crypto
}
