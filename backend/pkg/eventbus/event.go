package eventbus

import (
	"encoding/json"
	"time"
)

// Event represents a generic event that can carry any type of data
type Event struct {
	// ID is a unique identifier for the event
	ID string `json:"id"`

	// Type identifies the kind of event (e.g., "email.received", "user.created")
	Type string `json:"type"`

	// Source identifies where the event originated from
	Source string `json:"source"`

	// Data contains the event payload (can be any type)
	Data any `json:"data"`

	// Metadata contains additional information about the event
	Metadata map[string]string `json:"metadata,omitempty"`

	// Timestamp when the event was created
	Timestamp time.Time `json:"timestamp"`

	// Priority of the event (higher = more important)
	Priority int `json:"priority"`
}

// NewEvent creates a new event with the given type and data
func NewEvent(eventType string, data any) *Event {
	return &Event{
		ID:        generateEventID(),
		Type:      eventType,
		Data:      data,
		Timestamp: time.Now(),
		Priority:  0,
		Metadata:  make(map[string]string),
	}
}

// WithSource sets the source of the event
func (e *Event) WithSource(source string) *Event {
	e.Source = source
	return e
}

// WithPriority sets the priority of the event
func (e *Event) WithPriority(priority int) *Event {
	e.Priority = priority
	return e
}

// WithMetadata adds metadata to the event
func (e *Event) WithMetadata(key, value string) *Event {
	if e.Metadata == nil {
		e.Metadata = make(map[string]string)
	}
	e.Metadata[key] = value
	return e
}

// GetData unmarshals the event data into the provided type
func (e *Event) GetData(v any) error {
	// If data is already the right type, try direct assignment
	if e.Data == nil {
		return nil
	}

	// Marshal and unmarshal through JSON for type conversion
	jsonData, err := json.Marshal(e.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, v)
}

// Clone creates a deep copy of the event
func (e *Event) Clone() *Event {
	metadata := make(map[string]string)
	for k, v := range e.Metadata {
		metadata[k] = v
	}

	return &Event{
		ID:        e.ID,
		Type:      e.Type,
		Source:    e.Source,
		Data:      e.Data,
		Metadata:  metadata,
		Timestamp: e.Timestamp,
		Priority:  e.Priority,
	}
}

// generateEventID generates a unique event ID
func generateEventID() string {
	return time.Now().Format("20060102150405.000000") + randomString(8)
}

// randomString generates a random string of the given length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
