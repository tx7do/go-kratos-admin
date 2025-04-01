package data

import (
	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/noop"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"github.com/tx7do/go-utils/entgo"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"

	"kratos-admin/app/admin/service/internal/data/ent"

	"kratos-admin/pkg/oss"
)

// Data .
type Data struct {
	log *log.Helper

	rdb *redis.Client
	db  *entgo.EntClient[*ent.Client]

	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Engine
}

// NewData .
func NewData(
	logger log.Logger,
	entClient *entgo.EntClient[*ent.Client],
	redisClient *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		log: l,

		db:  entClient,
		rdb: redisClient,

		authenticator: authenticator,
		authorizer:    authorizer,
	}

	return d, func() {
		l.Info("closing the data resources")

		_ = d.db.Close()

		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	//l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))
	return redisClient.NewClient(cfg.Data)
}

// NewAuthenticator 创建认证器
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Rest.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Rest.Middleware.Auth.Method),
	)
	return authenticator
}

// NewAuthorizer 创建权鉴器
func NewAuthorizer() authzEngine.Engine {
	return noop.State{}
}

func NewUserTokenRepo(data *Data, authenticator authnEngine.Authenticator, logger log.Logger) *UserToken {
	const (
		userAccessTokenKeyPrefix  = "uat_"
		userRefreshTokenKeyPrefix = "urt_"
	)
	return NewUserToken(data.rdb, authenticator, logger, userAccessTokenKeyPrefix, userRefreshTokenKeyPrefix)
}

func NewMinIoClient(cfg *conf.Bootstrap, logger log.Logger) *oss.MinIOClient {
	return oss.NewMinIoClient(cfg, logger)
}
