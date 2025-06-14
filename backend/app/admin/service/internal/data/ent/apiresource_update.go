// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/apiresource"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiResourceUpdate is the builder for updating ApiResource entities.
type ApiResourceUpdate struct {
	config
	hooks     []Hook
	mutation  *ApiResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ApiResourceUpdate builder.
func (aru *ApiResourceUpdate) Where(ps ...predicate.ApiResource) *ApiResourceUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetUpdateTime sets the "update_time" field.
func (aru *ApiResourceUpdate) SetUpdateTime(t time.Time) *ApiResourceUpdate {
	aru.mutation.SetUpdateTime(t)
	return aru
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableUpdateTime(t *time.Time) *ApiResourceUpdate {
	if t != nil {
		aru.SetUpdateTime(*t)
	}
	return aru
}

// ClearUpdateTime clears the value of the "update_time" field.
func (aru *ApiResourceUpdate) ClearUpdateTime() *ApiResourceUpdate {
	aru.mutation.ClearUpdateTime()
	return aru
}

// SetDeleteTime sets the "delete_time" field.
func (aru *ApiResourceUpdate) SetDeleteTime(t time.Time) *ApiResourceUpdate {
	aru.mutation.SetDeleteTime(t)
	return aru
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableDeleteTime(t *time.Time) *ApiResourceUpdate {
	if t != nil {
		aru.SetDeleteTime(*t)
	}
	return aru
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (aru *ApiResourceUpdate) ClearDeleteTime() *ApiResourceUpdate {
	aru.mutation.ClearDeleteTime()
	return aru
}

// SetCreateBy sets the "create_by" field.
func (aru *ApiResourceUpdate) SetCreateBy(u uint32) *ApiResourceUpdate {
	aru.mutation.ResetCreateBy()
	aru.mutation.SetCreateBy(u)
	return aru
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableCreateBy(u *uint32) *ApiResourceUpdate {
	if u != nil {
		aru.SetCreateBy(*u)
	}
	return aru
}

// AddCreateBy adds u to the "create_by" field.
func (aru *ApiResourceUpdate) AddCreateBy(u int32) *ApiResourceUpdate {
	aru.mutation.AddCreateBy(u)
	return aru
}

// ClearCreateBy clears the value of the "create_by" field.
func (aru *ApiResourceUpdate) ClearCreateBy() *ApiResourceUpdate {
	aru.mutation.ClearCreateBy()
	return aru
}

// SetUpdateBy sets the "update_by" field.
func (aru *ApiResourceUpdate) SetUpdateBy(u uint32) *ApiResourceUpdate {
	aru.mutation.ResetUpdateBy()
	aru.mutation.SetUpdateBy(u)
	return aru
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableUpdateBy(u *uint32) *ApiResourceUpdate {
	if u != nil {
		aru.SetUpdateBy(*u)
	}
	return aru
}

// AddUpdateBy adds u to the "update_by" field.
func (aru *ApiResourceUpdate) AddUpdateBy(u int32) *ApiResourceUpdate {
	aru.mutation.AddUpdateBy(u)
	return aru
}

// ClearUpdateBy clears the value of the "update_by" field.
func (aru *ApiResourceUpdate) ClearUpdateBy() *ApiResourceUpdate {
	aru.mutation.ClearUpdateBy()
	return aru
}

// SetDescription sets the "description" field.
func (aru *ApiResourceUpdate) SetDescription(s string) *ApiResourceUpdate {
	aru.mutation.SetDescription(s)
	return aru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableDescription(s *string) *ApiResourceUpdate {
	if s != nil {
		aru.SetDescription(*s)
	}
	return aru
}

// ClearDescription clears the value of the "description" field.
func (aru *ApiResourceUpdate) ClearDescription() *ApiResourceUpdate {
	aru.mutation.ClearDescription()
	return aru
}

// SetModule sets the "module" field.
func (aru *ApiResourceUpdate) SetModule(s string) *ApiResourceUpdate {
	aru.mutation.SetModule(s)
	return aru
}

// SetNillableModule sets the "module" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableModule(s *string) *ApiResourceUpdate {
	if s != nil {
		aru.SetModule(*s)
	}
	return aru
}

// ClearModule clears the value of the "module" field.
func (aru *ApiResourceUpdate) ClearModule() *ApiResourceUpdate {
	aru.mutation.ClearModule()
	return aru
}

// SetModuleDescription sets the "module_description" field.
func (aru *ApiResourceUpdate) SetModuleDescription(s string) *ApiResourceUpdate {
	aru.mutation.SetModuleDescription(s)
	return aru
}

// SetNillableModuleDescription sets the "module_description" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableModuleDescription(s *string) *ApiResourceUpdate {
	if s != nil {
		aru.SetModuleDescription(*s)
	}
	return aru
}

// ClearModuleDescription clears the value of the "module_description" field.
func (aru *ApiResourceUpdate) ClearModuleDescription() *ApiResourceUpdate {
	aru.mutation.ClearModuleDescription()
	return aru
}

// SetOperation sets the "operation" field.
func (aru *ApiResourceUpdate) SetOperation(s string) *ApiResourceUpdate {
	aru.mutation.SetOperation(s)
	return aru
}

// SetNillableOperation sets the "operation" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableOperation(s *string) *ApiResourceUpdate {
	if s != nil {
		aru.SetOperation(*s)
	}
	return aru
}

// ClearOperation clears the value of the "operation" field.
func (aru *ApiResourceUpdate) ClearOperation() *ApiResourceUpdate {
	aru.mutation.ClearOperation()
	return aru
}

// SetPath sets the "path" field.
func (aru *ApiResourceUpdate) SetPath(s string) *ApiResourceUpdate {
	aru.mutation.SetPath(s)
	return aru
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillablePath(s *string) *ApiResourceUpdate {
	if s != nil {
		aru.SetPath(*s)
	}
	return aru
}

// ClearPath clears the value of the "path" field.
func (aru *ApiResourceUpdate) ClearPath() *ApiResourceUpdate {
	aru.mutation.ClearPath()
	return aru
}

// SetMethod sets the "method" field.
func (aru *ApiResourceUpdate) SetMethod(s string) *ApiResourceUpdate {
	aru.mutation.SetMethod(s)
	return aru
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (aru *ApiResourceUpdate) SetNillableMethod(s *string) *ApiResourceUpdate {
	if s != nil {
		aru.SetMethod(*s)
	}
	return aru
}

// ClearMethod clears the value of the "method" field.
func (aru *ApiResourceUpdate) ClearMethod() *ApiResourceUpdate {
	aru.mutation.ClearMethod()
	return aru
}

// Mutation returns the ApiResourceMutation object of the builder.
func (aru *ApiResourceUpdate) Mutation() *ApiResourceMutation {
	return aru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *ApiResourceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *ApiResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *ApiResourceUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *ApiResourceUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aru *ApiResourceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApiResourceUpdate {
	aru.modifiers = append(aru.modifiers, modifiers...)
	return aru
}

func (aru *ApiResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apiresource.Table, apiresource.Columns, sqlgraph.NewFieldSpec(apiresource.FieldID, field.TypeUint32))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if aru.mutation.CreateTimeCleared() {
		_spec.ClearField(apiresource.FieldCreateTime, field.TypeTime)
	}
	if value, ok := aru.mutation.UpdateTime(); ok {
		_spec.SetField(apiresource.FieldUpdateTime, field.TypeTime, value)
	}
	if aru.mutation.UpdateTimeCleared() {
		_spec.ClearField(apiresource.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := aru.mutation.DeleteTime(); ok {
		_spec.SetField(apiresource.FieldDeleteTime, field.TypeTime, value)
	}
	if aru.mutation.DeleteTimeCleared() {
		_spec.ClearField(apiresource.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := aru.mutation.CreateBy(); ok {
		_spec.SetField(apiresource.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := aru.mutation.AddedCreateBy(); ok {
		_spec.AddField(apiresource.FieldCreateBy, field.TypeUint32, value)
	}
	if aru.mutation.CreateByCleared() {
		_spec.ClearField(apiresource.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := aru.mutation.UpdateBy(); ok {
		_spec.SetField(apiresource.FieldUpdateBy, field.TypeUint32, value)
	}
	if value, ok := aru.mutation.AddedUpdateBy(); ok {
		_spec.AddField(apiresource.FieldUpdateBy, field.TypeUint32, value)
	}
	if aru.mutation.UpdateByCleared() {
		_spec.ClearField(apiresource.FieldUpdateBy, field.TypeUint32)
	}
	if value, ok := aru.mutation.Description(); ok {
		_spec.SetField(apiresource.FieldDescription, field.TypeString, value)
	}
	if aru.mutation.DescriptionCleared() {
		_spec.ClearField(apiresource.FieldDescription, field.TypeString)
	}
	if value, ok := aru.mutation.Module(); ok {
		_spec.SetField(apiresource.FieldModule, field.TypeString, value)
	}
	if aru.mutation.ModuleCleared() {
		_spec.ClearField(apiresource.FieldModule, field.TypeString)
	}
	if value, ok := aru.mutation.ModuleDescription(); ok {
		_spec.SetField(apiresource.FieldModuleDescription, field.TypeString, value)
	}
	if aru.mutation.ModuleDescriptionCleared() {
		_spec.ClearField(apiresource.FieldModuleDescription, field.TypeString)
	}
	if value, ok := aru.mutation.Operation(); ok {
		_spec.SetField(apiresource.FieldOperation, field.TypeString, value)
	}
	if aru.mutation.OperationCleared() {
		_spec.ClearField(apiresource.FieldOperation, field.TypeString)
	}
	if value, ok := aru.mutation.Path(); ok {
		_spec.SetField(apiresource.FieldPath, field.TypeString, value)
	}
	if aru.mutation.PathCleared() {
		_spec.ClearField(apiresource.FieldPath, field.TypeString)
	}
	if value, ok := aru.mutation.Method(); ok {
		_spec.SetField(apiresource.FieldMethod, field.TypeString, value)
	}
	if aru.mutation.MethodCleared() {
		_spec.ClearField(apiresource.FieldMethod, field.TypeString)
	}
	_spec.AddModifiers(aru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apiresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// ApiResourceUpdateOne is the builder for updating a single ApiResource entity.
type ApiResourceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ApiResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (aruo *ApiResourceUpdateOne) SetUpdateTime(t time.Time) *ApiResourceUpdateOne {
	aruo.mutation.SetUpdateTime(t)
	return aruo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableUpdateTime(t *time.Time) *ApiResourceUpdateOne {
	if t != nil {
		aruo.SetUpdateTime(*t)
	}
	return aruo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (aruo *ApiResourceUpdateOne) ClearUpdateTime() *ApiResourceUpdateOne {
	aruo.mutation.ClearUpdateTime()
	return aruo
}

// SetDeleteTime sets the "delete_time" field.
func (aruo *ApiResourceUpdateOne) SetDeleteTime(t time.Time) *ApiResourceUpdateOne {
	aruo.mutation.SetDeleteTime(t)
	return aruo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableDeleteTime(t *time.Time) *ApiResourceUpdateOne {
	if t != nil {
		aruo.SetDeleteTime(*t)
	}
	return aruo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (aruo *ApiResourceUpdateOne) ClearDeleteTime() *ApiResourceUpdateOne {
	aruo.mutation.ClearDeleteTime()
	return aruo
}

// SetCreateBy sets the "create_by" field.
func (aruo *ApiResourceUpdateOne) SetCreateBy(u uint32) *ApiResourceUpdateOne {
	aruo.mutation.ResetCreateBy()
	aruo.mutation.SetCreateBy(u)
	return aruo
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableCreateBy(u *uint32) *ApiResourceUpdateOne {
	if u != nil {
		aruo.SetCreateBy(*u)
	}
	return aruo
}

// AddCreateBy adds u to the "create_by" field.
func (aruo *ApiResourceUpdateOne) AddCreateBy(u int32) *ApiResourceUpdateOne {
	aruo.mutation.AddCreateBy(u)
	return aruo
}

// ClearCreateBy clears the value of the "create_by" field.
func (aruo *ApiResourceUpdateOne) ClearCreateBy() *ApiResourceUpdateOne {
	aruo.mutation.ClearCreateBy()
	return aruo
}

// SetUpdateBy sets the "update_by" field.
func (aruo *ApiResourceUpdateOne) SetUpdateBy(u uint32) *ApiResourceUpdateOne {
	aruo.mutation.ResetUpdateBy()
	aruo.mutation.SetUpdateBy(u)
	return aruo
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableUpdateBy(u *uint32) *ApiResourceUpdateOne {
	if u != nil {
		aruo.SetUpdateBy(*u)
	}
	return aruo
}

// AddUpdateBy adds u to the "update_by" field.
func (aruo *ApiResourceUpdateOne) AddUpdateBy(u int32) *ApiResourceUpdateOne {
	aruo.mutation.AddUpdateBy(u)
	return aruo
}

// ClearUpdateBy clears the value of the "update_by" field.
func (aruo *ApiResourceUpdateOne) ClearUpdateBy() *ApiResourceUpdateOne {
	aruo.mutation.ClearUpdateBy()
	return aruo
}

// SetDescription sets the "description" field.
func (aruo *ApiResourceUpdateOne) SetDescription(s string) *ApiResourceUpdateOne {
	aruo.mutation.SetDescription(s)
	return aruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableDescription(s *string) *ApiResourceUpdateOne {
	if s != nil {
		aruo.SetDescription(*s)
	}
	return aruo
}

// ClearDescription clears the value of the "description" field.
func (aruo *ApiResourceUpdateOne) ClearDescription() *ApiResourceUpdateOne {
	aruo.mutation.ClearDescription()
	return aruo
}

// SetModule sets the "module" field.
func (aruo *ApiResourceUpdateOne) SetModule(s string) *ApiResourceUpdateOne {
	aruo.mutation.SetModule(s)
	return aruo
}

// SetNillableModule sets the "module" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableModule(s *string) *ApiResourceUpdateOne {
	if s != nil {
		aruo.SetModule(*s)
	}
	return aruo
}

// ClearModule clears the value of the "module" field.
func (aruo *ApiResourceUpdateOne) ClearModule() *ApiResourceUpdateOne {
	aruo.mutation.ClearModule()
	return aruo
}

// SetModuleDescription sets the "module_description" field.
func (aruo *ApiResourceUpdateOne) SetModuleDescription(s string) *ApiResourceUpdateOne {
	aruo.mutation.SetModuleDescription(s)
	return aruo
}

// SetNillableModuleDescription sets the "module_description" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableModuleDescription(s *string) *ApiResourceUpdateOne {
	if s != nil {
		aruo.SetModuleDescription(*s)
	}
	return aruo
}

// ClearModuleDescription clears the value of the "module_description" field.
func (aruo *ApiResourceUpdateOne) ClearModuleDescription() *ApiResourceUpdateOne {
	aruo.mutation.ClearModuleDescription()
	return aruo
}

// SetOperation sets the "operation" field.
func (aruo *ApiResourceUpdateOne) SetOperation(s string) *ApiResourceUpdateOne {
	aruo.mutation.SetOperation(s)
	return aruo
}

// SetNillableOperation sets the "operation" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableOperation(s *string) *ApiResourceUpdateOne {
	if s != nil {
		aruo.SetOperation(*s)
	}
	return aruo
}

// ClearOperation clears the value of the "operation" field.
func (aruo *ApiResourceUpdateOne) ClearOperation() *ApiResourceUpdateOne {
	aruo.mutation.ClearOperation()
	return aruo
}

// SetPath sets the "path" field.
func (aruo *ApiResourceUpdateOne) SetPath(s string) *ApiResourceUpdateOne {
	aruo.mutation.SetPath(s)
	return aruo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillablePath(s *string) *ApiResourceUpdateOne {
	if s != nil {
		aruo.SetPath(*s)
	}
	return aruo
}

// ClearPath clears the value of the "path" field.
func (aruo *ApiResourceUpdateOne) ClearPath() *ApiResourceUpdateOne {
	aruo.mutation.ClearPath()
	return aruo
}

// SetMethod sets the "method" field.
func (aruo *ApiResourceUpdateOne) SetMethod(s string) *ApiResourceUpdateOne {
	aruo.mutation.SetMethod(s)
	return aruo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (aruo *ApiResourceUpdateOne) SetNillableMethod(s *string) *ApiResourceUpdateOne {
	if s != nil {
		aruo.SetMethod(*s)
	}
	return aruo
}

// ClearMethod clears the value of the "method" field.
func (aruo *ApiResourceUpdateOne) ClearMethod() *ApiResourceUpdateOne {
	aruo.mutation.ClearMethod()
	return aruo
}

// Mutation returns the ApiResourceMutation object of the builder.
func (aruo *ApiResourceUpdateOne) Mutation() *ApiResourceMutation {
	return aruo.mutation
}

// Where appends a list predicates to the ApiResourceUpdate builder.
func (aruo *ApiResourceUpdateOne) Where(ps ...predicate.ApiResource) *ApiResourceUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *ApiResourceUpdateOne) Select(field string, fields ...string) *ApiResourceUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated ApiResource entity.
func (aruo *ApiResourceUpdateOne) Save(ctx context.Context) (*ApiResource, error) {
	return withHooks(ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *ApiResourceUpdateOne) SaveX(ctx context.Context) *ApiResource {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *ApiResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *ApiResourceUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aruo *ApiResourceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApiResourceUpdateOne {
	aruo.modifiers = append(aruo.modifiers, modifiers...)
	return aruo
}

func (aruo *ApiResourceUpdateOne) sqlSave(ctx context.Context) (_node *ApiResource, err error) {
	_spec := sqlgraph.NewUpdateSpec(apiresource.Table, apiresource.Columns, sqlgraph.NewFieldSpec(apiresource.FieldID, field.TypeUint32))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApiResource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apiresource.FieldID)
		for _, f := range fields {
			if !apiresource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apiresource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if aruo.mutation.CreateTimeCleared() {
		_spec.ClearField(apiresource.FieldCreateTime, field.TypeTime)
	}
	if value, ok := aruo.mutation.UpdateTime(); ok {
		_spec.SetField(apiresource.FieldUpdateTime, field.TypeTime, value)
	}
	if aruo.mutation.UpdateTimeCleared() {
		_spec.ClearField(apiresource.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := aruo.mutation.DeleteTime(); ok {
		_spec.SetField(apiresource.FieldDeleteTime, field.TypeTime, value)
	}
	if aruo.mutation.DeleteTimeCleared() {
		_spec.ClearField(apiresource.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := aruo.mutation.CreateBy(); ok {
		_spec.SetField(apiresource.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := aruo.mutation.AddedCreateBy(); ok {
		_spec.AddField(apiresource.FieldCreateBy, field.TypeUint32, value)
	}
	if aruo.mutation.CreateByCleared() {
		_spec.ClearField(apiresource.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := aruo.mutation.UpdateBy(); ok {
		_spec.SetField(apiresource.FieldUpdateBy, field.TypeUint32, value)
	}
	if value, ok := aruo.mutation.AddedUpdateBy(); ok {
		_spec.AddField(apiresource.FieldUpdateBy, field.TypeUint32, value)
	}
	if aruo.mutation.UpdateByCleared() {
		_spec.ClearField(apiresource.FieldUpdateBy, field.TypeUint32)
	}
	if value, ok := aruo.mutation.Description(); ok {
		_spec.SetField(apiresource.FieldDescription, field.TypeString, value)
	}
	if aruo.mutation.DescriptionCleared() {
		_spec.ClearField(apiresource.FieldDescription, field.TypeString)
	}
	if value, ok := aruo.mutation.Module(); ok {
		_spec.SetField(apiresource.FieldModule, field.TypeString, value)
	}
	if aruo.mutation.ModuleCleared() {
		_spec.ClearField(apiresource.FieldModule, field.TypeString)
	}
	if value, ok := aruo.mutation.ModuleDescription(); ok {
		_spec.SetField(apiresource.FieldModuleDescription, field.TypeString, value)
	}
	if aruo.mutation.ModuleDescriptionCleared() {
		_spec.ClearField(apiresource.FieldModuleDescription, field.TypeString)
	}
	if value, ok := aruo.mutation.Operation(); ok {
		_spec.SetField(apiresource.FieldOperation, field.TypeString, value)
	}
	if aruo.mutation.OperationCleared() {
		_spec.ClearField(apiresource.FieldOperation, field.TypeString)
	}
	if value, ok := aruo.mutation.Path(); ok {
		_spec.SetField(apiresource.FieldPath, field.TypeString, value)
	}
	if aruo.mutation.PathCleared() {
		_spec.ClearField(apiresource.FieldPath, field.TypeString)
	}
	if value, ok := aruo.mutation.Method(); ok {
		_spec.SetField(apiresource.FieldMethod, field.TypeString, value)
	}
	if aruo.mutation.MethodCleared() {
		_spec.ClearField(apiresource.FieldMethod, field.TypeString)
	}
	_spec.AddModifiers(aruo.modifiers...)
	_node = &ApiResource{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apiresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
