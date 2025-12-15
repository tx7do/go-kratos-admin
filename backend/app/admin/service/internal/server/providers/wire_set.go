//go:build wireinject
// +build wireinject

// 仅用于定义依赖注入的集合，不包含任何业务逻辑
// Only used to define the dependency injection provider set; contains no business logic.

package providers

import (
	"github.com/google/wire"

	"go-wind-admin/app/admin/service/internal/server"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	server.NewRESTServer,
	server.NewAsynqServer,
	server.NewSseServer,
)
