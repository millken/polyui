//go:build glfw
// +build glfw

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"

	"github.com/millken/polyui/core"
	glimpl "github.com/millken/polyui/impl/gl"
	glfwimpl "github.com/millken/polyui/impl/glfw"
)

func main() {
	// Lock OS thread for GLFW
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	w, err := glfwimpl.NewGLFWWindow("PolyUI Go GLFW", 800, 600)
	if err != nil {
		panic(err)
	}

	// Make GL context current and initialize GL bindings before creating renderer
	w.MakeContextCurrent()
	if err := w.InitGL(); err != nil {
		panic(err)
	}

	// create a temporary PNG to test image loading (no external asset required)
	tmpFile, err := os.CreateTemp("", "polyui_go_img_*.png")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	img := image.NewRGBA(image.Rect(0, 0, 128, 128))
	// fill with a simple gradient/color
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			img.Set(x, y, color.RGBA{uint8(x * 2), uint8(y * 2), 128, 255})
		}
	}
	if err := png.Encode(tmpFile, img); err != nil {
		panic(err)
	}

	imageRes := core.Image{Path: tmpFile.Name()}

	block := &core.Block{X: 10, Y: 20, W: 200, H: 100, Color: core.Color{R: 0.2, G: 0.6, B: 0.9, A: 1}}

	// small Image component instance
	imgComp := &ImageComp{Img: imageRes, X: 240, Y: 20, W: 128, H: 128}

	// Initialize and use the real OpenGL renderer implementation
	glr := &glimpl.SimpleGLRenderer{}
	// our GLFW window already has a current context and InitGL was called above
	if err := glr.Init(); err != nil {
		panic(err)
	}
	ui := core.NewPolyUI(&core.Group{Children: []core.Component{block, imgComp}}, glr)
	if err := ui.Init(); err != nil {
		panic(err)
	}

	if err := w.Open(ui); err != nil {
		panic(err)
	}
}

// small Image component implemented at package level
type ImageComp struct {
	Img        core.Image
	X, Y, W, H float32
}

func (c *ImageComp) Setup(ui *core.PolyUI)         {}
func (c *ImageComp) PreRender(r core.Renderer)     {}
func (c *ImageComp) Render(r core.Renderer)        { r.Image(c.Img, c.X, c.Y, c.W, c.H) }
func (c *ImageComp) PostRender(r core.Renderer)    {}
func (c *ImageComp) HandleEvent(e core.Event) bool { return false }
