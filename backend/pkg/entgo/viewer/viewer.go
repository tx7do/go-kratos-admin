package viewer

import (
	"context"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	// Admin 是否是管理员
	Admin() bool

	// SystemAdmin 系统管理员
	SystemAdmin() bool

	// TenantAdmin 租户管理员
	TenantAdmin() bool

	// Tenant 返回租户ID
	Tenant() (uint32, bool)
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	TenantId *uint32              // Tenant ID
	Role     userV1.UserAuthority // Attached roles.
}

func (v UserViewer) Admin() bool {
	return v.SystemAdmin() || v.TenantAdmin()
}

func (v UserViewer) SystemAdmin() bool {
	return v.Role == userV1.UserAuthority_SYS_ADMIN
}

func (v UserViewer) TenantAdmin() bool {
	return v.Role == userV1.UserAuthority_TENANT_ADMIN
}

func (v UserViewer) Tenant() (uint32, bool) {
	if v.TenantId != nil {
		return *v.TenantId, true
	}
	return 0, false
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
