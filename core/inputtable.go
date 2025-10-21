package core

// Inputtable extends existing Component with input-related capabilities.
type Inputtable interface {
	Component
	IsEnabled() bool
	IsInside(x, y float32) bool
	AcceptsInput() bool
	Accept(event Event) bool

	Focusable() bool
	Focused() bool
	SetFocused(bool)

	Parent() Component
	Children() []Component

	InputState() int
	SetInputState(int)
}
