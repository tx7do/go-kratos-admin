package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_departments",
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

		field.Uint32("organization_id").
			Comment("所属组织ID").
			Nillable(),

		field.Uint32("manager_id").
			Comment("负责人ID").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("部门状态").
			NamedValues(
				"On", "ON",
				"Off", "OFF",
			).
			Default("ON").
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
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.SortOrder{},
		mixin.Remark{},
		mixin.Tree[Department]{},
		mixin.TenantID{},
	}
}

// Indexes of the Department.
func (Department) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").StorageKey("idx_sys_department_name"),
	}
}
