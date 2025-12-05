package models

import "github.com/tx7do/go-crud/gorm/mixin"

// RoleApi 对应表 sys_role_api，角色 - API 关联表
type RoleApi struct {
	mixin.AutoIncrementID

	RoleID *uint32 `gorm:"column:role_id;type:int unsigned;comment:角色ID;index:idx_sys_role_api_role_id;uniqueIndex:idx_sys_role_api_role_id_api_id"`
	ApiID  *uint32 `gorm:"column:api_id;type:int unsigned;comment:API ID;uniqueIndex:idx_sys_role_api_role_id_api_id"`

	mixin.TimeAt
	mixin.OperatorID
}

// TableName 指定表名
func (RoleApi) TableName() string {
	return "sys_role_api"
}
