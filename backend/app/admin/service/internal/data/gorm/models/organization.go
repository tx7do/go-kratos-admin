package models

import "github.com/tx7do/go-crud/gorm/mixin"

// Organization 对应表 sys_organizations
type Organization struct {
	mixin.AutoIncrementID

	Name             *string `gorm:"column:name;type:varchar(255);comment:组织名称"`
	Status           *string `gorm:"column:status;type:varchar(32);default:ON;comment:组织状态"`
	OrganizationType *string `gorm:"column:organization_type;type:varchar(64);comment:组织类型"`
	CreditCode       *string `gorm:"column:credit_code;type:varchar(255);comment:统一社会信用代码"`
	Address          *string `gorm:"column:address;type:varchar(1024);comment:注册地址"`
	BusinessScope    *string `gorm:"column:business_scope;type:varchar(1024);comment:核心业务范围"`
	IsLegalEntity    *bool   `gorm:"column:is_legal_entity;type:tinyint(1);default:0;comment:是否法人实体"`
	ManagerID        *uint32 `gorm:"column:manager_id;type:int unsigned;comment:负责人ID"`

	// 简单树结构字段
	ParentID *uint32 `gorm:"column:parent_id;type:int unsigned;comment:父级ID"`
	Path     *string `gorm:"column:path;type:varchar(1024);comment:路径"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.Remark
	mixin.SortOrder
	mixin.TenantID
}

// TableName 指定表名
func (Organization) TableName() string {
	return "sys_organizations"
}
