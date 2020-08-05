package RenderinG

import (
	"fmt"
	"math/rand"
)

type Animatable interface {
	GenerateID(int)
	CalculateCenter()
	Translate([]float64)
}

//Printing of Animation relevant objects
//Uses indentation
type GPrintable interface {
	Print(int)
}

type GProjectConfig struct {
	Name      string
	StageSize []float64
	Scenes    []string
}

type GProject struct {
	Name      string
	StageSize []float64
	Scenes    []GScene
}

//Object Configuration
type GObject struct {
	ID              string
	GeometricCenter []float64
	Vertices        [][]float64
	Colors          [][]float64
}

//Scene Configuration
//Maps to JSON
type GScene struct {
	Frames  int
	Objects []GObject
}

//Animation information
type GAnimation struct {
	StartFrame      float64
	EndFrame        float64
	Target          *GObject
	FunctionsParams map[string][]interface{}
}

//--------------------
//Translatable interface implementation
//--------------------

var letterBytes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (obj *GObject) GenerateID(n int) {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	obj.ID = string(b)
}

func (obj *GObject) CalculateCenter() {
	avgX := 0.0
	avgY := 0.0

	for i := range obj.Vertices {
		avgX += obj.Vertices[i][0]
		avgY += obj.Vertices[i][1]
	}

	vertexCount := float64(len(obj.Vertices))
	avgX /= vertexCount
	avgY /= vertexCount

	obj.GeometricCenter = []float64{
		avgX,
		avgY,
	}
}

func (obj *GObject) Translate(targetP []float64) {
	for i := range obj.Vertices {
		offset := []float64{
			obj.Vertices[i][0] - obj.GeometricCenter[0],
			obj.Vertices[i][1] - obj.GeometricCenter[1],
		}

		obj.Vertices[i][0] = targetP[0] + offset[0]
		obj.Vertices[i][1] = targetP[1] + offset[1]
	}
	obj.GeometricCenter = targetP
}

//---------------
//Various debug printing related methods
//---------------

func printSpacer(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("  ")
	}
}

func (project GProject) Print(depth int) {
	printSpacer(depth)
	fmt.Printf("GProject %s\n", project.Name)

	printSpacer(depth)
	fmt.Printf("Stage size: %.2f by %.2f\n", project.StageSize[0], project.StageSize[1])

	printSpacer(depth)
	for i := range project.Scenes {
		project.Scenes[i].Print(depth + 1)
	}
}

func (object GObject) Print(depth int) {
	printSpacer(depth)
	fmt.Printf("ID: %s\n", object.ID)

	printSpacer(depth)
	fmt.Printf("Geometric Center: ")
	fmt.Println(object.GeometricCenter)

	printSpacer(depth)
	fmt.Println("Vertices:")

	printSpacer(depth)
	for i := range object.Vertices {
		fmt.Print(object.Vertices[i])
	}
	fmt.Println()

	printSpacer(depth)
	fmt.Println("Colors:")

	printSpacer(depth)
	for i := range object.Colors {
		fmt.Print(object.Colors[i])
	}
	fmt.Println()
}

func (scene GScene) Print(depth int) {
	printSpacer(depth)
	fmt.Printf("GScene\n")

	printSpacer(depth + 1)
	fmt.Printf("Frames: %d\n", scene.Frames)

	printSpacer(depth + 1)
	fmt.Println("Objects:")
	for i := range scene.Objects {
		scene.Objects[i].Print(depth + 2)
	}
}
