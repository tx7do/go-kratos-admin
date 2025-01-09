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
	timeutil "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/dict"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
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
		CreateTime:    timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime:    timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime:    timeutil.TimeToTimestamppb(in.DeleteTime),
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
	if err != nil && !ent.IsNotFound(err) {
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
		SetNillableCreateBy(req.OperatorId)

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	} else {
		builder.SetCreateTime(*timeutil.TimestamppbToTime(req.Data.CreateTime))
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
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
			return r.Create(ctx, &systemV1.CreateDictRequest{Data: req.Data, OperatorId: req.OperatorId})
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
		SetNillableUpdateBy(req.OperatorId)

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	} else {
		builder.SetUpdateTime(*timeutil.TimestamppbToTime(req.Data.UpdateTime))
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.Data, req.GetUpdateMask().GetPaths())
		nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *DictRepo) Delete(ctx context.Context, req *systemV1.DeleteDictRequest) (bool, error) {
	err := r.data.db.Client().Dict.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
