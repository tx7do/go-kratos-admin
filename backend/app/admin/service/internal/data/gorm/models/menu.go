package models

import (
	"gorm.io/datatypes"

	"github.com/tx7do/go-crud/gorm/mixin"
)

// Menu 对应表 sys_menus
type Menu struct {
	mixin.AutoIncrementID

	Status    *string        `gorm:"column:status;type:varchar(32);default:ON;comment:菜单状态"`
	Type      *string        `gorm:"column:type;type:varchar(32);default:MENU;comment:菜单类型 FOLDER: 目录 MENU: 菜单 BUTTON: 按钮 EMBEDDED: 内嵌 LINK: 外链"`
	Path      *string        `gorm:"column:path;type:varchar(1024);comment:路径，类型为按钮时为操作名"`
	Redirect  *string        `gorm:"column:redirect;type:varchar(1024);comment:重定向地址"`
	Alias     *string        `gorm:"column:alias;type:varchar(255);comment:路由别名"`
	Name      *string        `gorm:"column:name;type:varchar(255);comment:路由命名"`
	Component *string        `gorm:"column:component;type:varchar(255);comment:前端页面组件"`
	Meta      datatypes.JSON `gorm:"column:meta;type:json;comment:前端页面组件元数据"`

	// 简单树结构字段（保留父级关系与路径）
	ParentID *uint32 `gorm:"column:parent_id;type:int unsigned;comment:父级ID"`
	TreePath *string `gorm:"column:tree_path;type:varchar(1024);comment:节点路径"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.Remark
}

// TableName 指定表名
func (Menu) TableName() string {
	return "sys_menus"
}
