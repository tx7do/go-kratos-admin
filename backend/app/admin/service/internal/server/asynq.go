package server

import (
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"github.com/tx7do/kratos-bootstrap/transport/asynq"

	asynqServer "github.com/tx7do/kratos-transport/transport/asynq"
)

// NewAsynqServer creates a new asynq server.
func NewAsynqServer(ctx *bootstrap.Context) *asynqServer.Server {
	cfg := ctx.GetConfig()

	if cfg == nil || cfg.Server == nil || cfg.Server.Asynq == nil {
		return nil
	}

	srv := asynq.NewAsynqServer(cfg.Server.Asynq)

	return srv
}
