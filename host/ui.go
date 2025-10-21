package host

type UI interface {
	Init() error
	RenderFrame(width, height, pixelRatio float32)
}
