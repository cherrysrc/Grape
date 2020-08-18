package main

import (
	"github.com/cherrysrc/Grape/Components/ProgramInterface"
	"github.com/faiface/pixel/pixelgl"
	"os"
)

func main() {
	pixelgl.Run(func() {
		export := false

		if len(os.Args) > 1{
			if os.Args[1] == "export"{
				export = true
			}
		}

		ProgramInterface.PixelMain("TestP", export)
	})
}
