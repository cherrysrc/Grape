package main

import (
	"RenderinG/RenderinG"
)

func main() {
	project := RenderinG.LoadProject("TestP")
	project.Print(0)
}
