package event

// Lightweight port of the Kotlin event types used by the Go prototype.
// The file intentionally keeps dependencies to a minimum to avoid import
// cycles during incremental migration.

// Event is the top-level marker interface for input/events.
type Event interface {
	Type() string
	FocusRequired() bool
}

// coreModifiers is a local alias matching the eventual core.Modifiers byte
// type. This avoids importing the core package during the refactor; replace
// usages with core.Modifiers when ready.
type coreModifiers byte

// Mouse events. Use concrete types rather than embedding to make Go method
// receivers straightforward.
type MousePressed struct {
	Button int
	X, Y   float32
	Mods   coreModifiers
}

type MouseReleased struct {
	Button int
	X, Y   float32
	Mods   coreModifiers
}

type MouseMoved struct{ X, Y float32 }

type MouseDragged struct{ X, Y float32 }

type MouseClicked struct {
	Button int
	X, Y   float32
	Amount int
	Mods   coreModifiers
}

type MouseScrolled struct {
	DX, DY float32
	Mods   coreModifiers
}

func (MousePressed) Type() string         { return "MousePressed" }
func (MousePressed) FocusRequired() bool  { return false }
func (MouseReleased) Type() string        { return "MouseReleased" }
func (MouseReleased) FocusRequired() bool { return false }
func (MouseMoved) Type() string           { return "MouseMoved" }
func (MouseMoved) FocusRequired() bool    { return false }
func (MouseDragged) Type() string         { return "MouseDragged" }
func (MouseDragged) FocusRequired() bool  { return false }
func (MouseClicked) Type() string         { return "MouseClicked" }
func (MouseClicked) FocusRequired() bool  { return false }
func (MouseScrolled) Type() string        { return "MouseScrolled" }
func (MouseScrolled) FocusRequired() bool { return false }

// Key/focus events (simplified)
type KeyPressed struct {
	Key  int
	Mods coreModifiers
}

type KeyReleased struct {
	Key  int
	Mods coreModifiers
}

type KeyTyped struct{ Rune rune }

type UnmappedInput struct {
	KeyCode, ScanCode int
	Down              bool
	Mods              coreModifiers
}

type FileDrop struct{ Paths []string }

func (KeyPressed) Type() string           { return "KeyPressed" }
func (KeyPressed) FocusRequired() bool    { return true }
func (KeyReleased) Type() string          { return "KeyReleased" }
func (KeyReleased) FocusRequired() bool   { return true }
func (KeyTyped) Type() string             { return "KeyTyped" }
func (KeyTyped) FocusRequired() bool      { return true }
func (UnmappedInput) Type() string        { return "UnmappedInput" }
func (UnmappedInput) FocusRequired() bool { return true }
func (FileDrop) Type() string             { return "FileDrop" }
func (FileDrop) FocusRequired() bool      { return true }
