package Interface

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

//Main loop for pixel engine
func PixelMain(projectName string) {
	project := LoadProject(projectName)
	project.Print(0)

	cfg := pixelgl.WindowConfig{
		Title:  projectName,
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(0, 0), atlas)

	frame := 0
	for !win.Closed() {
		win.Clear(colornames.Black)
		txt.Clear()

		fmt.Fprintln(txt, frame)
		frame++
		project.Update()

		project.Vertices.Draw(win)
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 4))

		win.Update()
	}
}
