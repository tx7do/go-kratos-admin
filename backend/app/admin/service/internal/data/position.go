package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/position"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type PositionRepo struct {
	data *Data
	log  *log.Helper
}

func NewPositionRepo(data *Data, logger log.Logger) *PositionRepo {
	l := log.NewHelper(log.With(logger, "module", "position/repo/admin-service"))
	return &PositionRepo{
		data: data,
		log:  l,
	}
}

func (r *PositionRepo) convertEntToProto(in *ent.Position) *userV1.Position {
	if in == nil {
		return nil
	}
	return &userV1.Position{
		Id:         trans.Ptr(in.ID),
		Name:       &in.Name,
		Code:       &in.Code,
		Remark:     in.Remark,
		SortId:     &in.SortID,
		ParentId:   &in.ParentID,
		Status:     (*string)(in.Status),
		CreateBy:   in.CreateBy,
		UpdateBy:   in.UpdateBy,
		CreateTime: timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime: timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *PositionRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Position.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *PositionRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListPositionResponse, error) {
	builder := r.data.db.Client().Position.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), position.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*userV1.Position, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListPositionResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *PositionRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().Position.Query().
		Where(position.IDEQ(id)).
		Exist(ctx)
}

func (r *PositionRepo) Get(ctx context.Context, req *userV1.GetPositionRequest) (*userV1.Position, error) {
	ret, err := r.data.db.Client().Position.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorPositionNotFound("position not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *PositionRepo) Create(ctx context.Context, req *userV1.CreatePositionRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().Position.Create().
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableCode(req.Data.Code).
		SetNillableStatus((*position.Status)(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *PositionRepo) Update(ctx context.Context, req *userV1.UpdatePositionRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &userV1.CreatePositionRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().Position.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableCode(req.Data.Code).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*position.Status)(req.Data.Status)).
		SetNillableUpdateBy(req.Data.UpdateBy).
		SetNillableUpdateTime(timeutil.StringTimeToTime(req.Data.UpdateTime))

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

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *PositionRepo) Delete(ctx context.Context, req *userV1.DeletePositionRequest) (bool, error) {
	err := r.data.db.Client().Position.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
