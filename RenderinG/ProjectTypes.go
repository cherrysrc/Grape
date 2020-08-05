package RenderinG

import "fmt"

type Translatable interface {
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
	StageSize []int
	Scenes    []string
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

func printSpacer(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf(" ")
	}
}

func (config GProjectConfig) Print(depth int) {
	printSpacer(depth)
	fmt.Printf("GConfig\n")

	printSpacer(depth)
	fmt.Printf("Name: %s\n", config.Name)

	printSpacer(depth)
	fmt.Printf("Stage size: %d by %d\n", config.StageSize[0], config.StageSize[1])

	printSpacer(depth)
	for i := range config.Scenes {
		fmt.Print(config.Scenes[i])
	}
}

func (object GObject) Print(depth int) {
	printSpacer(depth)
	fmt.Printf("ID: %s\n", object.ID)

	printSpacer(depth)
	fmt.Println("Vertices:")
	for i := range object.Vertices {
		fmt.Print(object.Vertices[i])
	}

	printSpacer(depth)
	fmt.Println("Colors:")
	for i := range object.Colors {
		fmt.Print(object.Colors[i])
	}
}

func (scene GScene) Print(depth int) {
	printSpacer(depth)
	fmt.Printf("Frames: %d", scene.Frames)

	printSpacer(depth)
	fmt.Println("Objects:")
	for i := range scene.Objects {
		scene.Objects[i].Print(depth + 1)
	}
}
