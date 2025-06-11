package data

import (
	"context"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/tx7do/go-utils/trans"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestUserTokenCache(t *testing.T) {
	ctx := context.Background()

	l := log.DefaultLogger
	//l := log.NewHelper(log.With(log.DefaultLogger, "module", "test"))

	var cfg = &conf.Bootstrap{
		Server: &conf.Server{
			Rest: &conf.Server_REST{
				Middleware: &conf.Middleware{
					Auth: &conf.Middleware_Auth{
						Method:                "HS256",
						Key:                   "some_api_key",
						AccessTokenKeyPrefix:  trans.Ptr("aat_"),
						RefreshTokenKeyPrefix: trans.Ptr("art_"),
						AccessTokenExpires:    durationpb.New(0 * time.Second),
						RefreshTokenExpires:   durationpb.New(0 * time.Second),
					},
				},
			},
		},
		Data: &conf.Data{
			Redis: &conf.Data_Redis{
				Addr:     "redis:6379",
				Password: "*Abcd123456",
			},
		},
	}

	authenticator := NewAuthenticator(cfg)
	assert.NotNil(t, authenticator)

	rdb := NewRedisClient(cfg, l)
	assert.NotNil(t, rdb)

	repo := NewUserTokenRepo(l, rdb, authenticator, cfg)
	assert.NotNil(t, repo)

	var userId uint32 = 0
	var err error

	err = repo.setAccessTokenToRedis(ctx, userId, "access_token", 0)
	assert.Nil(t, err)
	exist := repo.IsExistAccessToken(ctx, userId, "access_token")
	assert.True(t, exist)
	err = repo.RemoveAccessToken(ctx, userId, "access_token")
	assert.Nil(t, err)

	err = repo.setRefreshTokenToRedis(ctx, userId, "refresh_token", 0)
	assert.Nil(t, err)
	exist = repo.IsExistRefreshToken(ctx, userId, "refresh_token")
	assert.True(t, exist)
	err = repo.RemoveRefreshToken(ctx, userId, "refresh_token")
	assert.Nil(t, err)
}
