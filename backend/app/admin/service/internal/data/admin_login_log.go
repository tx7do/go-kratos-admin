package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	"github.com/tx7do/go-utils/copierutil"
	entgo "github.com/tx7do/go-utils/entgo/query"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/adminloginlog"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminLoginLogRepo struct {
	data         *Data
	log          *log.Helper
	copierOption copier.Option
}

func NewAdminLoginLogRepo(data *Data, logger log.Logger) *AdminLoginLogRepo {
	repo := &AdminLoginLogRepo{
		log:  log.NewHelper(log.With(logger, "module", "admin-login-log/repo/admin-service")),
		data: data,
	}

	repo.init()

	return repo
}

func (r *AdminLoginLogRepo) init() {
	r.copierOption = copier.Option{
		Converters: []copier.TypeConverter{},
	}

	r.copierOption.Converters = append(r.copierOption.Converters, copierutil.NewTimeStringConverterPair()...)
	r.copierOption.Converters = append(r.copierOption.Converters, copierutil.NewTimeTimestamppbConverterPair()...)
}

func (r *AdminLoginLogRepo) toProto(in *ent.AdminLoginLog) *adminV1.AdminLoginLog {
	if in == nil {
		return nil
	}

	var out adminV1.AdminLoginLog
	_ = copier.CopyWithOption(&out, in, r.copierOption)

	//out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)
	//out.LoginTime = timeutil.TimeToTimestamppb(in.LoginTime)

	return &out
}

func (r *AdminLoginLogRepo) toEnt(in *adminV1.AdminLoginLog) *ent.AdminLoginLog {
	if in == nil {
		return nil
	}

	var out ent.AdminLoginLog
	_ = copier.CopyWithOption(&out, in, r.copierOption)

	return &out
}

func (r *AdminLoginLogRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().AdminLoginLog.Query()
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

func (r *AdminLoginLogRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminLoginLogResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().AdminLoginLog.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), adminloginlog.FieldCreateTime,
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

	items := make([]*adminV1.AdminLoginLog, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListAdminLoginLogResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *AdminLoginLogRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().AdminLoginLog.Query().
		Where(adminloginlog.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *AdminLoginLogRepo) Get(ctx context.Context, req *adminV1.GetAdminLoginLogRequest) (*adminV1.AdminLoginLog, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().AdminLoginLog.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("admin login log not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *AdminLoginLogRepo) Create(ctx context.Context, req *adminV1.CreateAdminLoginLogRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().AdminLoginLog.
		Create().
		SetNillableLoginIP(req.Data.LoginIp).
		SetNillableLoginMAC(req.Data.LoginMac).
		SetNillableUserAgent(req.Data.UserAgent).
		SetNillableBrowserName(req.Data.BrowserName).
		SetNillableBrowserVersion(req.Data.BrowserVersion).
		SetNillableClientID(req.Data.ClientId).
		SetNillableClientName(req.Data.ClientName).
		SetNillableOsName(req.Data.OsName).
		SetNillableOsVersion(req.Data.OsVersion).
		SetNillableUserID(req.Data.UserId).
		SetNillableUsername(req.Data.Username).
		SetNillableStatusCode(req.Data.StatusCode).
		SetNillableSuccess(req.Data.Success).
		SetNillableReason(req.Data.Reason).
		SetNillableLocation(req.Data.Location).
		SetNillableLoginTime(timeutil.TimestamppbToTime(req.Data.LoginTime)).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.LoginTime == nil {
		builder.SetLoginTime(time.Now())
	}

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}
