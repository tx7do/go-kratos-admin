package data

import (
	"context"
	"time"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"kratos-admin/app/admin/service/internal/data/ent/role"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	entCrud "github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[userV1.Role, ent.Role]
	statusConverter *mapper.EnumTypeConverter[userV1.Role_Status, role.Status]

	repository *entCrud.Repository[
		ent.RoleQuery, ent.RoleSelect,
		ent.RoleCreate, ent.RoleCreateBulk,
		ent.RoleUpdate, ent.RoleUpdateOne,
		ent.RoleDelete,
		predicate.Role,
		userV1.Role, ent.Role,
	]
}

func NewRoleRepo(data *Data, logger log.Logger) *RoleRepo {
	repo := &RoleRepo{
		log:             log.NewHelper(log.With(logger, "module", "role/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[userV1.Role, ent.Role](),
		statusConverter: mapper.NewEnumTypeConverter[userV1.Role_Status, role.Status](userV1.Role_Status_name, userV1.Role_Status_value),
	}

	repo.init()

	return repo
}

func (r *RoleRepo) init() {
	r.repository = entCrud.NewRepository[
		ent.RoleQuery, ent.RoleSelect,
		ent.RoleCreate, ent.RoleCreateBulk,
		ent.RoleUpdate, ent.RoleUpdateOne,
		ent.RoleDelete,
		predicate.Role,
		userV1.Role, ent.Role,
	](r.mapper)

	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *RoleRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Role.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, userV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *RoleRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListRoleResponse, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Role.Query()

	ret, err := r.repository.ListWithPaging(ctx, builder, builder.Clone(), req)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return &userV1.ListRoleResponse{Total: 0, Items: nil}, nil
	}

	return &userV1.ListRoleResponse{
		Total: ret.Total,
		Items: ret.Items,
	}, nil
}

func (r *RoleRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Role.Query().
		Where(role.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, userV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *RoleRepo) Get(ctx context.Context, req *userV1.GetRoleRequest) (*userV1.Role, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Role.Query()

	var whereCond []func(s *sql.Selector)
	switch req.QueryBy.(type) {
	default:
	case *userV1.GetRoleRequest_Id:
		whereCond = append(whereCond, role.IDEQ(req.GetId()))
	}

	dto, err := r.repository.Get(ctx, builder, req.GetViewMask(), whereCond...)
	if err != nil {
		return nil, err
	}

	return dto, err
}

// GetRoleByCode 通过角色编码获取角色信息
func (r *RoleRepo) GetRoleByCode(ctx context.Context, code string) (*userV1.Role, error) {
	if code == "" {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().Role.Query().
		Where(role.CodeEQ(code)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorRoleNotFound("role not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

// GetRolesByRoleCodes 通过角色编码列表获取角色列表
func (r *RoleRepo) GetRolesByRoleCodes(ctx context.Context, codes []string) ([]*userV1.Role, error) {
	if len(codes) == 0 {
		return []*userV1.Role{}, nil
	}

	entities, err := r.data.db.Client().Role.Query().
		Where(role.CodeIn(codes...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query roles by codes failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query roles by codes failed")
	}

	dtos := make([]*userV1.Role, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

// GetRolesByRoleIds 通过角色ID列表获取角色列表
func (r *RoleRepo) GetRolesByRoleIds(ctx context.Context, ids []uint32) ([]*userV1.Role, error) {
	if len(ids) == 0 {
		return []*userV1.Role{}, nil
	}

	entities, err := r.data.db.Client().Role.Query().
		Where(role.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query roles by ids failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query roles by ids failed")
	}

	dtos := make([]*userV1.Role, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

// GetRoleCodesByRoleIds 通过角色ID列表获取角色编码列表
func (r *RoleRepo) GetRoleCodesByRoleIds(ctx context.Context, ids []uint32) ([]string, error) {
	if len(ids) == 0 {
		return []string{}, nil
	}

	entities, err := r.data.db.Client().Role.Query().
		Where(role.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query role codes failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query role codes failed")
	}

	codes := make([]string, 0, len(entities))
	for _, entity := range entities {
		if entity.Code != nil {
			codes = append(codes, *entity.Code)
		}
	}

	return codes, nil
}

func (r *RoleRepo) Create(ctx context.Context, req *userV1.CreateRoleRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Role.Create().
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortOrder(req.Data.SortOrder).
		SetNillableCode(req.Data.Code).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
		SetNillableCreatedBy(req.Data.CreatedBy).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.TenantId == nil {
		builder.SetTenantID(req.Data.GetTenantId())
	}
	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	if req.Data.Menus != nil {
		builder.SetMenus(req.Data.Menus)
	}
	if req.Data.Apis != nil {
		builder.SetApis(req.Data.Apis)
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return userV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *RoleRepo) Update(ctx context.Context, req *userV1.UpdateRoleRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &userV1.CreateRoleRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			return r.Create(ctx, createReq)
		}
	}

	builder := r.data.db.Client().Debug().Role.Update()
	err := r.repository.UpdateX(ctx, builder, req.Data, req.GetUpdateMask(),
		func(dto *userV1.Role) {
			builder.
				SetNillableName(req.Data.Name).
				SetNillableParentID(req.Data.ParentId).
				SetNillableSortOrder(req.Data.SortOrder).
				SetNillableCode(req.Data.Code).
				SetNillableRemark(req.Data.Remark).
				SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
				SetNillableUpdatedBy(req.Data.UpdatedBy).
				SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

			if req.Data.UpdatedAt == nil {
				builder.SetUpdatedAt(time.Now())
			}

			if req.Data.Menus != nil {
				builder.SetMenus(req.Data.Menus)
			}
			if req.Data.Apis != nil {
				builder.SetApis(req.Data.Apis)
			}
		},
		func(s *sql.Selector) {
			s.Where(sql.EQ(role.FieldID, req.Data.GetId()))
		},
	)

	return err
}

func (r *RoleRepo) Delete(ctx context.Context, req *userV1.DeleteRoleRequest) error {
	if req == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := entCrud.QueryAllChildrenIds(ctx, r.data.db, "sys_roles", req.GetId())
	if err != nil {
		r.log.Errorf("query child roles failed: %s", err.Error())
		return userV1.ErrorInternalServerError("query child roles failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Info("roles ids to delete: ", ids)

	if _, err = r.data.db.Client().Role.Delete().
		Where(role.IDIn(ids...)).
		Exec(ctx); err != nil {
		r.log.Errorf("delete roles failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete roles failed")
	}

	return nil
}
