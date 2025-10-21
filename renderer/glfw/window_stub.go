//go:build glfw
// +build glfw
package glfw

import (
	"fmt"

	"github.com/millken/polyui/host"
	"github.com/millken/polyui/renderer"
)

// Ensure GLFWWindow implements renderer.Window (stub fallback)
var _ renderer.Window = (*GLFWWindow)(nil)

// GLFWWindow fallback when build tag 'glfw' is not present.
type GLFWWindow struct{}

func (w *GLFWWindow) Open(ui host.UI) error { return fmt.Errorf("glfw support not enabled") }
func (w *GLFWWindow) Close()                {}
func (w *GLFWWindow) Width() int            { return 0 }
func (w *GLFWWindow) Height() int           { return 0 }
func (w *GLFWWindow) PixelRatio() float32   { return 1.0 }
func (w *GLFWWindow) SetClipboard(s string) {}
func (w *GLFWWindow) GetClipboard() string  { return "" }
