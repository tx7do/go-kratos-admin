package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"github.com/tx7do/kratos-transport/transport/sse"
)

// NewSseServer creates a new SSE server.
func NewSseServer(ctx *bootstrap.Context) *sse.Server {
	if ctx.Config == nil || ctx.Config.Server == nil || ctx.Config.Server.Sse == nil {
		return nil
	}

	l := log.NewHelper(log.With(ctx.Logger, "module", "sse-server/admin-service"))

	s := sse.NewServer(
		sse.WithAddress(ctx.Config.Server.Sse.GetAddr()),
		sse.WithCodec(ctx.Config.Server.Sse.GetCodec()),
		sse.WithPath(ctx.Config.Server.Sse.GetPath()),
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
