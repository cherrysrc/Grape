package main

import (
	"RenderinG/RenderinG"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
)

func main() {
	pixelgl.Run(RenderinG.PixelRun)
}
