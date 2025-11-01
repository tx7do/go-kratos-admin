package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/copierutil"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/adminloginrestriction"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminLoginRestrictionRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[adminV1.AdminLoginRestriction, ent.AdminLoginRestriction]
	typeConverter   *mapper.EnumTypeConverter[adminV1.AdminLoginRestriction_Type, adminloginrestriction.Type]
	methodConverter *mapper.EnumTypeConverter[adminV1.AdminLoginRestriction_Method, adminloginrestriction.Method]
}

func NewAdminLoginRestrictionRepo(data *Data, logger log.Logger) *AdminLoginRestrictionRepo {
	repo := &AdminLoginRestrictionRepo{
		log:             log.NewHelper(log.With(logger, "module", "admin-login-restriction/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[adminV1.AdminLoginRestriction, ent.AdminLoginRestriction](),
		typeConverter:   mapper.NewEnumTypeConverter[adminV1.AdminLoginRestriction_Type, adminloginrestriction.Type](adminV1.AdminLoginRestriction_Type_name, adminV1.AdminLoginRestriction_Type_value),
		methodConverter: mapper.NewEnumTypeConverter[adminV1.AdminLoginRestriction_Method, adminloginrestriction.Method](adminV1.AdminLoginRestriction_Method_name, adminV1.AdminLoginRestriction_Method_value),
	}

	repo.init()

	return repo
}

func (r *AdminLoginRestrictionRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.typeConverter.NewConverterPair())
	r.mapper.AppendConverters(r.methodConverter.NewConverterPair())
}

func (r *AdminLoginRestrictionRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().AdminLoginRestriction.Query()
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

func (r *AdminLoginRestrictionRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminLoginRestrictionResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().AdminLoginRestriction.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), adminloginrestriction.FieldCreatedAt,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, adminV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, adminV1.ErrorInternalServerError("query list failed")
	}

	dtos := make([]*adminV1.AdminLoginRestriction, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListAdminLoginRestrictionResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *AdminLoginRestrictionRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().AdminLoginRestriction.Query().
		Where(adminloginrestriction.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *AdminLoginRestrictionRepo) Get(ctx context.Context, req *adminV1.GetAdminLoginRestrictionRequest) (*adminV1.AdminLoginRestriction, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().AdminLoginRestriction.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("admin login restriction not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *AdminLoginRestrictionRepo) Create(ctx context.Context, req *adminV1.CreateAdminLoginRestrictionRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid request")
	}

	builder := r.data.db.Client().AdminLoginRestriction.Create().
		SetNillableTargetID(req.Data.TargetId).
		SetNillableType(r.typeConverter.ToEntity(req.Data.Type)).
		SetNillableMethod(r.methodConverter.ToEntity(req.Data.Method)).
		SetNillableValue(req.Data.Value).
		SetNillableReason(req.Data.Reason).
		SetNillableCreatedBy(req.Data.CreatedBy).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *AdminLoginRestrictionRepo) Update(ctx context.Context, req *adminV1.UpdateAdminLoginRestrictionRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &adminV1.CreateAdminLoginRestrictionRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
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

	builder := r.data.db.Client().AdminLoginRestriction.UpdateOneID(req.Data.GetId()).
		SetNillableTargetID(req.Data.TargetId).
		SetNillableType(r.typeConverter.ToEntity(req.Data.Type)).
		SetNillableMethod(r.methodConverter.ToEntity(req.Data.Method)).
		SetNillableValue(req.Data.Value).
		SetNillableReason(req.Data.Reason).
		SetNillableUpdatedBy(req.Data.UpdatedBy).
		SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

	if req.Data.UpdatedAt == nil {
		builder.SetUpdatedAt(time.Now())
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

func (r *AdminLoginRestrictionRepo) Delete(ctx context.Context, req *adminV1.DeleteAdminLoginRestrictionRequest) error {
	if req == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().AdminLoginRestriction.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return adminV1.ErrorNotFound("admin login restriction not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return adminV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
