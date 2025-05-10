package auth

import "github.com/go-kratos/kratos/v2/errors"

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrWrongContext          = errors.Unauthorized(reason, "wrong context for middleware")
	ErrMissingJwtToken       = errors.Unauthorized(reason, "no jwt token in context")
	ErrExtractUserInfoFailed = errors.Unauthorized(reason, "extract user info failed")
	ErrExtractSubjectFailed  = errors.Unauthorized(reason, "extract subject failed")
	ErrAccessTokenExpired    = errors.Unauthorized(reason, "access token expired")
	ErrInvalidRequest        = errors.Unauthorized(reason, "invalid request")
)
