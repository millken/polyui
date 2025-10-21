package input

import (
	"github.com/millken/polyui/event"
)

// Placeholder interfaces to avoid import cycles during incremental migration.
// Replace these with references to the real component package when ready.
type Inputtable interface{}
type Component interface{}

// InputManager is a small, pragmatic port of the Kotlin InputManager used in
// the polyui prototype. It deliberately omits advanced behaviors and focuses
// on the APIs the Go implementation uses.
type InputManager struct {
	master    Inputtable
	KeyBinder *KeyBinder
	settings  interface{}

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

func NewInputManager(master Inputtable, kb *KeyBinder, settings interface{}) *InputManager {
	return &InputManager{master: master, KeyBinder: kb, settings: settings}
}

func (m *InputManager) With(master Inputtable) *InputManager  { m.master = master; return m }
func (m *InputManager) FilesDropped(paths []string)           {}
func (m *InputManager) KeyTyped(r rune)                       {}
func (m *InputManager) KeyDown(key Key)                       {}
func (m *InputManager) KeyUp(key Key)                         {}
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

func (m *InputManager) Dispatch(e event.Event, bindable bool, to Inputtable) bool { return false }
func (m *InputManager) DropAll()                                                  {}
