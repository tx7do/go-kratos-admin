package data

import (
	"context"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/menu"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type MenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) *MenuRepo {
	l := log.NewHelper(log.With(logger, "module", "menu/repo/admin-service"))
	return &MenuRepo{
		data: data,
		log:  l,
	}
}

func (r *MenuRepo) convertMenuTypeToProto(in *menu.Type) *systemV1.MenuType {
	if in == nil {
		return nil
	}
	find, ok := systemV1.MenuType_value[string(*in)]
	if !ok {
		return nil
	}
	return (*systemV1.MenuType)(trans.Ptr(find))
}
func (r *MenuRepo) convertMenuTypeToEnt(in *systemV1.MenuType) *menu.Type {
	if in == nil {
		return nil
	}
	find, ok := systemV1.MenuType_name[int32(*in)]
	if !ok {
		return nil
	}
	return (*menu.Type)(trans.Ptr(find))
}

func (r *MenuRepo) convertUserStatusToEnt(status *systemV1.MenuStatus) *menu.Status {
	if status == nil {
		return nil
	}
	find, ok := systemV1.MenuStatus_name[int32(*status)]
	if !ok {
		return nil
	}
	return (*menu.Status)(trans.Ptr(find))
}

func (r *MenuRepo) convertUserStatusToProto(status *menu.Status) *systemV1.MenuStatus {
	if status == nil {
		return nil
	}
	find, ok := systemV1.MenuStatus_value[string(*status)]
	if !ok {
		return nil
	}
	return (*systemV1.MenuStatus)(trans.Ptr(find))
}

func (r *MenuRepo) convertEntToProto(in *ent.Menu) *systemV1.Menu {
	if in == nil {
		return nil
	}

	return &systemV1.Menu{
		Id:         trans.Ptr(in.ID),
		ParentId:   in.ParentID,
		Path:       in.Path,
		Redirect:   in.Redirect,
		Alias:      in.Alias,
		Name:       in.Name,
		Component:  in.Component,
		Meta:       in.Meta,
		Type:       r.convertMenuTypeToProto(in.Type),
		Status:     r.convertUserStatusToProto(in.Status),
		CreateBy:   in.CreateBy,
		UpdateBy:   in.UpdateBy,
		CreateTime: timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime: timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *MenuRepo) travelChild(nodes []*systemV1.Menu, node *systemV1.Menu) bool {
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
		return 0, systemV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *MenuRepo) List(ctx context.Context, req *pagination.PagingRequest, treeTravel bool) (*systemV1.ListMenuResponse, error) {
	if req == nil {
		return nil, systemV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Menu.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), menu.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, systemV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, systemV1.ErrorInternalServerError("query list failed")
	}

	items := make([]*systemV1.Menu, 0, len(results))
	if treeTravel {
		for _, m := range results {
			if m.ParentID == nil {
				item := r.convertEntToProto(m)
				items = append(items, item)
			}
		}
		for _, m := range results {
			if m.ParentID != nil {
				item := r.convertEntToProto(m)

				if r.travelChild(items, item) {
					continue
				}

				items = append(items, item)
			}
		}
	} else {
		for _, res := range results {
			item := r.convertEntToProto(res)
			items = append(items, item)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListMenuResponse{
		Total: uint32(count),
		Items: items,
	}, nil
}

func (r *MenuRepo) IsExist(ctx context.Context, id int32) (bool, error) {
	exist, err := r.data.db.Client().Menu.Query().
		Where(menu.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, systemV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *MenuRepo) Get(ctx context.Context, req *systemV1.GetMenuRequest) (*systemV1.Menu, error) {
	if req == nil {
		return nil, systemV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().Menu.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("menu not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, systemV1.ErrorInternalServerError("query data failed")
	}

	return r.convertEntToProto(ret), nil
}

func (r *MenuRepo) Create(ctx context.Context, req *systemV1.CreateMenuRequest) error {
	if req == nil || req.Data == nil {
		return systemV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Menu.Create().
		SetNillableParentID(req.Data.ParentId).
		SetNillableType(r.convertMenuTypeToEnt(req.Data.Type)).
		SetNillablePath(req.Data.Path).
		SetNillableRedirect(req.Data.Redirect).
		SetNillableAlias(req.Data.Alias).
		SetNillableName(req.Data.Name).
		SetNillableComponent(req.Data.Component).
		SetNillableStatus(r.convertUserStatusToEnt(req.Data.Status)).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.Meta != nil {
		builder.SetMeta(req.Data.Meta)
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return systemV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *MenuRepo) Update(ctx context.Context, req *systemV1.UpdateMenuRequest) error {
	if req == nil || req.Data == nil {
		return systemV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &systemV1.CreateMenuRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
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
			return systemV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().
		//Debug().
		Menu.UpdateOneID(req.Data.GetId()).
		SetNillableParentID(req.Data.ParentId).
		SetNillableType(r.convertMenuTypeToEnt(req.Data.Type)).
		SetNillablePath(req.Data.Path).
		SetNillableRedirect(req.Data.Redirect).
		SetNillableAlias(req.Data.Alias).
		SetNillableName(req.Data.Name).
		SetNillableComponent(req.Data.Component).
		SetNillableStatus(r.convertUserStatusToEnt(req.Data.Status)).
		SetNillableUpdateBy(req.Data.UpdateBy).
		SetNillableUpdateTime(timeutil.StringTimeToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
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
		return systemV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *MenuRepo) updateMetaField(builder *ent.MenuUpdateOne, meta *systemV1.RouteMeta, metaPaths []string) {
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

func (r *MenuRepo) Delete(ctx context.Context, req *systemV1.DeleteMenuRequest) error {
	if req == nil {
		return systemV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().Menu.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return systemV1.ErrorResourceNotFound("menu not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return systemV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
