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
	"kratos-admin/app/admin/service/internal/data/ent/dict"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type DictRepo struct {
	data *Data
	log  *log.Helper
}

func NewDictRepo(data *Data, logger log.Logger) *DictRepo {
	l := log.NewHelper(log.With(logger, "module", "dict/repo/admin-service"))
	return &DictRepo{
		data: data,
		log:  l,
	}
}

func (r *DictRepo) convertEntToProto(in *ent.Dict) *systemV1.Dict {
	if in == nil {
		return nil
	}
	return &systemV1.Dict{
		Id:            trans.Ptr(in.ID),
		Category:      in.Category,
		CategoryDesc:  in.CategoryDesc,
		Key:           in.Key,
		Value:         in.Value,
		ValueDesc:     in.ValueDesc,
		ValueDataType: in.ValueDataType,
		SortId:        in.SortID,
		Remark:        in.Remark,
		Status:        (*string)(in.Status),
		CreateBy:      in.CreateBy,
		UpdateBy:      in.UpdateBy,
		CreateTime:    timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime:    timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime:    timeutil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *DictRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Dict.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *DictRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListDictResponse, error) {
	builder := r.data.db.Client().Dict.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), dict.FieldCreateTime,
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

	items := make([]*systemV1.Dict, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListDictResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *DictRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().Dict.Query().
		Where(dict.IDEQ(id)).
		Exist(ctx)
}

func (r *DictRepo) Get(ctx context.Context, req *systemV1.GetDictRequest) (*systemV1.Dict, error) {
	ret, err := r.data.db.Client().Dict.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("dict not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *DictRepo) Create(ctx context.Context, req *systemV1.CreateDictRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().Dict.Create().
		SetNillableKey(req.Data.Key).
		SetNillableCategory(req.Data.Category).
		SetNillableCategoryDesc(req.Data.CategoryDesc).
		SetNillableValue(req.Data.Value).
		SetNillableValueDesc(req.Data.ValueDesc).
		SetNillableValueDataType(req.Data.ValueDataType).
		SetNillableSortID(req.Data.SortId).
		SetNillableStatus((*dict.Status)(req.Data.Status)).
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

func (r *DictRepo) Update(ctx context.Context, req *systemV1.UpdateDictRequest) error {
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
			createReq := &systemV1.CreateDictRequest{Data: req.Data}
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

	builder := r.data.db.Client().Dict.
		UpdateOneID(req.Data.GetId()).
		SetNillableKey(req.Data.Key).
		SetNillableCategory(req.Data.Category).
		SetNillableCategoryDesc(req.Data.CategoryDesc).
		SetNillableValue(req.Data.Value).
		SetNillableValueDesc(req.Data.ValueDesc).
		SetNillableValueDataType(req.Data.ValueDataType).
		SetNillableSortID(req.Data.SortId).
		SetNillableStatus((*dict.Status)(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
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

func (r *DictRepo) Delete(ctx context.Context, req *systemV1.DeleteDictRequest) (bool, error) {
	err := r.data.db.Client().Dict.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
