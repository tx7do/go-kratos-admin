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
		field.String("description").
			Comment("描述").
			Optional().
			Nillable(),

		field.String("module").
			Comment("所属业务模块").
			Optional().
			Nillable(),

		field.String("module_description").
			Comment("业务模块描述").
			Optional().
			Nillable(),

		field.String("operation").
			Comment("接口操作名").
			Optional().
			Nillable(),

		field.String("path").
			Comment("接口路径").
			Optional().
			Nillable(),

		field.String("method").
			Comment("请求方法").
			Optional().
			Nillable(),

		field.Enum("scope").
			Comment("作用域").
			NamedValues(
				"API_SCOPE_ADMIN", "ADMIN",
				"API_SCOPE_APP", "APP",
			).
			Default("ADMIN").
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
