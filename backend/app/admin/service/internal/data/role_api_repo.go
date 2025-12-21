package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"go-wind-admin/app/admin/service/internal/data/ent"
	"go-wind-admin/app/admin/service/internal/data/ent/roleapi"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type RoleApiRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleApiRepo(ctx *bootstrap.Context, data *Data) *RoleApiRepo {
	return &RoleApiRepo{
		log:  ctx.NewLoggerHelper("role-api/repo/admin-service"),
		data: data,
	}
}

// AssignApis 给角色分配API
func (r *RoleApiRepo) AssignApis(ctx context.Context, roleId uint32, apiIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleApi.Delete().Where(roleapi.RoleID(roleId)).Exec(ctx); err != nil {
		err = entgo.Rollback(tx, err)
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
			SetCreatedBy(operatorId).
			SetCreatedAt(time.Now())
		roleApis = append(roleApis, rm)
	}

	_, err = r.data.db.Client().RoleApi.CreateBulk(roleApis...).Save(ctx)
	if err != nil {
		err = entgo.Rollback(tx, err)
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

// ListApiIdsByRoleId 获取角色分配的API ID列表
func (r *RoleApiRepo) ListApiIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
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
func (r *RoleApiRepo) RemoveApis(ctx context.Context, roleId uint32, apiIds []uint32) error {
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
