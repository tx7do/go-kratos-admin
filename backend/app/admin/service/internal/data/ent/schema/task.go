package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
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
	}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Comment("任务类型").
			Values("Periodic", "Delay", "WaitResult").
			Optional().
			Nillable(),

		field.String("type_name").
			Comment("任务执行类型名").
			Unique().
			Optional().
			Nillable(),

		field.String("task_payload").
			Comment("任务的参数，以 JSON 格式存储，方便存储不同类型和数量的参数").
			Optional().
			Nillable(),

		field.String("cron_spec").
			Comment("cron表达式，用于定义任务的调度时间").
			Optional().
			Nillable(),

		field.Uint32("retry_count").
			Comment("任务最多可以重试的次数").
			Optional().
			Nillable(),

		field.Uint64("timeout").
			Comment("任务超时时间").
			Optional().
			Nillable(),

		field.Time("deadline").
			Comment("任务超时时间").
			Optional().
			Nillable(),

		field.Uint64("process_in").
			Comment("任务延迟处理时间").
			Optional().
			Nillable(),

		field.Time("process_at").
			Comment("任务执行时间点").
			Optional().
			Nillable(),

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
	}
}
