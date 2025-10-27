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

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

func (Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_positions",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("职位表"),
	}
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("职位名称").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("唯一编码").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.Uint32("parent_id").
			Comment("上一层职位ID").
			Optional().
			Nillable(),

		field.Int32("sort_id").
			Comment("排序ID").
			Default(0).
			Optional().
			Nillable(),

		field.Uint32("organization_id").
			Comment("所属组织ID").
			Nillable(),

		field.Uint32("department_id").
			Comment("所属部门ID").
			Nillable(),

		field.Enum("status").
			Comment("职位状态").
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

		field.Uint32("quota").
			Comment("编制人数").
			Optional().
			Nillable(),
	}
}

// Mixin of the Position.
func (Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Indexes of the Position.
func (Position) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("idx_sys_position_code"),
		index.Fields("name").StorageKey("idx_sys_position_name"),
	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("children", Position.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}
