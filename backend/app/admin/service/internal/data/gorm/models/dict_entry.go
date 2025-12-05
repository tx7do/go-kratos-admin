package models

import "github.com/tx7do/go-crud/gorm/mixin"

// DictEntry 对应表 sys_dict_entries
type DictEntry struct {
	mixin.AutoIncrementID

	EntryLabel   *string `gorm:"column:entry_label;type:varchar(255);comment:字典项的显示标签"`
	EntryValue   *string `gorm:"column:entry_value;type:varchar(255);comment:字典项的实际值"`
	NumericValue *int32  `gorm:"column:numeric_value;type:int;comment:数值型值"`
	LanguageCode *string `gorm:"column:language_code;type:varchar(32);comment:语言代码"`

	mixin.TimeAt
	mixin.OperatorID
	mixin.Description
	mixin.SortOrder
	mixin.IsEnabled
	mixin.TenantID
}

func (DictEntry) TableName() string {
	return "sys_dict_entries"
}
