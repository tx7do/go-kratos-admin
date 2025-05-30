// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"kratos-admin/app/admin/service/internal/data/ent/usercredential"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCredentialUpdate is the builder for updating UserCredential entities.
type UserCredentialUpdate struct {
	config
	hooks     []Hook
	mutation  *UserCredentialMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserCredentialUpdate builder.
func (ucu *UserCredentialUpdate) Where(ps ...predicate.UserCredential) *UserCredentialUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// SetUpdateTime sets the "update_time" field.
func (ucu *UserCredentialUpdate) SetUpdateTime(t time.Time) *UserCredentialUpdate {
	ucu.mutation.SetUpdateTime(t)
	return ucu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableUpdateTime(t *time.Time) *UserCredentialUpdate {
	if t != nil {
		ucu.SetUpdateTime(*t)
	}
	return ucu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ucu *UserCredentialUpdate) ClearUpdateTime() *UserCredentialUpdate {
	ucu.mutation.ClearUpdateTime()
	return ucu
}

// SetDeleteTime sets the "delete_time" field.
func (ucu *UserCredentialUpdate) SetDeleteTime(t time.Time) *UserCredentialUpdate {
	ucu.mutation.SetDeleteTime(t)
	return ucu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableDeleteTime(t *time.Time) *UserCredentialUpdate {
	if t != nil {
		ucu.SetDeleteTime(*t)
	}
	return ucu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ucu *UserCredentialUpdate) ClearDeleteTime() *UserCredentialUpdate {
	ucu.mutation.ClearDeleteTime()
	return ucu
}

// SetUserID sets the "user_id" field.
func (ucu *UserCredentialUpdate) SetUserID(u uint32) *UserCredentialUpdate {
	ucu.mutation.ResetUserID()
	ucu.mutation.SetUserID(u)
	return ucu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableUserID(u *uint32) *UserCredentialUpdate {
	if u != nil {
		ucu.SetUserID(*u)
	}
	return ucu
}

// AddUserID adds u to the "user_id" field.
func (ucu *UserCredentialUpdate) AddUserID(u int32) *UserCredentialUpdate {
	ucu.mutation.AddUserID(u)
	return ucu
}

// ClearUserID clears the value of the "user_id" field.
func (ucu *UserCredentialUpdate) ClearUserID() *UserCredentialUpdate {
	ucu.mutation.ClearUserID()
	return ucu
}

// SetIdentityType sets the "identity_type" field.
func (ucu *UserCredentialUpdate) SetIdentityType(ut usercredential.IdentityType) *UserCredentialUpdate {
	ucu.mutation.SetIdentityType(ut)
	return ucu
}

// SetNillableIdentityType sets the "identity_type" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableIdentityType(ut *usercredential.IdentityType) *UserCredentialUpdate {
	if ut != nil {
		ucu.SetIdentityType(*ut)
	}
	return ucu
}

// ClearIdentityType clears the value of the "identity_type" field.
func (ucu *UserCredentialUpdate) ClearIdentityType() *UserCredentialUpdate {
	ucu.mutation.ClearIdentityType()
	return ucu
}

// SetIdentifier sets the "identifier" field.
func (ucu *UserCredentialUpdate) SetIdentifier(s string) *UserCredentialUpdate {
	ucu.mutation.SetIdentifier(s)
	return ucu
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableIdentifier(s *string) *UserCredentialUpdate {
	if s != nil {
		ucu.SetIdentifier(*s)
	}
	return ucu
}

// ClearIdentifier clears the value of the "identifier" field.
func (ucu *UserCredentialUpdate) ClearIdentifier() *UserCredentialUpdate {
	ucu.mutation.ClearIdentifier()
	return ucu
}

// SetCredentialType sets the "credential_type" field.
func (ucu *UserCredentialUpdate) SetCredentialType(ut usercredential.CredentialType) *UserCredentialUpdate {
	ucu.mutation.SetCredentialType(ut)
	return ucu
}

// SetNillableCredentialType sets the "credential_type" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableCredentialType(ut *usercredential.CredentialType) *UserCredentialUpdate {
	if ut != nil {
		ucu.SetCredentialType(*ut)
	}
	return ucu
}

// ClearCredentialType clears the value of the "credential_type" field.
func (ucu *UserCredentialUpdate) ClearCredentialType() *UserCredentialUpdate {
	ucu.mutation.ClearCredentialType()
	return ucu
}

// SetCredential sets the "credential" field.
func (ucu *UserCredentialUpdate) SetCredential(s string) *UserCredentialUpdate {
	ucu.mutation.SetCredential(s)
	return ucu
}

// SetNillableCredential sets the "credential" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableCredential(s *string) *UserCredentialUpdate {
	if s != nil {
		ucu.SetCredential(*s)
	}
	return ucu
}

// ClearCredential clears the value of the "credential" field.
func (ucu *UserCredentialUpdate) ClearCredential() *UserCredentialUpdate {
	ucu.mutation.ClearCredential()
	return ucu
}

// SetIsPrimary sets the "is_primary" field.
func (ucu *UserCredentialUpdate) SetIsPrimary(b bool) *UserCredentialUpdate {
	ucu.mutation.SetIsPrimary(b)
	return ucu
}

// SetNillableIsPrimary sets the "is_primary" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableIsPrimary(b *bool) *UserCredentialUpdate {
	if b != nil {
		ucu.SetIsPrimary(*b)
	}
	return ucu
}

// ClearIsPrimary clears the value of the "is_primary" field.
func (ucu *UserCredentialUpdate) ClearIsPrimary() *UserCredentialUpdate {
	ucu.mutation.ClearIsPrimary()
	return ucu
}

// SetStatus sets the "status" field.
func (ucu *UserCredentialUpdate) SetStatus(u usercredential.Status) *UserCredentialUpdate {
	ucu.mutation.SetStatus(u)
	return ucu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableStatus(u *usercredential.Status) *UserCredentialUpdate {
	if u != nil {
		ucu.SetStatus(*u)
	}
	return ucu
}

// ClearStatus clears the value of the "status" field.
func (ucu *UserCredentialUpdate) ClearStatus() *UserCredentialUpdate {
	ucu.mutation.ClearStatus()
	return ucu
}

// SetExtraInfo sets the "extra_info" field.
func (ucu *UserCredentialUpdate) SetExtraInfo(s string) *UserCredentialUpdate {
	ucu.mutation.SetExtraInfo(s)
	return ucu
}

// SetNillableExtraInfo sets the "extra_info" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableExtraInfo(s *string) *UserCredentialUpdate {
	if s != nil {
		ucu.SetExtraInfo(*s)
	}
	return ucu
}

// ClearExtraInfo clears the value of the "extra_info" field.
func (ucu *UserCredentialUpdate) ClearExtraInfo() *UserCredentialUpdate {
	ucu.mutation.ClearExtraInfo()
	return ucu
}

// SetActivateToken sets the "activate_token" field.
func (ucu *UserCredentialUpdate) SetActivateToken(s string) *UserCredentialUpdate {
	ucu.mutation.SetActivateToken(s)
	return ucu
}

// SetNillableActivateToken sets the "activate_token" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableActivateToken(s *string) *UserCredentialUpdate {
	if s != nil {
		ucu.SetActivateToken(*s)
	}
	return ucu
}

