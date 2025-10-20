package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"

	appmixin "kratos-admin/pkg/entgo/mixin"
)

// NotificationMessageRecipient holds the schema definition for the NotificationMessageRecipient entity.
type NotificationMessageRecipient struct {
	ent.Schema
}

func (NotificationMessageRecipient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "notification_message_recipients",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("站内信通知消息接收者表"),
	}
}

// Fields of the NotificationMessageRecipient.
func (NotificationMessageRecipient) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("message_id").
			Comment("群发消息ID").
			Optional().
			Nillable(),

		field.Uint32("recipient_id").
			Comment("接收者用户ID").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("消息状态").
			NamedValues(
				"Unknown", "UNKNOWN",
				"Received", "RECEIVED",
				"Read", "READ",
				"Archived", "ARCHIVED",
				"Deleted", "DELETED",
			).
			Optional().
			Nillable(),
	}
}

// Mixin of the NotificationMessageRecipient.
func (NotificationMessageRecipient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		appmixin.TenantID{},
	}
}
