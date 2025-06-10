package ent

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"

	"kratos-admin/pkg/entgo/viewer"
	"kratos-admin/pkg/metadata"
)

// Server .
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			_, tTenantId, tAuthority := metadata.FromOperatorMetadata(ctx)

			if tTenantId != nil && tAuthority != nil {
				ctx = viewer.NewContext(ctx, viewer.UserViewer{
					Authority: *tAuthority,
					TenantId:  tTenantId,
				})
			}

			return handler(ctx, req)
		}
	}
}
