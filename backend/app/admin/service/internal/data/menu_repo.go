package data

import (
	"context"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/menu"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type MenuRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[adminV1.Menu, ent.Menu]
	statusConverter *mapper.EnumTypeConverter[adminV1.Menu_Status, menu.Status]
	typeConverter   *mapper.EnumTypeConverter[adminV1.Menu_Type, menu.Type]
}

func NewMenuRepo(data *Data, logger log.Logger) *MenuRepo {
	repo := &MenuRepo{
		log:             log.NewHelper(log.With(logger, "module", "menu/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[adminV1.Menu, ent.Menu](),
		statusConverter: mapper.NewEnumTypeConverter[adminV1.Menu_Status, menu.Status](adminV1.Menu_Status_name, adminV1.Menu_Status_value),
		typeConverter:   mapper.NewEnumTypeConverter[adminV1.Menu_Type, menu.Type](adminV1.Menu_Type_name, adminV1.Menu_Type_value),
	}

	repo.init()

	return repo
}

func (r *MenuRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
	r.mapper.AppendConverters(r.typeConverter.NewConverterPair())
}

func (r *MenuRepo) travelChild(nodes []*adminV1.Menu, node *adminV1.Menu) bool {
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

func (r *MenuRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Menu.Query()
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

func (r *MenuRepo) List(ctx context.Context, req *pagination.PagingRequest, treeTravel bool) (*adminV1.ListMenuResponse, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Menu.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), menu.FieldCreatedAt,
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

	dtos := make([]*adminV1.Menu, 0, len(entities))
	if treeTravel {
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
	} else {
		for _, entity := range entities {
			dto := r.mapper.ToDTO(entity)
			dtos = append(dtos, dto)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &adminV1.ListMenuResponse{
		Total: uint32(count),
		Items: dtos,
	}, nil
}

func (r *MenuRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Menu.Query().
		Where(menu.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, adminV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *MenuRepo) Get(ctx context.Context, req *adminV1.GetMenuRequest) (*adminV1.Menu, error) {
	if req == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().Menu.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, adminV1.ErrorNotFound("menu not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, adminV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *MenuRepo) Create(ctx context.Context, req *adminV1.CreateMenuRequest) error {
	if req == nil || req.Data == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Menu.Create().
		SetNillableParentID(req.Data.ParentId).
		SetNillableType(r.typeConverter.ToEntity(req.Data.Type)).
		SetNillablePath(req.Data.Path).
		SetNillableRedirect(req.Data.Redirect).
		SetNillableAlias(req.Data.Alias).
		SetNillableName(req.Data.Name).
		SetNillableComponent(req.Data.Component).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableCreatedBy(req.Data.CreatedBy).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	if req.Data.Meta != nil {
		builder.SetMeta(req.Data.Meta)
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

func (r *MenuRepo) Update(ctx context.Context, req *adminV1.UpdateMenuRequest) error {
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
			createReq := &adminV1.CreateMenuRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			return r.Create(ctx, createReq)
		}
	}

	var metaPaths []string
	if req.UpdateMask != nil {
		for _, v := range req.UpdateMask.GetPaths() {
			if strings.HasPrefix(v, "meta.") {
				metaPaths = append(metaPaths, strings.SplitAfter(v, "meta.")[1])
			}
		}

		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return adminV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().
		//Debug().
		Menu.UpdateOneID(req.Data.GetId()).
		SetNillableParentID(req.Data.ParentId).
		SetNillableType(r.typeConverter.ToEntity(req.Data.Type)).
		SetNillablePath(req.Data.Path).
		SetNillableRedirect(req.Data.Redirect).
		SetNillableAlias(req.Data.Alias).
		SetNillableName(req.Data.Name).
		SetNillableComponent(req.Data.Component).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableUpdatedBy(req.Data.UpdatedBy).
		SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

	if req.Data.UpdatedAt == nil {
		builder.SetUpdatedAt(time.Now())
	}

	if req.Data.Meta != nil {
		r.updateMetaField(builder, req.Data.Meta, metaPaths)
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

func (r *MenuRepo) updateMetaField(builder *ent.MenuUpdateOne, meta *adminV1.RouteMeta, metaPaths []string) {
	//builder.SetMeta(meta)

	// 删除空值
	nullUpdater := entgoUpdate.SetJsonFieldValueUpdateBuilder(menu.FieldMeta, meta, metaPaths, false)
	if nullUpdater != nil {
		builder.Modify(nullUpdater)
	}
	// 更新字段
	setUpdater := entgoUpdate.SetJsonNullFieldUpdateBuilder(menu.FieldMeta, meta, metaPaths)
	if setUpdater != nil {
		builder.Modify(setUpdater)
	}
}

func (r *MenuRepo) Delete(ctx context.Context, req *adminV1.DeleteMenuRequest) error {
	if req == nil {
		return adminV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := queryAllChildrenIDs(ctx, r.data.db, "sys_menus", req.GetId())
	if err != nil {
		r.log.Errorf("query child menus failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("query child menus failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Info("menu ids to delete: ", ids)

	if _, err = r.data.db.Client().Menu.Delete().
		Where(menu.IDIn(ids...)).
		Exec(ctx); err != nil {
		r.log.Errorf("delete menus failed: %s", err.Error())
		return adminV1.ErrorInternalServerError("delete menus failed")
	}

	return nil
}
