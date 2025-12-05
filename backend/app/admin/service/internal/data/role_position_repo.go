package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-crud/entgo"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/roleposition"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type RolePositionRepo struct {
	data *Data
	log  *log.Helper
}

func NewRolePositionRepo(data *Data, logger log.Logger) *RolePositionRepo {
	return &RolePositionRepo{
		log:  log.NewHelper(log.With(logger, "module", "role-position/repo/admin-service")),
		data: data,
	}
}

// AssignPositions 给角色分配岗位
func (r *RolePositionRepo) AssignPositions(ctx context.Context, roleId uint32, positionIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RolePosition.Delete().Where(roleposition.RoleID(roleId)).Exec(ctx); err != nil {
		err = entgo.Rollback(tx, err)
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
			SetCreatedBy(operatorId).
			SetCreatedAt(time.Now())
		rolePositions = append(rolePositions, rm)
	}

	_, err = r.data.db.Client().RolePosition.CreateBulk(rolePositions...).Save(ctx)
	if err != nil {
		err = entgo.Rollback(tx, err)
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

// ListPositionIdsByRoleId 获取角色分配的岗位ID列表
func (r *RolePositionRepo) ListPositionIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
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
func (r *RolePositionRepo) RemovePositions(ctx context.Context, roleId uint32, ids []uint32) error {
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
