package models

import "github.com/tx7do/go-crud/gorm/mixin"

// RolePosition 对应表 sys_role_position，角色 - 职位 关联表
type RolePosition struct {
	mixin.AutoIncrementID

	RoleID     *uint32 `gorm:"column:role_id;type:int unsigned;comment:角色ID;index:idx_sys_role_position_role_id;uniqueIndex:idx_sys_role_position_role_id_position_id,priority:1"`
	PositionID *uint32 `gorm:"column:position_id;type:int unsigned;comment:职位ID;uniqueIndex:idx_sys_role_position_role_id_position_id,priority:2"`

	mixin.TimeAt
	mixin.OperatorID
}

// TableName 指定表名
func (RolePosition) TableName() string {
	return "sys_role_position"
}
