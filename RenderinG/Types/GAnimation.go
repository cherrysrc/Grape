package Types

import (
	"strconv"
	"strings"
)

//Functions for GAnimation instances
type iAnimation interface {
	ParseFraming(string)
	ParseBody(string)
}

//Animation information
type GAnimation struct {
	StartFrame float64
	EndFrame   float64
	Target     *GObject
	Function   string
	Params     []interface{}
}

//--------------------
//Animation interface implementation
//--------------------

//Parse an animation framing (start and end)
func (animation *GAnimation) ParseFraming(framing string) {
	framing = framing[1 : len(framing)-1] //Strip enclosing parenthesis
	parts := strings.Split(framing, " ")

	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	animation.StartFrame = float64(start)
	animation.EndFrame = float64(end)
}

//Parse an animation body
func (animation *GAnimation) ParseBody(block string, project GProject) {
	block = block[1 : len(block)-1] //Strip enclosing curly brackets
	lines := strings.Split(block, "\n")

	for i := range lines {
		if lines[i] == "" || strings.Contains(lines[i], "#") {
			//Ignore empty lines, or lines containing #
			continue
		}

		elements := strings.Split(lines[i], " ")

		animation.Target = project.GetObjectByID(elements[0])

		animation.Function = elements[1]
		animation.Params = make([]interface{}, 0)

		for j := 2; j < len(elements); j++ {
			animation.Params = append(animation.Params, elements[j])
		}
	}
}
