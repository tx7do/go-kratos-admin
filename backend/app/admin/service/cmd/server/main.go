package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/tx7do/kratos-transport/transport/asynq"
	"github.com/tx7do/kratos-transport/transport/sse"

	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"kratos-admin/pkg/service"
)

var version string

// go build -ldflags "-X main.version=x.y.z"

func newApp(
	lg log.Logger,
	re registry.Registrar,
	hs *http.Server,
	as *asynq.Server,
	ss *sse.Server,
) *kratos.App {
	return bootstrap.NewApp(lg, re, hs, as, ss)
}

func main() {
	bootstrap.Bootstrap(initApp, trans.Ptr(service.AdminService), trans.Ptr(version))
}
