package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"

	appmixin "kratos-admin/pkg/entgo/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "users",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("用户表"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Comment("用户名").
			Unique().
			NotEmpty().
			Immutable().
			Optional().
			Nillable(),

		field.String("nickname").
			Comment("昵称").
			MaxLen(255).
			Optional().
			Nillable(),

		field.String("realname").
			Comment("真实名字").
			MaxLen(255).
			Optional().
			Nillable(),

		field.String("email").
			Comment("电子邮箱").
			MaxLen(320).
			Optional().
			Nillable(),

		field.String("mobile").
			Comment("手机号码").
			Default("").
			MaxLen(255).
			Optional().
			Nillable(),

		field.String("telephone").
			Comment("座机号码").
			Default("").
			MaxLen(255).
			Optional().
			Nillable(),

		field.String("avatar").
			Comment("头像").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("address").
			Comment("地址").
			Default("").
			MaxLen(2048).
			Optional().
			Nillable(),

		field.String("region").
			Comment("国家地区").
			Default("").
			MaxLen(255).
			Optional().
			Nillable(),

		field.String("description").
			Comment("个人说明").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.Enum("gender").
			Comment("性别").
			NamedValues(
				"Secret", "SECRET",
				"Male", "MALE",
				"Female", "FEMALE",
			).
			Optional().
			Nillable(),

		field.Enum("authority").
			Comment("授权").
			Optional().
			Nillable().
			NamedValues(
				"SysAdmin", "SYS_ADMIN",
				"TenantAdmin", "TENANT_ADMIN",
				"CustomerUser", "CUSTOMER_USER",
				"Guest", "GUEST",
			).
			Default("CUSTOMER_USER"),

		field.Time("last_login_time").
			Comment("最后一次登录的时间").
			Optional().
			Nillable(),

		field.String("last_login_ip").
			Comment("最后一次登录的IP").
			Default("").
			MaxLen(64).
			Optional().
			Nillable(),

		field.Uint32("org_id").
			Comment("部门ID").
			Optional().
			Nillable(),

		field.Uint32("position_id").
			Comment("职位ID").
			Optional().
			Nillable(),

		field.Uint32("work_id").
			Comment("员工工号").
			Optional().
			Nillable(),

		//field.Uint32("role_id").
		//	Comment("角色ID").
		//	Optional().
		//	Nillable(),

		field.Strings("roles").
			Comment("多角色角色码列表").
			Optional(),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Time{},
		mixin.Remark{},
		mixin.SwitchStatus{},
		appmixin.TenantID{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{}
}
