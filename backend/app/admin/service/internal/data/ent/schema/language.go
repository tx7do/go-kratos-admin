package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

func (Language) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_languages",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("语言表"),
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("language_code").
			Comment("标准语言代码").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("language_name").
			Comment("语言名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("native_name").
			Comment("本地语言名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.Bool("is_default").
			Comment("是否为默认语言").
			Optional().
			Nillable().
			Default(false),
	}
}

// Mixin of the Language.
func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.SortOrder{},
		mixin.IsEnabled{},
	}
}
