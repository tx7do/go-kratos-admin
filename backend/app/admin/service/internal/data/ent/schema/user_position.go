package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// UserPosition holds the schema definition for the UserPosition entity.
type UserPosition struct {
	ent.Schema
}

func (UserPosition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_user_position",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("用户 - 职位关联表"),
	}
}

// Fields of the UserPosition.
func (UserPosition) Fields() []ent.Field {
	return []ent.Field{

		field.Uint32("user_id").
			Comment("用户ID").
			Nillable(),

		field.Uint32("position_id").
			Comment("职位ID").
			Nillable(),
	}
}

// Mixin of the UserPosition.
func (UserPosition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
	}
}

// Indexes of the UserPosition.
func (UserPosition) Indexes() []ent.Index {
	return []ent.Index{
		// 避免用户重复分配同一角色
		index.Fields("user_id", "position_id").Unique(),
		index.Fields("user_id"),
		index.Fields("position_id"),
	}
}
