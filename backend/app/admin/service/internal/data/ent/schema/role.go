package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
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
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("角色标识").
			NotEmpty().
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
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.Remark{},
		mixin.SortOrder{},
		mixin.Tree[Role]{},
		mixin.TenantID{},
	}
}

// Indexes of the User.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique().StorageKey("idx_sys_role_name"),
		index.Fields("code").Unique().StorageKey("idx_sys_role_code"),
	}
}
