package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	"github.com/tx7do/go-utils/copierutil"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/dict"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type DictRepo struct {
	data         *Data
	log          *log.Helper
	copierOption copier.Option
}

func NewDictRepo(data *Data, logger log.Logger) *DictRepo {
	repo := &DictRepo{
		log:  log.NewHelper(log.With(logger, "module", "dict/repo/admin-service")),
		data: data,
	}

	repo.init()

	return repo
}

func (r *DictRepo) init() {
	r.copierOption = copier.Option{
		Converters: []copier.TypeConverter{
			copierutil.TimeToStringConverter,
			copierutil.StringToTimeConverter,
			copierutil.TimeToTimestamppbConverter,
			copierutil.TimestamppbToTimeConverter,
		},
	}
}

func (r *DictRepo) toProto(in *ent.Dict) *adminV1.Dict {
	if in == nil {
		return nil
	}

	var out adminV1.Dict
	_ = copier.Copy(&out, in)

	//out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)
	//out.UpdateTime = timeutil.TimeToTimeString(in.UpdateTime)
	//out.DeleteTime = timeutil.TimeToTimeString(in.DeleteTime)

	return &out
}

func (r *DictRepo) toEnt(in *adminV1.Dict) *ent.Dict {
	if in == nil {
		return nil
	}

	var out ent.Dict
	_ = copier.Copy(&out, in)

	return &out
}

func (r *DictRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Dict.Query()
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

func (r *DictRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListDictResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Dict.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), dict.FieldCreateTime,
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

	items := make([]*adminV1.Dict, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListDictResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *DictRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Dict.Query().
		Where(dict.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *DictRepo) Get(ctx context.Context, req *adminV1.GetDictRequest) (*adminV1.Dict, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().Dict.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *DictRepo) Create(ctx context.Context, req *adminV1.CreateDictRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
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
		return adminV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *DictRepo) Update(ctx context.Context, req *adminV1.UpdateDictRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &adminV1.CreateDictRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return adminV1.ErrorBadRequest("invalid field mask")
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
		return adminV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *DictRepo) Delete(ctx context.Context, req *adminV1.DeleteDictRequest) error {
	if req == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().Dict.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return adminV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return adminV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
