package server

import (
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"github.com/tx7do/kratos-transport/transport/sse"
)

// NewSseServer creates a new SSE server.
func NewSseServer(ctx *bootstrap.Context) *sse.Server {
	cfg := ctx.GetConfig()

	if cfg == nil || cfg.Server == nil || cfg.Server.Sse == nil {
		return nil
	}

	l := ctx.NewLoggerHelper("sse-server/admin-service")

	s := sse.NewServer(
		sse.WithAddress(cfg.Server.Sse.GetAddr()),
		sse.WithCodec(cfg.Server.Sse.GetCodec()),
		sse.WithPath(cfg.Server.Sse.GetPath()),
		sse.WithAutoStream(true),
		sse.WithAutoReply(false),
		sse.WithSubscriberFunction(func(streamID sse.StreamID, sub *sse.Subscriber) {
			//l.Infof("SSE: [%s]", sub.URL)
			l.Infof("subscriber [%s] connected", streamID)
		}),
	)

	//s.CreateStream("test")

	return s
}
