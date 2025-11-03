package server

import (
	"github.com/go-kratos/kratos/v2/log"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-transport/transport/sse"
)

// NewSseServer creates a new SSE server.
func NewSseServer(cfg *conf.Bootstrap, logger log.Logger) *sse.Server {
	if cfg == nil || cfg.Server == nil || cfg.Server.Sse == nil {
		return nil
	}

	l := log.NewHelper(log.With(logger, "module", "sse-server/admin-service"))

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
