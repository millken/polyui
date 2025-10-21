package glfw

import (
	"fmt"
	"log/slog"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/millken/polyui/host"
	"github.com/millken/polyui/renderer"
)

// Ensure GLFWWindow implements renderer.Window
var _ renderer.Window = (*GLFWWindow)(nil)

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
	slog.Debug("created glfw window", "title", title, "w", width, "h", height)
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
func (w *GLFWWindow) Open(ui host.UI) error {
	// Make context current on this thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	w.win.MakeContextCurrent()

	// Debug: print initial ShouldClose
	slog.Debug("initial", "shouldclose", w.win.ShouldClose())
	// Ensure window is not flagged closed before entering loop
	w.win.SetShouldClose(false)
	// install a close callback to debug external close requests
	w.win.SetCloseCallback(func(_ *glfw.Window) {
		slog.Debug("glfw close callback triggered")
	})
	slog.Debug("after SetShouldClose(false)", "shouldclose", w.win.ShouldClose())
	slog.Debug("entering main loop")
	slog.Debug("entering Open loop")
	// Simple loop: call ui.RenderFrame each frame
	frame := 0
	for !w.win.ShouldClose() {
		if frame == 0 {
			slog.Debug("first loop iteration")
		}
		width, height := w.win.GetSize()
		ui.RenderFrame(float32(width), float32(height), 1.0)
		glfw.PollEvents()
		w.win.SwapBuffers()
		frame++
		// debug: print should-close state after polling events
		slog.Debug("loop", "frame", frame, "shouldclose", w.win.ShouldClose())
		time.Sleep(time.Millisecond * 16) // ~60fps cap
	}
	slog.Debug("exiting Open loop")
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
