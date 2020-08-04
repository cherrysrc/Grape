package RenderinG

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//Function that translates an object
//id, delta x, delta y
func TranslateAnim(obj *GObject, dx float64, dy float64, duration int) {
	count := 0
	for count < duration {
		for i := range obj.Vertices {
			obj.Vertices[i][0] += dx
			obj.Vertices[i][1] += dy
		}
		count++
	}
}

type Animator interface {
	parseFraming(string)
	parseBlock(string)
}

//Container for animation information
type Animation struct {
	StartFrame int
	EndFrame   int
	Actions    map[string]interface{}
}

//Performs linear interpolation
func lerp(v0 float64, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

//Performs linear interpolation on a vector
func lerp2D(v0 []float64, v1 []float64, t float64) []float64 {
	return []float64{
		lerp(v0[0], v1[0], t),
		lerp(v0[1], v1[1], t),
	}
}

func (animation *Animation) parseFraming(framing string) {
	framing = framing[1 : len(framing)-1] //Chop of parenthesis
	parts := strings.Split(framing, " ")

	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	animation.StartFrame = start
	animation.EndFrame = end
}

func (animation *Animation) parseBlock(block string) {
	block = block[1 : len(block)-1]
	lines := strings.Split(block, "\n")

	for i := range lines {
		if strings.Contains(lines[i], "#") || lines[i] == "" {
			continue
		}
		fmt.Println(lines[i])
		//Todo parse lines
	}
}

func LoadAnimations(name string) []Animation {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	content := string(bytes)

	blocks := strings.Split(content, ";")

	animations := make([]Animation, 0)

	for i := 0; i < len(blocks)-1; i++ {
		var anim Animation
		framingRegex, _ := regexp.Compile("\\((.*?)\\)")
		blockRegex, _ := regexp.Compile("{([^}]*)}")

		framing := framingRegex.FindString(blocks[i])
		block := blockRegex.FindString(blocks[i])

		anim.parseFraming(framing)
		anim.parseBlock(block)

		animations = append(animations, anim)
	}

	return animations
}
