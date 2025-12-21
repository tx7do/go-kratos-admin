//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire

// 本文件包含 Wire 的 provider 组合，仅用于运行 Wire 代码生成器。
// 构建标签 `wireinject` 确保此文件不会被常规的 `go build` 包含到最终二进制中。
// Wire 运行后会生成不带该标签的文件（例如 `wire_gen.go`），生成文件会被包含到最终构建中。
//
// This file holds Wire provider setup used only by the Wire code generator.
// The build tag `wireinject` ensures this file is excluded from normal `go build`/final binaries.
// The generated file (e.g. `wire_gen.go`) does not have this tag and will be included in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	dataProviders "go-wind-admin/app/admin/service/internal/data/providers"
	serverProviders "go-wind-admin/app/admin/service/internal/server/providers"
	serviceProviders "go-wind-admin/app/admin/service/internal/service/providers"
)

// initApp 初始化 kratos 应用的 Wire provider 入口。
// initApp initializes the Wire provider entry for the kratos application.
//
// 参数 / Parameters:
//   - *bootstrap.Context: 引导上下文 / *bootstrap.Context: bootstrap context
//
// 返回 / Returns:
//   - *kratos.App: 已构建的应用实例 / *kratos.App: constructed application instance
//   - func(): 应用关闭时的清理函数 / func(): cleanup function to run on shutdown
//   - error: 构建过程中可能发生的错误 / error: possible construction error
func initApp(*bootstrap.Context) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			serverProviders.ProviderSet,
			serviceProviders.ProviderSet,
			dataProviders.ProviderSet,
			newApp,
		),
	)
}
