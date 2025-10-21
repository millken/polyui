//go:build glfw
// +build glfw

package glfw

import (
	"fmt"
	"runtime"
	"time"

	gl "github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/millken/polyui/core"
)

// GLFWWindow is a basic GLFW window that runs a render loop and calls into PolyUI.
type GLFWWindow struct {
	win *glfw.Window
}

// NewGLFWWindow must be called from the main OS thread (LockOSThread).
func NewGLFWWindow(title string, width, height int) (*GLFWWindow, error) {
	// Ensure caller locked the OS thread
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		return nil, fmt.Errorf("glfw init failed: %w", err)
	}

	// Request OpenGL 3.3 core profile by default (can be adjusted)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, fmt.Errorf("create window failed: %w", err)
	}
	return &GLFWWindow{win: win}, nil
}

// MakeContextCurrent makes the underlying GLFW context current on the calling thread.
func (w *GLFWWindow) MakeContextCurrent() {
	w.win.MakeContextCurrent()
}

// InitGL initializes the OpenGL function pointers. Call this after MakeContextCurrent.
func (w *GLFWWindow) InitGL() error {
	if err := gl.Init(); err != nil {
		return fmt.Errorf("gl init failed: %v", err)
	}
	return nil
}

// Open runs the window loop and blocks until the window is closed.
// Callers should invoke this from the main goroutine (LockOSThread).
func (w *GLFWWindow) Open(ui *core.PolyUI) error {
	// Make context current on this thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	w.win.MakeContextCurrent()

	// Simple loop: call ui.RenderFrame each frame
	for !w.win.ShouldClose() {
		width, height := w.win.GetSize()
		ui.RenderFrame(float32(width), float32(height), 1.0)
		glfw.PollEvents()
		w.win.SwapBuffers()
		time.Sleep(time.Millisecond * 16) // ~60fps cap
	}
	return nil
}

func (w *GLFWWindow) Close() { w.win.Destroy(); glfw.Terminate() }
func (w *GLFWWindow) Width() int {
	width, _ := w.win.GetSize()
	return width
}
func (w *GLFWWindow) Height() int {
	_, height := w.win.GetSize()
	return height
}
func (w *GLFWWindow) PixelRatio() float32   { return 1.0 }
func (w *GLFWWindow) SetClipboard(s string) { w.win.SetClipboardString(s) }
func (w *GLFWWindow) GetClipboard() string  { return w.win.GetClipboardString() }

// PlaceholderRenderer is a small renderer that prints draw calls; replace with NanoVG integration.
type PlaceholderRenderer struct{}

func (p *PlaceholderRenderer) Init() error { fmt.Println("PlaceholderRenderer: Init"); return nil }
func (p *PlaceholderRenderer) BeginFrame(width, height, pixelRatio float32) {
	fmt.Printf("BeginFrame %vx%v\n", width, height)
}
func (p *PlaceholderRenderer) EndFrame() { fmt.Println("EndFrame") }
func (p *PlaceholderRenderer) Rect(x, y, w, h float32, color core.Color) {
	fmt.Printf("Rect %v,%v %vx%v color=%+v\n", x, y, w, h, color)
}
func (p *PlaceholderRenderer) Text(x, y float32, text string, size float32, color core.Color) {
	fmt.Printf("Text '%s'\n", text)
}
func (p *PlaceholderRenderer) Image(img core.Image, x, y, w, h float32) {
	fmt.Printf("Image %s\n", img.Path)
}
func (p *PlaceholderRenderer) Push()                          {}
func (p *PlaceholderRenderer) Pop()                           {}
func (p *PlaceholderRenderer) PushScissor(x, y, w, h float32) {}
func (p *PlaceholderRenderer) PopScissor()                    {}
func (p *PlaceholderRenderer) DeleteImage(img core.Image)     {}
