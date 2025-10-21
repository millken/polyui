// Package core provides minimal component and rendering primitives used by the
// Go PolyUI prototype. It intentionally keeps APIs small for easy testing and
// iterative development.
package component

import (
	"github.com/millken/polyui/color"
	"github.com/millken/polyui/event"
	"github.com/millken/polyui/host"
	"github.com/millken/polyui/renderer"
	"github.com/millken/polyui/unit"
)

// Component is a minimal drawable / inputtable interface.
type Component interface {
	Setup(ui host.UI)
	PreRender(r renderer.Renderer)
	Render(r renderer.Renderer)
	PostRender(r renderer.Renderer)
	HandleEvent(e event.Event) bool
}

// BaseComponent provides a common implementation for basic component state.
type BaseComponent struct {
	Name     string
	Parent   Component
	Children []Component

	X, Y   float32
	Width  float32
	Height float32

	// Visible size and alignment helpers
	Visible unit.Vec2
	Align   unit.Align
}

func NewBaseComponent(at unit.Vec2, size unit.Vec2, alignment unit.Align) *BaseComponent {
	return &BaseComponent{
		Name:    "Component",
		X:       at.X,
		Y:       at.Y,
		Width:   size.X,
		Height:  size.Y,
		Visible: size,
		Align:   alignment,
	}
}

func (c *BaseComponent) Setup(ui host.UI) {
	for _, ch := range c.Children {
		ch.Setup(ui)
	}
}

func (c *BaseComponent) PreRender(r renderer.Renderer) {}

func (c *BaseComponent) Render(r renderer.Renderer) {
	// default: render children
	for _, ch := range c.Children {
		ch.PreRender(r)
		ch.Render(r)
		ch.PostRender(r)
	}
}

func (c *BaseComponent) PostRender(r renderer.Renderer) {}

func (c *BaseComponent) HandleEvent(e event.Event) bool { return false }

// AddChild appends a child and sets its parent.
func (c *BaseComponent) AddChild(child Component) {
	c.Children = append(c.Children, child)
	// set parent if child is a BaseComponent derivative
	if bc, ok := child.(*BaseComponent); ok {
		bc.Parent = c
	}
}

// RemoveChild removes a child by index.
func (c *BaseComponent) RemoveChild(index int) {
	if index < 0 || index >= len(c.Children) {
		return
	}
	c.Children = append(c.Children[:index], c.Children[index+1:]...)
}

// Group is a simple container component (backwards compatible).
type Group struct{ BaseComponent }

func NewGroup() *Group { return &Group{*NewBaseComponent(unit.Vec2{}, unit.Vec2{}, unit.Align{})} }

// Block is a simple drawable box
type Block struct {
	BaseComponent
	Color color.Color
}

func NewBlock(x, y, w, h float32, c color.Color) *Block {
	b := &Block{BaseComponent: *NewBaseComponent(unit.Vec2{X: x, Y: y}, unit.Vec2{X: w, Y: h}, unit.AlignDefault), Color: c}
	return b
}

func (b *Block) Render(r renderer.Renderer) { r.Rect(b.X, b.Y, b.Width, b.Height, b.Color) }

// Drawable provides higher-level rendering features and state.
type Drawable struct {
	BaseComponent
	NeedsRedraw    bool
	Rotation       float64
	SkewX, SkewY   float64
	ScaleX, ScaleY float32
	Alpha          float32
	Framebuffer    interface{}
}

func NewDrawable(at unit.Vec2, size unit.Vec2, alignment unit.Align) *Drawable {
	d := &Drawable{BaseComponent: *NewBaseComponent(at, size, alignment)}
	d.ScaleX = 1
	d.ScaleY = 1
	d.Alpha = 1
	d.NeedsRedraw = true
	return d
}

// PreRender applies transforms; basic implementation to satisfy interface.
func (d *Drawable) PreRender(r renderer.Renderer) {
	if !d.NeedsRedraw {
		return
	}
	r.Push()
}

func (d *Drawable) PostRender(r renderer.Renderer) {
	r.Pop()
}

// Inputtable provides event handling capabilities.
type Inputtable struct {
	BaseComponent
	AcceptsInput bool
	Focusable    bool
	Focused      bool
	Handlers     map[string][]func(event.Event) bool
}

func NewInputtable(at unit.Vec2, size unit.Vec2, alignment unit.Align, focusable bool) *Inputtable {
	return &Inputtable{BaseComponent: *NewBaseComponent(at, size, alignment), Focusable: focusable, Handlers: make(map[string][]func(event.Event) bool)}
}

func (it *Inputtable) HandleEvent(e event.Event) bool {
	if !it.AcceptsInput {
		return false
	}
	t := e.Type()
	if hs, ok := it.Handlers[t]; ok {
		for _, h := range hs {
			if h(e) {
				return true
			}
		}
	}
	return false
}

// Scrollable adds simple scroll state used by Drawable in Kotlin.
type Scrollable struct {
	Inputtable
	VisWidth, VisHeight float32
	XScroll, YScroll    interface{}
	ScreenX, ScreenY    float32
	ShouldScroll        bool
}

func NewScrollable(at unit.Vec2, size unit.Vec2, visible unit.Vec2, alignment unit.Align, focusable bool) *Scrollable {
	s := &Scrollable{Inputtable: *NewInputtable(at, size, alignment, focusable)}
	s.VisWidth = visible.X
	s.VisHeight = visible.Y
	s.ShouldScroll = true
	s.ScreenX = at.X
	s.ScreenY = at.Y
	return s
}
