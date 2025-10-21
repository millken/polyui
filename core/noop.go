package core

type NoOpRenderer struct{}

func (n *NoOpRenderer) Init() error { Debugln("NoOpRenderer: Init"); return nil }
func (n *NoOpRenderer) BeginFrame(width, height, pixelRatio float32) {
	Debugf("BeginFrame %vx%v @%v\n", width, height, pixelRatio)
}
func (n *NoOpRenderer) EndFrame() { Debugln("EndFrame") }
func (n *NoOpRenderer) Rect(x, y, w, h float32, color Color) {
	Debugf("Rect x=%v y=%v w=%v h=%v color=%+v\n", x, y, w, h, color)
}
func (n *NoOpRenderer) Text(x, y float32, text string, size float32, color Color) {
	Debugf("Text '%s' @%v,%v size=%v color=%+v\n", text, x, y, size, color)
}
func (n *NoOpRenderer) Image(img Image, x, y, w, h float32) {
	Debugf("Image %s @%v,%v %vx%v\n", img.Path, x, y, w, h)
}
func (n *NoOpRenderer) Push()                          {}
func (n *NoOpRenderer) Pop()                           {}
func (n *NoOpRenderer) PushScissor(x, y, w, h float32) {}
func (n *NoOpRenderer) PopScissor()                    {}
func (n *NoOpRenderer) DeleteImage(img Image)          {}
