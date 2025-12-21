package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-crud/entgo"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"go-wind-admin/app/admin/service/internal/data/ent"
	"go-wind-admin/app/admin/service/internal/data/ent/rolemenu"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type RoleMenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleMenuRepo(ctx *bootstrap.Context, data *Data) *RoleMenuRepo {
	return &RoleMenuRepo{
		log:  ctx.NewLoggerHelper("role-menu/repo/admin-service"),
		data: data,
	}
}

// AssignMenus 给角色分配菜单
func (r *RoleMenuRepo) AssignMenus(ctx context.Context, roleId uint32, menuIds []uint32, operatorId uint32) error {
	// 开启事务
	tx, err := r.data.db.Client().Tx(ctx)
	if err != nil {
		r.log.Errorf("start transaction failed: %s", err.Error())
		return userV1.ErrorInternalServerError("start transaction failed")
	}

	// 删除该角色的所有旧关联
	if _, err = tx.RoleMenu.Delete().Where(rolemenu.RoleID(roleId)).Exec(ctx); err != nil {
		err = entgo.Rollback(tx, err)
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
			SetCreatedBy(operatorId).
			SetCreatedAt(time.Now())
		roleMenus = append(roleMenus, rm)
	}

	_, err = r.data.db.Client().RoleMenu.CreateBulk(roleMenus...).Save(ctx)
	if err != nil {
		err = entgo.Rollback(tx, err)
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

// ListMenuIdsByRoleId 获取角色分配的菜单ID列表
func (r *RoleMenuRepo) ListMenuIdsByRoleId(ctx context.Context, roleId uint32) ([]uint32, error) {
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
func (r *RoleMenuRepo) RemoveMenus(ctx context.Context, roleId uint32, menuIds []uint32) error {
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
