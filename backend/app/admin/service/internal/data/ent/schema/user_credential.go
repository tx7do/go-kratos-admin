package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"

	appmixin "kratos-admin/pkg/entgo/mixin"
)

// UserCredential holds the schema definition for the UserCredential entity.
type UserCredential struct {
	ent.Schema
}

func (UserCredential) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user_credentials",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("用户认证信息表"),
	}
}

// Fields of the UserCredential.
func (UserCredential) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("user_id").
			Comment("关联主表的用户ID").
			Nillable().
			Optional(),

		field.Enum("identity_type").
			Comment("认证方式类型").
			NamedValues(
				"Password", "PASSWORD",
				"Email", "EMAIL",
				"Phone", "PHONE",

				"Wechat", "WECHAT",
				"QQ", "QQ",
				"Google", "GOOGLE",
				"Facebook", "FACEBOOK",
				"Apple", "APPLE",
				"Telegram", "TELEGRAM",
			).
			Default("PASSWORD").
			Nillable().
			Optional(),

		field.String("identifier").
			Comment("身份唯一标识符").
			NotEmpty().
			Nillable().
			Optional(),

		field.Enum("credential_type").
			Comment("凭证类型").
			NamedValues(
				"PasswordHash", "PASSWORD_HASH",
				"AccessToken", "ACCESS_TOKEN",
				"RefreshToken", "REFRESH_TOKEN",
			).
			Default("PASSWORD_HASH").
			Nillable().
			Optional(),

		field.String("credential").
			Comment("凭证").
			NotEmpty().
			Nillable().
			Optional(),

		field.Bool("is_primary").
			Comment("是否主认证方式").
			Default(false).
			Nillable().
			Optional(),

		field.Enum("status").
			Comment("凭证状态").
			NamedValues(
				"Disabled", "DISABLED",
				"Enabled", "ENABLED",
				"Expired", "EXPIRED",
				"Unverified", "UNVERIFIED",
				"Removed", "REMOVED",
				"Blocked", "BLOCKED",
				"Temporary", "TEMPORARY",
			).
			Default("ENABLED").
			Nillable().
			Optional(),

		field.String("extra_info").
			Comment("扩展信息").
			SchemaType(map[string]string{
				dialect.MySQL:    "json",
				dialect.Postgres: "jsonb",
			}).
			Nillable().
			Optional(),

		field.String("activate_token").
			Comment("激活账号用的令牌").
			MaxLen(255).
			Unique().
			Optional().
			Nillable(),

		field.String("reset_token").
			Comment("重置密码用的令牌").
			MaxLen(255).
			Unique().
			Optional().
			Nillable(),
	}
}

// Edges of the UserCredential.
func (UserCredential) Edges() []ent.Edge {
	return nil
}

// Mixin of the UserCredential.
func (UserCredential) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixin.AutoIncrementId{},
		appmixin.TenantID{},
	}
}

// Indexes of the UserCredential.
func (UserCredential) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "identity_type", "identifier").Unique(),
		index.Fields("identifier"),
		index.Fields("user_id"),
	}
}
