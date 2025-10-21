package core

// Window represents a minimal window contract for the prototype.
type Window interface {
	Open(ui *PolyUI) error
	Close()
	Width() int
	Height() int
	PixelRatio() float32

	SetClipboard(s string)
	GetClipboard() string
}
