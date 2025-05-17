package auth

import (
	"context"
)

type IsExistAccessToken func(ctx context.Context, userId uint32) bool

type options struct {
	isExistAccessToken IsExistAccessToken
	injectOperatorId   bool
	injectTenantId     bool
	enableAuthz        bool
	injectEnt          bool
	injectMetadata     bool
}

type Option func(*options)

func WithIsExistAccessTokenFunc(fc IsExistAccessToken) Option {
	return func(opts *options) {
		opts.isExistAccessToken = fc
	}
}

func WithInjectOperatorId(enable bool) Option {
	return func(opts *options) {
		opts.injectOperatorId = enable
	}
}

func WithInjectTenantId(enable bool) Option {
	return func(opts *options) {
		opts.injectTenantId = enable
	}
}

func WithInjectEnt(enable bool) Option {
	return func(opts *options) {
		opts.injectEnt = enable
	}
}

func WithInjectMetadata(enable bool) Option {
	return func(opts *options) {
		opts.injectMetadata = enable
	}
}

func WithEnableAuthority(enable bool) Option {
	return func(opts *options) {
		opts.enableAuthz = enable
	}
}
