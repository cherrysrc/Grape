package Structures

import (
	"github.com/cherrysrc/Grape/Components/Utils"
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
	executeAnimation(*GAnimation)
	checkHooks()
	broadcastFrameToAnimations()

	Update()

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

	animChannels []chan float64

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
	project.animChannels = make([]chan float64, 0)
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

//Fills the projects map of points in time and corresponding animations
func (project *GProject) GenerateAnimationHooks(animations []*GAnimation) {
	for i := range animations {
		project.animationHooks[animations[i].StartFrame] = animations[i]
	}
}

//Runs an animation
//Adds the animation and a channel to the parameters
//Appends the animation first, then the channel
func (project *GProject) executeAnimation(animation *GAnimation) {
	channel := make(chan float64)

	aInterface := interface{}(animation)

	params := animation.Params

	params = append(params, aInterface)
	params = append(params, channel)

	//Remember channel
	project.animChannels = append(project.animChannels, channel)
	//Call function as goroutine
	go AnimFunctions[animation.Function].(func([]interface{}))(params)
}

//Checks if theres an animation supposed to start at the current frame
//If there is, it calls executeAnimation to deal with further handling
//Called each frame through Update()
func (project *GProject) checkHooks() {
	if animation, exists := project.animationHooks[project.frameIdx]; exists {
		project.executeAnimation(animation)
	}
}

//Loops over every channel of the project
//Sends the current frame trough the channels
func (project *GProject) broadcastFrameToAnimations() {
	for i := range project.animChannels {
		project.animChannels[i] <- project.frameIdx

		status := <-project.animChannels[i]
		if status == 0.0 {
			//Remove channel, its dead
			project.animChannels = Utils.RemoveChannel(project.animChannels, i)
		}

		project.CalculateVertices()
	}
}

//General function combining all actions that need to performed done every frame
func (project *GProject) Update() {
	project.checkHooks()
	project.broadcastFrameToAnimations()
	project.NextFrame()
}

//Switches to the next scene
func (project *GProject) NextScene() {
	project.sceneIdx++
	//Switching to a new scene requires recalculation of vertices
	project.CalculateVertices()
}

//Switches to the next frame
func (project *GProject) NextFrame() {
	project.frameIdx++
}
