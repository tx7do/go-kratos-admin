package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// UserCredential holds the schema definition for the UserCredential entity.
type UserCredential struct {
	ent.Schema
}

func (UserCredential) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_user_credentials",
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
				"Username", "USERNAME",
				"UserId", "USERID",
				"Email", "EMAIL",
				"Phone", "PHONE",

				"SocialOauth", "SOCIAL_OAUTH",
				"EnterpriseSso", "ENTERPRISE_SSO",
				"IdentityApiKey", "IDENTITY_API_KEY",
				"DeviceId", "DEVICE_ID",
				"Custom", "CUSTOM",
			).
			Default("USERNAME").
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

				"ApiKey", "API_KEY",
				"ApiSecret", "API_SECRET",

				"AccessToken", "ACCESS_TOKEN",
				"RefreshToken", "REFRESH_TOKEN",
				"JWT", "JWT",

				"OauthToken", "OAUTH_TOKEN",
				"OauthAuthorizationCode", "OAUTH_AUTHORIZATION_CODE",
				"OauthClientCredentials", "OAUTH_CLIENT_CREDENTIALS",

				"OTP", "OTP",
				"TOTP", "TOTP",
				"SmsOtp", "SMS_OTP",
				"EmailOtp", "EMAIL_OTP",

				"HardwareToken", "HARDWARE_TOKEN",
				"SoftwareToken", "SOFTWARE_TOKEN",
				"SecurityQuestion", "SECURITY_QUESTION",

				"Biometric", "BIOMETRIC",
				"BiometricToken", "BIOMETRIC_TOKEN",

				"SsoToken", "SSO_TOKEN",
				"SamlAssertion", "SAML_ASSERTION",
				"OpenidConnectIdToken", "OPENID_CONNECT_ID_TOKEN",

				"SessionCookie", "SESSION_COOKIE",
				"TemporaryCredential", "TEMPORARY_CREDENTIAL",

				"Custom", "CUSTOM",
				"ReservedForFuture", "RESERVED_FOR_FUTURE",
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

		field.String("provider").
			Comment("第三方平台标识").
			Nillable().
			Optional(),

		field.String("provider_account_id").
			Comment("第三方平台的账号唯一ID").
			Nillable().
			Optional(),

		field.String("activate_token_hash").
			Comment("激活令牌哈希（不要存明文）").
			MaxLen(255).
			Nillable().
			Optional(),

		field.Time("activate_token_expires_at").
			Comment("激活令牌到期时间").
			Nillable().
			Optional(),

		field.Time("activate_token_used_at").
			Comment("激活令牌使用时间，单次使用时记录").
			Nillable().
			Optional(),

		field.String("reset_token_hash").
			Comment("重置密码令牌哈希（不要存明文）").
			MaxLen(255).
			Nillable().
			Optional(),

		field.Time("reset_token_expires_at").
			Comment("重置令牌到期时间").
			Nillable().
			Optional(),

		field.Time("reset_token_used_at").
			Comment("重置令牌使用时间").
			Nillable().
			Optional(),
	}
}

// Mixin of the UserCredential.
func (UserCredential) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeAt{},
		mixin.AutoIncrementId{},
		mixin.TenantID{},
	}
}

// Edges of the UserCredential.
func (UserCredential) Edges() []ent.Edge {
	return nil
}

// Indexes of the UserCredential.
func (UserCredential) Indexes() []ent.Index {
	return []ent.Index{
		// 组合唯一索引：注意 identifier 可能为 NULL，视数据库行为（Postgres 多个 NULL 允许）考虑在迁移层创建 partial unique index。
		index.Fields("user_id", "identity_type", "identifier").Unique().StorageKey("idx_sys_user_credential_uid_identity_identifier"),
		index.Fields("identifier").StorageKey("idx_sys_user_credential_identifier"),
		index.Fields("user_id").StorageKey("idx_sys_user_credential_user_id"),

		// 联合唯一索引：确保同一第三方平台账号不重复
		// 注意：若 provider/provider_account_id 为 NULL，Postgres 上可能允许多个 NULL。若需严格唯一，请在迁移脚本中为 Postgres 创建 partial unique index:
		// CREATE UNIQUE INDEX idx_user_credentials_provider_account ON sys_user_credentials (provider, provider_account_id) WHERE provider IS NOT NULL AND provider_account_id IS NOT NULL;
		index.Fields("provider", "provider_account_id").Unique().StorageKey("idx_sys_user_credential_provider_account"),

		index.Fields("provider").StorageKey("idx_sys_user_credential_provider"),
		index.Fields("user_id", "provider").StorageKey("idx_sys_user_credential_user_provider"),
	}
}
