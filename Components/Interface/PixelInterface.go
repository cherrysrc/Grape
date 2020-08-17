package Interface

//#include "../C/Rendering.h"
import "C"
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

	rendering := C.createRendering(C.int(project.StageSize[0]), C.int(project.StageSize[1]))
	defer C.freeRendering(rendering)

	cfg := pixelgl.WindowConfig{
		Title:  projectName,
		Bounds: pixel.R(0, 0, project.StageSize[0], project.StageSize[1]),
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

		//TODO fix image flip
		for y := project.StageSize[1] - 1; y >= 0; y-- {
			for x := 0.0; x < project.StageSize[0]; x++ {
				rgba := win.Color(pixel.V(x, y))
				C.setPixel(rendering, C.int(x), C.int(y), C.uchar(rgba.R*255), C.uchar(rgba.G*255), C.uchar(rgba.B*255))
			}
		}

		C.writeRendering(rendering)
		win.Update()
	}
}
