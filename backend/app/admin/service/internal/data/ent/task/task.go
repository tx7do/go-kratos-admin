// Code generated by ent, DO NOT EDIT.

package task

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldCreateBy holds the string denoting the create_by field in the database.
	FieldCreateBy = "create_by"
	// FieldUpdateBy holds the string denoting the update_by field in the database.
	FieldUpdateBy = "update_by"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTypeName holds the string denoting the type_name field in the database.
	FieldTypeName = "type_name"
	// FieldTaskPayload holds the string denoting the task_payload field in the database.
	FieldTaskPayload = "task_payload"
	// FieldCronSpec holds the string denoting the cron_spec field in the database.
	FieldCronSpec = "cron_spec"
	// FieldTaskOptions holds the string denoting the task_options field in the database.
	FieldTaskOptions = "task_options"
	// FieldEnable holds the string denoting the enable field in the database.
	FieldEnable = "enable"
	// Table holds the table name of the task in the database.
	Table = "sys_tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldCreateBy,
	FieldUpdateBy,
	FieldRemark,
	FieldTenantID,
	FieldType,
	FieldTypeName,
	FieldTaskPayload,
	FieldCronSpec,
	FieldTaskOptions,
	FieldEnable,
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
	// TenantIDValidator is a validator for the "tenant_id" field. It is called by the builders before save.
	TenantIDValidator func(uint32) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypePeriodic   Type = "PERIODIC"
	TypeDelay      Type = "DELAY"
	TypeWaitResult Type = "WAIT_RESULT"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypePeriodic, TypeDelay, TypeWaitResult:
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Task queries.
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

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByTypeName orders the results by the type_name field.
func ByTypeName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTypeName, opts...).ToFunc()
}

// ByTaskPayload orders the results by the task_payload field.
func ByTaskPayload(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaskPayload, opts...).ToFunc()
}

// ByCronSpec orders the results by the cron_spec field.
func ByCronSpec(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCronSpec, opts...).ToFunc()
}

// ByEnable orders the results by the enable field.
func ByEnable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnable, opts...).ToFunc()
}
