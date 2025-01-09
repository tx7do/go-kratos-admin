// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"kratos-admin/app/admin/service/internal/data"
	"kratos-admin/app/admin/service/internal/server"
	"kratos-admin/app/admin/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *v1.Bootstrap) (*kratos.App, func(), error) {
	authenticator := data.NewAuthenticator(bootstrap)
	engine := data.NewAuthorizer()
	entClient := data.NewEntClient(bootstrap, logger)
	client := data.NewRedisClient(bootstrap, logger)
	dataData, cleanup, err := data.NewData(entClient, client, authenticator, engine, logger)
	if err != nil {
		return nil, nil, err
	}
	userToken := data.NewUserTokenRepo(dataData, authenticator, logger)
	userRepo := data.NewUserRepo(dataData, logger)
	authenticationService := service.NewAuthenticationService(logger, userRepo, userToken)
	userService := service.NewUserService(logger, userRepo)
	menuRepo := data.NewMenuRepo(dataData, logger)
	menuService := service.NewMenuService(menuRepo, logger)
	routerService := service.NewRouterService(logger, menuRepo)
	organizationRepo := data.NewOrganizationRepo(dataData, logger)
	organizationService := service.NewOrganizationService(organizationRepo, logger)
	roleRepo := data.NewRoleRepo(dataData, logger)
	roleService := service.NewRoleService(roleRepo, logger)
	positionRepo := data.NewPositionRepo(dataData, logger)
	positionService := service.NewPositionService(positionRepo, logger)
	dictRepo := data.NewDictRepo(dataData, logger)
	dictService := service.NewDictService(dictRepo, logger)
	httpServer := server.NewRESTServer(bootstrap, logger, authenticator, engine, userToken, authenticationService, userService, menuService, routerService, organizationService, roleService, positionService, dictService)
	app := newApp(logger, registrar, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
