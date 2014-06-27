package main

import (
	"github.com/Laremere/mtsl/engine"
)

func main() {
	eng := engine.NewEngine()
	_ = eng

	x := 0
	y := 0
	for {
		eng.DrawColor(uint8(x), uint8(x+128), 255, 255)
		eng.DrawLine(0, 0, x, y)
		eng.Present()
		x += 1
		if x > 500 {
			x = 0
			y += 10
		}
	}
}
