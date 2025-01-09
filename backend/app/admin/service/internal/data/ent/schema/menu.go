package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "menus",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").
			Comment("id").
			StructTag(`json:"id,omitempty"`).
			Positive().
			Immutable().
			Unique(),

		field.Int32("parent_id").
			Comment("上一层菜单ID").
			Optional().
			Nillable(),

		field.Enum("type").
			Comment("菜单类型 FOLDER: 目录 MENU: 菜单 BUTTON: 按钮").
			Values(
				"FOLDER",
				"MENU",
				"BUTTON",
			).
			Default("MENU").
			Optional().
			Nillable(),

		field.String("path").
			Comment("路径,当其类型为'按钮'的时候对应的数据操作名,例如:/user.service.v1.UserService/Login").
			Default("").
			Optional().
			Nillable(),

		field.String("redirect").
			Comment("重定向地址").
			Optional().
			Nillable(),

		field.String("alias").
			Comment("路由别名").
			Optional().
			Nillable(),

		field.String("name").
			Comment("路由命名，然后我们可以使用 name 而不是 path 来传递 to 属性给 <router-link>。").
			Optional().
			Nillable(),

		field.String("component").
			Comment("前端页面组件").
			Default("").
			Optional().
			Nillable(),

		field.JSON("meta", &systemV1.RouteMeta{}).
			Comment("前端页面组件").
			Optional(),
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SwitchStatus{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Menu.Type).Annotations(entproto.Field(10)).
			From("parent").Unique().Field("parent_id").Annotations(entproto.Field(11)),
	}
}
