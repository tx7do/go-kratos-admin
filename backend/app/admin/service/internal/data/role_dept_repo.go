package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	entCrud "github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"go-wind-admin/app/admin/service/internal/data/ent"
	"go-wind-admin/app/admin/service/internal/data/ent/roledept"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type RoleDeptRepo struct {
	entClient *entCrud.EntClient[*ent.Client]
	log       *log.Helper
}

func NewRoleDeptRepo(ctx *bootstrap.Context, entClient *entCrud.EntClient[*ent.Client]) *RoleDeptRepo {
	return &RoleDeptRepo{
		log:       ctx.NewLoggerHelper("role-dept/repo/admin-service"),
		entClient: entClient,
	}
}

// AssignDepartments 给角色分配部门
func (r *RoleDeptRepo) AssignDepartments(ctx context.Context, roleId uint32, deptIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.entClient.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleDept.Delete().Where(roledept.RoleID(roleId)).Exec(ctx); err != nil {
		err = entCrud.Rollback(tx, err)
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
		rm := r.entClient.Client().RoleDept.
			Create().
			SetRoleID(roleId).
			SetDeptID(deptId).
			SetCreatedBy(operatorId).
			SetCreatedAt(time.Now())
		roleDepts = append(roleDepts, rm)
	}

	_, err = r.entClient.Client().RoleDept.CreateBulk(roleDepts...).Save(ctx)
	if err != nil {
		err = entCrud.Rollback(tx, err)
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

// ListDepartmentIdsByRoleId 获取角色分配的部门ID列表
func (r *RoleDeptRepo) ListDepartmentIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
	ids, err := r.entClient.Client().RoleDept.Query().
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
func (r *RoleDeptRepo) RemoveDepartments(ctx context.Context, roleId uint32, ids []uint32) error {
	_, err := r.entClient.Client().RoleDept.Delete().
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
