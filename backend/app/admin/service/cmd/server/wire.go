//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"

	dataProviders "go-wind-admin/app/admin/service/internal/data/providers"
	serverProviders "go-wind-admin/app/admin/service/internal/server/providers"
	serviceProviders "go-wind-admin/app/admin/service/internal/service/providers"
)

// initApp init kratos application.
func initApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			serverProviders.ProviderSet,
			serviceProviders.ProviderSet,
			dataProviders.ProviderSet,
			newApp,
		),
	)
}
