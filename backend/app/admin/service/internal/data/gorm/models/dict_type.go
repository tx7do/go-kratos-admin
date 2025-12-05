package models

import "github.com/tx7do/go-crud/gorm/mixin"

// DictType 对应表 sys_dict_types
type DictType struct {
	mixin.AutoIncrementID

	TypeCode *string `gorm:"column:type_code;type:varchar(128);comment:字典类型唯一代码"`
	TypeName *string `gorm:"column:type_name;type:varchar(255);comment:字典类型名称"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.IsEnabled
	mixin.SortOrder
	mixin.Description
	mixin.TenantID

	// 关联字典项，DictEntry 需在同一包中定义并包含 TypeID 字段
	Entries []DictEntry `gorm:"foreignKey:TypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName 指定表名
func (DictType) TableName() string {
	return "sys_dict_types"
}
