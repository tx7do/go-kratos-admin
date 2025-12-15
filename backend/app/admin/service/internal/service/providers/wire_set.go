//go:build wireinject
// +build wireinject

// 仅用于定义依赖注入的集合，不包含任何业务逻辑
// Only used to define the dependency injection provider set; contains no business logic.

package providers

import (
	"go-wind-admin/app/admin/service/internal/service"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	service.NewAuthenticationService,
	service.NewUserService,
	service.NewMenuService,
	service.NewRouterService,
	service.NewTaskService,
	service.NewRoleService,
	service.NewOrganizationService,
	service.NewDepartmentService,
	service.NewPositionService,
	service.NewDictService,
	service.NewAdminLoginLogService,
	service.NewAdminOperationLogService,
	service.NewOssService,
	service.NewUEditorService,
	service.NewFileService,
	service.NewTenantService,
	service.NewInternalMessageService,
	service.NewInternalMessageCategoryService,
	service.NewInternalMessageRecipientService,
	service.NewAdminLoginRestrictionService,
	service.NewUserProfileService,
	service.NewUserCredentialService,
	service.NewApiResourceService,
)
