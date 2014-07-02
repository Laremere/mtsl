package engine

import (
	"github.com/Laremere/mtsl/engine/input"
	"github.com/Laremere/mtsl/engine/output"
	"github.com/Laremere/mtsl/engine/sdl2"
	"image/png"
	"os"
)

type Engine struct {
	window   *sdl.Window
	keyDown  map[string]bool
	graphics map[*output.Graphic]*image
}

type image struct {
	tex           sdl.Texture
	width, height int
}

func NewEngine() *Engine {
	err := sdl.SdlInit()
	if err != nil {
		panic(err)
	}

	var eng Engine
	eng.keyDown = make(map[string]bool)
	eng.window, err = sdl.CreateWindowRenderer("Melcandra The Sky Lands", 20, 40, 1280, 720, sdl.WindowShown)
	if err != nil {
		panic(err)
	}

	eng.graphics = make(map[*output.Graphic]*image)
	for _, graphic := range output.Graphics {
		eng.graphics[graphic], err = eng.loadImage(graphic.Filepath)
		if err != nil {
			panic(err)
		}
	}

	return &eng
}

func (eng *Engine) loadImage(filePath string) (*image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	PNGimage, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	tex, err := eng.window.CreateTexture(PNGimage.Bounds().Dx(), PNGimage.Bounds().Dx())
	if err != nil {
		return nil, err
	}

	width := PNGimage.Bounds().Dx()
	height := PNGimage.Bounds().Dy()

	pixels := make([]byte, 0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := PNGimage.At(x, y).RGBA()
			pixels = append(pixels, byte(a), byte(b), byte(g), byte(r))
		}
	}

	err = sdl.UpdateTexture(tex, pixels, PNGimage.Bounds().Dx())

	i := image{tex, width, height}

	return &i, err
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
	eng.window.Clear()

	for _, graph := range out.Graphs {
		i := eng.graphics[graph.Image]
		err := eng.window.RenderCopy(i.tex, graph.X, graph.Y, i.width, i.height)
		if err != nil {
			panic(err)
		}
	}

	out.Reset()
	eng.window.Present()
}
