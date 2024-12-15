package auth

import "github.com/go-kratos/kratos/v2/errors"

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrWrongContext         = errors.Unauthorized(reason, "wrong context for middleware")
	ErrMissingJwtToken      = errors.Unauthorized(reason, "no jwt token in context")
	ErrExtractSubjectFailed = errors.Unauthorized(reason, "extract subject failed")
	ErrAccessTokenNotExist  = errors.Unauthorized(reason, "access token do not exists")
	ErrAccessTokenExpired   = errors.Unauthorized(reason, "access token expired")
)
