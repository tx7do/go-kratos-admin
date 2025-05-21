package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	entgo "github.com/tx7do/go-utils/entgo/query"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/adminloginlog"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type AdminLoginLogRepo struct {
	data *Data
	log  *log.Helper
}

func NewAdminLoginLogRepo(data *Data, logger log.Logger) *AdminLoginLogRepo {
	l := log.NewHelper(log.With(logger, "module", "admin-login-log/repo/admin-service"))
	return &AdminLoginLogRepo{
		data: data,
		log:  l,
	}
}

func (r *AdminLoginLogRepo) toProto(in *ent.AdminLoginLog) *systemV1.AdminLoginLog {
	if in == nil {
		return nil
	}

	var out systemV1.AdminLoginLog
	_ = copier.Copy(&out, in)

	out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)
	out.LoginTime = timeutil.TimeToTimestamppb(in.LoginTime)

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
		return 0, systemV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *AdminLoginLogRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListAdminLoginLogResponse, error) {
	if req == nil {
		return nil, systemV1.ErrorBadRequest("invalid parameter")
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
		return nil, systemV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, systemV1.ErrorInternalServerError("query list failed")
	}

	items := make([]*systemV1.AdminLoginLog, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListAdminLoginLogResponse{
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
		return false, systemV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *AdminLoginLogRepo) Get(ctx context.Context, req *systemV1.GetAdminLoginLogRequest) (*systemV1.AdminLoginLog, error) {
	if req == nil {
		return nil, systemV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().AdminLoginLog.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorNotFound("admin login log not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, systemV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *AdminLoginLogRepo) Create(ctx context.Context, req *systemV1.CreateAdminLoginLogRequest) error {
	if req == nil || req.Data == nil {
		return systemV1.ErrorBadRequest("invalid parameter")
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
		SetNillableUserName(req.Data.UserName).
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
		return systemV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}
