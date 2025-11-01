package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// NotificationMessageCategory holds the schema definition for the NotificationMessageCategory entity.
type NotificationMessageCategory struct {
	ent.Schema
}

func (NotificationMessageCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "notification_message_categories",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("站内信通知消息分类表"),
	}
}

// Fields of the NotificationMessageCategory.
func (NotificationMessageCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("名称").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("编码").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.Int32("sort_order").
			Comment("排序顺序，值越小越靠前").
			Default(0).
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

// Mixin of the NotificationMessageCategory.
func (NotificationMessageCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.Remark{},
		mixin.TenantID{},
	}
}

// Indexes of the NotificationMessageCategory.
func (NotificationMessageCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("idx_notification_message_category_code"),
		index.Fields("name").StorageKey("idx_notification_message_category_name"),
	}
}

// Edges of the NotificationMessageCategory.
func (NotificationMessageCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", NotificationMessageCategory.Type).Annotations(entproto.Field(10)).
			From("parent").Unique().Field("parent_id").Annotations(entproto.Field(11)),
	}
}
