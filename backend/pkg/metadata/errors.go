package metadata

import "github.com/go-kratos/kratos/v2/errors"

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingMetadata = errors.Unauthorized(reason, "missing metadata in context")
)
