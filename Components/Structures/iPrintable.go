package Structures

import (
	"fmt"
)

//Printing of Animation relevant objects
//Uses indentation
type iPrintable interface {
	Print(int)
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
	fmt.Println("Animation Hooks: ")
	fmt.Println(project.animationHooks)

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

func (animation GAnimation) Print(depth int) {
	printSpacer(depth)
	fmt.Printf("GAnimation %.2f to %.2f\n", animation.StartFrame, animation.EndFrame)

	printSpacer(depth + 1)
	fmt.Printf("Target: %s\n", animation.Target.ID)

	printSpacer(depth + 1)
	fmt.Print("FunctionsParams: ")
	fmt.Println(animation.Function)
	fmt.Println(animation.Params)

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
		fmt.Println()
	}

	printSpacer(depth + 1)
	fmt.Println("Animations")
	for i := range scene.Animations {
		scene.Animations[i].Print(depth + 2)
	}
}
