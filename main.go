package main

import (
	"github.com/cherrysrc/Grape/Components/ProgramInterface"
	"github.com/faiface/pixel/pixelgl"
	"os"
)

func main() {
	pixelgl.Run(func() {
		if len(os.Args) < 2 {
			panic("No Project Given!")
		}

		ProgramInterface.PixelMain(os.Args[1])
	})
}
