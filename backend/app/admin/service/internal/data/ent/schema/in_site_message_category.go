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

// InSiteMessageCategory holds the schema definition for the InSiteMessageCategory entity.
type InSiteMessageCategory struct {
	ent.Schema
}

func (InSiteMessageCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "in_site_message_categories",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the InSiteMessageCategory.
func (InSiteMessageCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("名称").
			Optional().
			Nillable(),

		field.String("code").
			Comment("编码").
			Optional().
			Nillable(),

		field.Int32("sort_id").
			Comment("排序编号").
			Optional().
			Nillable(),

		field.Bool("enable").
			Comment("是否启用").
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("父节点ID").
			Optional().
			Nillable(),
	}
}

// Mixin of the InSiteMessageCategory.
func (InSiteMessageCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
	}
}

// Edges of the InSiteMessageCategory.
func (InSiteMessageCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", InSiteMessageCategory.Type).Annotations(entproto.Field(10)).
			From("parent").Unique().Field("parent_id").Annotations(entproto.Field(11)),
	}
}
