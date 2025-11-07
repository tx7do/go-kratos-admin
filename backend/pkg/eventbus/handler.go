package eventbus

import (
	"context"
)

// EventHandler is a function that handles an event
type EventHandler func(ctx context.Context, event *Event) error

// EventHandlerFunc is an adapter to allow ordinary functions to be used as EventHandlers
type EventHandlerFunc func(ctx context.Context, event *Event) error

// Handle calls the handler function
func (f EventHandlerFunc) Handle(ctx context.Context, event *Event) error {
	return f(ctx, event)
}

// Handler interface for event handlers
type Handler interface {
	Handle(ctx context.Context, event *Event) error
}

// AsyncHandler wraps a handler to execute asynchronously
type AsyncHandler struct {
	handler Handler
}

// NewAsyncHandler creates a new async handler
func NewAsyncHandler(handler Handler) *AsyncHandler {
	return &AsyncHandler{handler: handler}
}

// Handle executes the handler asynchronously
func (h *AsyncHandler) Handle(ctx context.Context, event *Event) error {
	go func() {
		_ = h.handler.Handle(ctx, event)
	}()
	return nil
}

// ChainHandler chains multiple handlers together
type ChainHandler struct {
	handlers []Handler
}

// NewChainHandler creates a new chain handler
func NewChainHandler(handlers ...Handler) *ChainHandler {
	return &ChainHandler{handlers: handlers}
}

// Handle executes all handlers in sequence
func (h *ChainHandler) Handle(ctx context.Context, event *Event) error {
	for _, handler := range h.handlers {
		if err := handler.Handle(ctx, event); err != nil {
			return err
		}
	}
	return nil
}

// FilterHandler wraps a handler with a filter function
type FilterHandler struct {
	filter  func(*Event) bool
	handler Handler
}

// NewFilterHandler creates a new filter handler
func NewFilterHandler(filter func(*Event) bool, handler Handler) *FilterHandler {
	return &FilterHandler{
		filter:  filter,
		handler: handler,
	}
}

// Handle executes the handler only if the filter returns true
func (h *FilterHandler) Handle(ctx context.Context, event *Event) error {
	if h.filter(event) {
		return h.handler.Handle(ctx, event)
	}
	return nil
}
