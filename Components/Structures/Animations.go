package Structures

import (
	"strconv"
)

var AnimFunctions = map[string]interface{}{
	"move_to":   TranslateAnim,
	"rotate_to": RotateAnim,
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
	channel <- 0.0
}

//Performs a rotation animation
//RotateAnim(angle float64, animation *GAnimation, channel chan float64)
func RotateAnim(params []interface{}) {
	targetAngle, err := strconv.ParseFloat(params[0].(string), 64)
	if err != nil {
		panic("Wrong argument type in animation")
	}

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
	channel <- 0.0
}

func lerp(v0 float64, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

func lerp2D(v0 []float64, v1 []float64, t float64) []float64 {
	return []float64{
		lerp(v0[0], v1[0], t),
		lerp(v0[1], v1[1], t),
	}
}
