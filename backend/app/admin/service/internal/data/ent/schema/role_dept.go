package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// RoleDept holds the schema definition for the RoleDept entity.
type RoleDept struct {
	ent.Schema
}

func (RoleDept) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_role_dept",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("角色 - 部门关联表"),
	}
}

// Fields of the RoleDept.
func (RoleDept) Fields() []ent.Field {
	return []ent.Field{

		field.Uint32("role_id").
			Comment("角色ID").
			Nillable(),

		field.Uint32("dept_id").
			Comment("部门ID").
			Nillable(),
	}
}

// Mixin of the RoleDept.
func (RoleDept) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
	}
}

// Indexes of the RoleDept.
func (RoleDept) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "dept_id").Unique().StorageKey("idx_sys_role_dept_role_id_dept_id"),
		index.Fields("role_id").StorageKey("idx_sys_role_dept_role_id"),
	}
}
