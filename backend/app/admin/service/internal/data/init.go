//go:build wireinject
// +build wireinject

package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewRedisClient,
	NewEntClient,

	NewAuthenticator,
	NewAuthorizer,

	NewPasswordCrypto,

	NewMinIoClient,

	NewMenuRepo,
	NewDictMainRepo,
	NewDictItemRepo,
	NewTaskRepo,
	NewAdminLoginRestrictionRepo,
	NewApiResourceRepo,

	NewOrganizationRepo,
	NewDepartmentRepo,
	NewPositionRepo,
	NewRoleRepo,
	NewUserRepo,
	NewTenantRepo,
	NewUserCredentialRepo,

	NewRoleApiRepo,
	NewRoleDeptRepo,
	NewRoleMenuRepo,
	NewRoleOrgRepo,
	NewRolePositionRepo,
	NewUserRoleRepo,
	NewUserPositionRepo,

	NewAdminLoginLogRepo,
	NewAdminOperationLogRepo,

	NewFileRepo,

	NewNotificationMessageRepo,
	NewNotificationMessageCategoryRepo,
	NewNotificationMessageRecipientRepo,
	NewPrivateMessageRepo,

	NewUserTokenRepo,
)
