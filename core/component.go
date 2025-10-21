// Package core provides minimal component and rendering primitives used by the
// Go PolyUI prototype. It intentionally keeps APIs small for easy testing and
// iterative development.
package core

// Component is a minimal drawable / inputtable interface.
type Component interface {
	Setup(ui *PolyUI)
	PreRender(r Renderer)
	Render(r Renderer)
	PostRender(r Renderer)
	HandleEvent(e Event) bool
}

// Group is a simple container component.
type Group struct {
	Children []Component
}

func (g *Group) Setup(ui *PolyUI) {
	for _, c := range g.Children {
		c.Setup(ui)
	}
}

func (g *Group) PreRender(r Renderer) {}
func (g *Group) Render(r Renderer) {
	for _, c := range g.Children {
		c.Render(r)
	}
}
func (g *Group) PostRender(r Renderer)    {}
func (g *Group) HandleEvent(e Event) bool { return false }

// Block is a simple drawable box
type Block struct {
	X, Y, W, H float32
	Color      Color
}

func (b *Block) Setup(ui *PolyUI)         {}
func (b *Block) PreRender(r Renderer)     {}
func (b *Block) Render(r Renderer)        { r.Rect(b.X, b.Y, b.W, b.H, b.Color) }
func (b *Block) PostRender(r Renderer)    {}
func (b *Block) HandleEvent(e Event) bool { return false }
