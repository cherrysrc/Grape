package main

import (
	"RenderinG/RenderinG"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(func() {
		RenderinG.PixelMain("TestP")
	})
}
