package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/proto"

	"github.com/tx7do/go-utils/copierutil"
	entgoQuery "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/tenant"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type TenantRepo struct {
	data *Data
	log  *log.Helper

	mapper               *mapper.CopierMapper[userV1.Tenant, ent.Tenant]
	statusConverter      *mapper.EnumTypeConverter[userV1.Tenant_Status, tenant.Status]
	typeConverter        *mapper.EnumTypeConverter[userV1.Tenant_Type, tenant.Type]
	auditStatusConverter *mapper.EnumTypeConverter[userV1.Tenant_AuditStatus, tenant.AuditStatus]
}

func NewTenantRepo(data *Data, logger log.Logger) *TenantRepo {
	repo := &TenantRepo{
		log:                  log.NewHelper(log.With(logger, "module", "tenant/repo/admin-service")),
		data:                 data,
		mapper:               mapper.NewCopierMapper[userV1.Tenant, ent.Tenant](),
		statusConverter:      mapper.NewEnumTypeConverter[userV1.Tenant_Status, tenant.Status](userV1.Tenant_Status_name, userV1.Tenant_Status_value),
		typeConverter:        mapper.NewEnumTypeConverter[userV1.Tenant_Type, tenant.Type](userV1.Tenant_Type_name, userV1.Tenant_Type_value),
		auditStatusConverter: mapper.NewEnumTypeConverter[userV1.Tenant_AuditStatus, tenant.AuditStatus](userV1.Tenant_AuditStatus_name, userV1.Tenant_AuditStatus_value),
	}

	repo.init()

	return repo
}

func (r *TenantRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
	r.mapper.AppendConverters(r.typeConverter.NewConverterPair())
	r.mapper.AppendConverters(r.auditStatusConverter.NewConverterPair())
}

