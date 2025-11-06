package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// DictType holds the schema definition for the DictType entity.
type DictType struct {
	ent.Schema
}

func (DictType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_dict_types",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("字典类型表"),
	}
}

// Fields of the DictType.
func (DictType) Fields() []ent.Field {
	return []ent.Field{
		field.String("type_code").
			Comment("字典类型唯一代码").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("type_name").
			Comment("字典类型名称").
			NotEmpty().
			Optional().
			Nillable(),
	}
}

// Mixin of the DictType.
func (DictType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.IsEnabled{},
		mixin.SortOrder{},
		mixin.Description{},
		mixin.TenantID{},
	}
}

// Indexes of the DictType.
func (DictType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type_code").
			Unique().
			StorageKey("idx_sys_dict_types_type_code"),
	}
}

// Edges of the DictType.
func (DictType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("entries", DictEntry.Type).
			Ref("sys_dict_types"),
	}
}
