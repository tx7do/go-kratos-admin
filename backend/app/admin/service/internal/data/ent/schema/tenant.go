package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/tx7do/go-utils/entgo/mixin"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_tenants",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("租户表"),
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("租户名称").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("租户编号").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("logo_url").
			Comment("租户logo地址").
			Optional().
			Nillable(),

		field.String("industry").
			Comment("所属行业").
			Optional().
			Nillable(),

		field.Uint32("admin_user_id").
			Comment("管理员用户ID").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("租户状态").
			NamedValues(
				"TENANT_STATUS_ON", "ON",
				"TENANT_STATUS_OFF", "OFF",
				"TENANT_STATUS_EXPIRED", "EXPIRED",
				"TENANT_STATUS_FREEZE", "FREEZE",
			).
			Default("ON").
			Optional().
			Nillable(),

		field.Enum("type").
			Comment("租户类型").
			NamedValues(
				"TENANT_TYPE_TRIAL", "TRIAL",
				"TENANT_TYPE_PAID", "PAID",
				"TENANT_TYPE_INTERNAL", "INTERNAL",
				"TENANT_TYPE_PARTNER", "PARTNER",
				"TENANT_TYPE_CUSTOM", "CUSTOM",
			).
			Optional().
			Nillable(),

		field.Enum("audit_status").
			Comment("审核状态").
			NamedValues(
				"TENANT_AUDIT_STATUS_PENDING", "PENDING",
				"TENANT_AUDIT_STATUS_APPROVED", "APPROVED",
				"TENANT_AUDIT_STATUS_REJECTED", "REJECTED",
			).
			Optional().
			Nillable(),

		field.Time("subscription_at").
			Comment("订阅时间").
			Optional().
			Nillable(),

		field.Time("unsubscribe_at").
			Comment("取消订阅时间").
			Optional().
			Nillable(),

		field.String("subscription_plan").
			Comment("订阅套餐").
			Optional().
			Nillable(),

		field.Time("expired_at").
			Comment("租户有效期").
			Optional().
			Nillable(),

		field.Time("last_login_time").
			Comment("最后一次登录的时间").
			Optional().
			Nillable(),

		field.String("last_login_ip").
			Comment("最后一次登录的IP").
			Optional().
			Nillable(),
	}
}

// Mixin of the Tenant.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
	}
}

// Indexes of the Tenant.
func (Tenant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique().StorageKey("idx_sys_tenant_name"),
		index.Fields("code").Unique().StorageKey("idx_sys_tenant_code"),
		index.Fields("status", "audit_status").StorageKey("idx_sys_tenant_status_audit_status"),
		index.Fields("expired_at").StorageKey("idx_sys_tenant_expired_at"),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}
