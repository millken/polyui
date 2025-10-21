// Package renderer contains lightweight renderer and window contracts used by the prototype.
package renderer

import "github.com/millken/polyui/color"

// Package renderer contains lightweight renderer and window contracts used by the prototype.

// Image is an opaque handle to an image resource
type Image struct {
	Path string
	// Data holds optional raw image bytes (PNG/JPEG). If non-empty, renderer should prefer this
	// over loading from Path.
	Data []byte
}

// Font is an opaque handle to a font resource.
type Font struct {
	Path string
	Data []byte
}

// Vec2 represents a 2D vector or size (width, height).
type Vec2 struct {
	X, Y float32
}

// Renderer defines the rendering contract. Method names mirror the Kotlin API where practical,
// but adapted to Go naming and without overloads. Implementations should provide the behaviour
// described in the Kotlin `Renderer` documentation.
type Renderer interface {
	// Initialization / frame lifecycle
	Init() error
	BeginFrame(width, height, pixelRatio float32)
	EndFrame()

	// Global alpha
	GlobalAlpha(alpha float32)
	ResetGlobalAlpha()

	// Transforms (must be used inside Push/Pop)
	Translate(x, y float32)
	Scale(sx, sy, px, py float32)
	Rotate(angleRadians float64, px, py float32)
	SkewX(angleRadians float64, px, py float32)
	SkewY(angleRadians float64, px, py float32)

	// Scissor / clipping
	PushScissor(x, y, width, height float32)
	PushScissorIntersecting(x, y, width, height float32)
	PopScissor()

	// State stack
	Push()
	Pop()

	// Text
	Text(font Font, x, y float32, text string, color color.Color, fontSize float32)
	TextBounds(font Font, text string, fontSize float32) Vec2

	// Images
	InitImage(img Image, size Vec2)
	// Basic image draw
	Image(img Image, x, y, width, height float32)
	// Extended image draw with color mask and corner radii
	ImageEx(img Image, x, y, width, height float32, colorMask int, topLeftRadius, topRightRadius, bottomLeftRadius, bottomRightRadius float32)

	// Rectangles
	Rect(x, y, width, height float32, color color.Color)
	RectR(x, y, width, height float32, color color.Color, topLeftRadius, topRightRadius, bottomLeftRadius, bottomRightRadius float32)
	HollowRect(x, y, width, height float32, color color.Color, lineWidth float32)
	HollowRectR(x, y, width, height float32, color color.Color, lineWidth float32, topLeftRadius, topRightRadius, bottomLeftRadius, bottomRightRadius float32)

	// Line
	Line(x1, y1, x2, y2 float32, color color.Color, width float32)

	// Drop shadow
	DropShadow(x, y, width, height, blur, spread, radius float32)

	// Capability check
	TransformsWithPoint() bool

	// Resource deletion
	DeleteFont(font *Font)
	DeleteImage(img *Image)

	// Cleanup
	Cleanup()
}

type NoOpRenderer struct{}

func (r *NoOpRenderer) Init() error                                         { return nil }
func (r *NoOpRenderer) BeginFrame(width, height, pixelRatio float32)        {}
func (r *NoOpRenderer) EndFrame()                                           {}
func (r *NoOpRenderer) GlobalAlpha(alpha float32)                           {}
func (r *NoOpRenderer) ResetGlobalAlpha()                                   {}
func (r *NoOpRenderer) Translate(x, y float32)                              {}
func (r *NoOpRenderer) Scale(sx, sy, px, py float32)                        {}
func (r *NoOpRenderer) Rotate(angleRadians float64, px, py float32)         {}
func (r *NoOpRenderer) SkewX(angleRadians float64, px, py float32)          {}
func (r *NoOpRenderer) SkewY(angleRadians float64, px, py float32)          {}
func (r *NoOpRenderer) PushScissor(x, y, width, height float32)             {}
func (r *NoOpRenderer) PushScissorIntersecting(x, y, width, height float32) {}
func (r *NoOpRenderer) PopScissor()                                         {}
func (r *NoOpRenderer) Push()                                               {}
func (r *NoOpRenderer) Pop()                                                {}
func (r *NoOpRenderer) Text(font Font, x, y float32, text string, color color.Color, fontSize float32) {
}
func (r *NoOpRenderer) TextBounds(font Font, text string, fontSize float32) Vec2 {
	return Vec2{}
}
func (r *NoOpRenderer) InitImage(img Image, size Vec2) {}
func (r *NoOpRenderer) Image(img Image, x, y, width, height float32) {
}
func (r *NoOpRenderer) ImageEx(img Image, x, y, width, height float32, colorMask int, topLeftRadius, topRightRadius, bottomLeftRadius, bottomRightRadius float32) {
}
func (r *NoOpRenderer) Rect(x, y, width, height float32, color color.Color) {}
func (r *NoOpRenderer) RectR(x, y, width, height float32, color color.Color, topLeftRadius, topRightRadius, bottomLeftRadius, bottomRightRadius float32) {
}
func (r *NoOpRenderer) HollowRect(x, y, width, height float32, color color.Color, lineWidth float32) {
}
func (r *NoOpRenderer) HollowRectR(x, y, width, height float32, color color.Color, lineWidth float32, topLeftRadius, topRightRadius, bottomLeftRadius, bottomRightRadius float32) {
}
func (r *NoOpRenderer) Line(x1, y1, x2, y2 float32, color color.Color, width float32) {}
func (r *NoOpRenderer) DropShadow(x, y, width, height, blur, spread, radius float32) {
}
func (r *NoOpRenderer) TransformsWithPoint() bool { return false }
func (r *NoOpRenderer) DeleteFont(font *Font)     {}
func (r *NoOpRenderer) DeleteImage(img *Image)    {}
func (r *NoOpRenderer) Cleanup()                  {}
