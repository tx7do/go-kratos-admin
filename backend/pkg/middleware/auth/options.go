package auth

import (
	"context"
)

type IsExistAccessToken func(ctx context.Context, userId uint32) bool

type options struct {
	isExistAccessToken IsExistAccessToken
	setOperatorId      bool
	setTenantId        bool
}

type Option func(*options)

func WithIsExistAccessTokenFunc(fc IsExistAccessToken) Option {
	return func(opts *options) {
		opts.isExistAccessToken = fc
	}
}

func WithEnableSetOperatorId(enable bool) Option {
	return func(opts *options) {
		opts.setOperatorId = enable
	}
}

func WithEnableTenantId(enable bool) Option {
	return func(opts *options) {
		opts.setTenantId = enable
	}
}
