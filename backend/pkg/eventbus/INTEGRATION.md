# EventBus Integration Guide

## Integration with Email Processor

Here's how to integrate the EventBus with the email processing system:

### 1. Setup EventBus in Your Application

```go
// In your main.go or initialization code
import (
    "github.com/go-kratos/kratos/v2/log"
    "go-wind-admin/pkg/eventbus"
)

// Create event bus manager
eventBusManager := eventbus.NewManager(logger)
defer eventBusManager.Close()

// Get or create email event bus
emailEventBus := eventBusManager.GetBus("email")
```

### 2. Update Email Processor Handler

Modify the email processor to publish events:

```go
// In email_processor.go
type EmailProcessorTaskHandler struct {
    *BaseTaskHandler
    emailRepo *email.Repository
    eventBus  eventbus.EventBus  // Add this field
}

func NewEmailProcessorTaskHandler(emailRepo email.Repository, eventBus eventbus.EventBus) task.TaskHandler {
    // ... existing code ...
    return &EmailProcessorTaskHandler{
        BaseTaskHandler: NewBaseTaskHandlerWithMetadata(...),
        emailRepo:       emailRepo,
        eventBus:        eventBus,  // Add this
    }
}

// In processAndStoreEmail method
func (h *EmailProcessorTaskHandler) processAndStoreEmail(
    ctx context.Context,
    msg *imap.EmailMessage,
    accountID string,
    mailbox string,
    tenantID uint32,
    taskCtx *task.TaskContext,
    stats *ProcessingStats,
) error {
    stats.Processed++

    // Convert and store email
    email := h.convertToEmail(msg, accountID, mailbox, tenantID)
    if err := h.emailRepo.IndexEmail(ctx, email); err != nil {
        taskCtx.Logger.Errorf("Failed to store email UID=%d: %v", msg.UID, err)
        stats.Failed++

        // Publish email failed event
        h.publishEmailFailedEvent(ctx, email.ID, accountID, mailbox, err)

        return fmt.Errorf("failed to store email: %w", err)
    }

    stats.Stored++

    // Publish email received event
    h.publishEmailReceivedEvent(ctx, email)

    return nil
}

// Helper methods for publishing events
func (h *EmailProcessorTaskHandler) publishEmailReceivedEvent(ctx context.Context, email *email.Email) {
    if h.eventBus == nil {
        return
    }

    event := eventbus.NewEvent(eventbus.EventEmailReceived, eventbus.EmailReceivedEvent{
        EmailID:   email.ID,
        From:      h.getFirstAddress(email.From),
        To:        h.getFirstAddress(email.To),
        Subject:   email.Subject,
        Mailbox:   email.Mailbox,
        AccountID: email.AccountID,
        TenantID:  h.getTenantID(email.TenantID),
    }).WithSource("email-processor")

    // Publish asynchronously to not block email processing
    h.eventBus.PublishAsync(ctx, event)
}

func (h *EmailProcessorTaskHandler) publishEmailFailedEvent(ctx context.Context, emailID, accountID, mailbox string, err error) {
    if h.eventBus == nil {
        return
    }

    event := eventbus.NewEvent(eventbus.EventEmailFailed, map[string]interface{}{
        "email_id":   emailID,
        "account_id": accountID,
        "mailbox":    mailbox,
        "error":      err.Error(),
    }).WithSource("email-processor").WithPriority(5)

    h.eventBus.PublishAsync(ctx, event)
}

func (h *EmailProcessorTaskHandler) getFirstAddress(addrs []email.EmailAddress) string {
    if len(addrs) > 0 {
        return addrs[0].Address
    }
    return ""
}

func (h *EmailProcessorTaskHandler) getTenantID(tenantID *uint32) uint32 {
    if tenantID != nil {
        return *tenantID
    }
    return 0
}
```

### 3. Subscribe to Email Events

Create email event handlers:

