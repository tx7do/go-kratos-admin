package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"

	appmixin "kratos-admin/pkg/entgo/mixin"
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

		field.Uint32("parent_id").
			Comment("上一层组织ID").
			Optional().
			Nillable(),

		field.Int32("sort_id").
			Comment("排序ID").
			Default(0).
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
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Indexes of the Organization.
func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").StorageKey("idx_sys_organization_name"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Organization.Type).
			From("parent").Unique().Field("parent_id"),
	}
}
