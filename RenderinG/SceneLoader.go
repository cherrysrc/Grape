package RenderinG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//
//Interface for printable g structs
//
type GPrintable interface {
	Print(depth int)
}

//
//Scene configuration struct
//
type GConfig struct {
	Frames int
}

//
//Object data struct
//
type GObject struct {
	Vertices [][]float64
	Colors   [][]float64
}

//
//Scene information struct
//
type GScene struct {
	Config  GConfig
	Objects []GObject
}

//
//Loads a scene by its filename
//param name: scene name
//returns: Newly loaded GScene
func LoadScene(name string) GScene {
	content, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	}

	var scene GScene

	err = json.Unmarshal(content, &scene)

	if err != nil {
		log.Fatal(err)
	}

	return scene
}

//
//Function offsetting a print by a given amount
//
func printSpacer(count int) {
	for i := 0; i < count; i++ {
		fmt.Print("  ")
	}
}

//
//GPrintable interface (GObject) implementation
//
func (g GObject) Print(depth int) {
	printSpacer(depth)
	fmt.Println("GObject")
	for i := range g.Colors {
		printSpacer(depth)
		fmt.Printf("[%f, %f, %f, %f]\n", g.Colors[i][0], g.Colors[i][2], g.Colors[i][3], g.Colors[i][3])
	}
	printSpacer(depth)
	fmt.Printf("Vertices:\n")
	for i := range g.Vertices {
		printSpacer(depth)
		fmt.Printf("[%f, %f]\n", g.Vertices[i][0], g.Vertices[i][1])
	}
	printSpacer(depth)
	fmt.Println(" ")
}

//
//GPrintable interface (GConfig) implementation
//
func (g GConfig) Print(depth int) {
	printSpacer(depth)
	fmt.Println("GConfig")
	printSpacer(depth)
	fmt.Printf("Frames: %d\n", g.Frames)
	printSpacer(depth)
	fmt.Println(" ")
}

//
//GPrintable interface (GScene) implementation
//
func (g GScene) Print(depth int) {
	printSpacer(depth)
	fmt.Println("GScene")
	g.Config.Print(depth + 1)
	for i := range g.Objects {
		g.Objects[i].Print(depth + 1)
	}
	printSpacer(depth)
	fmt.Println(" ")
}
