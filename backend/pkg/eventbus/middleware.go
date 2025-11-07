package eventbus

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Middleware is a function that wraps a Handler
type Middleware func(Handler) Handler

// LoggingMiddleware logs event handling
func LoggingMiddleware(logger *log.Helper) Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			logger.Infof("Handling event: type=%s, id=%s, source=%s", event.Type, event.ID, event.Source)
			err := next.Handle(ctx, event)
			if err != nil {
				logger.Errorf("Error handling event %s: %v", event.ID, err)
			} else {
				logger.Debugf("Successfully handled event: %s", event.ID)
			}
			return err
		})
	}
}

// RecoveryMiddleware recovers from panics in handlers
func RecoveryMiddleware(logger *log.Helper) Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) (err error) {
			defer func() {
				if r := recover(); r != nil {
					logger.Errorf("Panic recovered in event handler for %s: %v", event.Type, r)
					err = &PanicError{Value: r}
				}
			}()
			return next.Handle(ctx, event)
		})
	}
}

// TimeoutMiddleware adds timeout to event handling
func TimeoutMiddleware(timeout time.Duration) Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			done := make(chan error, 1)
			go func() {
				done <- next.Handle(ctx, event)
			}()

			select {
			case err := <-done:
				return err
			case <-ctx.Done():
				return &TimeoutError{
					EventID:   event.ID,
					EventType: event.Type,
					Timeout:   timeout,
				}
			}
		})
	}
}

// RetryMiddleware retries failed event handling
func RetryMiddleware(maxRetries int, delay time.Duration) Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			var err error
			for i := 0; i <= maxRetries; i++ {
				err = next.Handle(ctx, event)
				if err == nil {
					return nil
				}

				if i < maxRetries {
					time.Sleep(delay)
				}
			}
			return err
		})
	}
}

// MetricsMiddleware collects metrics for event handling
func MetricsMiddleware(logger *log.Helper) Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			start := time.Now()
			err := next.Handle(ctx, event)
			duration := time.Since(start)

			logger.Infof("Event handling metrics: type=%s, duration=%s, success=%v",
				event.Type, duration, err == nil)

			return err
		})
	}
}

// Chain chains multiple middlewares together
func Chain(middlewares ...Middleware) Middleware {
	return func(handler Handler) Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

// PanicError represents a panic that occurred during event handling
type PanicError struct {
	Value interface{}
}

func (e *PanicError) Error() string {
	return "panic in event handler"
}

// TimeoutError represents a timeout during event handling
type TimeoutError struct {
	EventID   string
	EventType string
	Timeout   time.Duration
}

func (e *TimeoutError) Error() string {
	return "event handling timeout"
}
