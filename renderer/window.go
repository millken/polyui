package renderer

import "github.com/millken/polyui/host"

// Window represents a minimal window contract for the prototype.
type Window interface {
	// Open the window and begin the event/render loop. ui is intentionally an empty interface
	// to avoid importing higher-level packages and creating import cycles.
	Open(ui host.UI) error
	Close()
	Width() int
	Height() int
	PixelRatio() float32

	SetClipboard(s string)
	GetClipboard() string
}