// ClearActivateToken clears the value of the "activate_token" field.
func (ucu *UserCredentialUpdate) ClearActivateToken() *UserCredentialUpdate {
	ucu.mutation.ClearActivateToken()
	return ucu
}

// SetResetToken sets the "reset_token" field.
func (ucu *UserCredentialUpdate) SetResetToken(s string) *UserCredentialUpdate {
	ucu.mutation.SetResetToken(s)
	return ucu
}

// SetNillableResetToken sets the "reset_token" field if the given value is not nil.
func (ucu *UserCredentialUpdate) SetNillableResetToken(s *string) *UserCredentialUpdate {
	if s != nil {
		ucu.SetResetToken(*s)
	}
	return ucu
}

// ClearResetToken clears the value of the "reset_token" field.
func (ucu *UserCredentialUpdate) ClearResetToken() *UserCredentialUpdate {
	ucu.mutation.ClearResetToken()
	return ucu
}

// Mutation returns the UserCredentialMutation object of the builder.
func (ucu *UserCredentialUpdate) Mutation() *UserCredentialMutation {
	return ucu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UserCredentialUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ucu.sqlSave, ucu.mutation, ucu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UserCredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UserCredentialUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UserCredentialUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucu *UserCredentialUpdate) check() error {
	if v, ok := ucu.mutation.IdentityType(); ok {
		if err := usercredential.IdentityTypeValidator(v); err != nil {
			return &ValidationError{Name: "identity_type", err: fmt.Errorf(`ent: validator failed for field "UserCredential.identity_type": %w`, err)}
		}
	}
	if v, ok := ucu.mutation.Identifier(); ok {
		if err := usercredential.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "UserCredential.identifier": %w`, err)}
		}
	}
	if v, ok := ucu.mutation.CredentialType(); ok {
		if err := usercredential.CredentialTypeValidator(v); err != nil {
			return &ValidationError{Name: "credential_type", err: fmt.Errorf(`ent: validator failed for field "UserCredential.credential_type": %w`, err)}
		}
	}
	if v, ok := ucu.mutation.Credential(); ok {
		if err := usercredential.CredentialValidator(v); err != nil {
			return &ValidationError{Name: "credential", err: fmt.Errorf(`ent: validator failed for field "UserCredential.credential": %w`, err)}
		}
	}
	if v, ok := ucu.mutation.Status(); ok {
		if err := usercredential.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserCredential.status": %w`, err)}
		}
	}
	if v, ok := ucu.mutation.ActivateToken(); ok {
		if err := usercredential.ActivateTokenValidator(v); err != nil {
			return &ValidationError{Name: "activate_token", err: fmt.Errorf(`ent: validator failed for field "UserCredential.activate_token": %w`, err)}
		}
	}
	if v, ok := ucu.mutation.ResetToken(); ok {
		if err := usercredential.ResetTokenValidator(v); err != nil {
			return &ValidationError{Name: "reset_token", err: fmt.Errorf(`ent: validator failed for field "UserCredential.reset_token": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ucu *UserCredentialUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserCredentialUpdate {
	ucu.modifiers = append(ucu.modifiers, modifiers...)
	return ucu
}

func (ucu *UserCredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ucu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercredential.Table, usercredential.Columns, sqlgraph.NewFieldSpec(usercredential.FieldID, field.TypeUint32))
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ucu.mutation.CreateTimeCleared() {
		_spec.ClearField(usercredential.FieldCreateTime, field.TypeTime)
	}
	if value, ok := ucu.mutation.UpdateTime(); ok {
		_spec.SetField(usercredential.FieldUpdateTime, field.TypeTime, value)
	}
	if ucu.mutation.UpdateTimeCleared() {
		_spec.ClearField(usercredential.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := ucu.mutation.DeleteTime(); ok {
		_spec.SetField(usercredential.FieldDeleteTime, field.TypeTime, value)
	}
	if ucu.mutation.DeleteTimeCleared() {
		_spec.ClearField(usercredential.FieldDeleteTime, field.TypeTime)
	}
	if ucu.mutation.TenantIDCleared() {
		_spec.ClearField(usercredential.FieldTenantID, field.TypeUint32)
	}
	if value, ok := ucu.mutation.UserID(); ok {
		_spec.SetField(usercredential.FieldUserID, field.TypeUint32, value)
	}
	if value, ok := ucu.mutation.AddedUserID(); ok {
		_spec.AddField(usercredential.FieldUserID, field.TypeUint32, value)
	}
	if ucu.mutation.UserIDCleared() {
		_spec.ClearField(usercredential.FieldUserID, field.TypeUint32)
	}
	if value, ok := ucu.mutation.IdentityType(); ok {
		_spec.SetField(usercredential.FieldIdentityType, field.TypeEnum, value)
	}
	if ucu.mutation.IdentityTypeCleared() {
		_spec.ClearField(usercredential.FieldIdentityType, field.TypeEnum)
	}
	if value, ok := ucu.mutation.Identifier(); ok {
		_spec.SetField(usercredential.FieldIdentifier, field.TypeString, value)
	}
	if ucu.mutation.IdentifierCleared() {
		_spec.ClearField(usercredential.FieldIdentifier, field.TypeString)
	}
	if value, ok := ucu.mutation.CredentialType(); ok {
		_spec.SetField(usercredential.FieldCredentialType, field.TypeEnum, value)
	}
	if ucu.mutation.CredentialTypeCleared() {
		_spec.ClearField(usercredential.FieldCredentialType, field.TypeEnum)
	}
	if value, ok := ucu.mutation.Credential(); ok {
		_spec.SetField(usercredential.FieldCredential, field.TypeString, value)
	}
	if ucu.mutation.CredentialCleared() {
		_spec.ClearField(usercredential.FieldCredential, field.TypeString)
	}
	if value, ok := ucu.mutation.IsPrimary(); ok {
		_spec.SetField(usercredential.FieldIsPrimary, field.TypeBool, value)
	}
	if ucu.mutation.IsPrimaryCleared() {
		_spec.ClearField(usercredential.FieldIsPrimary, field.TypeBool)
	}
	if value, ok := ucu.mutation.Status(); ok {
		_spec.SetField(usercredential.FieldStatus, field.TypeEnum, value)
	}
	if ucu.mutation.StatusCleared() {
		_spec.ClearField(usercredential.FieldStatus, field.TypeEnum)
	}
	if value, ok := ucu.mutation.ExtraInfo(); ok {
		_spec.SetField(usercredential.FieldExtraInfo, field.TypeString, value)
	}
	if ucu.mutation.ExtraInfoCleared() {
		_spec.ClearField(usercredential.FieldExtraInfo, field.TypeString)
	}
	if value, ok := ucu.mutation.ActivateToken(); ok {
		_spec.SetField(usercredential.FieldActivateToken, field.TypeString, value)
	}
	if ucu.mutation.ActivateTokenCleared() {
		_spec.ClearField(usercredential.FieldActivateToken, field.TypeString)
	}
	if value, ok := ucu.mutation.ResetToken(); ok {
		_spec.SetField(usercredential.FieldResetToken, field.TypeString, value)
	}
	if ucu.mutation.ResetTokenCleared() {
		_spec.ClearField(usercredential.FieldResetToken, field.TypeString)
	}
	_spec.AddModifiers(ucu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucu.mutation.done = true
	return n, nil
}

// UserCredentialUpdateOne is the builder for updating a single UserCredential entity.
type UserCredentialUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserCredentialMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (ucuo *UserCredentialUpdateOne) SetUpdateTime(t time.Time) *UserCredentialUpdateOne {
	ucuo.mutation.SetUpdateTime(t)
	return ucuo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableUpdateTime(t *time.Time) *UserCredentialUpdateOne {
	if t != nil {
		ucuo.SetUpdateTime(*t)
	}
	return ucuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ucuo *UserCredentialUpdateOne) ClearUpdateTime() *UserCredentialUpdateOne {
	ucuo.mutation.ClearUpdateTime()
	return ucuo
}

// SetDeleteTime sets the "delete_time" field.
func (ucuo *UserCredentialUpdateOne) SetDeleteTime(t time.Time) *UserCredentialUpdateOne {
	ucuo.mutation.SetDeleteTime(t)
	return ucuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableDeleteTime(t *time.Time) *UserCredentialUpdateOne {
	if t != nil {
		ucuo.SetDeleteTime(*t)
	}
	return ucuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ucuo *UserCredentialUpdateOne) ClearDeleteTime() *UserCredentialUpdateOne {
	ucuo.mutation.ClearDeleteTime()
	return ucuo
}

// SetUserID sets the "user_id" field.
func (ucuo *UserCredentialUpdateOne) SetUserID(u uint32) *UserCredentialUpdateOne {
	ucuo.mutation.ResetUserID()
	ucuo.mutation.SetUserID(u)
	return ucuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableUserID(u *uint32) *UserCredentialUpdateOne {
	if u != nil {
		ucuo.SetUserID(*u)
	}
	return ucuo
}

// AddUserID adds u to the "user_id" field.
func (ucuo *UserCredentialUpdateOne) AddUserID(u int32) *UserCredentialUpdateOne {
	ucuo.mutation.AddUserID(u)
	return ucuo
}

// ClearUserID clears the value of the "user_id" field.
func (ucuo *UserCredentialUpdateOne) ClearUserID() *UserCredentialUpdateOne {
	ucuo.mutation.ClearUserID()
	return ucuo
}

// SetIdentityType sets the "identity_type" field.
func (ucuo *UserCredentialUpdateOne) SetIdentityType(ut usercredential.IdentityType) *UserCredentialUpdateOne {
	ucuo.mutation.SetIdentityType(ut)
	return ucuo
}

// SetNillableIdentityType sets the "identity_type" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableIdentityType(ut *usercredential.IdentityType) *UserCredentialUpdateOne {
	if ut != nil {
		ucuo.SetIdentityType(*ut)
	}
	return ucuo
}

// ClearIdentityType clears the value of the "identity_type" field.
func (ucuo *UserCredentialUpdateOne) ClearIdentityType() *UserCredentialUpdateOne {
	ucuo.mutation.ClearIdentityType()
	return ucuo
}

// SetIdentifier sets the "identifier" field.
func (ucuo *UserCredentialUpdateOne) SetIdentifier(s string) *UserCredentialUpdateOne {
	ucuo.mutation.SetIdentifier(s)
	return ucuo
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableIdentifier(s *string) *UserCredentialUpdateOne {
	if s != nil {
		ucuo.SetIdentifier(*s)
	}
	return ucuo
}

// ClearIdentifier clears the value of the "identifier" field.
func (ucuo *UserCredentialUpdateOne) ClearIdentifier() *UserCredentialUpdateOne {
	ucuo.mutation.ClearIdentifier()
	return ucuo
}

// SetCredentialType sets the "credential_type" field.
func (ucuo *UserCredentialUpdateOne) SetCredentialType(ut usercredential.CredentialType) *UserCredentialUpdateOne {
	ucuo.mutation.SetCredentialType(ut)
	return ucuo
}

// SetNillableCredentialType sets the "credential_type" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableCredentialType(ut *usercredential.CredentialType) *UserCredentialUpdateOne {
	if ut != nil {
		ucuo.SetCredentialType(*ut)
	}
	return ucuo
}

// ClearCredentialType clears the value of the "credential_type" field.
func (ucuo *UserCredentialUpdateOne) ClearCredentialType() *UserCredentialUpdateOne {
	ucuo.mutation.ClearCredentialType()
	return ucuo
}

// SetCredential sets the "credential" field.
func (ucuo *UserCredentialUpdateOne) SetCredential(s string) *UserCredentialUpdateOne {
	ucuo.mutation.SetCredential(s)
	return ucuo
}

// SetNillableCredential sets the "credential" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableCredential(s *string) *UserCredentialUpdateOne {
	if s != nil {
		ucuo.SetCredential(*s)
	}
	return ucuo
}

// ClearCredential clears the value of the "credential" field.
func (ucuo *UserCredentialUpdateOne) ClearCredential() *UserCredentialUpdateOne {
	ucuo.mutation.ClearCredential()
	return ucuo
}

// SetIsPrimary sets the "is_primary" field.
func (ucuo *UserCredentialUpdateOne) SetIsPrimary(b bool) *UserCredentialUpdateOne {
	ucuo.mutation.SetIsPrimary(b)
	return ucuo
}

// SetNillableIsPrimary sets the "is_primary" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableIsPrimary(b *bool) *UserCredentialUpdateOne {
	if b != nil {
		ucuo.SetIsPrimary(*b)
	}
	return ucuo
}

// ClearIsPrimary clears the value of the "is_primary" field.
func (ucuo *UserCredentialUpdateOne) ClearIsPrimary() *UserCredentialUpdateOne {
	ucuo.mutation.ClearIsPrimary()
	return ucuo
}

// SetStatus sets the "status" field.
func (ucuo *UserCredentialUpdateOne) SetStatus(u usercredential.Status) *UserCredentialUpdateOne {
	ucuo.mutation.SetStatus(u)
	return ucuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableStatus(u *usercredential.Status) *UserCredentialUpdateOne {
	if u != nil {
		ucuo.SetStatus(*u)
	}
	return ucuo
}

// ClearStatus clears the value of the "status" field.
func (ucuo *UserCredentialUpdateOne) ClearStatus() *UserCredentialUpdateOne {
	ucuo.mutation.ClearStatus()
	return ucuo
}

// SetExtraInfo sets the "extra_info" field.
func (ucuo *UserCredentialUpdateOne) SetExtraInfo(s string) *UserCredentialUpdateOne {
	ucuo.mutation.SetExtraInfo(s)
	return ucuo
}

// SetNillableExtraInfo sets the "extra_info" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableExtraInfo(s *string) *UserCredentialUpdateOne {
	if s != nil {
		ucuo.SetExtraInfo(*s)
	}
	return ucuo
}

// ClearExtraInfo clears the value of the "extra_info" field.
func (ucuo *UserCredentialUpdateOne) ClearExtraInfo() *UserCredentialUpdateOne {
	ucuo.mutation.ClearExtraInfo()
	return ucuo
}

// SetActivateToken sets the "activate_token" field.
func (ucuo *UserCredentialUpdateOne) SetActivateToken(s string) *UserCredentialUpdateOne {
	ucuo.mutation.SetActivateToken(s)
	return ucuo
}

// SetNillableActivateToken sets the "activate_token" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableActivateToken(s *string) *UserCredentialUpdateOne {
	if s != nil {
		ucuo.SetActivateToken(*s)
	}
	return ucuo
}

// ClearActivateToken clears the value of the "activate_token" field.
func (ucuo *UserCredentialUpdateOne) ClearActivateToken() *UserCredentialUpdateOne {
	ucuo.mutation.ClearActivateToken()
	return ucuo
}

// SetResetToken sets the "reset_token" field.
func (ucuo *UserCredentialUpdateOne) SetResetToken(s string) *UserCredentialUpdateOne {
	ucuo.mutation.SetResetToken(s)
	return ucuo
}

// SetNillableResetToken sets the "reset_token" field if the given value is not nil.
func (ucuo *UserCredentialUpdateOne) SetNillableResetToken(s *string) *UserCredentialUpdateOne {
	if s != nil {
		ucuo.SetResetToken(*s)
	}
	return ucuo
}

// ClearResetToken clears the value of the "reset_token" field.
func (ucuo *UserCredentialUpdateOne) ClearResetToken() *UserCredentialUpdateOne {
	ucuo.mutation.ClearResetToken()
	return ucuo
}

// Mutation returns the UserCredentialMutation object of the builder.
func (ucuo *UserCredentialUpdateOne) Mutation() *UserCredentialMutation {
	return ucuo.mutation
}

// Where appends a list predicates to the UserCredentialUpdate builder.
func (ucuo *UserCredentialUpdateOne) Where(ps ...predicate.UserCredential) *UserCredentialUpdateOne {
	ucuo.mutation.Where(ps...)
	return ucuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UserCredentialUpdateOne) Select(field string, fields ...string) *UserCredentialUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UserCredential entity.
func (ucuo *UserCredentialUpdateOne) Save(ctx context.Context) (*UserCredential, error) {
	return withHooks(ctx, ucuo.sqlSave, ucuo.mutation, ucuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UserCredentialUpdateOne) SaveX(ctx context.Context) *UserCredential {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UserCredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UserCredentialUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucuo *UserCredentialUpdateOne) check() error {
	if v, ok := ucuo.mutation.IdentityType(); ok {
		if err := usercredential.IdentityTypeValidator(v); err != nil {
			return &ValidationError{Name: "identity_type", err: fmt.Errorf(`ent: validator failed for field "UserCredential.identity_type": %w`, err)}
		}
	}
	if v, ok := ucuo.mutation.Identifier(); ok {
		if err := usercredential.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "UserCredential.identifier": %w`, err)}
		}
	}
	if v, ok := ucuo.mutation.CredentialType(); ok {
		if err := usercredential.CredentialTypeValidator(v); err != nil {
			return &ValidationError{Name: "credential_type", err: fmt.Errorf(`ent: validator failed for field "UserCredential.credential_type": %w`, err)}
		}
	}
	if v, ok := ucuo.mutation.Credential(); ok {
		if err := usercredential.CredentialValidator(v); err != nil {
			return &ValidationError{Name: "credential", err: fmt.Errorf(`ent: validator failed for field "UserCredential.credential": %w`, err)}
		}
	}
	if v, ok := ucuo.mutation.Status(); ok {
		if err := usercredential.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserCredential.status": %w`, err)}
		}
	}
	if v, ok := ucuo.mutation.ActivateToken(); ok {
		if err := usercredential.ActivateTokenValidator(v); err != nil {
			return &ValidationError{Name: "activate_token", err: fmt.Errorf(`ent: validator failed for field "UserCredential.activate_token": %w`, err)}
		}
	}
	if v, ok := ucuo.mutation.ResetToken(); ok {
		if err := usercredential.ResetTokenValidator(v); err != nil {
			return &ValidationError{Name: "reset_token", err: fmt.Errorf(`ent: validator failed for field "UserCredential.reset_token": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ucuo *UserCredentialUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserCredentialUpdateOne {
	ucuo.modifiers = append(ucuo.modifiers, modifiers...)
	return ucuo
}

func (ucuo *UserCredentialUpdateOne) sqlSave(ctx context.Context) (_node *UserCredential, err error) {
	if err := ucuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercredential.Table, usercredential.Columns, sqlgraph.NewFieldSpec(usercredential.FieldID, field.TypeUint32))
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserCredential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercredential.FieldID)
		for _, f := range fields {
			if !usercredential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ucuo.mutation.CreateTimeCleared() {
		_spec.ClearField(usercredential.FieldCreateTime, field.TypeTime)
	}
	if value, ok := ucuo.mutation.UpdateTime(); ok {
		_spec.SetField(usercredential.FieldUpdateTime, field.TypeTime, value)
	}
	if ucuo.mutation.UpdateTimeCleared() {
		_spec.ClearField(usercredential.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := ucuo.mutation.DeleteTime(); ok {
		_spec.SetField(usercredential.FieldDeleteTime, field.TypeTime, value)
	}
	if ucuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(usercredential.FieldDeleteTime, field.TypeTime)
	}
	if ucuo.mutation.TenantIDCleared() {
		_spec.ClearField(usercredential.FieldTenantID, field.TypeUint32)
	}
	if value, ok := ucuo.mutation.UserID(); ok {
		_spec.SetField(usercredential.FieldUserID, field.TypeUint32, value)
	}
	if value, ok := ucuo.mutation.AddedUserID(); ok {
		_spec.AddField(usercredential.FieldUserID, field.TypeUint32, value)
	}
	if ucuo.mutation.UserIDCleared() {
		_spec.ClearField(usercredential.FieldUserID, field.TypeUint32)
	}
	if value, ok := ucuo.mutation.IdentityType(); ok {
		_spec.SetField(usercredential.FieldIdentityType, field.TypeEnum, value)
	}
	if ucuo.mutation.IdentityTypeCleared() {
		_spec.ClearField(usercredential.FieldIdentityType, field.TypeEnum)
	}
	if value, ok := ucuo.mutation.Identifier(); ok {
		_spec.SetField(usercredential.FieldIdentifier, field.TypeString, value)
	}
	if ucuo.mutation.IdentifierCleared() {
		_spec.ClearField(usercredential.FieldIdentifier, field.TypeString)
	}
	if value, ok := ucuo.mutation.CredentialType(); ok {
		_spec.SetField(usercredential.FieldCredentialType, field.TypeEnum, value)
	}
	if ucuo.mutation.CredentialTypeCleared() {
		_spec.ClearField(usercredential.FieldCredentialType, field.TypeEnum)
	}
	if value, ok := ucuo.mutation.Credential(); ok {
		_spec.SetField(usercredential.FieldCredential, field.TypeString, value)
	}
	if ucuo.mutation.CredentialCleared() {
		_spec.ClearField(usercredential.FieldCredential, field.TypeString)
	}
	if value, ok := ucuo.mutation.IsPrimary(); ok {
		_spec.SetField(usercredential.FieldIsPrimary, field.TypeBool, value)
	}
	if ucuo.mutation.IsPrimaryCleared() {
		_spec.ClearField(usercredential.FieldIsPrimary, field.TypeBool)
	}
	if value, ok := ucuo.mutation.Status(); ok {
		_spec.SetField(usercredential.FieldStatus, field.TypeEnum, value)
	}
	if ucuo.mutation.StatusCleared() {
		_spec.ClearField(usercredential.FieldStatus, field.TypeEnum)
	}
	if value, ok := ucuo.mutation.ExtraInfo(); ok {
		_spec.SetField(usercredential.FieldExtraInfo, field.TypeString, value)
	}
	if ucuo.mutation.ExtraInfoCleared() {
		_spec.ClearField(usercredential.FieldExtraInfo, field.TypeString)
	}
	if value, ok := ucuo.mutation.ActivateToken(); ok {
		_spec.SetField(usercredential.FieldActivateToken, field.TypeString, value)
	}
	if ucuo.mutation.ActivateTokenCleared() {
		_spec.ClearField(usercredential.FieldActivateToken, field.TypeString)
	}
	if value, ok := ucuo.mutation.ResetToken(); ok {
		_spec.SetField(usercredential.FieldResetToken, field.TypeString, value)
	}
	if ucuo.mutation.ResetTokenCleared() {
		_spec.ClearField(usercredential.FieldResetToken, field.TypeString)
	}
	_spec.AddModifiers(ucuo.modifiers...)
	_node = &UserCredential{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucuo.mutation.done = true
	return _node, nil
}
