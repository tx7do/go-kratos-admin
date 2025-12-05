package models

import (
	"gorm.io/datatypes"

	"github.com/tx7do/go-crud/gorm/mixin"
)

// Role 对应表 sys_roles
type Role struct {
	mixin.AutoIncrementID

	Name      *string        `gorm:"column:name;type:varchar(255);comment:角色名称"`
	Code      *string        `gorm:"column:code;type:varchar(128);comment:角色标识"`
	Menus     datatypes.JSON `gorm:"column:menus;type:json;comment:分配的菜单列表"`
	Apis      datatypes.JSON `gorm:"column:apis;type:json;comment:分配的API列表"`
	DataScope *string        `gorm:"column:data_scope;type:varchar(32);comment:数据权限范围"`
	Status    *string        `gorm:"column:status;type:varchar(32);default:ON;comment:角色状态"`

	// 简单树结构字段（保留父级关系与路径）
	ParentID *uint32 `gorm:"column:parent_id;type:int unsigned;comment:父级ID"`
	Path     *string `gorm:"column:path;type:varchar(1024);comment:节点路径"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.Remark
	mixin.SortOrder
	mixin.TenantID
}

// TableName 指定表名
func (Role) TableName() string {
	return "sys_roles"
}