```go
// In a new file: email_event_handlers.go
package handlers

import (
    "context"
    "go-wind-admin/pkg/eventbus"
    "github.com/go-kratos/kratos/v2/log"
)

// RegisterEmailEventHandlers registers all email-related event handlers
func RegisterEmailEventHandlers(bus eventbus.EventBus, logger *log.Helper) {
    // Log all received emails
    bus.Subscribe(eventbus.EventEmailReceived, eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
        var data eventbus.EmailReceivedEvent
        if err := event.GetData(&data); err != nil {
            return err
        }

        logger.Infof("📧 Email received: From=%s, Subject=%s, Mailbox=%s",
            data.From, data.Subject, data.Mailbox)
        return nil
    }))

    // Handle email failures
    bus.Subscribe(eventbus.EventEmailFailed, eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
        logger.Errorf("❌ Email processing failed: %v", event.Data)
        // Could send alerts, update metrics, etc.
        return nil
    }))

    // Trigger notifications for new emails
    bus.Subscribe(eventbus.EventEmailReceived, eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
        var data eventbus.EmailReceivedEvent
        if err := event.GetData(&data); err != nil {
            return err
        }

        // Publish notification event
        notificationEvent := eventbus.NewEvent("notification.send", map[string]interface{}{
            "type":    "email_received",
            "message": "New email from " + data.From,
            "user_id": data.TenantID,
        })

        return bus.Publish(ctx, notificationEvent)
    }))

    // Update email statistics
    bus.SubscribeAsync(eventbus.EventEmailReceived, eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
        // Update metrics asynchronously
        logger.Debugf("Updating email statistics...")
        return nil
    }))
}
```

### 4. Wire Everything Together

In your service initialization:

```go
// In task_service.go or similar
func NewTaskService(
    logger log.Logger,
    taskRepo *data.TaskRepo,
    userRepo *data.UserRepo,
    taskAuditRepo *data.TaskAuditLogRepo,
    emailRepo *data.EmailRepo,
    eventBusManager *eventbus.Manager,  // Add this
) *TaskService {
    l := log.NewHelper(log.With(logger, "module", "task/service"))

    service := &TaskService{
        log:           l,
        taskRepo:      taskRepo,
        userRepo:      userRepo,
        taskAuditRepo: taskAuditRepo,
        emailRepo:     emailRepo,
        eventBusManager: eventBusManager,  // Add this
    }

    return service
}

// In registerHandlersWithAudit
func (s *TaskService) registerHandlersWithAudit() error {
    registry := s.taskManager.Registry()

    // Get email event bus
    emailEventBus := s.eventBusManager.GetBus("email")

    // Register email event handlers
    handlers.RegisterEmailEventHandlers(emailEventBus, s.log)

    // Create audit logger
    auditLogger := &SimpleAuditLogger{
        taskAuditRepo: s.taskAuditRepo,
        logger:        s.log,
    }

    // Register task handlers with event bus
    taskHandlers := []task.TaskHandler{
        handlers.NewEmailProcessorTaskHandler(s.emailRepo, emailEventBus),  // Pass event bus
        handlers.NewBackupTaskHandler(),
        handlers.NewNotificationTaskHandler(),
    }

    // ... rest of registration code
}
```

### 5. Add EventBus to Wire Dependencies

In `data/init.go`:

```go
var ProviderSet = wire.NewSet(
    NewData,
    // ... existing providers ...
    NewEventBusManager,  // Add this
)

// NewEventBusManager creates a new event bus manager
func NewEventBusManager(logger log.Logger) *eventbus.Manager {
    return eventbus.NewManager(logger)
}
```

## Event Flow Example

```
IMAP Client → Email Processor → EventBus
                                    ↓
                    ┌───────────────┼───────────────┐
                    ↓               ↓               ↓
            Logger Handler   Notification    Statistics
                              Handler         Handler
```

## Benefits

1. **Loose Coupling**: Email processor doesn't need to know about logging, notifications, etc.
2. **Scalability**: Easy to add new handlers without modifying email processor
3. **Testability**: Can test each handler independently
4. **Flexibility**: Enable/disable handlers at runtime
5. **Async Processing**: Non-critical tasks (logging, stats) run asynchronously

## Testing

```go
func TestEmailProcessor_WithEventBus(t *testing.T) {
    // Create mock event bus
    bus := eventbus.NewEventBus(log.DefaultLogger)
    defer bus.Close()

    // Track published events
    var receivedEvents []*eventbus.Event
    bus.Subscribe(eventbus.EventEmailReceived, eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
        receivedEvents = append(receivedEvents, event)
        return nil
    }))

    // Create email processor with event bus
    processor := handlers.NewEmailProcessorTaskHandler(mockEmailRepo, bus)

    // Process email
    // ... test code ...

    // Verify event was published
    assert.Equal(t, 1, len(receivedEvents))
    assert.Equal(t, eventbus.EventEmailReceived, receivedEvents[0].Type)
}
```

## Monitoring

Get event bus statistics:

```go
stats := eventBusManager.GetStats()
fmt.Printf("EventBus Stats: %+v\n", stats)

// Output:
// {
//   "total_buses": 2,
//   "buses": {
//     "email": {
//       "event_types": ["email.received", "email.failed", "email.processed"]
//     }
//   },
//   "global_bus": {
//     "event_types": ["system.started", "user.created"]
//   }
// }
```
