package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// RoleApi holds the schema definition for the RoleApi entity.
type RoleApi struct {
	ent.Schema
}

func (RoleApi) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_role_api",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("角色 - API关联表"),
	}
}

// Fields of the RoleApi.
func (RoleApi) Fields() []ent.Field {
	return []ent.Field{

		field.Uint32("role_id").
			Comment("角色ID").
			Nillable(),

		field.Uint32("api_id").
			Comment("API ID").
			Nillable(),
	}
}

// Mixin of the RoleApi.
func (RoleApi) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
	}
}

// Indexes of the RoleApi.
func (RoleApi) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "api_id").Unique().StorageKey("idx_sys_role_api_role_id_api_id"),
		index.Fields("role_id").StorageKey("idx_sys_role_api_role_id"),
	}
}
