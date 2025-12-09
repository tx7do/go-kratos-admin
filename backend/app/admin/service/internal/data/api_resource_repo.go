package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	entCrud "github.com/tx7do/go-crud/entgo"

	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/apiresource"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type ApiResourceRepo struct {
	data *Data
	log  *log.Helper

	mapper         *mapper.CopierMapper[adminV1.ApiResource, ent.ApiResource]
	scopeConverter *mapper.EnumTypeConverter[adminV1.ApiResource_Scope, apiresource.Scope]

	repository *entCrud.Repository[
		ent.ApiResourceQuery, ent.ApiResourceSelect,
		ent.ApiResourceCreate, ent.ApiResourceCreateBulk,
		ent.ApiResourceUpdate, ent.ApiResourceUpdateOne,
		ent.ApiResourceDelete,
		predicate.ApiResource,
		adminV1.ApiResource, ent.ApiResource,
	]
}

func NewApiResourceRepo(data *Data, logger log.Logger) *ApiResourceRepo {
	repo := &ApiResourceRepo{
		log:            log.NewHelper(log.With(logger, "module", "api-resource/repo/admin-service")),
		data:           data,
		mapper:         mapper.NewCopierMapper[adminV1.ApiResource, ent.ApiResource](),
		scopeConverter: mapper.NewEnumTypeConverter[adminV1.ApiResource_Scope, apiresource.Scope](adminV1.ApiResource_Scope_name, adminV1.ApiResource_Scope_value),
	}

	repo.init()

	return repo
}

func (r *ApiResourceRepo) init() {
	r.repository = entCrud.NewRepository[
		ent.ApiResourceQuery, ent.ApiResourceSelect,
		ent.ApiResourceCreate, ent.ApiResourceCreateBulk,
		ent.ApiResourceUpdate, ent.ApiResourceUpdateOne,
		ent.ApiResourceDelete,
		predicate.ApiResource,
		adminV1.ApiResource, ent.ApiResource,
	](r.mapper)

	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.scopeConverter.NewConverterPair())
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

	ret, err := r.repository.ListWithPaging(ctx, builder, builder.Clone(), req)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return &adminV1.ListApiResourceResponse{Total: 0, Items: nil}, nil
	}

	return &adminV1.ListApiResourceResponse{
		Total: ret.Total,
		Items: ret.Items,
	}, nil
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

	builder := r.data.db.Client().ApiResource.Query()

	var whereCond []func(s *sql.Selector)
	switch req.QueryBy.(type) {
	default:
	case *adminV1.GetApiResourceRequest_Id:
		whereCond = append(whereCond, apiresource.IDEQ(req.GetId()))
	}

	dto, err := r.repository.Get(ctx, builder, req.GetViewMask(), whereCond...)
	if err != nil {
		return nil, err
	}

	return dto, err
}

// GetApiResourceByEndpoint 根据路径和方法获取API资源
func (r *ApiResourceRepo) GetApiResourceByEndpoint(ctx context.Context, path, method string) (*adminV1.ApiResource, error) {
	if path == "" || method == "" {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().ApiResource.Query().
		Where(
			apiresource.PathEQ(path),
			apiresource.MethodEQ(method),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("api resource not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *ApiResourceRepo) Create(ctx context.Context, req *adminV1.CreateApiResourceRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().ApiResource.Create().
		SetNillableDescription(req.Data.Description).
		SetNillableModule(req.Data.Module).
		SetNillableModuleDescription(req.Data.ModuleDescription).
		SetNillableOperation(req.Data.Operation).
		SetNillablePath(req.Data.Path).
		SetNillableMethod(req.Data.Method).
		SetNillableScope(r.scopeConverter.ToEntity(req.Data.Scope)).
		SetNillableCreatedBy(req.Data.CreatedBy).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.GetData().GetId())
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
		exist, err := r.IsExist(ctx, req.GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &adminV1.CreateApiResourceRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			return r.Create(ctx, createReq)
		}
	}

	builder := r.data.db.Client().Debug().ApiResource.Update()
	err := r.repository.UpdateX(ctx, builder, req.Data, req.GetUpdateMask(),
		func(dto *adminV1.ApiResource) {
			builder.
				SetNillableDescription(req.Data.Description).
				SetNillableModule(req.Data.Module).
				SetNillableModuleDescription(req.Data.ModuleDescription).
				SetNillableOperation(req.Data.Operation).
				SetNillablePath(req.Data.Path).
				SetNillableMethod(req.Data.Method).
				SetNillableScope(r.scopeConverter.ToEntity(req.Data.Scope)).
				SetNillableUpdatedBy(req.Data.UpdatedBy).
				SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

			if req.Data.UpdatedAt == nil {
				builder.SetUpdatedAt(time.Now())
			}
		},
		func(s *sql.Selector) {
			s.Where(sql.EQ(apiresource.FieldID, req.GetId()))
		},
	)

	return err
}

func (r *ApiResourceRepo) Delete(ctx context.Context, req *adminV1.DeleteApiResourceRequest) error {
	if req == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Debug().ApiResource.Delete()

	_, err := r.repository.Delete(ctx, builder, func(s *sql.Selector) {
		s.Where(sql.EQ(apiresource.FieldID, req.GetId()))
	})
	if err != nil {
		r.log.Errorf("delete api resource failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("delete api resource failed")
	}

	return nil
}

// Truncate 清空表数据
func (r *ApiResourceRepo) Truncate(ctx context.Context) error {
	if _, err := r.data.db.Client().ApiResource.Delete().Exec(ctx); err != nil {
		r.log.Errorf("failed to truncate api_resources table: %s", err.Error())
		return adminV1.ErrorInternalServerError("truncate failed")
	}
	return nil
}
