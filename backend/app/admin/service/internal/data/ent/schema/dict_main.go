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

// DictMain holds the schema definition for the DictMain entity.
type DictMain struct {
	ent.Schema
}

func (DictMain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_dict_mains",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("主字典表"),
	}
}

// Fields of the DictMain.
func (DictMain) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			Comment("主字典编码").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("name").
			Comment("主字典名称").
			NotEmpty().
			Optional().
			Nillable(),
	}
}

// Mixin of the DictMain.
func (DictMain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}

// Indexes of the DictMain.
func (DictMain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("idx_sys_dict_mains_code"),
	}
}
