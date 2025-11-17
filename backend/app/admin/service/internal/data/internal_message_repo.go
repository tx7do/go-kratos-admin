package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/proto"

	"github.com/tx7do/go-utils/copierutil"
	entgoQuery "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/internalmessage"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"
)

type InternalMessageRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[internalMessageV1.InternalMessage, ent.InternalMessage]
	statusConverter *mapper.EnumTypeConverter[internalMessageV1.InternalMessage_Status, internalmessage.Status]
	typeConverter   *mapper.EnumTypeConverter[internalMessageV1.InternalMessage_Type, internalmessage.Type]
}

func NewInternalMessageRepo(data *Data, logger log.Logger) *InternalMessageRepo {
	repo := &InternalMessageRepo{
		log:             log.NewHelper(log.With(logger, "module", "internal-message/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[internalMessageV1.InternalMessage, ent.InternalMessage](),
		statusConverter: mapper.NewEnumTypeConverter[internalMessageV1.InternalMessage_Status, internalmessage.Status](internalMessageV1.InternalMessage_Status_name, internalMessageV1.InternalMessage_Status_value),
		typeConverter:   mapper.NewEnumTypeConverter[internalMessageV1.InternalMessage_Type, internalmessage.Type](internalMessageV1.InternalMessage_Type_name, internalMessageV1.InternalMessage_Type_value),
	}

	repo.init()

	return repo
}

func (r *InternalMessageRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
	r.mapper.AppendConverters(r.typeConverter.NewConverterPair())
}

func (r *InternalMessageRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().InternalMessage.Query()
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

func (r *InternalMessageRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListInternalMessageResponse, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().InternalMessage.Query()

	err, whereSelectors, querySelectors := entgoQuery.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), internalmessage.FieldCreatedAt,
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

	dtos := make([]*internalMessageV1.InternalMessage, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &internalMessageV1.ListInternalMessageResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *InternalMessageRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().InternalMessage.Query().
		Where(internalmessage.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, internalMessageV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *InternalMessageRepo) Get(ctx context.Context, req *internalMessageV1.GetInternalMessageRequest) (*internalMessageV1.InternalMessage, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().InternalMessage.Query()

	builder.Where(internalmessage.IDEQ(req.GetId()))

	entgoQuery.ApplyFieldMaskToBuilder(builder, req.ViewMask)

	entity, err := builder.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorNotFound("message not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, internalMessageV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *InternalMessageRepo) Create(ctx context.Context, req *internalMessageV1.CreateInternalMessageRequest) (*internalMessageV1.InternalMessage, error) {
	if req == nil || req.Data == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().InternalMessage.Create().
		SetNillableTitle(req.Data.Title).
		SetNillableContent(req.Data.Content).
		SetNillableSenderID(req.Data.SenderId).
		SetNillableCategoryID(req.Data.CategoryId).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableType(r.typeConverter.ToEntity(req.Data.Type)).
		SetNillableCreatedBy(req.Data.CreatedBy).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	var err error
	var entity *ent.InternalMessage
	if entity, err = builder.Save(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, internalMessageV1.ErrorInternalServerError("insert data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *InternalMessageRepo) Update(ctx context.Context, req *internalMessageV1.UpdateInternalMessageRequest) error {
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
			createReq := &internalMessageV1.CreateInternalMessageRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			_, err = r.Create(ctx, createReq)
			return err
		}
	}

	if err := fieldmaskutil.FilterByFieldMask(trans.Ptr(proto.Message(req.GetData())), req.UpdateMask); err != nil {
		r.log.Errorf("invalid field mask [%v], error: %s", req.UpdateMask, err.Error())
		return internalMessageV1.ErrorBadRequest("invalid field mask")
	}

	builder := r.data.db.Client().InternalMessage.UpdateOneID(req.Data.GetId()).
		SetNillableTitle(req.Data.Title).
		SetNillableContent(req.Data.Content).
		SetNillableSenderID(req.Data.SenderId).
		SetNillableCategoryID(req.Data.CategoryId).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableType(r.typeConverter.ToEntity(req.Data.Type)).
		SetNillableUpdatedBy(req.Data.UpdatedBy).
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

func (r *InternalMessageRepo) Delete(ctx context.Context, id uint32) error {
	if id == 0 {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().InternalMessage.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return internalMessageV1.ErrorNotFound("internal message not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return internalMessageV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
