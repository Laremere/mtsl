package output

type Graphic struct {
	Filepath string
}

var Graphics = make([]*Graphic, 0)

func newGraphic(Filepath string) *Graphic {
	var graphic Graphic
	graphic.Filepath = "content/img/" + Filepath + ".png"
	Graphics = append(Graphics, &graphic)
	return &graphic
}

type Graph struct {
	Image *Graphic
	X, Y  float32
}

type Output struct {
	Graphs           []Graph
	XCenter, YCenter float32
}

func NewOutput() *Output {
	var out Output
	out.Graphs = make([]Graph, 0)

	return &out
}

func (out *Output) Draw(Image *Graphic, X, Y float32) {
	out.Graphs = append(out.Graphs, Graph{Image, X, Y})
}

func (out *Output) Reset() {
	out.Graphs = out.Graphs[0:0]
}

var Grass = newGraphic("grass")
