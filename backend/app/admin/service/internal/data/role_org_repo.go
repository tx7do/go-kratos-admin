package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-crud/entgo"

	"go-wind-admin/app/admin/service/internal/data/ent"
	"go-wind-admin/app/admin/service/internal/data/ent/roleorg"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type RoleOrgRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleOrgRepo(data *Data, logger log.Logger) *RoleOrgRepo {
	return &RoleOrgRepo{
		log:  log.NewHelper(log.With(logger, "module", "role-org/repo/admin-service")),
		data: data,
	}
}

// AssignOrganizations 给角色分配组织
func (r *RoleOrgRepo) AssignOrganizations(ctx context.Context, roleId uint32, orgIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleOrg.Delete().Where(roleorg.RoleID(roleId)).Exec(ctx); err != nil {
		err = entgo.Rollback(tx, err)
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
			SetCreatedBy(operatorId).
			SetCreatedAt(time.Now())
		roleOrgs = append(roleOrgs, rm)
	}

	_, err = r.data.db.Client().RoleOrg.CreateBulk(roleOrgs...).Save(ctx)
	if err != nil {
		err = entgo.Rollback(tx, err)
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

// ListOrganizationIdsByRoleId 获取角色分配的组织ID列表
func (r *RoleOrgRepo) ListOrganizationIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
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
func (r *RoleOrgRepo) RemoveOrganizations(ctx context.Context, roleId uint32, ids []uint32) error {
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
