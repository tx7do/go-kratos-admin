package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	"github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/task"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
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

func (r *TaskRepo) toEntType(in *adminV1.TaskType) *task.Type {
	if in == nil {
		return nil
	}

	find, ok := adminV1.TaskType_name[int32(*in)]
	if !ok {
		return nil
	}

	return (*task.Type)(trans.Ptr(find))
}

func (r *TaskRepo) toProtoType(in *task.Type) *adminV1.TaskType {
	if in == nil {
		return nil
	}

	find, ok := adminV1.TaskType_value[string(*in)]
	if !ok {
		return nil
	}

	return (*adminV1.TaskType)(trans.Ptr(find))
}

func (r *TaskRepo) toProto(in *ent.Task) *adminV1.Task {
	if in == nil {
		return nil
	}

	var out adminV1.Task
	_ = copier.Copy(&out, in)

	out.Type = r.toProtoType(in.Type)
	out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)
	out.UpdateTime = timeutil.TimeToTimeString(in.UpdateTime)
	out.DeleteTime = timeutil.TimeToTimeString(in.DeleteTime)

	return &out
}

func (r *TaskRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Task.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, adminV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *TaskRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListTaskResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Task.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), task.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, adminV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, adminV1.ErrorInternalServerError("query list failed")
	}

	items := make([]*adminV1.Task, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListTaskResponse{
		Total: uint32(count),
		Items: items,
	}, nil
}

func (r *TaskRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Task.Query().
		Where(task.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *TaskRepo) Get(ctx context.Context, id uint32) (*adminV1.Task, error) {
	if id == 0 {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().Task.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("task not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *TaskRepo) GetByTypeName(ctx context.Context, typeName string) (*adminV1.Task, error) {
	if typeName == "" {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().Task.Query().
		Where(task.TypeNameEQ(typeName)).
		First(ctx)
	if err != nil {
		r.log.Errorf("query one data failed: %s", err.Error())

		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("task not found")
		}

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *TaskRepo) Create(ctx context.Context, req *adminV1.CreateTaskRequest) (*adminV1.Task, error) {
	if req == nil || req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Task.Create().
		SetNillableType(r.toEntType(req.Data.Type)).
		SetNillableTypeName(req.Data.TypeName).
		SetNillableTaskPayload(req.Data.TaskPayload).
		SetNillableTaskID(req.Data.TaskId).
		SetNillableCronSpec(req.Data.CronSpec).
		SetNillableEnable(req.Data.Enable).
		SetNillableRemark(req.Data.Remark).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.TaskOptions != nil {
		builder.SetTaskOptions(req.Data.TaskOptions)
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	t, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, adminV1.ErrorInternalServerError("insert data failed")
	}

	return r.toProto(t), nil
}

func (r *TaskRepo) Update(ctx context.Context, req *adminV1.UpdateTaskRequest) (*adminV1.Task, error) {
	if req == nil || req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return nil, err
		}
		if !exist {
			createReq := &adminV1.CreateTaskRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return nil, adminV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().
		//Debug().
		Task.UpdateOneID(req.Data.GetId()).
		SetNillableType(r.toEntType(req.Data.Type)).
		SetNillableTypeName(req.Data.TypeName).
		SetNillableTaskPayload(req.Data.TaskPayload).
		SetNillableTaskID(req.Data.TaskId).
		SetNillableCronSpec(req.Data.CronSpec).
		SetNillableEnable(req.Data.Enable).
		SetNillableRemark(req.Data.Remark).
		SetNillableUpdateBy(req.Data.UpdateBy).
		SetNillableUpdateTime(timeutil.StringTimeToTime(req.Data.UpdateTime))

	if req.Data.TaskOptions != nil {
		builder.SetTaskOptions(req.Data.TaskOptions)
	}

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
		return nil, adminV1.ErrorInternalServerError("update data failed")
	}

	return r.toProto(t), nil
}

func (r *TaskRepo) Delete(ctx context.Context, req *adminV1.DeleteTaskRequest) error {
	if req == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().Task.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return adminV1.ErrorNotFound("task not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return adminV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
