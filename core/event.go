// Package core contains the minimal core types used by the Go PolyUI prototype,
// including events and small helpers used by the input system.
package core

// EventType enumerates simple prototype event categories.
type EventType int

const (
	EventMouseClick EventType = iota
	EventKeyPress
)

// SimpleEvent is a lightweight event struct used for prototypes and tests.
// It's intentionally named SimpleEvent to avoid colliding with the core Event
// interface defined in input.go.
type SimpleEvent struct {
	Type EventType
	X, Y int
	Key  int
}
