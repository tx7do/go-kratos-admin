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
	timeutil "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/role"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) *RoleRepo {
	l := log.NewHelper(log.With(logger, "module", "role/repo/admin-service"))
	return &RoleRepo{
		data: data,
		log:  l,
	}
}

func (r *RoleRepo) convertEntToProto(in *ent.Role) *userV1.Role {
	if in == nil {
		return nil
	}
	return &userV1.Role{
		Id:         trans.Ptr(in.ID),
		Name:       in.Name,
		Code:       in.Code,
		Remark:     in.Remark,
		SortId:     in.SortID,
		ParentId:   in.ParentID,
		Menus:      in.Menus,
		Status:     (*string)(in.Status),
		CreateBy:   in.CreateBy,
		UpdateBy:   in.UpdateBy,
		CreateTime: timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime: timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *RoleRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Role.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *RoleRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListRoleResponse, error) {
	builder := r.data.db.Client().Role.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), role.FieldCreateTime,
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

	items := make([]*userV1.Role, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListRoleResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *RoleRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().Role.Query().
		Where(role.IDEQ(id)).
		Exist(ctx)
}

func (r *RoleRepo) GetRole(ctx context.Context, id uint32) (*userV1.Role, error) {
	ret, err := r.data.db.Client().Role.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorRoleNotFound("role not found")
		}

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *RoleRepo) CreateRole(ctx context.Context, req *userV1.CreateRoleRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().Role.Create().
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableCode(req.Data.Code).
		SetNillableStatus((*role.Status)(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
		SetNillableCreateBy(req.OperatorId).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.Menus != nil {
		builder.SetMenus(req.Data.Menus)
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *RoleRepo) UpdateRole(ctx context.Context, req *userV1.UpdateRoleRequest) error {
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
			return r.CreateRole(ctx, &userV1.CreateRoleRequest{Data: req.Data, OperatorId: req.OperatorId})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().Role.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableCode(req.Data.Code).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*role.Status)(req.Data.Status)).
		SetNillableUpdateBy(req.OperatorId).
		SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if req.Data.Menus != nil {
		builder.SetMenus(req.Data.Menus)
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

func (r *RoleRepo) DeleteRole(ctx context.Context, req *userV1.DeleteRoleRequest) (bool, error) {
	err := r.data.db.Client().Role.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
