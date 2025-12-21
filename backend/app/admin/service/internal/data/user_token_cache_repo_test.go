package data

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
)

func TestUserTokenCache(t *testing.T) {
	ctx := context.Background()
	l := log.DefaultLogger

	var cfg = &conf.Bootstrap{
		Authn: &conf.Authentication{
			Type: "jwt",
			Jwt: &conf.Authentication_Jwt{
				Method: "HS256",
				Key:    "some_api_key",
			},
		},
		Data: &conf.Data{
			Redis: &conf.Data_Redis{
				Addr:     "127.0.0.1:6379",
				Password: "*Abcd123456",
			},
		},
	}
	bctx := bootstrap.NewContextWithParam(ctx, &conf.AppInfo{}, cfg, l)

	authenticator := NewAuthenticator(bctx)
	assert.NotNil(t, authenticator)

	rdb := NewRedisClient(bctx)
	assert.NotNil(t, rdb)

	repo := NewUserTokenRepo(bctx, rdb, authenticator)
	assert.NotNil(t, repo)

	var userId uint32 = 0
	var err error

	err = repo.setAccessTokenToRedis(ctx, userId, "access_token", 0)
	assert.Nil(t, err)
	exist := repo.IsExistAccessToken(ctx, userId, "access_token")
	assert.True(t, exist)
	accessTokens := repo.GetAccessToken(ctx, userId)
	assert.NotEmpty(t, accessTokens)
	assert.Equal(t, "access_token", accessTokens[0])
	err = repo.RemoveAccessToken(ctx, userId, "access_token")
	assert.Nil(t, err)

	err = repo.setRefreshTokenToRedis(ctx, userId, "refresh_token", 0)
	assert.Nil(t, err)
	exist = repo.IsExistRefreshToken(ctx, userId, "refresh_token")
	assert.True(t, exist)
	refreshTokens := repo.GetRefreshToken(ctx, userId)
	assert.NotEmpty(t, refreshTokens)
	assert.Equal(t, "refresh_token", refreshTokens[0])
	err = repo.RemoveRefreshToken(ctx, userId, "refresh_token")
	assert.Nil(t, err)
}
