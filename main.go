package main

import (
	"github.com/cherrysrc/Grape/Components/Interface"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(func() {
		Interface.PixelMain("TestP")
	})
}
