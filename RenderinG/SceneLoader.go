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
	Size   []int
}

//
//Object data struct
//
type GObject struct {
	Vertices [][]int
	Color    string
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
//
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
	printSpacer(depth)
	fmt.Printf("Color: %s\n", g.Color)
	printSpacer(depth)
	fmt.Printf("Vertices:\n")
	for i := 0; i < len(g.Vertices); i++ {
		printSpacer(depth)
		fmt.Printf("[%d, %d]\n", g.Vertices[i][0], g.Vertices[i][1])
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
	fmt.Printf("Frames: %d, Size: [%d, %d]\n", g.Frames, g.Size[0], g.Size[1])
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
