package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("角色标识").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("上一层角色ID").
			Nillable().
			Optional(),

		field.Int32("sort_id").
			Comment("排序ID").
			Default(0).
			Optional().
			Nillable(),

		field.JSON("menus", []uint32{}).
			Comment("分配的菜单列表").
			Optional(),

		field.JSON("apis", []uint32{}).
			Comment("分配的API列表").
			Optional(),

		field.Enum("data_scope").
			Comment("数据权限范围").
			NamedValues(
				"All", "ALL",
				"Custom", "CUSTOM",
				"Self", "SELF",
				"Org", "ORG",
				"OrgAndChild", "ORG_AND_CHILD",
				"Dept", "DEPT",
				"DeptAndChild", "DEPT_AND_CHILD",
			).
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("角色状态").
			NamedValues(
				"On", "ON",
				"Off", "OFF",
			).
			Default("ON").
			Optional().
			Nillable(),
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Indexes of the User.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique().StorageKey("idx_sys_role_name"),
		index.Fields("code").Unique().StorageKey("idx_sys_role_code"),
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
