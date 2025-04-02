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

	NewMinIoClient,

	NewMenuRepo,
	NewDictRepo,
	NewTaskRepo,

	NewOrganizationRepo,
	NewDepartmentRepo,
	NewPositionRepo,
	NewRoleRepo,
	NewUserRepo,
	NewTenantRepo,

	NewAdminLoginLogRepo,
	NewAdminOperationLogRepo,

	NewFileRepo,

	NewNotificationMessageRepo,
	NewNotificationMessageCategoryRepo,
	NewNotificationMessageRecipientRepo,
	NewPrivateMessageRepo,

	NewUserTokenRepo,
)
