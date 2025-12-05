package models

import "github.com/tx7do/go-crud/gorm/mixin"

// RoleMenu 对应表 sys_role_menu，角色 - 菜单 关联表
type RoleMenu struct {
	mixin.AutoIncrementID

	RoleID *uint32 `gorm:"column:role_id;type:int unsigned;comment:角色ID;index:idx_sys_role_menu_role_id;uniqueIndex:idx_sys_role_menu_role_id_menu_id,priority:1"`
	MenuID *uint32 `gorm:"column:menu_id;type:int unsigned;comment:菜单ID;uniqueIndex:idx_sys_role_menu_role_id_menu_id,priority:2"`

	mixin.TimeAt
	mixin.OperatorID
}

// TableName 指定表名
func (RoleMenu) TableName() string {
	return "sys_role_menu"
}
