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
	"kratos-admin/app/admin/service/internal/data/ent/apiresource"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type ApiResourceRepo struct {
	data         *Data
	log          *log.Helper
	copierOption copier.Option
}

func NewApiResourceRepo(data *Data, logger log.Logger) *ApiResourceRepo {
	repo := &ApiResourceRepo{
		log:  log.NewHelper(log.With(logger, "module", "api-resource/repo/admin-service")),
		data: data,
	}

	repo.init()

	return repo
}

func (r *ApiResourceRepo) init() {
	r.copierOption = copier.Option{
		Converters: []copier.TypeConverter{},
	}

	r.copierOption.Converters = append(r.copierOption.Converters, copierutil.NewTimeStringConverterPair()...)
	r.copierOption.Converters = append(r.copierOption.Converters, copierutil.NewTimeTimestamppbConverterPair()...)
}

func (r *ApiResourceRepo) toProto(in *ent.ApiResource) *adminV1.ApiResource {
	if in == nil {
		return nil
	}

	var out adminV1.ApiResource
	_ = copier.CopyWithOption(&out, in, r.copierOption)

	//out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)
	//out.UpdateTime = timeutil.TimeToTimeString(in.UpdateTime)
	//out.DeleteTime = timeutil.TimeToTimeString(in.DeleteTime)

	return &out
}

func (r *ApiResourceRepo) toEnt(in *adminV1.ApiResource) *ent.ApiResource {
	if in == nil {
		return nil
	}

	var out ent.ApiResource
	_ = copier.CopyWithOption(&out, in, r.copierOption)

	return &out
}

func (r *ApiResourceRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().ApiResource.Query()
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

func (r *ApiResourceRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListApiResourceResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().ApiResource.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), apiresource.FieldCreateTime,
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

	items := make([]*adminV1.ApiResource, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListApiResourceResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *ApiResourceRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().ApiResource.Query().
		Where(apiresource.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *ApiResourceRepo) Get(ctx context.Context, req *adminV1.GetApiResourceRequest) (*adminV1.ApiResource, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().ApiResource.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("api resource not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *ApiResourceRepo) Create(ctx context.Context, req *adminV1.CreateApiResourceRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().ApiResource.Create().
		SetNillableDescription(req.Data.Description).
		SetNillableModule(req.Data.Module).
		SetNillableOperation(req.Data.Operation).
		SetNillablePath(req.Data.Path).
		SetNillableMethod(req.Data.Method).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

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

func (r *ApiResourceRepo) Update(ctx context.Context, req *adminV1.UpdateApiResourceRequest) error {
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
			createReq := &adminV1.CreateApiResourceRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return adminV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().ApiResource.
		UpdateOneID(req.Data.GetId()).
		SetNillableDescription(req.Data.Description).
		SetNillableModule(req.Data.Module).
		SetNillableOperation(req.Data.Operation).
		SetNillablePath(req.Data.Path).
		SetNillableMethod(req.Data.Method).
		SetNillableUpdateBy(req.Data.UpdateBy).
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

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *ApiResourceRepo) Delete(ctx context.Context, req *adminV1.DeleteApiResourceRequest) error {
	if req == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().ApiResource.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return adminV1.ErrorNotFound("api resource not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return adminV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
