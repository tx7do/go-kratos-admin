package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/copierutil"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/privatemessage"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"
)

type PrivateMessageRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[ent.PrivateMessage, internalMessageV1.PrivateMessage]
	statusConverter *mapper.EnumTypeConverter[internalMessageV1.MessageStatus, privatemessage.Status]
}

func NewPrivateMessageRepo(data *Data, logger log.Logger) *PrivateMessageRepo {
	repo := &PrivateMessageRepo{
		log:             log.NewHelper(log.With(logger, "module", "private-message/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[ent.PrivateMessage, internalMessageV1.PrivateMessage](),
		statusConverter: mapper.NewEnumTypeConverter[internalMessageV1.MessageStatus, privatemessage.Status](internalMessageV1.MessageStatus_name, internalMessageV1.MessageStatus_value),
	}

	repo.init()

	return repo
}

func (r *PrivateMessageRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *PrivateMessageRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().PrivateMessage.Query()
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

func (r *PrivateMessageRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListPrivateMessageResponse, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().PrivateMessage.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), privatemessage.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, internalMessageV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, internalMessageV1.ErrorInternalServerError("query list failed")
	}

	models := make([]*internalMessageV1.PrivateMessage, 0, len(results))
	for _, dto := range results {
		model := r.mapper.ToModel(dto)
		models = append(models, model)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &internalMessageV1.ListPrivateMessageResponse{
		Total: uint32(count),
		Items: models,
	}, err
}

func (r *PrivateMessageRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().PrivateMessage.Query().
		Where(privatemessage.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, internalMessageV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *PrivateMessageRepo) Get(ctx context.Context, req *internalMessageV1.GetPrivateMessageRequest) (*internalMessageV1.PrivateMessage, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	dto, err := r.data.db.Client().PrivateMessage.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorNotFound("message not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, internalMessageV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToModel(dto), nil
}

func (r *PrivateMessageRepo) Create(ctx context.Context, req *internalMessageV1.CreatePrivateMessageRequest) error {
	if req == nil || req.Data == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().PrivateMessage.Create().
		SetNillableSubject(req.Data.Subject).
		SetNillableContent(req.Data.Content).
		SetNillableSenderID(req.Data.SenderId).
		SetNillableReceiverID(req.Data.ReceiverId).
		SetNillableStatus(r.statusConverter.ToModel(req.Data.Status)).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *PrivateMessageRepo) Update(ctx context.Context, req *internalMessageV1.UpdatePrivateMessageRequest) error {
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
			createReq := &internalMessageV1.CreatePrivateMessageRequest{Data: req.Data}
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return internalMessageV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().PrivateMessage.UpdateOneID(req.Data.GetId()).
		SetNillableSubject(req.Data.Subject).
		SetNillableContent(req.Data.Content).
		SetNillableSenderID(req.Data.SenderId).
		SetNillableReceiverID(req.Data.ReceiverId).
		SetNillableStatus(r.statusConverter.ToModel(req.Data.Status)).
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

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *PrivateMessageRepo) Delete(ctx context.Context, req *internalMessageV1.DeletePrivateMessageRequest) error {
	if req == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().PrivateMessage.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return internalMessageV1.ErrorNotFound("private message not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return internalMessageV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
