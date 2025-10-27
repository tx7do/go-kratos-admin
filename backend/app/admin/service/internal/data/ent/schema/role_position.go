package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// RolePosition holds the schema definition for the RolePosition entity.
type RolePosition struct {
	ent.Schema
}

func (RolePosition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_role_position",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("角色 - 职位关联表"),
	}
}

// Fields of the RolePosition.
func (RolePosition) Fields() []ent.Field {
	return []ent.Field{

		field.Uint32("role_id").
			Comment("角色ID").
			Nillable(),

		field.Uint32("position_id").
			Comment("职位ID").
			Nillable(),
	}
}

// Mixin of the RolePosition.
func (RolePosition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
	}
}

// Indexes of the RolePosition.
func (RolePosition) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "position_id").Unique().StorageKey("idx_sys_role_position_role_id_position_id"),
		index.Fields("role_id").StorageKey("idx_sys_role_position_role_id"),
	}
}
