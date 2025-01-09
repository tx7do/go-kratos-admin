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
	"kratos-admin/app/admin/service/internal/service"

	"kratos-admin/pkg/cache"
	"kratos-admin/pkg/middleware/auth"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
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
	userToken *cache.UserToken,
) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(userToken),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())
	return ms
}

// NewRESTServer new an HTTP server.
func NewRESTServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	userToken *cache.UserToken,
	authnSvc *service.AuthenticationService,
	userSvc *service.UserService,
	menuSvc *service.MenuService,
	routerSvc *service.RouterService,
	orgSvc *service.OrganizationService,
	roleSvc *service.RoleService,
	positionSvc *service.PositionService,
	dictSvc *service.DictService,
) *http.Server {
	srv := rpc.CreateRestServer(cfg, newRestMiddleware(logger, authenticator, authorizer, userToken)...)

	adminV1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)
	adminV1.RegisterUserServiceHTTPServer(srv, userSvc)
	adminV1.RegisterMenuServiceHTTPServer(srv, menuSvc)
	adminV1.RegisterRouterServiceHTTPServer(srv, routerSvc)
	adminV1.RegisterOrganizationServiceHTTPServer(srv, orgSvc)
	adminV1.RegisterRoleServiceHTTPServer(srv, roleSvc)
	adminV1.RegisterPositionServiceHTTPServer(srv, positionSvc)
	adminV1.RegisterDictServiceHTTPServer(srv, dictSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Kratos Admin"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
