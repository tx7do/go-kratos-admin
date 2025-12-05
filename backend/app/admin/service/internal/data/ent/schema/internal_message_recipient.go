package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// InternalMessageRecipient holds the schema definition for the InternalMessageRecipient entity.
type InternalMessageRecipient struct {
	ent.Schema
}

func (InternalMessageRecipient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "internal_message_recipients",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("站内信消息用户接收信息表"),
	}
}

// Fields of the InternalMessageRecipient.
func (InternalMessageRecipient) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("message_id").
			Comment("站内信内容ID").
			Optional().
			Nillable(),

		field.Uint32("recipient_user_id").
			Comment("接收者用户ID").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("消息状态").
			NamedValues(
				"Sent", "SENT",
				"Received", "RECEIVED",
				"Read", "READ",
				"Revoked", "REVOKED",
				"Deleted", "DELETED",
			).
			Optional().
			Nillable(),

		field.Time("received_at").
			Comment("消息到达用户收件箱的时间").
			Optional().
			Nillable(),

		field.Time("read_at").
			Comment("用户阅读消息的时间").
			Optional().
			Nillable(),
	}
}

// Mixin of the InternalMessageRecipient.
func (InternalMessageRecipient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.TenantID{},
	}
}
