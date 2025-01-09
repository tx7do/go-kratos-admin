// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-admin/app/admin/service/internal/data/ent/dict"
	"kratos-admin/app/admin/service/internal/data/ent/menu"
	"kratos-admin/app/admin/service/internal/data/ent/organization"
	"kratos-admin/app/admin/service/internal/data/ent/position"
	"kratos-admin/app/admin/service/internal/data/ent/role"
	"kratos-admin/app/admin/service/internal/data/ent/schema"
	"kratos-admin/app/admin/service/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	dictMixin := schema.Dict{}.Mixin()
	dictMixinFields0 := dictMixin[0].Fields()
	_ = dictMixinFields0
	dictMixinFields2 := dictMixin[2].Fields()
	_ = dictMixinFields2
	dictMixinFields5 := dictMixin[5].Fields()
	_ = dictMixinFields5
	dictFields := schema.Dict{}.Fields()
	_ = dictFields
	// dictDescRemark is the schema descriptor for remark field.
	dictDescRemark := dictMixinFields5[0].Descriptor()
	// dict.DefaultRemark holds the default value on creation for the remark field.
	dict.DefaultRemark = dictDescRemark.Default.(string)
	// dictDescSortID is the schema descriptor for sort_id field.
	dictDescSortID := dictFields[6].Descriptor()
	// dict.DefaultSortID holds the default value on creation for the sort_id field.
	dict.DefaultSortID = dictDescSortID.Default.(int32)
	// dictDescID is the schema descriptor for id field.
	dictDescID := dictMixinFields0[0].Descriptor()
	// dict.IDValidator is a validator for the "id" field. It is called by the builders before save.
	dict.IDValidator = dictDescID.Validators[0].(func(uint32) error)
	menuMixin := schema.Menu{}.Mixin()
	menuMixinFields0 := menuMixin[0].Fields()
	_ = menuMixinFields0
	menuMixinFields4 := menuMixin[4].Fields()
	_ = menuMixinFields4
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescRemark is the schema descriptor for remark field.
	menuDescRemark := menuMixinFields4[0].Descriptor()
	// menu.DefaultRemark holds the default value on creation for the remark field.
	menu.DefaultRemark = menuDescRemark.Default.(string)
	// menuDescPath is the schema descriptor for path field.
	menuDescPath := menuFields[3].Descriptor()
	// menu.DefaultPath holds the default value on creation for the path field.
	menu.DefaultPath = menuDescPath.Default.(string)
	// menuDescComponent is the schema descriptor for component field.
	menuDescComponent := menuFields[7].Descriptor()
	// menu.DefaultComponent holds the default value on creation for the component field.
	menu.DefaultComponent = menuDescComponent.Default.(string)
	// menuDescID is the schema descriptor for id field.
	menuDescID := menuFields[0].Descriptor()
	// menu.IDValidator is a validator for the "id" field. It is called by the builders before save.
	menu.IDValidator = menuDescID.Validators[0].(func(int32) error)
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationMixinFields2 := organizationMixin[2].Fields()
	_ = organizationMixinFields2
	organizationMixinFields5 := organizationMixin[5].Fields()
	_ = organizationMixinFields5
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescRemark is the schema descriptor for remark field.
	organizationDescRemark := organizationMixinFields5[0].Descriptor()
	// organization.DefaultRemark holds the default value on creation for the remark field.
	organization.DefaultRemark = organizationDescRemark.Default.(string)
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[0].Descriptor()
	// organization.DefaultName holds the default value on creation for the name field.
	organization.DefaultName = organizationDescName.Default.(string)
	// organizationDescSortID is the schema descriptor for sort_id field.
	organizationDescSortID := organizationFields[2].Descriptor()
	// organization.DefaultSortID holds the default value on creation for the sort_id field.
	organization.DefaultSortID = organizationDescSortID.Default.(int32)
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationMixinFields0[0].Descriptor()
	// organization.IDValidator is a validator for the "id" field. It is called by the builders before save.
	organization.IDValidator = organizationDescID.Validators[0].(func(uint32) error)
	positionMixin := schema.Position{}.Mixin()
	positionMixinFields0 := positionMixin[0].Fields()
	_ = positionMixinFields0
	positionMixinFields2 := positionMixin[2].Fields()
	_ = positionMixinFields2
	positionMixinFields5 := positionMixin[5].Fields()
	_ = positionMixinFields5
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescRemark is the schema descriptor for remark field.
	positionDescRemark := positionMixinFields5[0].Descriptor()
	// position.DefaultRemark holds the default value on creation for the remark field.
	position.DefaultRemark = positionDescRemark.Default.(string)
	// positionDescName is the schema descriptor for name field.
	positionDescName := positionFields[0].Descriptor()
	// position.DefaultName holds the default value on creation for the name field.
	position.DefaultName = positionDescName.Default.(string)
	// position.NameValidator is a validator for the "name" field. It is called by the builders before save.
	position.NameValidator = positionDescName.Validators[0].(func(string) error)
	// positionDescCode is the schema descriptor for code field.
	positionDescCode := positionFields[1].Descriptor()
	// position.DefaultCode holds the default value on creation for the code field.
	position.DefaultCode = positionDescCode.Default.(string)
	// position.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	position.CodeValidator = positionDescCode.Validators[0].(func(string) error)
	// positionDescParentID is the schema descriptor for parent_id field.
	positionDescParentID := positionFields[2].Descriptor()
	// position.DefaultParentID holds the default value on creation for the parent_id field.
	position.DefaultParentID = positionDescParentID.Default.(uint32)
	// positionDescSortID is the schema descriptor for sort_id field.
	positionDescSortID := positionFields[3].Descriptor()
	// position.DefaultSortID holds the default value on creation for the sort_id field.
	position.DefaultSortID = positionDescSortID.Default.(int32)
	// positionDescID is the schema descriptor for id field.
	positionDescID := positionMixinFields0[0].Descriptor()
	// position.IDValidator is a validator for the "id" field. It is called by the builders before save.
	position.IDValidator = positionDescID.Validators[0].(func(uint32) error)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields2 := roleMixin[2].Fields()
	_ = roleMixinFields2
	roleMixinFields5 := roleMixin[5].Fields()
	_ = roleMixinFields5
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescRemark is the schema descriptor for remark field.
	roleDescRemark := roleMixinFields5[0].Descriptor()
	// role.DefaultRemark holds the default value on creation for the remark field.
	role.DefaultRemark = roleDescRemark.Default.(string)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescCode is the schema descriptor for code field.
	roleDescCode := roleFields[1].Descriptor()
	// role.DefaultCode holds the default value on creation for the code field.
	role.DefaultCode = roleDescCode.Default.(string)
	// role.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	role.CodeValidator = roleDescCode.Validators[0].(func(string) error)
	// roleDescSortID is the schema descriptor for sort_id field.
	roleDescSortID := roleFields[3].Descriptor()
	// role.DefaultSortID holds the default value on creation for the sort_id field.
	role.DefaultSortID = roleDescSortID.Default.(int32)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleMixinFields0[0].Descriptor()
	// role.IDValidator is a validator for the "id" field. It is called by the builders before save.
	role.IDValidator = roleDescID.Validators[0].(func(uint32) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields4 := userMixin[4].Fields()
	_ = userMixinFields4
	userMixinFields5 := userMixin[5].Fields()
	_ = userMixinFields5
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescRemark is the schema descriptor for remark field.
	userDescRemark := userMixinFields4[0].Descriptor()
	// user.DefaultRemark holds the default value on creation for the remark field.
	user.DefaultRemark = userDescRemark.Default.(string)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[2].Descriptor()
	// user.NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	user.NickNameValidator = userDescNickName.Validators[0].(func(string) error)
	// userDescRealName is the schema descriptor for real_name field.
	userDescRealName := userFields[3].Descriptor()
	// user.RealNameValidator is a validator for the "real_name" field. It is called by the builders before save.
	user.RealNameValidator = userDescRealName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[4].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[5].Descriptor()
	// user.DefaultMobile holds the default value on creation for the mobile field.
	user.DefaultMobile = userDescMobile.Default.(string)
	// user.MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	user.MobileValidator = userDescMobile.Validators[0].(func(string) error)
	// userDescTelephone is the schema descriptor for telephone field.
	userDescTelephone := userFields[6].Descriptor()
	// user.DefaultTelephone holds the default value on creation for the telephone field.
	user.DefaultTelephone = userDescTelephone.Default.(string)
	// user.TelephoneValidator is a validator for the "telephone" field. It is called by the builders before save.
	user.TelephoneValidator = userDescTelephone.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[7].Descriptor()
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[9].Descriptor()
	// user.DefaultAddress holds the default value on creation for the address field.
	user.DefaultAddress = userDescAddress.Default.(string)
	// user.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	user.AddressValidator = userDescAddress.Validators[0].(func(string) error)
	// userDescRegion is the schema descriptor for region field.
	userDescRegion := userFields[10].Descriptor()
	// user.DefaultRegion holds the default value on creation for the region field.
	user.DefaultRegion = userDescRegion.Default.(string)
	// user.RegionValidator is a validator for the "region" field. It is called by the builders before save.
	user.RegionValidator = userDescRegion.Validators[0].(func(string) error)
	// userDescDescription is the schema descriptor for description field.
	userDescDescription := userFields[11].Descriptor()
	// user.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	user.DescriptionValidator = userDescDescription.Validators[0].(func(string) error)
	// userDescLastLoginIP is the schema descriptor for last_login_ip field.
	userDescLastLoginIP := userFields[14].Descriptor()
	// user.DefaultLastLoginIP holds the default value on creation for the last_login_ip field.
	user.DefaultLastLoginIP = userDescLastLoginIP.Default.(string)
	// user.LastLoginIPValidator is a validator for the "last_login_ip" field. It is called by the builders before save.
	user.LastLoginIPValidator = userDescLastLoginIP.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(uint32) error)
}
