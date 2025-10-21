//go:build glfw
// +build glfw

package glfw

import (
	"fmt"
	_ "image/gif"

	// NVG binding removed due to unavailability; this file is a placeholder.

	"github.com/millken/polyui/core"
)

// NVGRenderer implements core.Renderer using NanoVG via fyne-io/nanovg binding.
// NVGRenderer was removed because no maintained NanoVG binding is available.
// Keep this file as a historical placeholder; users should use the OpenGL
// renderer in ../gl or implement their own renderer.

type NVGRenderer struct{}

func (n *NVGRenderer) Init() error                                                    { return fmt.Errorf("NanoVG renderer not available") }
func (n *NVGRenderer) BeginFrame(width, height, pixelRatio float32)                   {}
func (n *NVGRenderer) EndFrame()                                                      {}
func (n *NVGRenderer) Rect(x, y, w, h float32, color core.Color)                      {}
func (n *NVGRenderer) Text(x, y float32, text string, size float32, color core.Color) {}
func (n *NVGRenderer) Image(img core.Image, x, y, w, h float32)                       {}
func (n *NVGRenderer) Push()                                                          {}
func (n *NVGRenderer) Pop()                                                           {}
func (n *NVGRenderer) PushScissor(x, y, w, h float32)                                 {}
func (n *NVGRenderer) PopScissor()                                                    {}
func (n *NVGRenderer) DeleteImage(img core.Image)                                     {}
