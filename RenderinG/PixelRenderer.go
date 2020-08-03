package RenderinG

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	"os"
)

//Loads a project and converts its objects into drawables to be used by pixel
//param name: name of the project to load
//returns: (1)the project loaded as a GProject, (2)a IMDraw, used for pixel engine
func initProject(name string) (GProject, *imdraw.IMDraw) {
	project := LoadProject(name)
	vertices := imdraw.New(nil)

	_, sceneObjects := GetSceneProperties(project, 0)

	for i := range sceneObjects {
		for vertex := range sceneObjects[i].Vertices {
			vertices.Color = pixel.RGBA{
				R: sceneObjects[i].Colors[vertex][0],
				G: sceneObjects[i].Colors[vertex][1],
				B: sceneObjects[i].Colors[vertex][2],
				A: sceneObjects[i].Colors[vertex][3],
			}
			vertices.Push(pixel.V(sceneObjects[i].Vertices[vertex][0], sceneObjects[i].Vertices[vertex][1]))
		}
		//Todo json adjustable thickness
		vertices.Polygon(1)
	}

	return project, vertices
}

//
//Pixel Engines main loop
//
func PixelRun() {
	project, vertices := initProject("TestP")
	project.Print(0)

	//Todo project name as window title
	cfg := pixelgl.WindowConfig{
		Title:  "ProjectNameHere",
		Bounds: pixel.R(0, 0, 800, 600),
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
		//Render here
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
