package data

import (
	"context"
	"kratos-admin/app/admin/service/internal/data/ent/roledept"
	"kratos-admin/app/admin/service/internal/data/ent/roleorg"
	"kratos-admin/app/admin/service/internal/data/ent/roleposition"

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
	"kratos-admin/app/admin/service/internal/data/ent/role"
	"kratos-admin/app/admin/service/internal/data/ent/roleapi"
	"kratos-admin/app/admin/service/internal/data/ent/rolemenu"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper

	mapper *mapper.CopierMapper[userV1.Role, ent.Role]
}

func NewRoleRepo(data *Data, logger log.Logger) *RoleRepo {
	repo := &RoleRepo{
		log:    log.NewHelper(log.With(logger, "module", "role/repo/admin-service")),
		data:   data,
		mapper: mapper.NewCopierMapper[userV1.Role, ent.Role](),
	}

	repo.init()

	return repo
}

func (r *RoleRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
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

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), role.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, userV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query list failed")
	}

	dtos := make([]*userV1.Role, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListRoleResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
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

func (r *RoleRepo) Get(ctx context.Context, id uint32) (*userV1.Role, error) {
	if id == 0 {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().Role.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorRoleNotFound("role not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
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
		SetNillableSortID(req.Data.SortId).
		SetNillableCode(req.Data.Code).
		SetNillableStatus((*role.Status)(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
		SetNillableTenantID(req.Data.TenantId).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
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
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return userV1.ErrorBadRequest("invalid field mask")
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
		SetNillableUpdateBy(req.Data.UpdateBy).
		SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if req.Data.Menus != nil {
		builder.SetMenus(req.Data.Menus)
	}
	if req.Data.Apis != nil {
		builder.SetApis(req.Data.Apis)
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
		return userV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *RoleRepo) Delete(ctx context.Context, req *userV1.DeleteRoleRequest) error {
	if req == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := queryAllChildrenIDs(ctx, r.data.db, "sys_roles", req.GetId())
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

// AssignMenus 给角色分配菜单
func (r *RoleRepo) AssignMenus(ctx context.Context, roleId uint32, menuIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleMenu.Delete().Where(rolemenu.RoleID(roleId)).Exec(ctx); err != nil {
		err = rollback(tx, err)
		r.log.Errorf("delete old role menus failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete old role menus failed")
	}

	// 如果没有分配任何菜单，则直接提交事务返回
	if len(menuIds) == 0 {
		// 提交事务
		if err = tx.Commit(); err != nil {
			r.log.Errorf("commit transaction failed: %s", err.Error())
			return userV1.ErrorInternalServerError("commit transaction failed")
		}
		return nil
	}

	var roleMenus []*ent.RoleMenuCreate
	for _, menuID := range menuIds {
		rm := r.data.db.Client().RoleMenu.
			Create().
			SetRoleID(roleId).
			SetMenuID(menuID).
			SetCreateBy(operatorId).
			SetCreateTime(time.Now())
		roleMenus = append(roleMenus, rm)
	}

	_, err = r.data.db.Client().RoleMenu.CreateBulk(roleMenus...).Save(ctx)
	if err != nil {
		err = rollback(tx, err)
		r.log.Errorf("assign menus to role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("assign menus to role failed")
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		r.log.Errorf("commit transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("commit transaction failed")
	}

	return nil
}

// GetMenuIdsByRoleId 获取角色分配的菜单ID列表
func (r *RoleRepo) GetMenuIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
	menuIds, err := r.data.db.Client().RoleMenu.Query().
		Where(rolemenu.RoleIDEQ(roleId)).
		Select(rolemenu.FieldMenuID).
		IDs(ctx)
	if err != nil {
		r.log.Errorf("query menu ids by role id failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query menu ids by role id failed")
	}
	return menuIds, nil
}

// RemoveMenus 从角色移除菜单
func (r *RoleRepo) RemoveMenus(ctx context.Context, roleId uint32, menuIds []uint32) error {
	_, err := r.data.db.Client().RoleMenu.Delete().
		Where(
			rolemenu.And(
				rolemenu.RoleIDEQ(roleId),
				rolemenu.MenuIDIn(menuIds...),
			),
		).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("remove menus from role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("remove menus from role failed")
	}
	return nil
}

// AssignApis 给角色分配API
func (r *RoleRepo) AssignApis(ctx context.Context, roleId uint32, apiIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleApi.Delete().Where(roleapi.RoleID(roleId)).Exec(ctx); err != nil {
		err = rollback(tx, err)
		r.log.Errorf("delete old role apis failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete old role apis failed")
	}

	// 如果没有分配任何，则直接提交事务返回
	if len(apiIds) == 0 {
		// 提交事务
		if err = tx.Commit(); err != nil {
			r.log.Errorf("commit transaction failed: %s", err.Error())
			return userV1.ErrorInternalServerError("commit transaction failed")
		}
		return nil
	}

	var roleApis []*ent.RoleApiCreate
	for _, apiId := range apiIds {
		rm := r.data.db.Client().RoleApi.
			Create().
			SetRoleID(roleId).
			SetAPIID(apiId).
			SetCreateBy(operatorId).
			SetCreateTime(time.Now())
		roleApis = append(roleApis, rm)
	}

	_, err = r.data.db.Client().RoleApi.CreateBulk(roleApis...).Save(ctx)
	if err != nil {
		err = rollback(tx, err)
		r.log.Errorf("assign apis to role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("assign apis to role failed")
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		r.log.Errorf("commit transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("commit transaction failed")
	}

	return nil
}

// GetApiIdsByRoleId 获取角色分配的API ID列表
func (r *RoleRepo) GetApiIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
	apiIds, err := r.data.db.Client().RoleApi.Query().
		Where(roleapi.IDEQ(roleId)).
		Select(roleapi.FieldAPIID).
		IDs(ctx)
	if err != nil {
		r.log.Errorf("query api ids by role id failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query api ids by role id failed")
	}
	return apiIds, nil
}

// RemoveApis 从角色移除API
func (r *RoleRepo) RemoveApis(ctx context.Context, roleId uint32, apiIds []uint32) error {
	_, err := r.data.db.Client().RoleApi.Delete().
		Where(
			roleapi.And(
				roleapi.RoleIDEQ(roleId),
				roleapi.APIIDIn(apiIds...),
			),
		).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("remove apis from role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("remove apis from role failed")
	}
	return nil
}

// AssignOrganizations 给角色分配组织
func (r *RoleRepo) AssignOrganizations(ctx context.Context, roleId uint32, orgIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleOrg.Delete().Where(roleorg.RoleID(roleId)).Exec(ctx); err != nil {
		err = rollback(tx, err)
		r.log.Errorf("delete old role organizations failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete old role organizations failed")
	}

	// 如果没有分配任何，则直接提交事务返回
	if len(orgIds) == 0 {
		// 提交事务
		if err = tx.Commit(); err != nil {
			r.log.Errorf("commit transaction failed: %s", err.Error())
			return userV1.ErrorInternalServerError("commit transaction failed")
		}
		return nil
	}

	var roleOrgs []*ent.RoleOrgCreate
	for _, orgId := range orgIds {
		rm := r.data.db.Client().RoleOrg.
			Create().
			SetRoleID(roleId).
			SetOrgID(orgId).
			SetCreateBy(operatorId).
			SetCreateTime(time.Now())
		roleOrgs = append(roleOrgs, rm)
	}

	_, err = r.data.db.Client().RoleOrg.CreateBulk(roleOrgs...).Save(ctx)
	if err != nil {
		err = rollback(tx, err)
		r.log.Errorf("assign organizations to role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("assign organizations to role failed")
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		r.log.Errorf("commit transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("commit transaction failed")
	}

	return nil
}

// GetOrganizationIdsByRoleId 获取角色分配的组织ID列表
func (r *RoleRepo) GetOrganizationIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
	ids, err := r.data.db.Client().RoleOrg.Query().
		Where(roleorg.RoleIDEQ(roleId)).
		Select(roleorg.FieldOrgID).
		IDs(ctx)
	if err != nil {
		r.log.Errorf("query organization ids by role id failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query organization ids by role id failed")
	}
	return ids, nil
}

// RemoveOrganizations 从角色移除组织
func (r *RoleRepo) RemoveOrganizations(ctx context.Context, roleId uint32, ids []uint32) error {
	_, err := r.data.db.Client().RoleOrg.Delete().
		Where(
			roleorg.And(
				roleorg.RoleIDEQ(roleId),
				roleorg.OrgIDIn(ids...),
			),
		).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("remove organizations from role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("remove organizations from role failed")
	}
	return nil
}

// AssignDepartments 给角色分配部门
func (r *RoleRepo) AssignDepartments(ctx context.Context, roleId uint32, deptIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleDept.Delete().Where(roledept.RoleID(roleId)).Exec(ctx); err != nil {
		err = rollback(tx, err)
		r.log.Errorf("delete old role departments failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete old role departments failed")
	}

	// 如果没有分配任何，则直接提交事务返回
	if len(deptIds) == 0 {
		// 提交事务
		if err = tx.Commit(); err != nil {
			r.log.Errorf("commit transaction failed: %s", err.Error())
			return userV1.ErrorInternalServerError("commit transaction failed")
		}
		return nil
	}

	var roleDepts []*ent.RoleDeptCreate
	for _, deptId := range deptIds {
		rm := r.data.db.Client().RoleDept.
			Create().
			SetRoleID(roleId).
			SetDeptID(deptId).
			SetCreateBy(operatorId).
			SetCreateTime(time.Now())
		roleDepts = append(roleDepts, rm)
	}

	_, err = r.data.db.Client().RoleDept.CreateBulk(roleDepts...).Save(ctx)
	if err != nil {
		err = rollback(tx, err)
		r.log.Errorf("assign departments to role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("assign departments to role failed")
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		r.log.Errorf("commit transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("commit transaction failed")
	}

	return nil
}

// GetDepartmentIdsByRoleId 获取角色分配的部门ID列表
func (r *RoleRepo) GetDepartmentIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
	ids, err := r.data.db.Client().RoleDept.Query().
		Where(roledept.RoleIDEQ(roleId)).
		Select(roledept.FieldDeptID).
		IDs(ctx)
	if err != nil {
		r.log.Errorf("query department ids by role id failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query department ids by role id failed")
	}
	return ids, nil
}

// RemoveDepartments 从角色移除部门
func (r *RoleRepo) RemoveDepartments(ctx context.Context, roleId uint32, ids []uint32) error {
	_, err := r.data.db.Client().RoleDept.Delete().
		Where(
			roledept.And(
				roledept.RoleIDEQ(roleId),
				roledept.DeptIDIn(ids...),
			),
		).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("remove departments from role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("remove departments from role failed")
	}
	return nil
}

// AssignPositions 给角色分配岗位
func (r *RoleRepo) AssignPositions(ctx context.Context, roleId uint32, positionIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RolePosition.Delete().Where(roleposition.RoleID(roleId)).Exec(ctx); err != nil {
		err = rollback(tx, err)
		r.log.Errorf("delete old role positions failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete old role positions failed")
	}

	// 如果没有分配任何，则直接提交事务返回
	if len(positionIds) == 0 {
		// 提交事务
		if err = tx.Commit(); err != nil {
			r.log.Errorf("commit transaction failed: %s", err.Error())
			return userV1.ErrorInternalServerError("commit transaction failed")
		}
		return nil
	}

	var rolePositions []*ent.RolePositionCreate
	for _, positionId := range positionIds {
		rm := r.data.db.Client().RolePosition.
			Create().
			SetRoleID(roleId).
			SetPositionID(positionId).
			SetCreateBy(operatorId).
			SetCreateTime(time.Now())
		rolePositions = append(rolePositions, rm)
	}

	_, err = r.data.db.Client().RolePosition.CreateBulk(rolePositions...).Save(ctx)
	if err != nil {
		err = rollback(tx, err)
		r.log.Errorf("assign positions to role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("assign positions to role failed")
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		r.log.Errorf("commit transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("commit transaction failed")
	}

	return nil
}

// GetPositionIdsByRoleId 获取角色分配的岗位ID列表
func (r *RoleRepo) GetPositionIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
	ids, err := r.data.db.Client().RolePosition.Query().
		Where(roleposition.RoleIDEQ(roleId)).
		Select(roleposition.FieldPositionID).
		IDs(ctx)
	if err != nil {
		r.log.Errorf("query position ids by role id failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query position ids by role id failed")
	}
	return ids, nil
}

// RemovePositions 从角色移除岗位
func (r *RoleRepo) RemovePositions(ctx context.Context, roleId uint32, ids []uint32) error {
	_, err := r.data.db.Client().RolePosition.Delete().
		Where(
			roleposition.And(
				roleposition.RoleIDEQ(roleId),
				roleposition.PositionIDIn(ids...),
			),
		).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("remove positions from role failed: %s", err.Error())
		return userV1.ErrorInternalServerError("remove positions from role failed")
	}
	return nil
}
