package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// InternalMessageCategory holds the schema definition for the InternalMessageCategory entity.
type InternalMessageCategory struct {
	ent.Schema
}

func (InternalMessageCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "internal_message_categories",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("站内信消息分类表"),
	}
}

// Fields of the InternalMessageCategory.
func (InternalMessageCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("编码").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("icon_url").
			Comment("图标URL").
			Optional().
			Nillable(),
	}
}

// Mixin of the InternalMessageCategory.
func (InternalMessageCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.IsEnabled{},
		mixin.SortOrder{},
		mixin.Remark{},
		mixin.TenantID{},
		mixin.Tree[InternalMessageCategory]{},
	}
}

// Indexes of the InternalMessageCategory.
func (InternalMessageCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("idx_internal_message_category_code"),
		index.Fields("name").StorageKey("idx_internal_message_category_name"),
	}
}
