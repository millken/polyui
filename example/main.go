package main

import (
	"runtime"

	polyapp "github.com/millken/polyui"
	pcolor "github.com/millken/polyui/color"
	"github.com/millken/polyui/component"
	glimpl "github.com/millken/polyui/renderer/gl"
	"github.com/millken/polyui/renderer/glfw"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	w, err := glfw.NewGLFWWindow("PolyUI Go Example", 640, 480)
	if err != nil {
		panic(err)
	}
	w.MakeContextCurrent()
	if err := w.InitGL(); err != nil {
		panic(err)
	}

	glr := &glimpl.SimpleGLRenderer{}
	if err := glr.Init(); err != nil {
		panic(err)
	}

	blk := component.NewBlock(20, 20, 200, 100, pcolor.RGBA(100, 180, 240, 1.0))
	ui := polyapp.NewPolyUI(blk, glr)
	if err := ui.Init(); err != nil {
		panic(err)
	}
	if err := w.Open(ui); err != nil {
		panic(err)
	}
}
