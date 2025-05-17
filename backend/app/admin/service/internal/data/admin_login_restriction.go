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

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/adminloginrestriction"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminLoginRestrictionRepo struct {
	data *Data
	log  *log.Helper
}

func NewAdminLoginRestrictionRepo(data *Data, logger log.Logger) *AdminLoginRestrictionRepo {
	l := log.NewHelper(log.With(logger, "module", "admin-login-restriction/repo/admin-service"))
	return &AdminLoginRestrictionRepo{
		data: data,
		log:  l,
	}
}

func (r *AdminLoginRestrictionRepo) toEntType(in *adminV1.AdminLoginRestrictionType) *adminloginrestriction.Type {
	if in == nil {
		return nil
	}

	switch *in {
	case adminV1.AdminLoginRestrictionType_LOGIN_RESTRICTION_TYPE_UNSPECIFIED:
		return trans.Ptr(adminloginrestriction.TypeUnspecified)

	case adminV1.AdminLoginRestrictionType_LOGIN_RESTRICTION_TYPE_BLACKLIST:
		return trans.Ptr(adminloginrestriction.TypeBlacklist)

	case adminV1.AdminLoginRestrictionType_LOGIN_RESTRICTION_TYPE_WHITELIST:
		return trans.Ptr(adminloginrestriction.TypeWhitelist)

	default:
		return nil
	}
}

func (r *AdminLoginRestrictionRepo) toProtoType(in *adminloginrestriction.Type) *adminV1.AdminLoginRestrictionType {
	if in == nil {
		return nil
	}

	switch *in {
	case adminloginrestriction.TypeUnspecified:
		return trans.Ptr(adminV1.AdminLoginRestrictionType_LOGIN_RESTRICTION_TYPE_UNSPECIFIED)

	case adminloginrestriction.TypeBlacklist:
		return trans.Ptr(adminV1.AdminLoginRestrictionType_LOGIN_RESTRICTION_TYPE_BLACKLIST)

	case adminloginrestriction.TypeWhitelist:
		return trans.Ptr(adminV1.AdminLoginRestrictionType_LOGIN_RESTRICTION_TYPE_WHITELIST)

	default:
		return nil
	}
}

func (r *AdminLoginRestrictionRepo) toEntMethod(in *adminV1.AdminLoginRestrictionMethod) *adminloginrestriction.Method {
	if in == nil {
		return nil
	}

	switch *in {
	case adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_UNSPECIFIED:
		return trans.Ptr(adminloginrestriction.MethodUnspecified)

	case adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_IP:
		return trans.Ptr(adminloginrestriction.MethodIp)

	case adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_MAC:
		return trans.Ptr(adminloginrestriction.MethodMac)

	case adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_REGION:
		return trans.Ptr(adminloginrestriction.MethodRegion)

	case adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_TIME:
		return trans.Ptr(adminloginrestriction.MethodTime)

	case adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_DEVICE:
		return trans.Ptr(adminloginrestriction.MethodDevice)

	default:
		return nil
	}
}

func (r *AdminLoginRestrictionRepo) toProtoMethod(in *adminloginrestriction.Method) *adminV1.AdminLoginRestrictionMethod {
	if in == nil {
		return nil
	}

	switch *in {
	case adminloginrestriction.MethodUnspecified:
		return trans.Ptr(adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_UNSPECIFIED)

	case adminloginrestriction.MethodIp:
		return trans.Ptr(adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_IP)

	case adminloginrestriction.MethodMac:
		return trans.Ptr(adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_MAC)

	case adminloginrestriction.MethodRegion:
		return trans.Ptr(adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_REGION)

	case adminloginrestriction.MethodTime:
		return trans.Ptr(adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_TIME)

	case adminloginrestriction.MethodDevice:
		return trans.Ptr(adminV1.AdminLoginRestrictionMethod_LOGIN_RESTRICTION_METHOD_DEVICE)

	default:
		return nil
	}
}

func (r *AdminLoginRestrictionRepo) convertEntToProto(in *ent.AdminLoginRestriction) *adminV1.AdminLoginRestriction {
	if in == nil {
		return nil
	}
	return &adminV1.AdminLoginRestriction{
		Id:         trans.Ptr(in.ID),
		AdminId:    in.AdminID,
		Type:       r.toProtoType(in.Type),
		Method:     r.toProtoMethod(in.Method),
		Value:      in.Value,
		Reason:     in.Reason,
		CreateBy:   in.CreateBy,
		UpdateBy:   in.UpdateBy,
		CreateTime: timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime: timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *AdminLoginRestrictionRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().AdminLoginRestriction.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *AdminLoginRestrictionRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminLoginRestrictionResponse, error) {
	builder := r.data.db.Client().AdminLoginRestriction.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), adminloginrestriction.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("list query param error [%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*adminV1.AdminLoginRestriction, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListAdminLoginRestrictionResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *AdminLoginRestrictionRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().AdminLoginRestriction.Query().
		Where(adminloginrestriction.IDEQ(id)).
		Exist(ctx)
}

func (r *AdminLoginRestrictionRepo) Get(ctx context.Context, req *adminV1.GetAdminLoginRestrictionRequest) (*adminV1.AdminLoginRestriction, error) {
	ret, err := r.data.db.Client().AdminLoginRestriction.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorResourceNotFound("admin login restriction not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *AdminLoginRestrictionRepo) Create(ctx context.Context, req *adminV1.CreateAdminLoginRestrictionRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid request")
	}

	builder := r.data.db.Client().AdminLoginRestriction.Create().
		SetNillableAdminID(req.Data.AdminId).
		SetNillableType(r.toEntType(req.Data.Type)).
		SetNillableMethod(r.toEntMethod(req.Data.Method)).
		SetNillableValue(req.Data.Value).
		SetNillableReason(req.Data.Reason).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
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

	builder := r.data.db.Client().AdminLoginRestriction.UpdateOneID(req.Data.GetId()).
		SetNillableAdminID(req.Data.AdminId).
		SetNillableType(r.toEntType(req.Data.Type)).
		SetNillableMethod(r.toEntMethod(req.Data.Method)).
		SetNillableValue(req.Data.Value).
		SetNillableReason(req.Data.Reason).
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

func (r *AdminLoginRestrictionRepo) Delete(ctx context.Context, req *adminV1.DeleteAdminLoginRestrictionRequest) error {
	if err := r.data.db.Client().AdminLoginRestriction.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return adminV1.ErrorResourceNotFound("admin login restriction not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return adminV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
