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
	"github.com/tx7do/go-utils/trans"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/adminloginrestriction"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminLoginRestrictionRepo struct {
	data         *Data
	log          *log.Helper
	copierOption copier.Option
}

func NewAdminLoginRestrictionRepo(data *Data, logger log.Logger) *AdminLoginRestrictionRepo {
	repo := &AdminLoginRestrictionRepo{
		log:  log.NewHelper(log.With(logger, "module", "admin-login-restriction/repo/admin-service")),
		data: data,
	}

	repo.init()

	return repo
}

func (r *AdminLoginRestrictionRepo) init() {
	r.copierOption = copier.Option{
		Converters: []copier.TypeConverter{},
	}

	r.copierOption.Converters = append(r.copierOption.Converters, copierutil.NewTimeStringConverterPair()...)
	r.copierOption.Converters = append(r.copierOption.Converters, copierutil.NewTimeTimestamppbConverterPair()...)
	r.copierOption.Converters = append(r.copierOption.Converters, r.NewTypeConverterPair()...)
	r.copierOption.Converters = append(r.copierOption.Converters, r.NewMethodConverterPair()...)
}

func (r *AdminLoginRestrictionRepo) NewTypeConverterPair() []copier.TypeConverter {
	srcType := trans.Ptr(adminV1.AdminLoginRestrictionType(0))
	dstType := trans.Ptr(adminloginrestriction.Type(""))

	fromFn := r.toEntType
	toFn := r.toProtoType

	return copierutil.NewGenericTypeConverterPair(srcType, dstType, fromFn, toFn)
}

func (r *AdminLoginRestrictionRepo) NewMethodConverterPair() []copier.TypeConverter {
	srcType := trans.Ptr(adminV1.AdminLoginRestrictionMethod(0))
	dstType := trans.Ptr(adminloginrestriction.Method(""))

	fromFn := r.toEntMethod
	toFn := r.toProtoMethod

	return copierutil.NewGenericTypeConverterPair(srcType, dstType, fromFn, toFn)
}

func (r *AdminLoginRestrictionRepo) toEntType(in *adminV1.AdminLoginRestrictionType) *adminloginrestriction.Type {
	if in == nil {
		return nil
	}

	find, ok := adminV1.AdminLoginRestrictionType_name[int32(*in)]
	if !ok {
		return nil
	}

	return (*adminloginrestriction.Type)(trans.Ptr(find))
}

func (r *AdminLoginRestrictionRepo) toProtoType(in *adminloginrestriction.Type) *adminV1.AdminLoginRestrictionType {
	if in == nil {
		return nil
	}

	find, ok := adminV1.AdminLoginRestrictionType_value[string(*in)]
	if !ok {
		return nil
	}

	return (*adminV1.AdminLoginRestrictionType)(trans.Ptr(find))
}

func (r *AdminLoginRestrictionRepo) toEntMethod(in *adminV1.AdminLoginRestrictionMethod) *adminloginrestriction.Method {
	if in == nil {
		return nil
	}

	find, ok := adminV1.AdminLoginRestrictionMethod_name[int32(*in)]
	if !ok {
		return nil
	}

	return (*adminloginrestriction.Method)(trans.Ptr(find))
}

func (r *AdminLoginRestrictionRepo) toProtoMethod(in *adminloginrestriction.Method) *adminV1.AdminLoginRestrictionMethod {
	if in == nil {
		return nil
	}

	find, ok := adminV1.AdminLoginRestrictionMethod_value[string(*in)]
	if !ok {
		return nil
	}

	return (*adminV1.AdminLoginRestrictionMethod)(trans.Ptr(find))
}

func (r *AdminLoginRestrictionRepo) toProto(in *ent.AdminLoginRestriction) *adminV1.AdminLoginRestriction {
	if in == nil {
		return nil
	}

	var out adminV1.AdminLoginRestriction
	_ = copier.CopyWithOption(&out, in, r.copierOption)

	//out.Type = r.toProtoType(in.Type)
	//out.Method = r.toProtoMethod(in.Method)
	//out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)
	//out.UpdateTime = timeutil.TimeToTimeString(in.UpdateTime)
	//out.DeleteTime = timeutil.TimeToTimeString(in.DeleteTime)

	return &out
}

func (r *AdminLoginRestrictionRepo) toEnt(in *adminV1.AdminLoginRestriction) *ent.AdminLoginRestriction {
	if in == nil {
		return nil
	}

	var out ent.AdminLoginRestriction
	_ = copier.CopyWithOption(&out, in, r.copierOption)

	return &out
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
		req.GetOrderBy(), adminloginrestriction.FieldCreateTime,
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

	items := make([]*adminV1.AdminLoginRestriction, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
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

	ret, err := r.data.db.Client().AdminLoginRestriction.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("admin login restriction not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
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
