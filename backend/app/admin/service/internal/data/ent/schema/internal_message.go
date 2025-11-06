package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// InternalMessage holds the schema definition for the InternalMessage entity.
type InternalMessage struct {
	ent.Schema
}

func (InternalMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "internal_messages",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("站内信消息表"),
	}
}

// Fields of the InternalMessage.
func (InternalMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("消息标题").
			Optional().
			Nillable(),

		field.String("content").
			Comment("消息内容").
			Optional().
			Nillable(),

		field.Uint32("sender_id").
			Comment("发送者用户ID").
			Optional().
			Nillable(),

		field.Uint32("category_id").
			Comment("分类ID").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("消息状态").
			NamedValues(
				"Draft", "DRAFT",
				"Published", "PUBLISHED",
				"Scheduled", "SCHEDULED",
				"Revoked", "REVOKED",
				"Archived", "ARCHIVED",
				"Deleted", "DELETED",
			).
			Default("DRAFT").
			Optional().
			Nillable(),

		field.Enum("type").
			Comment("消息类型").
			NamedValues(
				"Notification", "NOTIFICATION",
				"Private", "PRIVATE",
				"Group", "GROUP",
			).
			Default("NOTIFICATION").
			Optional().
			Nillable(),
	}
}

// Mixin of the InternalMessage.
func (InternalMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.TenantID{},
	}
}
