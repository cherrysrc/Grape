package Structures

import (
	"math"
	"strconv"
)

//Todo individual vertex translation
//Todo individual vertex fading

//Map containing references to every animation
//Keys correspond to the .anim file contents
var AnimFunctions = map[string]interface{}{
	"move":          TranslateAnim,
	"rotate":        RotateAnim,
	"fade":          FadeAnim,
	"scale":         ScaleAnim,
	"scene_transit": SceneTransit,
}

//Performs a translation animation
//(anim *GAnimation, channel chan float64, targetPos []float64)
func TranslateAnim(params []interface{}) {
	//Convert parameters

	dx, err := strconv.ParseFloat(params[0].(string), 64)
	if err != nil {
		panic("Wrong argument type in animation")
	}

	dy, err := strconv.ParseFloat(params[1].(string), 64)
	if err != nil {
		panic("Wrong argument type in animation")
	}

	targetPos := []float64{dx, dy}
	anim := params[2].(*GAnimation)
	channel := params[3].(chan float64)

	//Save original position
	originPos := anim.Target.GeometricCenter

	duration := anim.EndFrame - anim.StartFrame

	frame := <-channel
	//Calculate current interpolation progress based on the current frame received through the channel
	interp := (frame - anim.StartFrame) / duration
	for frame < anim.EndFrame {
		channel <- 1.0
		//Linearly interpolate between target and original position
		anim.Target.Translate(lerp2D(originPos, targetPos, interp))

		//Update frame and interpolation progress
		frame = <-channel
		interp = (frame - anim.StartFrame) / duration
	}
	interp = 1
	anim.Target.Translate(lerp2D(originPos, targetPos, interp))
	channel <- 0.0
}

//Performs a rotation animation
//RotateAnim(angle float64, animation *GAnimation, channel chan float64)
func RotateAnim(params []interface{}) {
	targetAngle, err := strconv.ParseFloat(params[0].(string), 64)
	if err != nil {
		panic("Wrong argument type in animation")
	}
	//Convert to radians
	targetAngle = targetAngle * math.Pi / 180

	anim := params[1].(*GAnimation)
	channel := params[2].(chan float64)

	//Save original rotation
	originRotation := anim.Target.Rotation

	duration := anim.EndFrame - anim.StartFrame

	frame := <-channel
	//Calculate current interpolation progress based on the current frame received through the channel
	interp := (frame - anim.StartFrame) / duration
	for frame < anim.EndFrame {
		channel <- 1.0
		//Linearly interpolate between target and original rotation
		anim.Target.Rotate(lerp(originRotation, targetAngle, interp))

		//Update frame and interpolation progress
		frame = <-channel
		interp = (frame - anim.StartFrame) / duration
	}
	interp = 1
	anim.Target.Rotate(lerp(originRotation, targetAngle, interp))
	channel <- 0.0
}

//Fades all vertices based on the first one
//FadeAnim(targetA float64, animation *GAnimation, channel chan float64)
func FadeAnim(params []interface{}) {
	targetA, err := strconv.ParseFloat(params[0].(string), 64)
	if err != nil {
		panic("Wrong argument type in animation")
	}

	anim := params[1].(*GAnimation)
	channel := params[2].(chan float64)

	//Save original transparency (RGBA-A)
	originA := anim.Target.Transparency

	duration := anim.EndFrame - anim.StartFrame

	frame := <-channel
	//Calculate current interpolation progress based on the current frame received through the channel
	interp := (frame - anim.StartFrame) / duration
	for frame < anim.EndFrame {
		channel <- 1.0
		//Linearly interpolate between target and original transparency
		anim.Target.Fade(lerp(originA, targetA, interp))

		//Update frame and interpolation progress
		frame = <-channel
		interp = (frame - anim.StartFrame) / duration
	}
	interp = 1
	anim.Target.Fade(lerp(originA, targetA, interp))
	channel <- 0.0
}

//Scales object up
//ScaleAnim(targetScl float64, animation *GAnimation, channel chan float64)
func ScaleAnim(params []interface{}) {
	targetScl, err := strconv.ParseFloat(params[0].(string), 64)
	if err != nil {
		panic("Wrong argument type in animation")
	}

	anim := params[1].(*GAnimation)
	channel := params[2].(chan float64)

	//Save original transparency (RGBA-A)
	originScl := anim.Target.Scale

	duration := anim.EndFrame - anim.StartFrame

	frame := <-channel
	//Calculate current interpolation progress based on the current frame received through the channel
	interp := (frame - anim.StartFrame) / duration
	for frame < anim.EndFrame {
		channel <- 1.0
		//Linearly interpolate between target and original transparency
		anim.Target.Scl(lerp(originScl, targetScl, interp))

		//Update frame and interpolation progress
		frame = <-channel
		interp = (frame - anim.StartFrame) / duration
	}
	interp = 1
	anim.Target.Scl(lerp(originScl, targetScl, interp))
	channel <- 0.0
}

//Special type of animation
//Gets auto inserted at the corresponding frames
//Executes a scene switch
func SceneTransit(params []interface{}) {
	project := params[0].(*GProject)
	channel := params[2].(chan float64)
	//Ignore frameIdx, doesnt matter here
	_ = <-channel

	project.NextScene()

	//Send back 0.0 to let the main thread know this animation is done
	channel <- 0.0
}

//Helper Function
//Performs 1D linear interpolation
func lerp(v0 float64, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

//Helper Function
//Performs 2D linear interpolation
//Uses 1D linear interpolation on each component
func lerp2D(v0 []float64, v1 []float64, t float64) []float64 {
	return []float64{
		lerp(v0[0], v1[0], t),
		lerp(v0[1], v1[1], t),
	}
}
