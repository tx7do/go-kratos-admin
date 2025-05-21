package logging

import (
	"context"

	authnEngine "github.com/tx7do/kratos-authn/engine"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type WriteOperationLogFunc func(ctx context.Context, data *adminV1.AdminOperationLog) error
type WriteLoginLogFunc func(ctx context.Context, data *adminV1.AdminLoginLog) error

type options struct {
	authenticator authnEngine.Authenticator

	writeOperationLogFunc WriteOperationLogFunc
	writeLoginLogFunc     WriteLoginLogFunc
}

type Option func(*options)

func WithAuthenticator(authenticator authnEngine.Authenticator) Option {
	return func(opts *options) {
		opts.authenticator = authenticator
	}
}

func WithWriteOperationLogFunc(fnc WriteOperationLogFunc) Option {
	return func(opts *options) {
		opts.writeOperationLogFunc = fnc
	}
}

func WithWriteLoginLogFunc(fnc WriteLoginLogFunc) Option {
	return func(opts *options) {
		opts.writeLoginLogFunc = fnc
	}
}
