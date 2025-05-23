package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	appmixin "kratos-admin/pkg/entgo/mixin"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "tasks",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("任务表"),
	}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Comment("任务类型").
			NamedValues(
				"Periodic", "PERIODIC",
				"Delay", "DELAY",
				"WaitResult", "WAIT_RESULT",
			).
			Optional().
			Nillable(),

		field.String("type_name").
			Comment("任务执行类型名").
			Unique().
			Optional().
			Nillable(),

		field.String("task_payload").
			Comment("任务数据").
			SchemaType(map[string]string{
				dialect.MySQL:    "json",
				dialect.Postgres: "jsonb",
			}).
			Optional().
			Nillable(),

		field.String("cron_spec").
			Comment("cron表达式").
			Optional().
			Nillable(),

		field.JSON("task_options", &adminV1.TaskOption{}).
			Comment("任务选项").
			Optional(),

		field.Bool("enable").
			Comment("启用/禁用任务").
			Optional().
			Nillable(),
	}
}

// Mixin of the Task.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
		appmixin.TenantID{},
	}
}
