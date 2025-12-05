package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// RoleOrg holds the schema definition for the RoleOrg entity.
type RoleOrg struct {
	ent.Schema
}

func (RoleOrg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_role_org",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("角色 - 组织关联表"),
	}
}

// Fields of the RoleOrg.
func (RoleOrg) Fields() []ent.Field {
	return []ent.Field{

		field.Uint32("role_id").
			Comment("角色ID").
			Nillable(),

		field.Uint32("org_id").
			Comment("组织ID").
			Nillable(),
	}
}

// Mixin of the RoleOrg.
func (RoleOrg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
	}
}

// Indexes of the RoleOrg.
func (RoleOrg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "org_id").Unique().StorageKey("idx_sys_role_org_role_id_org_id"),
		index.Fields("role_id").StorageKey("idx_sys_role_org_role_id"),
	}
}
