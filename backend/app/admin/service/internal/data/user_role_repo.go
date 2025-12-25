package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	entCrud "github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"go-wind-admin/app/admin/service/internal/data/ent"
	"go-wind-admin/app/admin/service/internal/data/ent/userrole"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type UserRoleRepo struct {
	entClient *entCrud.EntClient[*ent.Client]
	log       *log.Helper
}

func NewUserRoleRepo(ctx *bootstrap.Context, entClient *entCrud.EntClient[*ent.Client]) *UserRoleRepo {
	return &UserRoleRepo{
		log:       ctx.NewLoggerHelper("user-role/repo/admin-service"),
		entClient: entClient,
	}
}

// AssignRoles 分配角色给用户
func (r *UserRoleRepo) AssignRoles(ctx context.Context, userId uint32, ids []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.entClient.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该用户的所有旧关联
	if _, err = tx.UserRole.Delete().Where(userrole.UserID(userId)).Exec(ctx); err != nil {
		err = entCrud.Rollback(tx, err)
		r.log.Errorf("delete old user roles failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete old user roles failed")
	}

	// 如果没有分配任何，则直接提交事务返回
	if len(ids) == 0 {
		// 提交事务
		if err = tx.Commit(); err != nil {
			r.log.Errorf("commit transaction failed: %s", err.Error())
			return userV1.ErrorInternalServerError("commit transaction failed")
		}
		return nil
	}

	var userRoles []*ent.UserRoleCreate
	for _, id := range ids {
		rm := r.entClient.Client().UserRole.
			Create().
			SetUserID(userId).
			SetRoleID(id).
			SetCreatedBy(operatorId).
			SetCreatedAt(time.Now())
		userRoles = append(userRoles, rm)
	}

	_, err = r.entClient.Client().UserRole.CreateBulk(userRoles...).Save(ctx)
	if err != nil {
		err = entCrud.Rollback(tx, err)
		r.log.Errorf("assign roles to user failed: %s", err.Error())
		return userV1.ErrorInternalServerError("assign roles to user failed")
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		r.log.Errorf("commit transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("commit transaction failed")
	}

	return nil
}

// ListRoleIdsByUserId 获取用户关联的角色ID列表
func (r *UserRoleRepo) ListRoleIdsByUserId(ctx context.Context, userId uint32) ([]uint32, error) {
	ids, err := r.entClient.Client().UserRole.Query().
		Where(userrole.UserIDEQ(userId)).
		Select(userrole.FieldRoleID).
		IDs(ctx)
	if err != nil {
		r.log.Errorf("query role ids by user id failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query role ids by user id failed")
	}
	return ids, nil
}

// RemoveRoles 从用户移除角色
func (r *UserRoleRepo) RemoveRoles(ctx context.Context, userId uint32, ids []uint32) error {
	_, err := r.entClient.Client().UserRole.Delete().
		Where(
			userrole.And(
				userrole.UserIDEQ(userId),
				userrole.RoleIDIn(ids...),
			),
		).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("remove roles from user failed: %s", err.Error())
		return userV1.ErrorInternalServerError("remove roles from user failed")
	}
	return nil
}
