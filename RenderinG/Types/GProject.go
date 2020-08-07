package Types

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

//Functions for GProject instances
type iProject interface {
	Init()
	GetCurrentScene() *GScene
	SetCurrentScene(int)

	CalculateVertices()
	GetObjectByID(string, GProject) *GObject

	GenerateAnimationHooks([]*GAnimation)

	NextScene()
	NextFrame()
}

//Project struct
//Used for json parsing
type GProjectConfig struct {
	Name      string
	StageSize []float64
	Scenes    []string
}

//Actual project structure
type GProject struct {
	Name string

	StageSize []float64

	Scenes   []GScene
	sceneIdx int

	frameIdx       float64
	animationHooks map[float64]*GAnimation

	Vertices *imdraw.IMDraw
}

//--------------------
//Project interface implementation
//--------------------

//iProject GetCurrentScene implementation
func (project *GProject) Init() {
	project.frameIdx = 0
	project.sceneIdx = 0
	project.animationHooks = make(map[float64]*GAnimation)
}

func (project *GProject) GetCurrentScene() GScene {
	return project.Scenes[project.sceneIdx]
}

func (project *GProject) SetCurrentScene(idx int) {
	project.sceneIdx = idx
}

//iProject CalculateVertices implementation
func (project *GProject) CalculateVertices() {
	vertices := imdraw.New(nil)
	scene := project.GetCurrentScene()

	for i := range scene.Objects {

		colorCount := len(scene.Objects[i].Colors)

		for vertex := range scene.Objects[i].Vertices {

			//Set color
			//Results in the last specified color to be used in case there is no color for every vertex
			if vertex < colorCount {
				vertices.Color = pixel.RGBA{
					R: scene.Objects[i].Colors[vertex][0],
					G: scene.Objects[i].Colors[vertex][1],
					B: scene.Objects[i].Colors[vertex][2],
					A: scene.Objects[i].Colors[vertex][3],
				}
			}

			//Add vertex
			vertices.Push(pixel.V(scene.Objects[i].Vertices[vertex][0], scene.Objects[i].Vertices[vertex][1]))
		}
		//Finish up shape
		vertices.Polygon(0)
	}

	project.Vertices = vertices
}

//Retrieve an object using its ID
func (project GProject) GetObjectByID(id string) *GObject {
	scene := project.GetCurrentScene()

	for i := range scene.Objects {
		if scene.Objects[i].ID == id {
			return &scene.Objects[i]
		}
	}
	panic("Unknown ID specified in Animation")
}

func (project *GProject) GenerateAnimationHooks(animations []*GAnimation) {
	for i := range animations {
		project.animationHooks[animations[i].StartFrame] = animations[i]
	}
}

func (project *GProject) NextScene() {
	project.sceneIdx++
	//Switching to a new scene requires recalculation of vertices
	project.CalculateVertices()
}

func (project *GProject) NextFrame() {
	project.frameIdx++
}
