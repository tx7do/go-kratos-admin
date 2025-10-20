package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"

	appmixin "kratos-admin/pkg/entgo/mixin"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "departments",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("部门表"),
	}
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("部门名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("上一层部门ID").
			Optional().
			Nillable(),

		field.Uint32("organization_id").
			Comment("所属组织ID").
			Nillable(),

		field.Uint32("manager_id").
			Comment("负责人ID").
			Optional().
			Nillable(),

		field.Int32("sort_id").
			Comment("排序ID").
			Default(0).
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("部门状态").
			NamedValues(
				"DEPARTMENT_STATUS_ON", "ON",
				"DEPARTMENT_STATUS_OFF", "OFF",
			).
			Optional().
			Nillable(),

		field.String("description").
			Comment("职能描述").
			Optional().
			Nillable(),
	}
}

// Mixin of the Department.
func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Department.Type).
			From("parent").Unique().Field("parent_id"),
	}
}
