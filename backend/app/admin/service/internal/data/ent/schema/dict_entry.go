package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// DictEntry holds the schema definition for the DictEntry entity.
type DictEntry struct {
	ent.Schema
}

func (DictEntry) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_dict_entries",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("字典条目表"),
	}
}

// Fields of the DictEntry.
func (DictEntry) Fields() []ent.Field {
	return []ent.Field{
		field.String("entry_label").
			Comment("字典项的显示标签").
			NotEmpty().
			Optional().
			Nillable(),

		field.String("entry_value").
			Comment("字典项的实际值").
			NotEmpty().
			Optional().
			Nillable(),

		field.Int32("numeric_value").
			Comment("数值型值").
			Optional().
			Nillable(),

		field.String("language_code").
			Comment("语言代码").
			Optional().
			Nillable(),
	}
}

// Mixin of the DictEntry.
func (DictEntry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.Description{},
		mixin.SortOrder{},
		mixin.IsEnabled{},
		mixin.TenantID{},
	}
}

// Indexes of the DictEntry.
func (DictEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("entry_value").
			Edges("sys_dict_types").
			Unique().
			StorageKey("uk_sys_dict_entries_entry_value"),
	}
}

// Edges of the DictEntry.
func (DictEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sys_dict_types", DictType.Type).
			Unique().
			Required().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			StorageKey(edge.Column("type_id")),
	}
}
