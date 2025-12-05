package models

import "github.com/tx7do/go-crud/gorm/mixin"

// Department 对应表 sys_departments
type Department struct {
	mixin.AutoIncrementID

	Name           *string `gorm:"column:name;type:varchar(255);comment:部门名称"`
	OrganizationID *uint32 `gorm:"column:organization_id;type:int unsigned;comment:所属组织ID"`
	ManagerID      *uint32 `gorm:"column:manager_id;type:int unsigned;comment:负责人ID"`
	Status         *string `gorm:"column:status;type:varchar(32);default:ON;comment:部门状态"`
	Description    *string `gorm:"column:description;type:varchar(255);comment:职能描述"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.SortOrder
	mixin.Remark
	mixin.TenantID
}

// TableName 指定表名
func (Department) TableName() string {
	return "sys_departments"
}
