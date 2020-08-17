package Structures

import (
	"github.com/cherrysrc/Grape/Components/Utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math"
)

//Functions for GProject instances
type iProject interface {
	Init()
	GetCurrentScene() *GScene
	SetCurrentScene(int)

	CalculateVertices()
	GetObjectByID(string, GProject) *GObject

	GenerateAnimationHooks([]*GAnimation, float64)
	executeAnimation(*GAnimation)
	checkHooks()
	broadcastFrameToAnimations()

	Update()

	NextScene()
	NextFrame()

	PreProcess()
	PostProcess()
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
	animationHooks map[float64][]*GAnimation

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
	project.animationHooks = make(map[float64][]*GAnimation, 0)
	project.animChannels = make([]chan float64, 0)
}

func (project *GProject) GetCurrentScene() *GScene {
	if project.sceneIdx >= len(project.Scenes) {
		return nil
	}
	return &project.Scenes[project.sceneIdx]
}

func (project *GProject) SetCurrentScene(idx int) {
	project.sceneIdx = idx
}

//iProject CalculateVertices implementation
//Transformations like translation and rotation arent stored explicitly
//They're only calculated in pixel vertices
func (project *GProject) CalculateVertices() {
	vertices := imdraw.New(nil)
	scene := project.GetCurrentScene()

	if scene == nil {
		project.Vertices = vertices
		return
	}

	vertices.EndShape = imdraw.RoundEndShape
	for i := range scene.Objects {

		colorCount := len(scene.Objects[i].Colors)

		for vertex := range scene.Objects[i].Vertices {

			//Set color
			//Results in the last specified color to be used in case there is no color for every vertex
			if vertex < colorCount {
				color := pixel.RGB(scene.Objects[i].Colors[vertex][0], scene.Objects[i].Colors[vertex][1], scene.Objects[i].Colors[vertex][2])
				color = color.Mul(pixel.Alpha(scene.Objects[i].Transparency))
				vertices.Color = color
			}

			originX := scene.Objects[i].Vertices[vertex][0]
			originY := scene.Objects[i].Vertices[vertex][1]

			rotatedX := originX*math.Cos(scene.Objects[i].Rotation) - originY*math.Sin(scene.Objects[i].Rotation)
			rotatedY := originX*math.Sin(scene.Objects[i].Rotation) + originY*math.Cos(scene.Objects[i].Rotation)

			//Add vertex
			vertices.Push(pixel.V(scene.Objects[i].GeometricCenter[0]+rotatedX, scene.Objects[i].GeometricCenter[1]+rotatedY))
		}
		//Finish up shape
		vertices.Polygon(2)
	}

	project.Vertices = vertices
}

//Retrieve an object using its ID
//Returns a pointer
func (project GProject) GetObjectByID(id string) *GObject {
	for i := range project.Scenes {
		for j := range project.Scenes[i].Objects {

			if project.Scenes[i].Objects[j].ID == id {
				return &project.Scenes[i].Objects[j]
			}
		}
	}
	panic("Unknown ID specified in Animation")
}

//Fills the projects map of points in time and corresponding animations
func (project *GProject) GenerateAnimationHooks(animations []*GAnimation, sceneOffset float64) {
	for i := range animations {
		//Adjust animation start and end according to the scene they're in
		animations[i].StartFrame += sceneOffset
		animations[i].EndFrame += sceneOffset

		project.animationHooks[animations[i].StartFrame] = append(project.animationHooks[animations[i].StartFrame], animations[i])
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
	if animations, exists := project.animationHooks[project.frameIdx]; exists {
		for i := range animations {
			project.executeAnimation(animations[i])
		}
	}
}

//Loops over every channel of the project
//Sends the current frame trough the channels
func (project *GProject) broadcastFrameToAnimations() {
	for i := len(project.animChannels) - 1; i >= 0; i-- {
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

//Called before any important file parsing happens
func (project *GProject) PreProcess() {
	//Todo vertex template support
}

func (project *GProject) PostProcess() {
	//Generate scene transitions
	frameOffset := 0.0
	for i := range project.Scenes {

		var anim GAnimation

		anim.StartFrame = frameOffset + project.Scenes[i].Frames
		anim.EndFrame = anim.StartFrame + 1

		anim.Function = "scene_transit"
		anim.Params = []interface{}{project}

		project.animationHooks[anim.StartFrame] = append(project.animationHooks[anim.StartFrame], &anim)

		frameOffset += project.Scenes[i].Frames
	}
}
