package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_user_role",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("用户 - 角色关联表"),
	}
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{

		field.Uint32("user_id").
			Comment("用户ID").
			Nillable(),

		field.Uint32("role_id").
			Comment("角色ID").
			Nillable(),
	}
}

// Mixin of the UserRole.
func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
	}
}

// Indexes of the UserRole.
func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		// 避免用户重复分配同一角色
		index.Fields("user_id", "role_id").Unique().StorageKey("idx_sys_user_role_user_id_role_id"),
		index.Fields("user_id").StorageKey("idx_sys_user_role_user_id"),
		index.Fields("role_id").StorageKey("idx_sys_user_role_role_id"),
	}
}
