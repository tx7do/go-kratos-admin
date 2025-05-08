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
		schema.Comment("职位表"),
	}
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("职位名称").
			Default("").
			MaxLen(128),

		field.String("code").
			Comment("职位标识").
			Default("").
			MaxLen(128),

		field.Uint32("parent_id").
			Comment("上一层职位ID").
			Default(0).
			Optional(),

		field.Int32("sort_id").
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
		mixin.UpdateBy{},
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
