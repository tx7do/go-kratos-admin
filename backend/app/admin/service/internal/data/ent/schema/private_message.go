package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// PrivateMessage holds the schema definition for the PrivateMessage entity.
type PrivateMessage struct {
	ent.Schema
}

func (PrivateMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "private_messages",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("站内信私信消息表"),
	}
}

// Fields of the PrivateMessage.
func (PrivateMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").
			Comment("主题").
			Optional().
			Nillable(),

		field.String("content").
			Comment("内容").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("消息状态").
			NamedValues(
				"Unknown", "UNKNOWN",
				"Draft", "DRAFT",
				"Sent", "SENT",
				"Received", "RECEIVED",
				"Read", "READ",
				"Archived", "ARCHIVED",
				"Deleted", "DELETED",
			).
			Optional().
			Nillable(),

		field.Uint32("sender_id").
			Comment("发送者用户ID").
			Optional().
			Nillable(),

		field.Uint32("receiver_id").
			Comment("接收者用户ID").
			Optional().
			Nillable(),
	}
}

// Mixin of the PrivateMessage.
func (PrivateMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.TenantID{},
	}
}
