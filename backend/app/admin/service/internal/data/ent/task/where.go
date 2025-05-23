// Code generated by ent, DO NOT EDIT.

package task

import (
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeleteTime, v))
}

// CreateBy applies equality check predicate on the "create_by" field. It's identical to CreateByEQ.
func CreateBy(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateBy, v))
}

// UpdateBy applies equality check predicate on the "update_by" field. It's identical to UpdateByEQ.
func UpdateBy(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateBy, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldRemark, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTenantID, v))
}

// TypeName applies equality check predicate on the "type_name" field. It's identical to TypeNameEQ.
func TypeName(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTypeName, v))
}

// TaskPayload applies equality check predicate on the "task_payload" field. It's identical to TaskPayloadEQ.
func TaskPayload(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskPayload, v))
}

// CronSpec applies equality check predicate on the "cron_spec" field. It's identical to CronSpecEQ.
func CronSpec(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCronSpec, v))
}

// Enable applies equality check predicate on the "enable" field. It's identical to EnableEQ.
func Enable(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldEnable, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldDeleteTime))
}

// CreateByEQ applies the EQ predicate on the "create_by" field.
func CreateByEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateBy, v))
}

// CreateByNEQ applies the NEQ predicate on the "create_by" field.
func CreateByNEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreateBy, v))
}

// CreateByIn applies the In predicate on the "create_by" field.
func CreateByIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreateBy, vs...))
}

// CreateByNotIn applies the NotIn predicate on the "create_by" field.
func CreateByNotIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreateBy, vs...))
}

// CreateByGT applies the GT predicate on the "create_by" field.
func CreateByGT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreateBy, v))
}

// CreateByGTE applies the GTE predicate on the "create_by" field.
func CreateByGTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreateBy, v))
}

// CreateByLT applies the LT predicate on the "create_by" field.
func CreateByLT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreateBy, v))
}

// CreateByLTE applies the LTE predicate on the "create_by" field.
func CreateByLTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreateBy, v))
}

// CreateByIsNil applies the IsNil predicate on the "create_by" field.
func CreateByIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldCreateBy))
}

// CreateByNotNil applies the NotNil predicate on the "create_by" field.
func CreateByNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldCreateBy))
}

// UpdateByEQ applies the EQ predicate on the "update_by" field.
func UpdateByEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateBy, v))
}

// UpdateByNEQ applies the NEQ predicate on the "update_by" field.
func UpdateByNEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldUpdateBy, v))
}

// UpdateByIn applies the In predicate on the "update_by" field.
func UpdateByIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldUpdateBy, vs...))
}

// UpdateByNotIn applies the NotIn predicate on the "update_by" field.
func UpdateByNotIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldUpdateBy, vs...))
}

// UpdateByGT applies the GT predicate on the "update_by" field.
func UpdateByGT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldUpdateBy, v))
}

// UpdateByGTE applies the GTE predicate on the "update_by" field.
func UpdateByGTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldUpdateBy, v))
}

// UpdateByLT applies the LT predicate on the "update_by" field.
func UpdateByLT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldUpdateBy, v))
}

// UpdateByLTE applies the LTE predicate on the "update_by" field.
func UpdateByLTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldUpdateBy, v))
}

// UpdateByIsNil applies the IsNil predicate on the "update_by" field.
func UpdateByIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldUpdateBy))
}

// UpdateByNotNil applies the NotNil predicate on the "update_by" field.
func UpdateByNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldUpdateBy))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldRemark, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDIsNil applies the IsNil predicate on the "tenant_id" field.
func TenantIDIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldTenantID))
}

// TenantIDNotNil applies the NotNil predicate on the "tenant_id" field.
func TenantIDNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldTenantID))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldType, vs...))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldType))
}

// TypeNameEQ applies the EQ predicate on the "type_name" field.
func TypeNameEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTypeName, v))
}

// TypeNameNEQ applies the NEQ predicate on the "type_name" field.
func TypeNameNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTypeName, v))
}

// TypeNameIn applies the In predicate on the "type_name" field.
func TypeNameIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTypeName, vs...))
}

// TypeNameNotIn applies the NotIn predicate on the "type_name" field.
func TypeNameNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTypeName, vs...))
}

// TypeNameGT applies the GT predicate on the "type_name" field.
func TypeNameGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTypeName, v))
}

// TypeNameGTE applies the GTE predicate on the "type_name" field.
func TypeNameGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTypeName, v))
}

// TypeNameLT applies the LT predicate on the "type_name" field.
func TypeNameLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTypeName, v))
}

// TypeNameLTE applies the LTE predicate on the "type_name" field.
func TypeNameLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTypeName, v))
}

// TypeNameContains applies the Contains predicate on the "type_name" field.
func TypeNameContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTypeName, v))
}

// TypeNameHasPrefix applies the HasPrefix predicate on the "type_name" field.
func TypeNameHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTypeName, v))
}

// TypeNameHasSuffix applies the HasSuffix predicate on the "type_name" field.
func TypeNameHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTypeName, v))
}

// TypeNameIsNil applies the IsNil predicate on the "type_name" field.
func TypeNameIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldTypeName))
}

