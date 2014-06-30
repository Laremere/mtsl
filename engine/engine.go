package engine

import (
	"github.com/Laremere/mtsl/engine/input"
	"github.com/Laremere/mtsl/engine/output"
	"github.com/Laremere/mtsl/engine/sdl2"
)

type Engine struct {
	window  *sdl.Window
	keyDown map[string]bool
}

func NewEngine() *Engine {
	err := sdl.SdlInit()
	if err != nil {
		panic(err)
	}

	var eng Engine
	eng.keyDown = make(map[string]bool)
	eng.window, err = sdl.CreateWindowRenderer("Melcandra The Sky Lands", 10, 10, 1280, 720, sdl.WindowShown)
	if err != nil {
		panic(err)
	}
	return &eng
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

func (eng *Engine) Input(in *input.Input) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event := event.(type) {
		case *sdl.QuitEvent:
			in.Running = false
		case *sdl.MouseMoveEvent:
		case *sdl.KeyupEvent:
			eng.keyDown[event.Key] = false
		case *sdl.KeydownEvent:
			eng.keyDown[event.Key] = true
		default:
			//log.Println("Unkown event:", reflect.ValueOf(event).Type().Name(), event)
		}
	}

	in.Action.Update(eng.keyDown["Space"])

	in.North.Update(eng.keyDown["W"])
	in.West.Update(eng.keyDown["A"])
	in.South.Update(eng.keyDown["S"])
	in.East.Update(eng.keyDown["D"])

}

func (eng *Engine) Output(out *output.Output) {
	eng.window.Present()
}
