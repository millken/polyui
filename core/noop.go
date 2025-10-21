package core

import "fmt"

type NoOpRenderer struct{}

func (n *NoOpRenderer) Init() error { fmt.Println("NoOpRenderer: Init"); return nil }
func (n *NoOpRenderer) BeginFrame(width, height, pixelRatio float32) { fmt.Printf("BeginFrame %vx%v @%v\n", width, height, pixelRatio) }
func (n *NoOpRenderer) EndFrame() { fmt.Println("EndFrame") }
func (n *NoOpRenderer) Rect(x, y, w, h float32, color Color) { fmt.Printf("Rect x=%v y=%v w=%v h=%v color=%+v\n", x,y,w,h,color) }
func (n *NoOpRenderer) Text(x, y float32, text string, size float32, color Color) { fmt.Printf("Text '%s' @%v,%v size=%v color=%+v\n", text, x,y,size,color) }
func (n *NoOpRenderer) Image(img Image, x, y, w, h float32) { fmt.Printf("Image %s @%v,%v %vx%v\n", img.Path, x,y,w,h) }
func (n *NoOpRenderer) Push() {}
func (n *NoOpRenderer) Pop() {}
func (n *NoOpRenderer) PushScissor(x, y, w, h float32) {}
func (n *NoOpRenderer) PopScissor() {}
func (n *NoOpRenderer) DeleteImage(img Image) {}
