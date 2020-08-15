package Structures

//Scene Configuration
//Maps to JSON
type GScene struct {
	Frames     int
	Objects    []GObject
	Animations []*GAnimation
}
