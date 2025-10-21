package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("租户编号").
			NotEmpty().
			MaxLen(64).
			Optional().
			Nillable(),

		field.Int32("member_count").
			Comment("成员数").
			Default(0).
			Optional().
			Nillable(),

		field.Time("subscription_at").
			Comment("订阅时间").
			Default(time.Now).
			Optional().
			Nillable(),

		field.Time("unsubscribe_at").
			Comment("取消订阅时间").
			Default(time.Now).
			Optional().
			Nillable(),
	}
}

// Mixin of the Tenant.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.SwitchStatus{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}
