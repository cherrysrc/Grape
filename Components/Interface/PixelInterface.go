package Interface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

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

	for !win.Closed() {
		project.Update()
		
		win.Clear(colornames.Skyblue)
		project.Vertices.Draw(win)
		win.Update()
	}
}
