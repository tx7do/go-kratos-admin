package models

import "github.com/tx7do/go-crud/gorm/mixin"

// UserPosition 对应表 sys_user_position，用户 - 职位 关联表
type UserPosition struct {
	mixin.AutoIncrementID

	UserID     *uint32 `gorm:"column:user_id;type:int unsigned;comment:用户ID;index:idx_sys_user_position_user_id;uniqueIndex:idx_sys_user_position_user_id_position_id,priority:1"`
	PositionID *uint32 `gorm:"column:position_id;type:int unsigned;comment:职位ID;index:idx_sys_user_position_position_id;uniqueIndex:idx_sys_user_position_user_id_position_id,priority:2"`

	mixin.TimeAt
	mixin.OperatorID
}

// TableName 指定表名
func (UserPosition) TableName() string {
	return "sys_user_position"
}
