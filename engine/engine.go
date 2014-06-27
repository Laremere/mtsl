package engine

import (
	"github.com/Laremere/mtsl/engine/sdl2"
)

type Engine struct {
	window *sdl.Window
}

func NewEngine() *Engine {
	err := sdl.SdlInit()
	if err != nil {
		panic(err)
	}

	var eng Engine
	eng.window, err = sdl.CreateWindowRenderer("Melcandra The Sky Lands", 10, 10, 1280, 720, sdl.WindowShown)
	if err != nil {
		panic(err)
	}
	return &eng
}

func (eng *Engine) Present() {
	eng.window.Present()
}

func (eng *Engine) DrawColor(r, g, b, a uint8) {
	err := eng.window.SetDrawColor(r, g, b, a)
	if err != nil {
		panic(err)
	}
}

func (eng *Engine) DrawLine(x1, y1, x2, y2 int) {
	err := eng.window.RenderDrawLine(x1, y1, x2, y2)
	if err != nil {
		panic(err)
	}
}
