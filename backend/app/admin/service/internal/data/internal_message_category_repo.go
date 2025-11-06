package data

import (
	"context"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/entgo"

	"github.com/tx7do/go-utils/copierutil"
	entgoQuery "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/internalmessagecategory"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"
)

type InternalMessageCategoryRepo struct {
	data *Data
	log  *log.Helper

	mapper *mapper.CopierMapper[internalMessageV1.InternalMessageCategory, ent.InternalMessageCategory]
}

func NewInternalMessageCategoryRepo(data *Data, logger log.Logger) *InternalMessageCategoryRepo {
	repo := &InternalMessageCategoryRepo{
		log:    log.NewHelper(log.With(logger, "module", "internal-message-category/repo/admin-service")),
		data:   data,
		mapper: mapper.NewCopierMapper[internalMessageV1.InternalMessageCategory, ent.InternalMessageCategory](),
	}

	repo.init()

	return repo
}

func (r *InternalMessageCategoryRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
}

func (r *InternalMessageCategoryRepo) travelChild(nodes []*internalMessageV1.InternalMessageCategory, node *internalMessageV1.InternalMessageCategory) bool {
	if nodes == nil {
		return false
	}

	if node.ParentId == nil {
		nodes = append(nodes, node)
		return true
	}

	for _, n := range nodes {
		if node.ParentId == nil {
			continue
		}

		if n.GetId() == node.GetParentId() {
			n.Children = append(n.Children, node)
			return true
		} else {
			if r.travelChild(n.Children, node) {
				return true
			}
		}
	}
	return false
}

func (r *InternalMessageCategoryRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().InternalMessageCategory.Query()
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

func (r *InternalMessageCategoryRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListInternalMessageCategoryResponse, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().InternalMessageCategory.Query()

	err, whereSelectors, querySelectors := entgoQuery.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), internalmessagecategory.FieldCreatedAt,
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

	sort.SliceStable(entities, func(i, j int) bool {
		if entities[j].ParentID == nil {
			return true
		}
		if entities[i].ParentID == nil {
			return true
		}
		return *entities[i].ParentID < *entities[j].ParentID
	})

	dtos := make([]*internalMessageV1.InternalMessageCategory, 0, len(entities))
	for _, entity := range entities {
		if entity.ParentID == nil {
			dto := r.mapper.ToDTO(entity)
			dtos = append(dtos, dto)
		}
	}
	for _, entity := range entities {
		if entity.ParentID != nil {
			dto := r.mapper.ToDTO(entity)

			if r.travelChild(dtos, dto) {
				continue
			}

			dtos = append(dtos, dto)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &internalMessageV1.ListInternalMessageCategoryResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *InternalMessageCategoryRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().InternalMessageCategory.Query().
		Where(internalmessagecategory.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, internalMessageV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *InternalMessageCategoryRepo) Get(ctx context.Context, req *internalMessageV1.GetInternalMessageCategoryRequest) (*internalMessageV1.InternalMessageCategory, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().InternalMessageCategory.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorNotFound("message category not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, internalMessageV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

// GetCategoriesByIds 根据ID列表获取分类列表
func (r *InternalMessageCategoryRepo) GetCategoriesByIds(ctx context.Context, ids []uint32) ([]*internalMessageV1.InternalMessageCategory, error) {
	if len(ids) == 0 {
		return []*internalMessageV1.InternalMessageCategory{}, nil
	}

	entities, err := r.data.db.Client().InternalMessageCategory.Query().
		Where(internalmessagecategory.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query internal message category by ids failed: %s", err.Error())
		return nil, internalMessageV1.ErrorInternalServerError("query internal message category by ids failed")
	}

	dtos := make([]*internalMessageV1.InternalMessageCategory, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (r *InternalMessageCategoryRepo) Create(ctx context.Context, req *internalMessageV1.CreateInternalMessageCategoryRequest) error {
	if req == nil || req.Data == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().InternalMessageCategory.Create().
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableIconURL(req.Data.IconUrl).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortOrder(req.Data.SortOrder).
		SetNillableIsEnabled(req.Data.IsEnabled).
		SetNillableCreatedBy(req.Data.CreatedBy).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *InternalMessageCategoryRepo) Update(ctx context.Context, req *internalMessageV1.UpdateInternalMessageCategoryRequest) error {
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
			createReq := &internalMessageV1.CreateInternalMessageCategoryRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
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

	builder := r.data.db.Client().InternalMessageCategory.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableIconURL(req.Data.IconUrl).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortOrder(req.Data.SortOrder).
		SetNillableIsEnabled(req.Data.IsEnabled).
		SetNillableUpdatedBy(req.Data.UpdatedBy).
		SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

	if req.Data.UpdatedAt == nil {
		builder.SetUpdatedAt(time.Now())
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

func (r *InternalMessageCategoryRepo) Delete(ctx context.Context, req *internalMessageV1.DeleteInternalMessageCategoryRequest) error {
	if req == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := entgo.QueryAllChildrenIds(ctx, r.data.db, "internal_message_categories", req.GetId())
	if err != nil {
		r.log.Errorf("query child internal message categories failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("query child internal message categories failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Info("internal message category ids to delete: ", ids)

	if _, err = r.data.db.Client().InternalMessageCategory.Delete().
		Where(internalmessagecategory.IDIn(ids...)).
		Exec(ctx); err != nil {
		r.log.Errorf("delete internal message categories failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("delete internal message categories failed")
	}

	return nil
}
