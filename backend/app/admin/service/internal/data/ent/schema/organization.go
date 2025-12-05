package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_organizations",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("组织表"),
	}
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("组织名称").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("组织状态").
			NamedValues(
				"On", "ON",
				"Off", "OFF",
			).
			Default("ON").
			Optional().
			Nillable(),

		field.Enum("organization_type").
			Comment("组织类型").
			NamedValues(
				"Group", "GROUP",
				"Subsidiary", "SUBSIDIARY",
				"Filiale", "FILIALE",
				"Division", "DIVISION",
			).
			Optional().
			Nillable(),

		field.String("credit_code").
			Comment("统一社会信用代码").
			Optional().
			Nillable(),

		field.String("address").
			Comment("注册地址").
			Optional().
			Nillable(),

		field.String("business_scope").
			Comment("核心业务范围").
			Optional().
			Nillable(),

		field.Bool("is_legal_entity").
			Comment("是否法人实体").
			Optional().
			Nillable(),

		field.Uint32("manager_id").
			Comment("负责人ID").
			Optional().
			Nillable(),
	}
}

// Mixin of the Organization.
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.Remark{},
		mixin.SortOrder{},
		mixin.Tree[Organization]{},
		mixin.TenantID{},
	}
}

// Indexes of the Organization.
func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").StorageKey("idx_sys_organization_name"),
	}
}
