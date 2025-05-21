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
	"kratos-admin/app/admin/service/internal/data/ent/adminoperationlog"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type AdminOperationLogRepo struct {
	data *Data
	log  *log.Helper
}

func NewAdminOperationLogRepo(data *Data, logger log.Logger) *AdminOperationLogRepo {
	l := log.NewHelper(log.With(logger, "module", "admin-operation-log/repo/admin-service"))
	return &AdminOperationLogRepo{
		data: data,
		log:  l,
	}
}

func (r *AdminOperationLogRepo) toProto(in *ent.AdminOperationLog) *systemV1.AdminOperationLog {
	if in == nil {
		return nil
	}

	var out systemV1.AdminOperationLog
	_ = copier.Copy(&out, in)

	out.CostTime = timeutil.SecondToDurationpb(in.CostTime)
	out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)

	return &out
}

func (r *AdminOperationLogRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().AdminOperationLog.Query()
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

func (r *AdminOperationLogRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListAdminOperationLogResponse, error) {
	if req == nil {
		return nil, systemV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().AdminOperationLog.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), adminoperationlog.FieldCreateTime,
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

	items := make([]*systemV1.AdminOperationLog, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListAdminOperationLogResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *AdminOperationLogRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().AdminOperationLog.Query().
		Where(adminoperationlog.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, systemV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *AdminOperationLogRepo) Get(ctx context.Context, req *systemV1.GetAdminOperationLogRequest) (*systemV1.AdminOperationLog, error) {
	if req == nil {
		return nil, systemV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().AdminOperationLog.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorNotFound("admin operation log not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, systemV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *AdminOperationLogRepo) Create(ctx context.Context, req *systemV1.CreateAdminOperationLogRequest) error {
	if req == nil || req.Data == nil {
		return systemV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().AdminOperationLog.
		Create().
		SetNillableRequestID(req.Data.RequestId).
		SetNillableMethod(req.Data.Method).
		SetNillableOperation(req.Data.Operation).
		SetNillablePath(req.Data.Path).
		SetNillableReferer(req.Data.Referer).
		SetNillableRequestURI(req.Data.RequestUri).
		SetNillableRequestBody(req.Data.RequestBody).
		SetNillableRequestHeader(req.Data.RequestHeader).
		SetNillableResponse(req.Data.Response).
		SetNillableCostTime(timeutil.DurationpbSecond(req.Data.CostTime)).
		SetNillableUserID(req.Data.UserId).
		SetNillableUserName(req.Data.UserName).
		SetNillableClientIP(req.Data.ClientIp).
		SetNillableUserAgent(req.Data.UserAgent).
		SetNillableBrowserName(req.Data.BrowserName).
		SetNillableBrowserVersion(req.Data.BrowserVersion).
		SetNillableClientID(req.Data.ClientId).
		SetNillableClientName(req.Data.ClientName).
		SetNillableOsName(req.Data.OsName).
		SetNillableOsVersion(req.Data.OsVersion).
		SetNillableStatusCode(req.Data.StatusCode).
		SetNillableSuccess(req.Data.Success).
		SetNillableReason(req.Data.Reason).
		SetNillableLocation(req.Data.Location).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return systemV1.ErrorInternalServerError("insert data failed")
	}

	return err
}
