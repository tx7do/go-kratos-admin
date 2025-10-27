package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-utils/entgo/mixin"

	appmixin "kratos-admin/pkg/entgo/mixin"
)

// Dict holds the schema definition for the Dict entity.
type Dict struct {
	ent.Schema
}

func (Dict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_dicts",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("字典表"),
	}
}

// Fields of the Dict.
func (Dict) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").
			Comment("字典键").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("category").
			Comment("字典类型").
			Optional().
			Nillable(),

		field.String("category_desc").
			Comment("字典类型名称").
			Optional().
			Nillable(),

		field.String("value").
			Comment("字典值").
			Optional().
			Nillable(),

		field.String("value_desc").
			Comment("字典值名称").
			Optional().
			Nillable(),

		field.String("value_data_type").
			Comment("字典值数据类型").
			Optional().
			Nillable(),

		field.Int32("sort_id").
			Comment("排序ID").
			Default(0).
			Optional().
			Nillable(),
	}
}

// Mixin of the Dict.
func (Dict) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.SwitchStatus{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Indexes of the Dict.
func (Dict) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique().StorageKey("idx_sys_dict_key"),
	}
}
