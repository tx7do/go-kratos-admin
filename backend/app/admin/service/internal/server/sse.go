package server

import (
	"github.com/go-kratos/kratos/v2/log"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-transport/transport/sse"
)

// NewSseServer creates a new SSE server.
func NewSseServer(cfg *conf.Bootstrap, _ log.Logger) *sse.Server {
	if cfg == nil || cfg.Server == nil || cfg.Server.Sse == nil {
		return nil
	}

	s := sse.NewServer(
		sse.WithAddress(cfg.Server.Sse.GetAddr()),
		sse.WithCodec(cfg.Server.Sse.GetCodec()),
		sse.WithPath(cfg.Server.Sse.GetPath()),
		sse.WithSubscriberFunction(func(streamID sse.StreamID, sub *sse.Subscriber) {
			//var token string
			//if sub.URL != nil {
			//	token = sub.URL.Query().Get("token")
			//}
		}),
	)

	return s
}
