package RenderinG

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	"os"
)

func calculateVertices(project GProject) *imdraw.IMDraw {
	vertices := imdraw.New(nil)
	scene := project.GetCurrentScene()

	for i := range scene.Objects {
		colorCount := len(scene.Objects[i].Colors)

		for vertex := range scene.Objects[i].Vertices {

			if vertex < colorCount {
				vertices.Color = pixel.RGBA{
					R: scene.Objects[i].Colors[vertex][0],
					G: scene.Objects[i].Colors[vertex][1],
					B: scene.Objects[i].Colors[vertex][2],
					A: scene.Objects[i].Colors[vertex][3],
				}
			}

			vertices.Push(pixel.V(scene.Objects[i].Vertices[vertex][0], scene.Objects[i].Vertices[vertex][1]))
		}

		//Todo json adjustable thickness
		vertices.Polygon(1)
	}

	return vertices
}

//
//Pixel Engines main loop
//
func PixelRun(projectName string) {
	//Todo actual parameter for project name instead of hardcoded 'TestP'
	project := LoadProject(projectName)
	vertices := calculateVertices(project)

	cfg := pixelgl.WindowConfig{
		Title:  project.Name,
		Bounds: pixel.R(0, 0, project.StageSize[0], project.StageSize[1]),
		VSync:  true,
	}

	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	window.SetSmooth(true)

	//last := time.Now()
	for !window.Closed() {
		//Δt := time.Since(last).Seconds()
		//last = time.Now()

		window.Clear(colornames.Firebrick)

		vertices.Draw(window)

		window.Update()
	}
}

//Loads a single image into memory to be used as a sprite
//param path: path of the file
//returns: the new picture
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
