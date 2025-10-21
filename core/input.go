package core

// Package core contains the minimal core types used by the Go PolyUI prototype,
// including events, input management, and basic components.

// Settings is a lightweight placeholder for input behavior. Can be expanded later.
type Settings struct {
	DragThreshold         float32
	ComboMaxIntervalNanos int64
	MaxComboSize          int
	ClearComboWhenMaxed   bool
	NaturalScrolling      bool
	ScrollMultiplier      [2]float32
}

// EventKind represents the kind of an input event.
type EventKind string

const (
	EventKindMousePressed  EventKind = "MousePressed"
	EventKindMouseReleased EventKind = "MouseReleased"
	EventKindMouseMoved    EventKind = "MouseMoved"
	EventKindMouseDragged  EventKind = "MouseDragged"
	EventKindMouseClicked  EventKind = "MouseClicked"
	EventKindMouseScrolled EventKind = "MouseScrolled"

	EventKindKeyPressed  EventKind = "KeyPressed"
	EventKindKeyReleased EventKind = "KeyReleased"
	EventKindKeyTyped    EventKind = "KeyTyped"
	EventKindUnmapped    EventKind = "UnmappedInput"
	EventKindFileDrop    EventKind = "FileDrop"
)

// Event represents an input event. Concrete types implement this.
// Use Type() to identify the event kind and FocusRequired() to indicate
// whether the event should be delivered only to a focused target.
type Event interface {
	Type() EventKind
	FocusRequired() bool
}

// MousePressed represents a mouse button press.
type MousePressed struct {
	Button int
	X, Y   float32
	Mods   Modifiers
}
type MouseReleased struct {
	Button int
	X, Y   float32
	Mods   Modifiers
}
type MouseMoved struct{ X, Y float32 }
type MouseDragged struct{ X, Y float32 }
type MouseClicked struct {
	Button int
	X, Y   float32
	Amount int
	Mods   Modifiers
}
type MouseScrolled struct {
	DX, DY float32
	Mods   Modifiers
}

func (MousePressed) Type() EventKind      { return EventKindMousePressed }
func (MousePressed) FocusRequired() bool  { return false }
func (MouseReleased) Type() EventKind     { return EventKindMouseReleased }
func (MouseReleased) FocusRequired() bool { return false }
func (MouseMoved) Type() EventKind        { return EventKindMouseMoved }
func (MouseMoved) FocusRequired() bool    { return false }
func (MouseDragged) Type() EventKind      { return EventKindMouseDragged }
func (MouseDragged) FocusRequired() bool  { return false }
func (MouseClicked) Type() EventKind      { return EventKindMouseClicked }
func (MouseClicked) FocusRequired() bool  { return false }
func (MouseScrolled) Type() EventKind     { return EventKindMouseScrolled }
func (MouseScrolled) FocusRequired() bool { return false }

// KeyPressed represents a key press event (focus-sensitive).
type KeyPressed struct {
	Key  int
	Mods Modifiers
}
type KeyReleased struct {
	Key  int
	Mods Modifiers
}
type KeyTyped struct{ Rune rune }
type UnmappedInput struct {
	KeyCode, ScanCode int
	Down              bool
	Mods              Modifiers
}
type FileDrop struct{ Paths []string }

func (KeyPressed) Type() EventKind        { return EventKindKeyPressed }
func (KeyPressed) FocusRequired() bool    { return true }
func (KeyReleased) Type() EventKind       { return EventKindKeyReleased }
func (KeyReleased) FocusRequired() bool   { return true }
func (KeyTyped) Type() EventKind          { return EventKindKeyTyped }
func (KeyTyped) FocusRequired() bool      { return true }
func (UnmappedInput) Type() EventKind     { return EventKindUnmapped }
func (UnmappedInput) FocusRequired() bool { return true }
func (FileDrop) Type() EventKind          { return EventKindFileDrop }
func (FileDrop) FocusRequired() bool      { return true }

// Modifiers is a bitmask for key modifiers
type Modifiers byte

const (
	ModShift Modifiers = 1 << iota
	ModCtrl
	ModAlt
	ModMeta
)

func (m Modifiers) HasShift() bool { return m&ModShift != 0 }
func (m Modifiers) IsEmpty() bool  { return m == 0 }

// InputManager handles input state and event dispatching. Methods mirror the Kotlin API.
type InputManager struct {
	master    Inputtable
	keyBinder *KeyBinder
	settings  *Settings

	// runtime state (exported only for tests/debug if needed)
	MouseOver      Inputtable
	MouseX, MouseY float32
	PressX, PressY float32
	Dragging       bool
	MouseDown      bool

	ClickTimer    int64
	ClickAmount   int
	ClickedButton int

	Focused Inputtable
	Mods    byte
}

func NewInputManager(master Inputtable, kb *KeyBinder, s *Settings) *InputManager {
	return &InputManager{master: master, keyBinder: kb, settings: s}
}

// With sets the master Inputtable for this manager and returns the manager.
func (m *InputManager) With(master Inputtable) *InputManager  { m.master = master; return m }
func (m *InputManager) FilesDropped(paths []string)           {}
func (m *InputManager) KeyTyped(r rune)                       {}
func (m *InputManager) KeyDown(key int)                       {}
func (m *InputManager) KeyUp(key int)                         {}
func (m *InputManager) KeyDownUnmapped(keyCode, scanCode int) {}
func (m *InputManager) KeyUpUnmapped(keyCode, scanCode int)   {}

func (m *InputManager) Recalculate()           {}
func (m *InputManager) Drop(target Inputtable) {}

func (m *InputManager) AddModifier(mod byte)    {}
func (m *InputManager) RemoveModifier(mod byte) {}
func (m *InputManager) ClearModifiers()         {}

func (m *InputManager) RayCheck(it Inputtable, x, y float32) Inputtable    { return nil }
func (m *InputManager) RayCheckUnsafe(c Component, x, y float32) Component { return nil }

func (m *InputManager) MouseMoved(x, y float32)      {}
func (m *InputManager) MousePressed(button int)      {}
func (m *InputManager) MouseReleased(button int)     {}
func (m *InputManager) MouseScrolled(dx, dy float32) {}

func (m *InputManager) SafeFocus(c Inputtable) bool { return false }
func (m *InputManager) Focus(c Inputtable) bool     { return false }
func (m *InputManager) Unfocus()                    {}
func (m *InputManager) UnfocusTarget(c Inputtable)  {}

func (m *InputManager) Dispatch(event Event, bindable bool, to Inputtable) bool { return false }
func (m *InputManager) DropAll()                                                {}
