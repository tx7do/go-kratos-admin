//go:build wireinject
// +build wireinject

// 仅用于定义依赖注入的集合，不包含任何业务逻辑
// Only used to define the dependency injection provider set; contains no business logic.

package providers

import (
	"github.com/google/wire"

	"go-wind-admin/app/admin/service/internal/data"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	data.NewData,

	data.NewRedisClient,
	data.NewEntClient,

	data.NewAuthenticator,
	data.NewAuthorizer,

	data.NewPasswordCrypto,

	data.NewMinIoClient,

	data.NewMenuRepo,
	data.NewDictTypeRepo,
	data.NewDictEntryRepo,
	data.NewTaskRepo,
	data.NewAdminLoginRestrictionRepo,
	data.NewApiResourceRepo,

	data.NewOrganizationRepo,
	data.NewDepartmentRepo,
	data.NewPositionRepo,
	data.NewRoleRepo,
	data.NewUserRepo,
	data.NewTenantRepo,
	data.NewUserCredentialRepo,

	data.NewRoleApiRepo,
	data.NewRoleDeptRepo,
	data.NewRoleMenuRepo,
	data.NewRoleOrgRepo,
	data.NewRolePositionRepo,
	data.NewUserRoleRepo,
	data.NewUserPositionRepo,

	data.NewAdminLoginLogRepo,
	data.NewAdminOperationLogRepo,

	data.NewFileRepo,

	data.NewInternalMessageRepo,
	data.NewInternalMessageCategoryRepo,
	data.NewInternalMessageRecipientRepo,

	data.NewUserTokenRepo,
)
