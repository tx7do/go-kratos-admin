// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/position"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionUpdate is the builder for updating Position entities.
type PositionUpdate struct {
	config
	hooks     []Hook
	mutation  *PositionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PositionUpdate builder.
func (pu *PositionUpdate) Where(ps ...predicate.Position) *PositionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PositionUpdate) SetUpdateTime(t time.Time) *PositionUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableUpdateTime(t *time.Time) *PositionUpdate {
	if t != nil {
		pu.SetUpdateTime(*t)
	}
	return pu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (pu *PositionUpdate) ClearUpdateTime() *PositionUpdate {
	pu.mutation.ClearUpdateTime()
	return pu
}

// SetDeleteTime sets the "delete_time" field.
func (pu *PositionUpdate) SetDeleteTime(t time.Time) *PositionUpdate {
	pu.mutation.SetDeleteTime(t)
	return pu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableDeleteTime(t *time.Time) *PositionUpdate {
	if t != nil {
		pu.SetDeleteTime(*t)
	}
	return pu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (pu *PositionUpdate) ClearDeleteTime() *PositionUpdate {
	pu.mutation.ClearDeleteTime()
	return pu
}

// SetStatus sets the "status" field.
func (pu *PositionUpdate) SetStatus(po position.Status) *PositionUpdate {
	pu.mutation.SetStatus(po)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableStatus(po *position.Status) *PositionUpdate {
	if po != nil {
		pu.SetStatus(*po)
	}
	return pu
}

// ClearStatus clears the value of the "status" field.
func (pu *PositionUpdate) ClearStatus() *PositionUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// SetCreateBy sets the "create_by" field.
func (pu *PositionUpdate) SetCreateBy(u uint32) *PositionUpdate {
	pu.mutation.ResetCreateBy()
	pu.mutation.SetCreateBy(u)
	return pu
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableCreateBy(u *uint32) *PositionUpdate {
	if u != nil {
		pu.SetCreateBy(*u)
	}
	return pu
}

// AddCreateBy adds u to the "create_by" field.
func (pu *PositionUpdate) AddCreateBy(u int32) *PositionUpdate {
	pu.mutation.AddCreateBy(u)
	return pu
}

// ClearCreateBy clears the value of the "create_by" field.
func (pu *PositionUpdate) ClearCreateBy() *PositionUpdate {
	pu.mutation.ClearCreateBy()
	return pu
}

// SetUpdateBy sets the "update_by" field.
func (pu *PositionUpdate) SetUpdateBy(u uint32) *PositionUpdate {
	pu.mutation.ResetUpdateBy()
	pu.mutation.SetUpdateBy(u)
	return pu
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableUpdateBy(u *uint32) *PositionUpdate {
	if u != nil {
		pu.SetUpdateBy(*u)
	}
	return pu
}

// AddUpdateBy adds u to the "update_by" field.
func (pu *PositionUpdate) AddUpdateBy(u int32) *PositionUpdate {
	pu.mutation.AddUpdateBy(u)
	return pu
}

// ClearUpdateBy clears the value of the "update_by" field.
func (pu *PositionUpdate) ClearUpdateBy() *PositionUpdate {
	pu.mutation.ClearUpdateBy()
	return pu
}

// SetRemark sets the "remark" field.
func (pu *PositionUpdate) SetRemark(s string) *PositionUpdate {
	pu.mutation.SetRemark(s)
	return pu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableRemark(s *string) *PositionUpdate {
	if s != nil {
		pu.SetRemark(*s)
	}
	return pu
}

// ClearRemark clears the value of the "remark" field.
func (pu *PositionUpdate) ClearRemark() *PositionUpdate {
	pu.mutation.ClearRemark()
	return pu
}

// SetName sets the "name" field.
func (pu *PositionUpdate) SetName(s string) *PositionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableName(s *string) *PositionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetCode sets the "code" field.
func (pu *PositionUpdate) SetCode(s string) *PositionUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableCode(s *string) *PositionUpdate {
	if s != nil {
		pu.SetCode(*s)
	}
	return pu
}

// SetParentID sets the "parent_id" field.
func (pu *PositionUpdate) SetParentID(u uint32) *PositionUpdate {
	pu.mutation.SetParentID(u)
	return pu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableParentID(u *uint32) *PositionUpdate {
	if u != nil {
		pu.SetParentID(*u)
	}
	return pu
}

// ClearParentID clears the value of the "parent_id" field.
func (pu *PositionUpdate) ClearParentID() *PositionUpdate {
	pu.mutation.ClearParentID()
	return pu
}

// SetSortID sets the "sort_id" field.
func (pu *PositionUpdate) SetSortID(i int32) *PositionUpdate {
	pu.mutation.ResetSortID()
	pu.mutation.SetSortID(i)
	return pu
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableSortID(i *int32) *PositionUpdate {
	if i != nil {
		pu.SetSortID(*i)
	}
	return pu
}

// AddSortID adds i to the "sort_id" field.
func (pu *PositionUpdate) AddSortID(i int32) *PositionUpdate {
	pu.mutation.AddSortID(i)
	return pu
}

// SetParent sets the "parent" edge to the Position entity.
func (pu *PositionUpdate) SetParent(p *Position) *PositionUpdate {
	return pu.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the Position entity by IDs.
func (pu *PositionUpdate) AddChildIDs(ids ...uint32) *PositionUpdate {
	pu.mutation.AddChildIDs(ids...)
	return pu
}

// AddChildren adds the "children" edges to the Position entity.
func (pu *PositionUpdate) AddChildren(p ...*Position) *PositionUpdate {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddChildIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pu *PositionUpdate) Mutation() *PositionMutation {
	return pu.mutation
}

// ClearParent clears the "parent" edge to the Position entity.
func (pu *PositionUpdate) ClearParent() *PositionUpdate {
	pu.mutation.ClearParent()
	return pu
}

// ClearChildren clears all "children" edges to the Position entity.
func (pu *PositionUpdate) ClearChildren() *PositionUpdate {
	pu.mutation.ClearChildren()
	return pu
}

// RemoveChildIDs removes the "children" edge to Position entities by IDs.
func (pu *PositionUpdate) RemoveChildIDs(ids ...uint32) *PositionUpdate {
	pu.mutation.RemoveChildIDs(ids...)
	return pu
}

// RemoveChildren removes "children" edges to Position entities.
func (pu *PositionUpdate) RemoveChildren(p ...*Position) *PositionUpdate {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PositionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PositionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PositionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PositionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PositionUpdate) check() error {
	if v, ok := pu.mutation.Status(); ok {
		if err := position.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Position.status": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Code(); ok {
		if err := position.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Position.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PositionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PositionUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.CreateTimeCleared() {
		_spec.ClearField(position.FieldCreateTime, field.TypeTime)
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.SetField(position.FieldUpdateTime, field.TypeTime, value)
	}
	if pu.mutation.UpdateTimeCleared() {
		_spec.ClearField(position.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := pu.mutation.DeleteTime(); ok {
		_spec.SetField(position.FieldDeleteTime, field.TypeTime, value)
	}
	if pu.mutation.DeleteTimeCleared() {
		_spec.ClearField(position.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.StatusCleared() {
		_spec.ClearField(position.FieldStatus, field.TypeEnum)
	}
	if value, ok := pu.mutation.CreateBy(); ok {
		_spec.SetField(position.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := pu.mutation.AddedCreateBy(); ok {
		_spec.AddField(position.FieldCreateBy, field.TypeUint32, value)
	}
	if pu.mutation.CreateByCleared() {
		_spec.ClearField(position.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := pu.mutation.UpdateBy(); ok {
		_spec.SetField(position.FieldUpdateBy, field.TypeUint32, value)
	}
	if value, ok := pu.mutation.AddedUpdateBy(); ok {
		_spec.AddField(position.FieldUpdateBy, field.TypeUint32, value)
	}
	if pu.mutation.UpdateByCleared() {
		_spec.ClearField(position.FieldUpdateBy, field.TypeUint32)
	}
	if value, ok := pu.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
	}
	if pu.mutation.RemarkCleared() {
		_spec.ClearField(position.FieldRemark, field.TypeString)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.SortID(); ok {
		_spec.SetField(position.FieldSortID, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedSortID(); ok {
		_spec.AddField(position.FieldSortID, field.TypeInt32, value)
	}
	if pu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.ParentTable,
			Columns: []string{position.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.ParentTable,
			Columns: []string{position.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !pu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PositionUpdateOne is the builder for updating a single Position entity.
type PositionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PositionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (puo *PositionUpdateOne) SetUpdateTime(t time.Time) *PositionUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableUpdateTime(t *time.Time) *PositionUpdateOne {
	if t != nil {
		puo.SetUpdateTime(*t)
	}
	return puo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (puo *PositionUpdateOne) ClearUpdateTime() *PositionUpdateOne {
	puo.mutation.ClearUpdateTime()
	return puo
}

// SetDeleteTime sets the "delete_time" field.
func (puo *PositionUpdateOne) SetDeleteTime(t time.Time) *PositionUpdateOne {
	puo.mutation.SetDeleteTime(t)
	return puo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableDeleteTime(t *time.Time) *PositionUpdateOne {
	if t != nil {
		puo.SetDeleteTime(*t)
	}
	return puo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (puo *PositionUpdateOne) ClearDeleteTime() *PositionUpdateOne {
	puo.mutation.ClearDeleteTime()
	return puo
}

// SetStatus sets the "status" field.
func (puo *PositionUpdateOne) SetStatus(po position.Status) *PositionUpdateOne {
	puo.mutation.SetStatus(po)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableStatus(po *position.Status) *PositionUpdateOne {
	if po != nil {
		puo.SetStatus(*po)
	}
	return puo
}

// ClearStatus clears the value of the "status" field.
func (puo *PositionUpdateOne) ClearStatus() *PositionUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// SetCreateBy sets the "create_by" field.
func (puo *PositionUpdateOne) SetCreateBy(u uint32) *PositionUpdateOne {
	puo.mutation.ResetCreateBy()
	puo.mutation.SetCreateBy(u)
	return puo
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableCreateBy(u *uint32) *PositionUpdateOne {
	if u != nil {
		puo.SetCreateBy(*u)
	}
	return puo
}

// AddCreateBy adds u to the "create_by" field.
func (puo *PositionUpdateOne) AddCreateBy(u int32) *PositionUpdateOne {
	puo.mutation.AddCreateBy(u)
	return puo
}

// ClearCreateBy clears the value of the "create_by" field.
func (puo *PositionUpdateOne) ClearCreateBy() *PositionUpdateOne {
	puo.mutation.ClearCreateBy()
	return puo
}

// SetUpdateBy sets the "update_by" field.
func (puo *PositionUpdateOne) SetUpdateBy(u uint32) *PositionUpdateOne {
	puo.mutation.ResetUpdateBy()
	puo.mutation.SetUpdateBy(u)
	return puo
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableUpdateBy(u *uint32) *PositionUpdateOne {
	if u != nil {
		puo.SetUpdateBy(*u)
	}
	return puo
}

// AddUpdateBy adds u to the "update_by" field.
func (puo *PositionUpdateOne) AddUpdateBy(u int32) *PositionUpdateOne {
	puo.mutation.AddUpdateBy(u)
	return puo
}

// ClearUpdateBy clears the value of the "update_by" field.
func (puo *PositionUpdateOne) ClearUpdateBy() *PositionUpdateOne {
	puo.mutation.ClearUpdateBy()
	return puo
}

// SetRemark sets the "remark" field.
func (puo *PositionUpdateOne) SetRemark(s string) *PositionUpdateOne {
	puo.mutation.SetRemark(s)
	return puo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableRemark(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetRemark(*s)
	}
	return puo
}

// ClearRemark clears the value of the "remark" field.
func (puo *PositionUpdateOne) ClearRemark() *PositionUpdateOne {
	puo.mutation.ClearRemark()
	return puo
}

// SetName sets the "name" field.
func (puo *PositionUpdateOne) SetName(s string) *PositionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableName(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetCode sets the "code" field.
func (puo *PositionUpdateOne) SetCode(s string) *PositionUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableCode(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetCode(*s)
	}
	return puo
}

// SetParentID sets the "parent_id" field.
func (puo *PositionUpdateOne) SetParentID(u uint32) *PositionUpdateOne {
	puo.mutation.SetParentID(u)
	return puo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableParentID(u *uint32) *PositionUpdateOne {
	if u != nil {
		puo.SetParentID(*u)
	}
	return puo
}

// ClearParentID clears the value of the "parent_id" field.
func (puo *PositionUpdateOne) ClearParentID() *PositionUpdateOne {
	puo.mutation.ClearParentID()
	return puo
}

// SetSortID sets the "sort_id" field.
func (puo *PositionUpdateOne) SetSortID(i int32) *PositionUpdateOne {
	puo.mutation.ResetSortID()
	puo.mutation.SetSortID(i)
	return puo
}

// SetNillableSortID sets the "sort_id" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableSortID(i *int32) *PositionUpdateOne {
	if i != nil {
		puo.SetSortID(*i)
	}
	return puo
}

// AddSortID adds i to the "sort_id" field.
func (puo *PositionUpdateOne) AddSortID(i int32) *PositionUpdateOne {
	puo.mutation.AddSortID(i)
	return puo
}

// SetParent sets the "parent" edge to the Position entity.
func (puo *PositionUpdateOne) SetParent(p *Position) *PositionUpdateOne {
	return puo.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the Position entity by IDs.
func (puo *PositionUpdateOne) AddChildIDs(ids ...uint32) *PositionUpdateOne {
	puo.mutation.AddChildIDs(ids...)
	return puo
}

// AddChildren adds the "children" edges to the Position entity.
func (puo *PositionUpdateOne) AddChildren(p ...*Position) *PositionUpdateOne {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddChildIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (puo *PositionUpdateOne) Mutation() *PositionMutation {
	return puo.mutation
}

// ClearParent clears the "parent" edge to the Position entity.
func (puo *PositionUpdateOne) ClearParent() *PositionUpdateOne {
	puo.mutation.ClearParent()
	return puo
}

// ClearChildren clears all "children" edges to the Position entity.
func (puo *PositionUpdateOne) ClearChildren() *PositionUpdateOne {
	puo.mutation.ClearChildren()
	return puo
}

// RemoveChildIDs removes the "children" edge to Position entities by IDs.
func (puo *PositionUpdateOne) RemoveChildIDs(ids ...uint32) *PositionUpdateOne {
	puo.mutation.RemoveChildIDs(ids...)
	return puo
}

// RemoveChildren removes "children" edges to Position entities.
func (puo *PositionUpdateOne) RemoveChildren(p ...*Position) *PositionUpdateOne {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the PositionUpdate builder.
func (puo *PositionUpdateOne) Where(ps ...predicate.Position) *PositionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PositionUpdateOne) Select(field string, fields ...string) *PositionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Position entity.
func (puo *PositionUpdateOne) Save(ctx context.Context) (*Position, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PositionUpdateOne) SaveX(ctx context.Context) *Position {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PositionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PositionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PositionUpdateOne) check() error {
	if v, ok := puo.mutation.Status(); ok {
		if err := position.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Position.status": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Code(); ok {
		if err := position.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Position.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PositionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PositionUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PositionUpdateOne) sqlSave(ctx context.Context) (_node *Position, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Position.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, position.FieldID)
		for _, f := range fields {
			if !position.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.CreateTimeCleared() {
		_spec.ClearField(position.FieldCreateTime, field.TypeTime)
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.SetField(position.FieldUpdateTime, field.TypeTime, value)
	}
	if puo.mutation.UpdateTimeCleared() {
		_spec.ClearField(position.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := puo.mutation.DeleteTime(); ok {
		_spec.SetField(position.FieldDeleteTime, field.TypeTime, value)
	}
	if puo.mutation.DeleteTimeCleared() {
		_spec.ClearField(position.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.StatusCleared() {
		_spec.ClearField(position.FieldStatus, field.TypeEnum)
	}
	if value, ok := puo.mutation.CreateBy(); ok {
		_spec.SetField(position.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := puo.mutation.AddedCreateBy(); ok {
		_spec.AddField(position.FieldCreateBy, field.TypeUint32, value)
	}
	if puo.mutation.CreateByCleared() {
		_spec.ClearField(position.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := puo.mutation.UpdateBy(); ok {
		_spec.SetField(position.FieldUpdateBy, field.TypeUint32, value)
	}
	if value, ok := puo.mutation.AddedUpdateBy(); ok {
		_spec.AddField(position.FieldUpdateBy, field.TypeUint32, value)
	}
	if puo.mutation.UpdateByCleared() {
		_spec.ClearField(position.FieldUpdateBy, field.TypeUint32)
	}
	if value, ok := puo.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
	}
	if puo.mutation.RemarkCleared() {
		_spec.ClearField(position.FieldRemark, field.TypeString)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.SortID(); ok {
		_spec.SetField(position.FieldSortID, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedSortID(); ok {
		_spec.AddField(position.FieldSortID, field.TypeInt32, value)
	}
	if puo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.ParentTable,
			Columns: []string{position.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.ParentTable,
			Columns: []string{position.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !puo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Position{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
