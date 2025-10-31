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

// DictItem holds the schema definition for the DictItem entity.
type DictItem struct {
	ent.Schema
}

func (DictItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_dict_items",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("子字典表"),
	}
}

// Fields of the DictItem.
func (DictItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			Comment("子项编码").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("name").
			Comment("子项名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.Uint32("main_id").
			Comment("主字典ID").
			Optional().
			Nillable(),

		field.Int32("sort_id").
			Comment("排序ID").
			Default(0).
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("字典状态").
			NamedValues(
				"On", "ON",
				"Off", "OFF",
			).
			Default("ON").
			Optional().
			Nillable(),
	}
}

// Mixin of the DictItem.
func (DictItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Indexes of the DictItem.
func (DictItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("idx_sys_dict_items_code"),
		index.Fields("status").StorageKey("idx_sys_dict_items_status"),
	}
}
