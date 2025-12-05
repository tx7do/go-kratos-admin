package models

import "github.com/tx7do/go-crud/gorm/mixin"

// RoleOrg 对应表 sys_role_org，角色 - 组织 关联表
type RoleOrg struct {
	mixin.AutoIncrementID

	RoleID *uint32 `gorm:"column:role_id;type:int unsigned;comment:角色ID;index:idx_sys_role_org_role_id;uniqueIndex:idx_sys_role_org_role_id_org_id,priority:1"`
	OrgID  *uint32 `gorm:"column:org_id;type:int unsigned;comment:组织ID;uniqueIndex:idx_sys_role_org_role_id_org_id,priority:2"`

	mixin.TimeAt
	mixin.OperatorID
}

// TableName 指定表名
func (RoleOrg) TableName() string {
	return "sys_role_org"
}
