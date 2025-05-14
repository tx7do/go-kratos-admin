package data

import (
	"context"
	"errors"
	"kratos-admin/pkg/middleware/auth"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/tenant"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type TenantRepo struct {
	data *Data
	log  *log.Helper
}

func NewTenantRepo(data *Data, logger log.Logger) *TenantRepo {
	l := log.NewHelper(log.With(logger, "module", "tenant/repo/admin-service"))
	return &TenantRepo{
		data: data,
		log:  l,
	}
}

func (r *TenantRepo) convertEntToProto(in *ent.Tenant) *userV1.Tenant {
	if in == nil {
		return nil
	}
	return &userV1.Tenant{
		Id:             trans.Ptr(in.ID),
		Name:           in.Name,
		Code:           in.Code,
		MemberCount:    in.MemberCount,
		CreateBy:       in.CreateBy,
		UpdateBy:       in.UpdateBy,
		Remark:         in.Remark,
		SubscriptionAt: timeutil.TimeToTimestamppb(in.SubscriptionAt),
		UnsubscribeAt:  timeutil.TimeToTimestamppb(in.UnsubscribeAt),
		Status:         (*string)(in.Status),
		CreateTime:     timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime:     timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime:     timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *TenantRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Tenant.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *TenantRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListTenantResponse, error) {
	builder := r.data.db.Client().Tenant.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), tenant.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析SELECT条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*userV1.Tenant, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	ret := userV1.ListTenantResponse{
		Total: uint32(count),
		Items: items,
	}

	return &ret, err
}

func (r *TenantRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().Tenant.Query().
		Where(tenant.IDEQ(id)).
		Exist(ctx)
}

func (r *TenantRepo) Get(ctx context.Context, req *userV1.GetTenantRequest) (*userV1.Tenant, error) {
	ret, err := r.data.db.Client().Tenant.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorTenantNotFound("tenant not found")
		}

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *TenantRepo) Create(ctx context.Context, req *userV1.CreateTenantRequest, operator *auth.UserTokenPayload) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().Tenant.Create().
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableMemberCount(req.Data.MemberCount).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*tenant.Status)(req.Data.Status)).
		SetNillableCreateBy(trans.Ptr(operator.UserId)).
		SetNillableSubscriptionAt(timeutil.TimestamppbToTime(req.Data.SubscriptionAt)).
		SetNillableUnsubscribeAt(timeutil.TimestamppbToTime(req.Data.UnsubscribeAt)).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *TenantRepo) Update(ctx context.Context, req *userV1.UpdateTenantRequest, operator *auth.UserTokenPayload) error {
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
			return r.Create(ctx, &userV1.CreateTenantRequest{Data: req.Data}, operator)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().Tenant.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableMemberCount(req.Data.MemberCount).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*tenant.Status)(req.Data.Status)).
		SetNillableCreateBy(trans.Ptr(operator.UserId)).
		SetNillableSubscriptionAt(timeutil.TimestamppbToTime(req.Data.SubscriptionAt)).
		SetNillableUnsubscribeAt(timeutil.TimestamppbToTime(req.Data.UnsubscribeAt)).
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

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *TenantRepo) Delete(ctx context.Context, req *userV1.DeleteTenantRequest) (bool, error) {
	err := r.data.db.Client().Tenant.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
