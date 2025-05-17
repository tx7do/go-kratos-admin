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
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/notificationmessage"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"
)

type NotificationMessageRepo struct {
	data *Data
	log  *log.Helper
}

func NewNotificationMessageRepo(data *Data, logger log.Logger) *NotificationMessageRepo {
	l := log.NewHelper(log.With(logger, "module", "notification-message/repo/admin-service"))
	return &NotificationMessageRepo{
		data: data,
		log:  l,
	}
}

func (r *NotificationMessageRepo) toProtoStatus(in *notificationmessage.Status) *internalMessageV1.NotificationMessageStatus {
	if in == nil {
		return nil
	}

	switch *in {
	case notificationmessage.StatusUnknown:
		return trans.Ptr(internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Unknown)

	case notificationmessage.StatusDraft:
		return trans.Ptr(internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Draft)

	case notificationmessage.StatusPublished:
		return trans.Ptr(internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Published)

	case notificationmessage.StatusScheduled:
		return trans.Ptr(internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Scheduled)

	case notificationmessage.StatusRevoked:
		return trans.Ptr(internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Revoked)

	case notificationmessage.StatusArchived:
		return trans.Ptr(internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Archived)

	case notificationmessage.StatusDeleted:
		return trans.Ptr(internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Deleted)

	default:
		return nil
	}
}

func (r *NotificationMessageRepo) toEntStatus(in *internalMessageV1.NotificationMessageStatus) *notificationmessage.Status {
	if in == nil {
		return nil
	}

	switch *in {
	case internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Unknown:
		return trans.Ptr(notificationmessage.StatusUnknown)

	case internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Draft:
		return trans.Ptr(notificationmessage.StatusDraft)

	case internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Published:
		return trans.Ptr(notificationmessage.StatusPublished)

	case internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Scheduled:
		return trans.Ptr(notificationmessage.StatusScheduled)

	case internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Revoked:
		return trans.Ptr(notificationmessage.StatusRevoked)

	case internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Archived:
		return trans.Ptr(notificationmessage.StatusArchived)

	case internalMessageV1.NotificationMessageStatus_NotificationMessageStatus_Deleted:
		return trans.Ptr(notificationmessage.StatusDeleted)

	default:
		return nil
	}
}

func (r *NotificationMessageRepo) convertEntToProto(in *ent.NotificationMessage) *internalMessageV1.NotificationMessage {
	if in == nil {
		return nil
	}
	return &internalMessageV1.NotificationMessage{
		Id:         trans.Ptr(in.ID),
		Subject:    in.Subject,
		Content:    in.Content,
		CategoryId: in.CategoryID,
		Status:     r.toProtoStatus(in.Status),
		CreateBy:   in.CreateBy,
		UpdateBy:   in.UpdateBy,
		CreateTime: timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime: timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *NotificationMessageRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().NotificationMessage.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *NotificationMessageRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageResponse, error) {
	builder := r.data.db.Client().NotificationMessage.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), notificationmessage.FieldCreateTime,
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

	items := make([]*internalMessageV1.NotificationMessage, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &internalMessageV1.ListNotificationMessageResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *NotificationMessageRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().NotificationMessage.Query().
		Where(notificationmessage.IDEQ(id)).
		Exist(ctx)
}

func (r *NotificationMessageRepo) Get(ctx context.Context, req *internalMessageV1.GetNotificationMessageRequest) (*internalMessageV1.NotificationMessage, error) {
	ret, err := r.data.db.Client().NotificationMessage.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorResourceNotFound("message not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *NotificationMessageRepo) Create(ctx context.Context, req *internalMessageV1.CreateNotificationMessageRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().NotificationMessage.Create().
		SetNillableSubject(req.Data.Subject).
		SetNillableContent(req.Data.Content).
		SetNillableCategoryID(req.Data.CategoryId).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *NotificationMessageRepo) Update(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageRequest) error {
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
			createReq := &internalMessageV1.CreateNotificationMessageRequest{Data: req.Data}
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

	builder := r.data.db.Client().NotificationMessage.UpdateOneID(req.Data.GetId()).
		SetNillableSubject(req.Data.Subject).
		SetNillableContent(req.Data.Content).
		SetNillableCategoryID(req.Data.CategoryId).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
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

func (r *NotificationMessageRepo) Delete(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageRequest) (bool, error) {
	err := r.data.db.Client().NotificationMessage.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
