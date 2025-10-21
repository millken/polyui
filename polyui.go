// Package polyui contains the top-level PolyUI application type moved out of the core package.
package polyui

import (
	"github.com/millken/polyui/component"
	"github.com/millken/polyui/renderer"
)

// PolyUI holds the root component and settings for the prototype.
type PolyUI struct {
	// Root component is commented during initial migration to avoid circular dependencies.
	Root     component.Component
	Renderer renderer.Renderer
	Window   renderer.Window
}

// NewPolyUI creates a new PolyUI instance with the supplied renderer.
func NewPolyUI(root component.Component, r renderer.Renderer) *PolyUI {
	return &PolyUI{Root: root, Renderer: r}
}

// Init initializes the renderer and any root component setup. Root handling is currently
// commented out while Component is migrated.
func (p *PolyUI) Init() error {
	if err := p.Renderer.Init(); err != nil {
		return err
	}
	if p.Root != nil {
		p.Root.Setup(p)
	}
	return nil
}

// RenderFrame runs a single frame render using the configured renderer.
func (p *PolyUI) RenderFrame(width, height, pixelRatio float32) {
	p.Renderer.BeginFrame(width, height, pixelRatio)
	if p.Root != nil {
		p.Root.PreRender(p.Renderer)
		p.Root.Render(p.Renderer)
		p.Root.PostRender(p.Renderer)
	}
	p.Renderer.EndFrame()
}
