package rule

import (
	"context"
	"entgo.io/ent/entql"

	"kratos-admin/app/admin/service/internal/data/ent/privacy"

	"kratos-admin/pkg/entgo/viewer"
)

// FilterTenantRule is a query/mutation rule that filters out entities that are not in the tenant.
func FilterTenantRule() privacy.QueryMutationRule {
	type TenantsFilter interface {
		WhereTenantID(p entql.Uint32P)
		Where(p entql.P)
	}

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Skip
		}

		// Skip if the viewer is a system admin
		if view.SystemAdmin() {
			return privacy.Skip
		}

		tid, ok := view.Tenant()
		if !ok {
			return privacy.Denyf("missing tenant information in viewer")
		}

		tf, ok := f.(TenantsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %TenantId", f)
		}

		tf.Where(
			entql.Uint32Or(
				entql.Uint32EQ(tid),
				entql.Uint32EQ(0),
				entql.Uint32Nil(),
			).Field("tenant_id"),
		)

		return privacy.Skip
	})
}