func (r *TenantRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Tenant.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, userV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *TenantRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListTenantResponse, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Tenant.Query()

	err, whereSelectors, querySelectors := entgoQuery.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), tenant.FieldCreatedAt,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, userV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query list failed")
	}

	dtos := make([]*userV1.Tenant, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListTenantResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *TenantRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Tenant.Query().
		Where(tenant.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, userV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *TenantRepo) Get(ctx context.Context, req *userV1.GetTenantRequest) (*userV1.Tenant, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Tenant.Query()

	builder.Where(tenant.IDEQ(req.GetId()))

	entgoQuery.ApplyFieldMaskToBuilder(builder, req.ViewMask)

	entity, err := builder.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorTenantNotFound("tenant not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *TenantRepo) Create(ctx context.Context, data *userV1.Tenant) (*userV1.Tenant, error) {
	if data == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Tenant.Create().
		SetNillableName(data.Name).
		SetNillableCode(data.Code).
		SetNillableLogoURL(data.LogoUrl).
		SetNillableRemark(data.Remark).
		SetNillableIndustry(data.Industry).
		SetNillableAdminUserID(data.AdminUserId).
		SetNillableStatus(r.statusConverter.ToEntity(data.Status)).
		SetNillableType(r.typeConverter.ToEntity(data.Type)).
		SetNillableAuditStatus(r.auditStatusConverter.ToEntity(data.AuditStatus)).
		SetNillableLastLoginTime(timeutil.TimestamppbToTime(data.LastLoginTime)).
		SetNillableLastLoginIP(data.LastLoginIp).
		SetNillableSubscriptionPlan(data.SubscriptionPlan).
		SetNillableExpiredAt(timeutil.TimestamppbToTime(data.ExpiredAt)).
		SetNillableSubscriptionAt(timeutil.TimestamppbToTime(data.SubscriptionAt)).
		SetNillableUnsubscribeAt(timeutil.TimestamppbToTime(data.UnsubscribeAt))

	builder.SetNillableCreatedBy(data.CreatedBy)

	if data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	} else {
		builder.SetNillableCreatedAt(timeutil.TimestamppbToTime(data.CreatedAt))
	}

	if data.Id != nil {
		builder.SetID(data.GetId())
	}

	if ret, err := builder.Save(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("insert data failed")
	} else {
		return r.mapper.ToDTO(ret), nil
	}
}

func (r *TenantRepo) Update(ctx context.Context, req *userV1.UpdateTenantRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &userV1.CreateTenantRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			_, err = r.Create(ctx, createReq.Data)
			return err
		}
	}

	if err := fieldmaskutil.FilterByFieldMask(trans.Ptr(proto.Message(req.GetData())), req.UpdateMask); err != nil {
		r.log.Errorf("invalid field mask [%v], error: %s", req.UpdateMask, err.Error())
		return userV1.ErrorBadRequest("invalid field mask")
	}

	builder := r.data.db.Client().Tenant.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableLogoURL(req.Data.LogoUrl).
		SetNillableRemark(req.Data.Remark).
		SetNillableIndustry(req.Data.Industry).
		SetNillableAdminUserID(req.Data.AdminUserId).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableType(r.typeConverter.ToEntity(req.Data.Type)).
		SetNillableAuditStatus(r.auditStatusConverter.ToEntity(req.Data.AuditStatus)).
		SetNillableLastLoginTime(timeutil.TimestamppbToTime(req.Data.LastLoginTime)).
		SetNillableLastLoginIP(req.Data.LastLoginIp).
		SetNillableSubscriptionPlan(req.Data.SubscriptionPlan).
		SetNillableExpiredAt(timeutil.TimestamppbToTime(req.Data.ExpiredAt)).
		SetNillableSubscriptionAt(timeutil.TimestamppbToTime(req.Data.SubscriptionAt)).
		SetNillableUnsubscribeAt(timeutil.TimestamppbToTime(req.Data.UnsubscribeAt))

	builder.SetNillableUpdatedBy(req.Data.UpdatedBy)

	if req.Data.UpdatedAt == nil {
		builder.SetUpdatedAt(time.Now())
	} else {
		builder.SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))
	}

	entgoUpdate.ApplyNilFieldMask(proto.Message(req.GetData()), req.UpdateMask, builder)

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return userV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *TenantRepo) Delete(ctx context.Context, req *userV1.DeleteTenantRequest) error {
	if req == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().Tenant.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return userV1.ErrorNotFound("tenant not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return userV1.ErrorInternalServerError("delete failed")
	}

	return nil
}

// GetTenantByTenantName gets tenant by tenant name.
func (r *TenantRepo) GetTenantByTenantName(ctx context.Context, userName string) (*userV1.Tenant, error) {
	entity, err := r.data.db.Client().Tenant.Query().
		Where(tenant.NameEQ(userName)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorNotFound("tenant not found")
		}

		r.log.Errorf("query user data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

// GetTenantByTenantCode gets tenant by tenant code.
func (r *TenantRepo) GetTenantByTenantCode(ctx context.Context, code string) (*userV1.Tenant, error) {
	entity, err := r.data.db.Client().Tenant.Query().
		Where(tenant.CodeEQ(code)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorNotFound("tenant not found")
		}

		r.log.Errorf("query user data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

// TenantExists checks if a tenant with the given username exists.
func (r *TenantRepo) TenantExists(ctx context.Context, req *userV1.TenantExistsRequest) (*userV1.TenantExistsResponse, error) {
	exist, err := r.data.db.Client().Tenant.Query().
		Where(tenant.CodeEQ(req.GetCode())).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query exist failed")
	}

	return &userV1.TenantExistsResponse{
		Exist: exist,
	}, nil
}

// GetTenantsByIds gets tenants by a list of IDs.
func (r *TenantRepo) GetTenantsByIds(ctx context.Context, ids []uint32) ([]*userV1.Tenant, error) {
	if len(ids) == 0 {
		return []*userV1.Tenant{}, nil
	}

	entities, err := r.data.db.Client().Tenant.Query().
		Where(tenant.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query tenant by ids failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query tenant by ids failed")
	}

	dtos := make([]*userV1.Tenant, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	return dtos, nil
}
