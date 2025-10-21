package main

import (
	"github.com/millken/polyui/core"
)

func main() {
	r := &core.NoOpRenderer{}
	block := &core.Block{X: 10, Y: 20, W: 200, H: 100, Color: core.Color{R: 0.2, G: 0.6, B: 0.9, A: 1}}
	ui := core.NewPolyUI(block, r)
	if err := ui.Init(); err != nil {
		panic(err)
	}
	ui.RenderFrame(800, 600, 1)
}
