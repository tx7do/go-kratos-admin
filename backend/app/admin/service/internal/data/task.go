package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/task"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type TaskRepo struct {
	data *Data
	log  *log.Helper
}

func NewTaskRepo(data *Data, logger log.Logger) *TaskRepo {
	l := log.NewHelper(log.With(logger, "module", "task/repo/admin-service"))
	return &TaskRepo{
		data: data,
		log:  l,
	}
}

func (r *TaskRepo) toEntType(in *systemV1.TaskType) *task.Type {
	if in == nil {
		return nil
	}

	switch *in {
	case systemV1.TaskType_TaskType_Periodic:
		return trans.Ptr(task.TypePeriodic)

	case systemV1.TaskType_TaskType_Delay:
		return trans.Ptr(task.TypeDelay)

	case systemV1.TaskType_TaskType_WaitResult:
		return trans.Ptr(task.TypeWaitResult)

	default:
		return nil
	}
}

func (r *TaskRepo) toProtoType(in *task.Type) *systemV1.TaskType {
	if in == nil {
		return nil
	}

	switch *in {
	case task.TypePeriodic:
		return trans.Ptr(systemV1.TaskType_TaskType_Periodic)

	case task.TypeDelay:
		return trans.Ptr(systemV1.TaskType_TaskType_Delay)

	case task.TypeWaitResult:
		return trans.Ptr(systemV1.TaskType_TaskType_WaitResult)

	default:
		return nil
	}
}

func (r *TaskRepo) convertEntToProto(in *ent.Task) *systemV1.Task {
	if in == nil {
		return nil
	}

	return &systemV1.Task{
		Id:          trans.Ptr(in.ID),
		Type:        r.toProtoType(in.Type),
		TypeName:    in.TypeName,
		TaskPayload: in.TaskPayload,
		CronSpec:    in.CronSpec,
		RetryCount:  in.RetryCount,
		Timeout:     timeutil.NumberToDurationpb(in.Timeout, time.Second),
		Deadline:    timeutil.TimeToTimestamppb(in.Deadline),
		ProcessIn:   timeutil.NumberToDurationpb(in.ProcessIn, time.Second),
		ProcessAt:   timeutil.TimeToTimestamppb(in.ProcessAt),
		Enable:      in.Enable,
		Remark:      in.Remark,
		CreateBy:    in.CreateBy,
		UpdateBy:    in.UpdateBy,
		CreateTime:  timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime:  timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime:  timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *TaskRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Task.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *TaskRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListTaskResponse, error) {
	builder := r.data.db.Client().Task.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), task.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析SELECT条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*systemV1.Task, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListTaskResponse{
		Total: uint32(count),
		Items: items,
	}, nil
}

func (r *TaskRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().Task.Query().
		Where(task.IDEQ(id)).
		Exist(ctx)
}

func (r *TaskRepo) Get(ctx context.Context, id uint32) (*systemV1.Task, error) {
	ret, err := r.data.db.Client().Task.Get(ctx, id)
	if err != nil {
		r.log.Errorf("query one data failed: %s", err.Error())

		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("task not found")
		}

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *TaskRepo) GetByTypeName(ctx context.Context, typeName string) (*systemV1.Task, error) {
	ret, err := r.data.db.Client().Task.Query().
		Where(task.TypeNameEQ(typeName)).
		First(ctx)
	if err != nil {
		r.log.Errorf("query one data failed: %s", err.Error())

		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("task not found")
		}

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *TaskRepo) Create(ctx context.Context, req *systemV1.CreateTaskRequest, operator *auth.UserTokenPayload) (*systemV1.Task, error) {
	if req.Data == nil {
		return nil, errors.New("invalid request")
	}

	builder := r.data.db.Client().Task.Create().
		SetNillableType(r.toEntType(req.Data.Type)).
		SetNillableTypeName(req.Data.TypeName).
		SetNillableTaskPayload(req.Data.TaskPayload).
		SetNillableCronSpec(req.Data.CronSpec).
		SetNillableRetryCount(req.Data.RetryCount).
		SetNillableTimeout(timeutil.DurationpbToNumber[uint64](req.Data.Timeout, time.Second)).
		SetNillableDeadline(timeutil.TimestamppbToTime(req.Data.Deadline)).
		SetNillableProcessIn(timeutil.DurationpbToNumber[uint64](req.Data.ProcessIn, time.Second)).
		SetNillableProcessAt(timeutil.TimestamppbToTime(req.Data.ProcessAt)).
		SetNillableEnable(req.Data.Enable).
		SetNillableRemark(req.Data.Remark).
		SetNillableCreateBy(trans.Ptr(operator.UserId)).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	t, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(t), nil
}

func (r *TaskRepo) Update(ctx context.Context, req *systemV1.UpdateTaskRequest, operator *auth.UserTokenPayload) (*systemV1.Task, error) {
	if req == nil || req.Data == nil {
		return nil, errors.New("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return nil, err
		}
		if !exist {
			return r.Create(ctx, &systemV1.CreateTaskRequest{Data: req.Data}, operator)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return nil, errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().
		//Debug().
		Task.UpdateOneID(req.Data.GetId()).
		SetNillableType(r.toEntType(req.Data.Type)).
		SetNillableTypeName(req.Data.TypeName).
		SetNillableTaskPayload(req.Data.TaskPayload).
		SetNillableCronSpec(req.Data.CronSpec).
		SetNillableRetryCount(req.Data.RetryCount).
		SetNillableTimeout(timeutil.DurationpbToNumber[uint64](req.Data.Timeout, time.Second)).
		SetNillableDeadline(timeutil.TimestamppbToTime(req.Data.Deadline)).
		SetNillableProcessIn(timeutil.DurationpbToNumber[uint64](req.Data.ProcessIn, time.Second)).
		SetNillableProcessAt(timeutil.TimestamppbToTime(req.Data.ProcessAt)).
		SetNillableEnable(req.Data.Enable).
		SetNillableRemark(req.Data.Remark).
		SetNillableUpdateBy(trans.Ptr(operator.UserId)).
		SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.Data, req.GetUpdateMask().GetPaths())
		nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

	t, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(t), nil
}

func (r *TaskRepo) Delete(ctx context.Context, req *systemV1.DeleteTaskRequest) (bool, error) {
	err := r.data.db.Client().Task.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
