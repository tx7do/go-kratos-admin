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
	"kratos-admin/app/admin/service/internal/data/ent/dictitem"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type DictItemRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[adminV1.DictItem, ent.DictItem]
	statusConverter *mapper.EnumTypeConverter[adminV1.DictItem_Status, dictitem.Status]
}

func NewDictItemRepo(data *Data, logger log.Logger) *DictItemRepo {
	repo := &DictItemRepo{
		log:             log.NewHelper(log.With(logger, "module", "dict-item/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[adminV1.DictItem, ent.DictItem](),
		statusConverter: mapper.NewEnumTypeConverter[adminV1.DictItem_Status, dictitem.Status](adminV1.DictItem_Status_name, adminV1.DictItem_Status_value),
	}

	repo.init()

	return repo
}

func (r *DictItemRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *DictItemRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().DictItem.Query()
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

func (r *DictItemRepo) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListDictItemResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().DictItem.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), dictitem.FieldCreateTime,
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

	dtos := make([]*adminV1.DictItem, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListDictItemResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *DictItemRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().DictItem.Query().
		Where(dictitem.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *DictItemRepo) Get(ctx context.Context, req *adminV1.GetDictItemRequest) (*adminV1.DictItem, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().DictItem.Query()

	switch req.GetQueryBy().(type) {
	case *adminV1.GetDictItemRequest_Id:
		builder.Where(dictitem.IDEQ(req.GetId()))
	case *adminV1.GetDictItemRequest_Code:
		builder.Where(dictitem.CodeEQ(req.GetCode()))
	default:
		return nil, adminV1.ErrorBadRequest("invalid query parameter")
	}

	entity, err := builder.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *DictItemRepo) Create(ctx context.Context, req *adminV1.CreateDictItemRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().DictItem.Create().
		SetNillableCode(req.Data.Code).
		SetNillableName(req.Data.Name).
		SetNillableSortID(req.Data.SortId).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *DictItemRepo) Update(ctx context.Context, req *adminV1.UpdateDictItemRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &adminV1.CreateDictItemRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return adminV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().DictItem.
		UpdateOneID(req.Data.GetId()).
		SetNillableCode(req.Data.Code).
		SetNillableName(req.Data.Name).
		SetNillableSortID(req.Data.SortId).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
		SetNillableUpdateBy(req.Data.UpdateBy).
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
		return adminV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *DictItemRepo) Delete(ctx context.Context, id uint32) error {
	if id == 0 {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().DictItem.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return adminV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return adminV1.ErrorInternalServerError("delete failed")
	}

	return nil
}

func (r *DictItemRepo) BatchDelete(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	if _, err := r.data.db.Client().DictItem.Delete().
		Where(dictitem.IDIn(ids...)).
		Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return adminV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return adminV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
