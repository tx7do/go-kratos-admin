package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	swaggerUI "github.com/tx7do/kratos-swagger-ui"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"kratos-admin/app/admin/service/cmd/server/assets"

	"kratos-admin/app/admin/service/internal/data"
	"kratos-admin/app/admin/service/internal/service"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"

	"kratos-admin/pkg/middleware/auth"
	applogging "kratos-admin/pkg/middleware/logging"
)

// NewWhiteListMatcher 创建jwt白名单
func newRestWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[adminV1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newRestMiddleware(
	logger log.Logger,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	userToken *data.UserToken,
	operationLogRepo *data.AdminOperationLogRepo,
	loginLogRepo *data.AdminLoginLogRepo,
) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))

	ms = append(ms, applogging.Server(
		applogging.WithAuthenticator(authenticator),
		applogging.WithWriteOperationLogFunc(func(ctx context.Context, data *adminV1.AdminOperationLog) error {
			// TODO 如果系统的负载比较小，可以同步写入数据库，否则，建议使用异步方式，即投递进队列。
			return operationLogRepo.Create(ctx, &adminV1.CreateAdminOperationLogRequest{Data: data})
		}),
		applogging.WithWriteLoginLogFunc(func(ctx context.Context, data *adminV1.AdminLoginLog) error {
			// TODO 如果系统的负载比较小，可以同步写入数据库，否则，建议使用异步方式，即投递进队列。
			return loginLogRepo.Create(ctx, &adminV1.CreateAdminLoginLogRequest{Data: data})
		}),
	))

	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())

	return ms
}

// NewRESTServer new an HTTP server.
func NewRESTServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	userToken *data.UserToken,
	operationLogRepo *data.AdminOperationLogRepo,
	loginLogRepo *data.AdminLoginLogRepo,
	authnSvc *service.AuthenticationService,
	userSvc *service.UserService,
	menuSvc *service.MenuService,
	routerSvc *service.RouterService,
	orgSvc *service.OrganizationService,
	roleSvc *service.RoleService,
	positionSvc *service.PositionService,
	dictSvc *service.DictService,
	deptSvc *service.DepartmentService,
	adminLoginLogSvc *service.AdminLoginLogService,
	adminOperationLogSvc *service.AdminOperationLogService,
	ossSvc *service.OssService,
	ueditorSvc *service.UEditorService,
	fileService *service.FileService,
	tenantService *service.TenantService,
	taskService *service.TaskService,
	notificationMessageService *service.NotificationMessageService,
	notificationMessageCategoryService *service.NotificationMessageCategoryService,
	notificationMessageRecipientService *service.NotificationMessageRecipientService,
	privateMessageService *service.PrivateMessageService,
	adminLoginRestrictionService *service.AdminLoginRestrictionService,
	userProfileService *service.UserProfileService,
) *http.Server {
	if cfg == nil || cfg.Server == nil || cfg.Server.Rest == nil {
		return nil
	}

	srv := rpc.CreateRestServer(cfg,
		newRestMiddleware(logger, authenticator, authorizer, userToken, operationLogRepo, loginLogRepo)...,
	)

	adminV1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)

	adminV1.RegisterUserProfileServiceHTTPServer(srv, userProfileService)

	adminV1.RegisterMenuServiceHTTPServer(srv, menuSvc)
	adminV1.RegisterRouterServiceHTTPServer(srv, routerSvc)
	adminV1.RegisterDictServiceHTTPServer(srv, dictSvc)
	adminV1.RegisterTaskServiceHTTPServer(srv, taskService)
	adminV1.RegisterAdminLoginRestrictionServiceHTTPServer(srv, adminLoginRestrictionService)

	adminV1.RegisterUserServiceHTTPServer(srv, userSvc)
	adminV1.RegisterOrganizationServiceHTTPServer(srv, orgSvc)
	adminV1.RegisterRoleServiceHTTPServer(srv, roleSvc)
	adminV1.RegisterPositionServiceHTTPServer(srv, positionSvc)
	adminV1.RegisterDepartmentServiceHTTPServer(srv, deptSvc)
	adminV1.RegisterTenantServiceHTTPServer(srv, tenantService)

	adminV1.RegisterAdminLoginLogServiceHTTPServer(srv, adminLoginLogSvc)
	adminV1.RegisterAdminOperationLogServiceHTTPServer(srv, adminOperationLogSvc)

	adminV1.RegisterOssServiceHTTPServer(srv, ossSvc)
	adminV1.RegisterFileServiceHTTPServer(srv, fileService)

	adminV1.RegisterUEditorServiceHTTPServer(srv, ueditorSvc)

	adminV1.RegisterNotificationMessageServiceHTTPServer(srv, notificationMessageService)
	adminV1.RegisterNotificationMessageCategoryServiceHTTPServer(srv, notificationMessageCategoryService)
	adminV1.RegisterNotificationMessageRecipientServiceHTTPServer(srv, notificationMessageRecipientService)
	adminV1.RegisterPrivateMessageServiceHTTPServer(srv, privateMessageService)

	registerFileUploadHandler(srv, ossSvc)
	registerUEditorUploadHandler(srv, ueditorSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Kratos Admin"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
