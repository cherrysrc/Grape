package ProgramInterface

//#include "../C/Rendering.h"
import "C"
import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//Main loop for pixel engine
func PixelMain(projectName string, exporting bool) {
	//Todo cli args
	project, lastFrame := LoadProject(projectName)

	//C Library for outputting NetPBM images to stdout
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

	frame := 0.0
	for !win.Closed() {
		win.Clear(colornames.Black)

		frame++
		project.Update()

		project.Vertices.Draw(win)

		if exporting {
			for y := 0.0; y < project.StageSize[1]; y++ {
				for x := 0.0; x < project.StageSize[0]; x++ {
					rgba := win.Color(pixel.V(x, project.StageSize[1]-y-1))
					C.setPixel(rendering, C.int(x), C.int(y), C.uchar(rgba.R*255), C.uchar(rgba.G*255), C.uchar(rgba.B*255))
				}
			}

			C.writeRendering(rendering)
		}

		win.Update()

		//Close window if all frames have been done
		if frame >= lastFrame{
			win.SetClosed(true)
		}
	}
}
