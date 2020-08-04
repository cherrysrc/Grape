package main

import (
	"RenderinG/RenderinG"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"path/filepath"
)

func main() {
	pixelgl.Run(func() {
		fp, _ := filepath.Abs("./Projects/TestP/Scene01.anim")
		anims := RenderinG.LoadAnimations(fp)
		fmt.Println("Animations")
		fmt.Println(anims)
		//Todo use cli args
		RenderinG.PixelRun("TestP")
	})
}
