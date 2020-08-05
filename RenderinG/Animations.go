package RenderinG

var AnimFunctions = map[string]interface{}{
	"move_to": TranslateAnim,
}

//Go routine
//Receives current frame via channel
func TranslateAnim(anim *GAnimation, targetPos []float64, iChannel chan float64) {
	originPos := anim.Target.GeometricCenter

	duration := anim.EndFrame - anim.StartFrame

	frame := <-iChannel
	interp := (frame - anim.StartFrame) / duration
	for frame < anim.EndFrame {

		anim.Target.Translate(lerp2D(originPos, targetPos, interp))

		frame = <-iChannel
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
