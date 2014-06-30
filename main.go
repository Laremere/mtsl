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

	for in.Running {
		eng.Input(&in)

		eng.Output(out)
	}
}
