package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
	appmixin "kratos-admin/pkg/entgo/mixin"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_roles",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("角色表"),
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("角色名称").
			Unique().
			Optional().
			Nillable().
			MaxLen(128),

		field.String("code").
			Comment("角色标识").
			Default("").
			Optional().
			Nillable().
			MaxLen(128),

		field.Uint32("parent_id").
			Comment("上一层角色ID").
			Nillable().
			Optional(),

		field.Int32("sort_id").
			Comment("排序ID").
			Optional().
			Nillable().
			Default(0),

		field.JSON("menus", []uint32{}).
			Comment("分配的菜单列表").
			Optional(),
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.SwitchStatus{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Role.Type).
			From("parent").Unique().Field("parent_id"),
	}
}
