// Code generated by ent, DO NOT EDIT.

package dict

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the dict type in the database.
	Label = "dict"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateBy holds the string denoting the create_by field in the database.
	FieldCreateBy = "create_by"
	// FieldUpdateBy holds the string denoting the update_by field in the database.
	FieldUpdateBy = "update_by"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldCategoryDesc holds the string denoting the category_desc field in the database.
	FieldCategoryDesc = "category_desc"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldValueDesc holds the string denoting the value_desc field in the database.
	FieldValueDesc = "value_desc"
	// FieldValueDataType holds the string denoting the value_data_type field in the database.
	FieldValueDataType = "value_data_type"
	// FieldSortID holds the string denoting the sort_id field in the database.
	FieldSortID = "sort_id"
	// Table holds the table name of the dict in the database.
	Table = "dict"
)

// Columns holds all SQL columns for dict fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldStatus,
	FieldCreateBy,
	FieldUpdateBy,
	FieldRemark,
	FieldKey,
	FieldCategory,
	FieldCategoryDesc,
	FieldValue,
	FieldValueDesc,
	FieldValueDataType,
	FieldSortID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
	// DefaultSortID holds the default value on creation for the "sort_id" field.
	DefaultSortID int32
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusON is the default value of the Status enum.
const DefaultStatus = StatusON

// Status values.
const (
	StatusOFF Status = "OFF"
	StatusON  Status = "ON"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOFF, StatusON:
		return nil
	default:
		return fmt.Errorf("dict: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Dict queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreateBy orders the results by the create_by field.
func ByCreateBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateBy, opts...).ToFunc()
}

// ByUpdateBy orders the results by the update_by field.
func ByUpdateBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateBy, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByCategoryDesc orders the results by the category_desc field.
func ByCategoryDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryDesc, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByValueDesc orders the results by the value_desc field.
func ByValueDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValueDesc, opts...).ToFunc()
}

// ByValueDataType orders the results by the value_data_type field.
func ByValueDataType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValueDataType, opts...).ToFunc()
}

// BySortID orders the results by the sort_id field.
func BySortID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSortID, opts...).ToFunc()
}
