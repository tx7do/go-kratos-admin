package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "organizations",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("名字").
			Default("").
			MaxLen(128).
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("上一层部门ID").
			Optional().
			Nillable(),

		field.Int32("order_no").
			Comment("排序ID").
			Default(0).
			Annotations(
				entproto.Field(6),
			).
			Optional().
			Nillable(),
	}
}

// Mixin of the Organization.
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.SwitchStatus{},
		mixin.CreateBy{},
		mixin.Remark{},
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Organization.Type).
			From("parent").Unique().Field("parent_id"),
	}
}
