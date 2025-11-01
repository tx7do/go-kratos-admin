package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/tx7do/go-utils/copierutil"
	entgo "github.com/tx7do/go-utils/entgo/query"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/adminoperationlog"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminOperationLogRepo struct {
	data *Data
	log  *log.Helper

	mapper *mapper.CopierMapper[adminV1.AdminOperationLog, ent.AdminOperationLog]
}

func NewAdminOperationLogRepo(data *Data, logger log.Logger) *AdminOperationLogRepo {
	repo := &AdminOperationLogRepo{
		log:    log.NewHelper(log.With(logger, "module", "admin-operation-log/repo/admin-service")),
		data:   data,
		mapper: mapper.NewCopierMapper[adminV1.AdminOperationLog, ent.AdminOperationLog](),
	}

	repo.init()

	return repo
}

func (r *AdminOperationLogRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.NewFloatSecondConverterPair())
}

func (r *AdminOperationLogRepo) NewFloatSecondConverterPair() []copier.TypeConverter {
	srcType := durationpb.New(0)
	dstType := trans.Ptr(float64(0))

	fromFn := timeutil.DurationpbToSecond
	toFn := timeutil.SecondToDurationpb

	return copierutil.NewGenericTypeConverterPair(srcType, dstType, fromFn, toFn)
}

func (r *AdminOperationLogRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().AdminOperationLog.Query()
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

func (r *AdminOperationLogRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminOperationLogResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().AdminOperationLog.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), adminoperationlog.FieldCreatedAt,
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

	dtos := make([]*adminV1.AdminOperationLog, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListAdminOperationLogResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *AdminOperationLogRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().AdminOperationLog.Query().
		Where(adminoperationlog.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *AdminOperationLogRepo) Get(ctx context.Context, req *adminV1.GetAdminOperationLogRequest) (*adminV1.AdminOperationLog, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().AdminOperationLog.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("admin operation log not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *AdminOperationLogRepo) Create(ctx context.Context, req *adminV1.CreateAdminOperationLogRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
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
		SetNillableCostTime(timeutil.DurationpbToSecond(req.Data.CostTime)).
		SetNillableUserID(req.Data.UserId).
		SetNillableUsername(req.Data.Username).
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
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("insert data failed")
	}

	return err
}
