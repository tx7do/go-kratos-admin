package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/copierutil"
	entgoQuery "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/proto"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/internalmessagerecipient"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"
)

type InternalMessageRecipientRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[internalMessageV1.InternalMessageRecipient, ent.InternalMessageRecipient]
	statusConverter *mapper.EnumTypeConverter[internalMessageV1.InternalMessageRecipient_Status, internalmessagerecipient.Status]
}

func NewInternalMessageRecipientRepo(data *Data, logger log.Logger) *InternalMessageRecipientRepo {
	repo := &InternalMessageRecipientRepo{
		log:             log.NewHelper(log.With(logger, "module", "internal-message-recipient/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[internalMessageV1.InternalMessageRecipient, ent.InternalMessageRecipient](),
		statusConverter: mapper.NewEnumTypeConverter[internalMessageV1.InternalMessageRecipient_Status, internalmessagerecipient.Status](internalMessageV1.InternalMessageRecipient_Status_name, internalMessageV1.InternalMessageRecipient_Status_value),
	}

	repo.init()

	return repo
}

func (r *InternalMessageRecipientRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *InternalMessageRecipientRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().InternalMessageRecipient.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, internalMessageV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *InternalMessageRecipientRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().InternalMessageRecipient.Query().
		Where(internalmessagerecipient.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, internalMessageV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *InternalMessageRecipientRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListUserInboxResponse, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().InternalMessageRecipient.Query()

	err, whereSelectors, querySelectors := entgoQuery.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), internalmessagerecipient.FieldCreatedAt,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, internalMessageV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, internalMessageV1.ErrorInternalServerError("query list failed")
	}

	dtos := make([]*internalMessageV1.InternalMessageRecipient, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &internalMessageV1.ListUserInboxResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *InternalMessageRecipientRepo) Get(ctx context.Context, id uint32) (*internalMessageV1.InternalMessageRecipient, error) {
	if id == 0 {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().InternalMessageRecipient.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorNotFound("message not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, internalMessageV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *InternalMessageRecipientRepo) Create(ctx context.Context, req *internalMessageV1.InternalMessageRecipient) (*internalMessageV1.InternalMessageRecipient, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().InternalMessageRecipient.Create().
		SetNillableMessageID(req.MessageId).
		SetNillableRecipientUserID(req.RecipientUserId).
		SetNillableStatus(r.statusConverter.ToEntity(req.Status)).
		SetNillableReceivedAt(timeutil.TimestamppbToTime(req.ReceivedAt)).
		SetNillableReadAt(timeutil.TimestamppbToTime(req.ReadAt)).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.CreatedAt))

	if req.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	var err error
	var entity *ent.InternalMessageRecipient
	if entity, err = builder.Save(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, internalMessageV1.ErrorInternalServerError("insert data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *InternalMessageRecipientRepo) Update(ctx context.Context, req *internalMessageV1.UpdateInternalMessageRecipientRequest) error {
	if req == nil || req.Data == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			req.Data.CreatedBy = req.Data.UpdatedBy
			req.Data.UpdatedBy = nil
			_, err = r.Create(ctx, req.Data)
			return err
		}
	}

	if err := fieldmaskutil.FilterByFieldMask(trans.Ptr(proto.Message(req.GetData())), req.UpdateMask); err != nil {
		r.log.Errorf("invalid field mask [%v], error: %s", req.UpdateMask, err.Error())
		return internalMessageV1.ErrorBadRequest("invalid field mask")
	}

	builder := r.data.db.Client().InternalMessageRecipient.UpdateOneID(req.Data.GetId()).
		SetNillableMessageID(req.Data.MessageId).
		SetNillableRecipientUserID(req.Data.RecipientUserId).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableReceivedAt(timeutil.TimestamppbToTime(req.Data.ReceivedAt)).
		SetNillableReadAt(timeutil.TimestamppbToTime(req.Data.ReadAt)).
		SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

	if req.Data.UpdatedAt == nil {
		builder.SetUpdatedAt(time.Now())
	}

	entgoUpdate.ApplyNilFieldMask(proto.Message(req.GetData()), req.UpdateMask, builder)

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *InternalMessageRecipientRepo) Delete(ctx context.Context, id uint32) error {
	if id == 0 {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().InternalMessageRecipient.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return internalMessageV1.ErrorNotFound("internal message recipient not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return internalMessageV1.ErrorInternalServerError("delete failed")
	}

	return nil
}

// MarkNotificationAsRead 将通知标记为已读
func (r *InternalMessageRecipientRepo) MarkNotificationAsRead(ctx context.Context, req *internalMessageV1.MarkNotificationAsReadRequest) error {
	if len(req.GetRecipientIds()) == 0 {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}
	if req.GetUserId() == 0 {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	now := time.Now()
	_, err := r.data.db.Client().InternalMessageRecipient.Update().
		Where(
			internalmessagerecipient.IDIn(req.GetRecipientIds()...),
			internalmessagerecipient.RecipientUserIDEQ(req.GetUserId()),
			internalmessagerecipient.StatusNEQ(internalmessagerecipient.StatusRead),
		).
		SetStatus(internalmessagerecipient.StatusRead).
		SetNillableReadAt(trans.Ptr(now)).
		SetNillableUpdatedAt(trans.Ptr(now)).
		Save(ctx)
	return err
}

// MarkNotificationsStatus 标记特定用户的某些或所有通知的状态
func (r *InternalMessageRecipientRepo) MarkNotificationsStatus(ctx context.Context, req *internalMessageV1.MarkNotificationsStatusRequest) error {
	if len(req.GetRecipientIds()) == 0 {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}
	if req.GetUserId() == 0 {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	now := time.Now()
	var readAt *time.Time
	var receiveAt *time.Time
	switch req.GetNewStatus() {
	case internalMessageV1.InternalMessageRecipient_READ:
		readAt = trans.Ptr(now)
	case internalMessageV1.InternalMessageRecipient_RECEIVED:
		receiveAt = trans.Ptr(now)
	}

	_, err := r.data.db.Client().InternalMessageRecipient.Update().
		Where(
			internalmessagerecipient.IDIn(req.GetRecipientIds()...),
			internalmessagerecipient.RecipientUserIDEQ(req.GetUserId()),
			internalmessagerecipient.StatusNEQ(*r.statusConverter.ToEntity(trans.Ptr(req.GetNewStatus()))),
		).
		SetNillableStatus(r.statusConverter.ToEntity(trans.Ptr(req.GetNewStatus()))).
		SetNillableReadAt(readAt).
		SetNillableReceivedAt(receiveAt).
		SetNillableUpdatedAt(trans.Ptr(now)).
		Save(ctx)
	return err
}

// RevokeMessage 撤销某条消息
func (r *InternalMessageRecipientRepo) RevokeMessage(ctx context.Context, req *internalMessageV1.RevokeMessageRequest) error {
	_, err := r.data.db.Client().InternalMessageRecipient.Delete().
		Where(
			internalmessagerecipient.MessageIDEQ(req.GetMessageId()),
			internalmessagerecipient.RecipientUserIDEQ(req.GetUserId()),
		).
		Exec(ctx)
	return err
}

func (r *InternalMessageRecipientRepo) DeleteNotificationFromInbox(ctx context.Context, req *internalMessageV1.DeleteNotificationFromInboxRequest) error {
	_, err := r.data.db.Client().InternalMessageRecipient.Delete().
		Where(
			internalmessagerecipient.IDIn(req.GetRecipientIds()...),
			internalmessagerecipient.RecipientUserIDEQ(req.GetUserId()),
		).
		Exec(ctx)
	return err
}
