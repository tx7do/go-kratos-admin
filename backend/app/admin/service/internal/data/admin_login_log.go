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
	"kratos-admin/app/admin/service/internal/data/ent/adminloginlog"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
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

func (r *AdminLoginLogRepo) convertEntToProto(in *ent.AdminLoginLog) *systemV1.AdminLoginLog {
	if in == nil {
		return nil
	}
	return &systemV1.AdminLoginLog{
		Id:             trans.Ptr(in.ID),
		LoginIp:        in.LoginIP,
		LoginMac:       in.LoginMAC,
		LoginTime:      timeutil.TimeToTimestamppb(in.LoginTime),
		UserAgent:      in.UserAgent,
		BrowserName:    in.BrowserName,
		BrowserVersion: in.BrowserVersion,
		ClientId:       in.ClientID,
		ClientName:     in.ClientName,
		OsName:         in.OsName,
		OsVersion:      in.OsVersion,
		UserId:         in.UserID,
		UserName:       in.UserName,
		StatusCode:     in.StatusCode,
		Success:        in.Success,
		Reason:         in.Reason,
		Location:       in.Location,
		CreateTime:     timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime:     timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime:     timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *AdminLoginLogRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().AdminLoginLog.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *AdminLoginLogRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListAdminLoginLogResponse, error) {
	builder := r.data.db.Client().AdminLoginLog.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), adminloginlog.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*systemV1.AdminLoginLog, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
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
	return r.data.db.Client().AdminLoginLog.Query().
		Where(adminloginlog.IDEQ(id)).
		Exist(ctx)
}

func (r *AdminLoginLogRepo) Get(ctx context.Context, req *systemV1.GetAdminLoginLogRequest) (*systemV1.AdminLoginLog, error) {
	ret, err := r.data.db.Client().AdminLoginLog.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("admin login log not found")
		}
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *AdminLoginLogRepo) Create(ctx context.Context, req *systemV1.CreateAdminLoginLogRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
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
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.LoginTime == nil {
		builder.SetLoginTime(time.Now())
	}

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

func (r *AdminLoginLogRepo) Update(ctx context.Context, req *systemV1.UpdateAdminLoginLogRequest) error {
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
			return r.Create(ctx, &systemV1.CreateAdminLoginLogRequest{Data: req.Data, OperatorId: req.OperatorId})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().AdminLoginLog.
		UpdateOneID(req.Data.GetId()).
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
		SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))

	if req.Data.LoginTime == nil {
		builder.SetLoginTime(time.Now())
	}

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

func (r *AdminLoginLogRepo) Delete(ctx context.Context, req *systemV1.DeleteAdminLoginLogRequest) (bool, error) {
	err := r.data.db.Client().AdminLoginLog.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
