package RenderinG

import "RenderinG/RenderinG/Types"

var AnimFunctions = map[string]interface{}{
	"move_to": TranslateAnim,
}

//Go routine
//Receives current frame via channel
//func TranslateAnim(anim *GAnimation, targetPos []float64, channel chan float64) {
func TranslateAnim(params []interface{}) {
	//Convert parameters
	anim := params[0].(*Types.GAnimation)
	targetPos := []float64{params[1].(float64), params[2].(float64)}
	channel := params[3].(chan float64)
	
	//Save original position
	originPos := anim.Target.GeometricCenter

	duration := anim.EndFrame - anim.StartFrame

	frame := <-channel
	//Calculate current interpolation progress based on the current frame received through the channel
	interp := (frame - anim.StartFrame) / duration
	for frame < anim.EndFrame {
		//Linearly interpolate between target and original position
		anim.Target.Translate(lerp2D(originPos, targetPos, interp))

		//Update frame and interpolation progress
		frame = <-channel
		interp = (frame - anim.StartFrame) / duration
	}
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
