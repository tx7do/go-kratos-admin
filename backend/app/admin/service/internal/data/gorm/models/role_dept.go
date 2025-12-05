package models

import "github.com/tx7do/go-crud/gorm/mixin"

// RoleDept 对应表 sys_role_dept，角色 - 部门 关联表
type RoleDept struct {
	mixin.AutoIncrementID

	RoleID *uint32 `gorm:"column:role_id;type:int unsigned;comment:角色ID;index:idx_sys_role_dept_role_id;uniqueIndex:idx_sys_role_dept_role_id_dept_id"`
	DeptID *uint32 `gorm:"column:dept_id;type:int unsigned;comment:部门ID;uniqueIndex:idx_sys_role_dept_role_id_dept_id"`

	mixin.TimeAt
	mixin.OperatorID
}

// TableName 指定表名
func (RoleDept) TableName() string {
	return "sys_role_dept"
}
