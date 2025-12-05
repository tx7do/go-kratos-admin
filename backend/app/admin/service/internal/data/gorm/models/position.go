package models

import "github.com/tx7do/go-crud/gorm/mixin"

// Position 对应表 sys_positions
type Position struct {
	mixin.AutoIncrementID

	Name           *string `gorm:"column:name;type:varchar(255);comment:职位名称"`
	Code           *string `gorm:"column:code;type:varchar(128);comment:唯一编码"`
	OrganizationID *uint32 `gorm:"column:organization_id;type:int unsigned;comment:所属组织ID"`
	DepartmentID   *uint32 `gorm:"column:department_id;type:int unsigned;comment:所属部门ID"`
	Status         *string `gorm:"column:status;type:varchar(32);default:ON;comment:职位状态"`
	Description    *string `gorm:"column:description;type:varchar(1024);comment:职能描述"`
	Quota          *uint32 `gorm:"column:quota;type:int unsigned;comment:编制人数"`

	// 简单树结构字段（保留父级关系与路径）
	ParentID *uint32 `gorm:"column:parent_id;type:int unsigned;comment:父级ID"`
	Path     *string `gorm:"column:path;type:varchar(1024);comment:路径"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.SortOrder
	mixin.Remark
	mixin.TenantID
}

// TableName 指定表名
func (Position) TableName() string {
	return "sys_positions"
}
