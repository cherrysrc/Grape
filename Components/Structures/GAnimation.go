package Structures

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

//Parse a line of the animation segments body
func (animation *GAnimation) ParseLine(line string, project GProject) {
	if line == "" || strings.Contains(line, "#") {
		//Ignore empty lines, or lines containing #
		return
	}

	elements := strings.Split(line, " ")

	animation.Target = project.GetObjectByID(elements[0])

	animation.Function = elements[1]

	animation.Params = make([]interface{}, len(elements)-2)

	for j := 2; j < len(elements); j++ {
		animation.Params[j-2] = elements[j]
	}
}
