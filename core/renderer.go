package core

// Color is a simple RGBA container
type Color struct {
	R, G, B, A float32
}

// Image is an opaque handle to an image resource
type Image struct {
	Path string
	// Data holds optional raw image bytes (PNG/JPEG). If non-empty, renderer should prefer this
	// over loading from Path.
	Data []byte
}

// Renderer defines the small rendering contract used by the prototype.
type Renderer interface {
	Init() error
	BeginFrame(width, height, pixelRatio float32)
	EndFrame()

	Rect(x, y, w, h float32, color Color)
	Text(x, y float32, text string, size float32, color Color)
	Image(img Image, x, y, w, h float32)

	Push()
	Pop()
	PushScissor(x, y, w, h float32)
	PopScissor()

	// Delete resources if any
	DeleteImage(img Image)
}
