package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-crud/entgo"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/userposition"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type UserPositionRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserPositionRepo(data *Data, logger log.Logger) *UserPositionRepo {
	return &UserPositionRepo{
		log:  log.NewHelper(log.With(logger, "module", "user-position/repo/admin-service")),
		data: data,
	}
}

// AssignPositions 分配岗位给用户
func (r *UserPositionRepo) AssignPositions(ctx context.Context, userId uint32, ids []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该用户的所有旧关联
	if _, err = tx.UserPosition.Delete().Where(userposition.UserID(userId)).Exec(ctx); err != nil {
		err = entgo.Rollback(tx, err)
		r.log.Errorf("delete old user positions failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete old user positions failed")
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

	var userPositions []*ent.UserPositionCreate
	for _, id := range ids {
		rm := r.data.db.Client().UserPosition.
			Create().
			SetUserID(userId).
			SetPositionID(id).
			SetCreatedBy(operatorId).
			SetCreatedAt(time.Now())
		userPositions = append(userPositions, rm)
	}

	_, err = r.data.db.Client().UserPosition.CreateBulk(userPositions...).Save(ctx)
	if err != nil {
		err = entgo.Rollback(tx, err)
		r.log.Errorf("assign positions to user failed: %s", err.Error())
		return userV1.ErrorInternalServerError("assign positions to user failed")
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		r.log.Errorf("commit transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("commit transaction failed")
	}

	return nil
}

// GetPositionIdsByUserId 获取用户的岗位ID列表
func (r *UserPositionRepo) GetPositionIdsByUserId(ctx context.Context, userId uint32) ([]uint32, error) {
	ids, err := r.data.db.Client().UserPosition.Query().
		Where(userposition.UserIDEQ(userId)).
		Select(userposition.FieldPositionID).
		IDs(ctx)
	if err != nil {
		r.log.Errorf("query position ids by user id failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query position ids by user id failed")
	}
	return ids, nil
}

// RemovePositions 从用户移除岗位
func (r *UserPositionRepo) RemovePositions(ctx context.Context, userId uint32, ids []uint32) error {
	_, err := r.data.db.Client().UserPosition.Delete().
		Where(
			userposition.And(
				userposition.UserIDEQ(userId),
				userposition.PositionIDIn(ids...),
			),
		).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("remove positions from user failed: %s", err.Error())
		return userV1.ErrorInternalServerError("remove positions from user failed")
	}
	return nil
}
