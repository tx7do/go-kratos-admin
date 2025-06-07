package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// ApiResource holds the schema definition for the ApiResource entity.
type ApiResource struct {
	ent.Schema
}

func (ApiResource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_api_resources",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("API资源表"),
	}
}

// Fields of the ApiResource.
func (ApiResource) Fields() []ent.Field {
	return []ent.Field{
		field.String("operation").
			Comment("操作路径").
			Unique().
			Optional().
			Nillable(),

		field.String("description").
			Comment("描述").
			Optional().
			Nillable(),

		field.String("module").
			Comment("所属业务模块").
			Optional().
			Nillable(),
	}
}

// Mixin of the ApiResource.
func (ApiResource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
	}
}
