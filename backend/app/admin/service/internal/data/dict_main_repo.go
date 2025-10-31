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
	"kratos-admin/app/admin/service/internal/data/ent/dictmain"

	dictV1 "kratos-admin/api/gen/go/dict/service/v1"
)

type DictMainRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[dictV1.DictMain, ent.DictMain]
	statusConverter *mapper.EnumTypeConverter[dictV1.DictMain_Status, dictmain.Status]
}

func NewDictMainRepo(data *Data, logger log.Logger) *DictMainRepo {
	repo := &DictMainRepo{
		log:             log.NewHelper(log.With(logger, "module", "dict-main/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[dictV1.DictMain, ent.DictMain](),
		statusConverter: mapper.NewEnumTypeConverter[dictV1.DictMain_Status, dictmain.Status](dictV1.DictMain_Status_name, dictV1.DictMain_Status_value),
	}

	repo.init()

	return repo
}

func (r *DictMainRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *DictMainRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().DictMain.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, dictV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *DictMainRepo) List(ctx context.Context, req *pagination.PagingRequest) (*dictV1.ListDictMainResponse, error) {
	if req == nil {
		return nil, dictV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().DictMain.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), dictmain.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, dictV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, dictV1.ErrorInternalServerError("query list failed")
	}

	dtos := make([]*dictV1.DictMain, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &dictV1.ListDictMainResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *DictMainRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().DictMain.Query().
		Where(dictmain.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, dictV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *DictMainRepo) Get(ctx context.Context, req *dictV1.GetDictMainRequest) (*dictV1.DictMain, error) {
	if req == nil {
		return nil, dictV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().DictMain.Query()

	switch req.GetQueryBy().(type) {
	case *dictV1.GetDictMainRequest_Id:
		builder.Where(dictmain.IDEQ(req.GetId()))
	case *dictV1.GetDictMainRequest_Code:
		builder.Where(dictmain.CodeEQ(req.GetCode()))
	default:
		return nil, dictV1.ErrorBadRequest("invalid query parameter")
	}

	entity, err := builder.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, dictV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, dictV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *DictMainRepo) Create(ctx context.Context, req *dictV1.CreateDictMainRequest) error {
	if req == nil || req.Data == nil {
		return dictV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().DictMain.Create().
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
		return dictV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *DictMainRepo) Update(ctx context.Context, req *dictV1.UpdateDictMainRequest) error {
	if req == nil || req.Data == nil {
		return dictV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &dictV1.CreateDictMainRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return dictV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().DictMain.
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
		return dictV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *DictMainRepo) Delete(ctx context.Context, id uint32) error {
	if id == 0 {
		return dictV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().DictMain.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return dictV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return dictV1.ErrorInternalServerError("delete failed")
	}

	return nil
}

func (r *DictMainRepo) BatchDelete(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return dictV1.ErrorBadRequest("invalid parameter")
	}

	if _, err := r.data.db.Client().DictMain.Delete().
		Where(dictmain.IDIn(ids...)).
		Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return dictV1.ErrorNotFound("dict not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return dictV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
