package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// AdminLoginRestriction holds the schema definition for the AdminLoginRestriction entity.
type AdminLoginRestriction struct {
	ent.Schema
}

func (AdminLoginRestriction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "admin_login_restrictions",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("后台登录限制表"),
	}
}

// Fields of the AdminLoginRestriction.
func (AdminLoginRestriction) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("admin_id").
			Comment("管理员ID").
			Unique().
			Optional().
			Nillable(),

		field.String("value").
			Comment("限制值（如IP地址、MAC地址或地区代码）").
			Optional().
			Nillable(),

		field.String("reason").
			Comment("限制原因").
			Optional().
			Nillable(),

		field.Enum("type").
			Comment("限制类型").
			NamedValues(
				"Unspecified", "UNSPECIFIED",
				"Blacklist", "BLACKLIST",
				"Whitelist", "WHITELIST",
			).
			Optional().
			Nillable(),

		field.Enum("method").
			Comment("限制方式").
			NamedValues(
				"Unspecified", "UNSPECIFIED",
				"Ip", "IP",
				"Mac", "MAC",
				"Region", "REGION",
				"Time", "TIME",
				"Device", "DEVICE",
			).
			Optional().
			Nillable(),
	}
}

// Mixin of the AdminLoginRestriction.
func (AdminLoginRestriction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
	}
}
