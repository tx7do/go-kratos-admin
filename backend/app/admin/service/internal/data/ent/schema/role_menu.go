package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// RoleMenu holds the schema definition for the RoleMenu entity.
type RoleMenu struct {
	ent.Schema
}

func (RoleMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_role_menu",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("角色 - 菜单关联表"),
	}
}

// Fields of the RoleMenu.
func (RoleMenu) Fields() []ent.Field {
	return []ent.Field{

		field.Uint32("role_id").
			Comment("角色ID").
			Nillable(),

		field.Uint32("menu_id").
			Comment("菜单ID").
			Nillable(),
	}
}

// Mixin of the RoleMenu.
func (RoleMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
	}
}

// Indexes of the RoleMenu.
func (RoleMenu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "menu_id").Unique().StorageKey("idx_sys_role_menu_role_id_menu_id"),
		index.Fields("role_id").StorageKey("idx_sys_role_menu_role_id"),
	}
}
