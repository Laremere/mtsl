package main

import (
	"github.com/Laremere/mtsl/engine"
	"github.com/Laremere/mtsl/engine/input"
	"github.com/Laremere/mtsl/engine/output"
)

func main() {
	eng := engine.NewEngine()
	in := input.Input{Running: true}
	out := output.NewOutput()

	var x, y int = -1, -1
	dx, dy := 1, 1

	for in.Running {
		eng.Input(&in)

		if x >= 1280-32 {
			dx = -1
		} else if x <= 0 {
			dx = 1
		}
		if y >= 720-32 {
			dy = -1
		} else if y <= 0 {
			dy = 1
		}

		x += dx
		y += dy

		out.Draw(output.Grass, x, y)
		eng.Output(out)
	}
}
