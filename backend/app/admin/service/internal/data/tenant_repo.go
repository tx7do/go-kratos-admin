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

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/tenant"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type TenantRepo struct {
	data *Data
	log  *log.Helper

	mapper *mapper.CopierMapper[ent.Tenant, userV1.Tenant]
}

func NewTenantRepo(data *Data, logger log.Logger) *TenantRepo {
	repo := &TenantRepo{
		log:    log.NewHelper(log.With(logger, "module", "tenant/repo/admin-service")),
		data:   data,
		mapper: mapper.NewCopierMapper[ent.Tenant, userV1.Tenant](),
	}

	repo.init()

	return repo
}

func (r *TenantRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
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

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), tenant.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, userV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query list failed")
	}

	models := make([]*userV1.Tenant, 0, len(results))
	for _, dto := range results {
		model := r.mapper.ToModel(dto)
		models = append(models, model)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListTenantResponse{
		Total: uint32(count),
		Items: models,
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

	dto, err := r.data.db.Client().Tenant.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorTenantNotFound("tenant not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToModel(dto), nil
}

func (r *TenantRepo) Create(ctx context.Context, req *userV1.CreateTenantRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Tenant.Create().
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableMemberCount(req.Data.MemberCount).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*tenant.Status)(req.Data.Status)).
		SetNillableSubscriptionAt(timeutil.TimestamppbToTime(req.Data.SubscriptionAt)).
		SetNillableUnsubscribeAt(timeutil.TimestamppbToTime(req.Data.UnsubscribeAt))

	builder.SetNillableCreateBy(req.Data.CreateBy)

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	} else {
		builder.SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return userV1.ErrorInternalServerError("insert data failed")
	}

	return nil
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
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return userV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().Tenant.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableMemberCount(req.Data.MemberCount).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*tenant.Status)(req.Data.Status)).
		SetNillableSubscriptionAt(timeutil.TimestamppbToTime(req.Data.SubscriptionAt)).
		SetNillableUnsubscribeAt(timeutil.TimestamppbToTime(req.Data.UnsubscribeAt))

	builder.SetNillableUpdateBy(req.Data.UpdateBy)

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	} else {
		builder.SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))
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

func (r *TenantRepo) GetTenantByTenantName(ctx context.Context, userName string) (*userV1.Tenant, error) {
	ret, err := r.data.db.Client().Tenant.Query().
		Where(tenant.NameEQ(userName)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorNotFound("tenant not found")
		}

		r.log.Errorf("query user data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToModel(ret), nil
}

func (r *TenantRepo) GetTenantByTenantCode(ctx context.Context, code string) (*userV1.Tenant, error) {
	ret, err := r.data.db.Client().Tenant.Query().
		Where(tenant.CodeEQ(code)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorNotFound("tenant not found")
		}

		r.log.Errorf("query user data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToModel(ret), nil
}
