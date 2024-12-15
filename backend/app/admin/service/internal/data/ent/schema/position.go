package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

func (Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "positions",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("角色名称").
			Default("").
			MaxLen(128),

		field.String("code").
			Comment("角色标识").
			Default("").
			MaxLen(128),

		field.Uint32("parent_id").
			Comment("上一层角色ID").
			Default(0).
			Optional(),

		field.Int32("order_no").
			Comment("排序ID").
			Default(0),
	}
}

// Mixin of the Position.
func (Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.SwitchStatus{},
		mixin.CreateBy{},
		mixin.Remark{},
	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Position.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}