// TypeNameNotNil applies the NotNil predicate on the "type_name" field.
func TypeNameNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldTypeName))
}

// TypeNameEqualFold applies the EqualFold predicate on the "type_name" field.
func TypeNameEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTypeName, v))
}

// TypeNameContainsFold applies the ContainsFold predicate on the "type_name" field.
func TypeNameContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTypeName, v))
}

// TaskPayloadEQ applies the EQ predicate on the "task_payload" field.
func TaskPayloadEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskPayload, v))
}

// TaskPayloadNEQ applies the NEQ predicate on the "task_payload" field.
func TaskPayloadNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTaskPayload, v))
}

// TaskPayloadIn applies the In predicate on the "task_payload" field.
func TaskPayloadIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTaskPayload, vs...))
}

// TaskPayloadNotIn applies the NotIn predicate on the "task_payload" field.
func TaskPayloadNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTaskPayload, vs...))
}

// TaskPayloadGT applies the GT predicate on the "task_payload" field.
func TaskPayloadGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTaskPayload, v))
}

// TaskPayloadGTE applies the GTE predicate on the "task_payload" field.
func TaskPayloadGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTaskPayload, v))
}

// TaskPayloadLT applies the LT predicate on the "task_payload" field.
func TaskPayloadLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTaskPayload, v))
}

// TaskPayloadLTE applies the LTE predicate on the "task_payload" field.
func TaskPayloadLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTaskPayload, v))
}

// TaskPayloadContains applies the Contains predicate on the "task_payload" field.
func TaskPayloadContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTaskPayload, v))
}

// TaskPayloadHasPrefix applies the HasPrefix predicate on the "task_payload" field.
func TaskPayloadHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTaskPayload, v))
}

// TaskPayloadHasSuffix applies the HasSuffix predicate on the "task_payload" field.
func TaskPayloadHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTaskPayload, v))
}

// TaskPayloadIsNil applies the IsNil predicate on the "task_payload" field.
func TaskPayloadIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldTaskPayload))
}

// TaskPayloadNotNil applies the NotNil predicate on the "task_payload" field.
func TaskPayloadNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldTaskPayload))
}

// TaskPayloadEqualFold applies the EqualFold predicate on the "task_payload" field.
func TaskPayloadEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTaskPayload, v))
}

// TaskPayloadContainsFold applies the ContainsFold predicate on the "task_payload" field.
func TaskPayloadContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTaskPayload, v))
}

// CronSpecEQ applies the EQ predicate on the "cron_spec" field.
func CronSpecEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCronSpec, v))
}

// CronSpecNEQ applies the NEQ predicate on the "cron_spec" field.
func CronSpecNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCronSpec, v))
}

// CronSpecIn applies the In predicate on the "cron_spec" field.
func CronSpecIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCronSpec, vs...))
}

// CronSpecNotIn applies the NotIn predicate on the "cron_spec" field.
func CronSpecNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCronSpec, vs...))
}

// CronSpecGT applies the GT predicate on the "cron_spec" field.
func CronSpecGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCronSpec, v))
}

// CronSpecGTE applies the GTE predicate on the "cron_spec" field.
func CronSpecGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCronSpec, v))
}

// CronSpecLT applies the LT predicate on the "cron_spec" field.
func CronSpecLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCronSpec, v))
}

// CronSpecLTE applies the LTE predicate on the "cron_spec" field.
func CronSpecLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCronSpec, v))
}

// CronSpecContains applies the Contains predicate on the "cron_spec" field.
func CronSpecContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldCronSpec, v))
}

// CronSpecHasPrefix applies the HasPrefix predicate on the "cron_spec" field.
func CronSpecHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldCronSpec, v))
}

// CronSpecHasSuffix applies the HasSuffix predicate on the "cron_spec" field.
func CronSpecHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldCronSpec, v))
}

// CronSpecIsNil applies the IsNil predicate on the "cron_spec" field.
func CronSpecIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldCronSpec))
}

// CronSpecNotNil applies the NotNil predicate on the "cron_spec" field.
func CronSpecNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldCronSpec))
}

// CronSpecEqualFold applies the EqualFold predicate on the "cron_spec" field.
func CronSpecEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldCronSpec, v))
}

// CronSpecContainsFold applies the ContainsFold predicate on the "cron_spec" field.
func CronSpecContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldCronSpec, v))
}

// TaskOptionsIsNil applies the IsNil predicate on the "task_options" field.
func TaskOptionsIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldTaskOptions))
}

// TaskOptionsNotNil applies the NotNil predicate on the "task_options" field.
func TaskOptionsNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldTaskOptions))
}

// EnableEQ applies the EQ predicate on the "enable" field.
func EnableEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldEnable, v))
}

// EnableNEQ applies the NEQ predicate on the "enable" field.
func EnableNEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldEnable, v))
}

// EnableIsNil applies the IsNil predicate on the "enable" field.
func EnableIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldEnable))
}

// EnableNotNil applies the NotNil predicate on the "enable" field.
func EnableNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldEnable))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(sql.NotPredicates(p))
}
