package Structures

import (
	"math/rand"
)

//Object Configuration
type GObject struct {
	ID string

	GeometricCenter []float64
	Rotation        float64
	Transparency    float64

	Vertices [][]float64
	Colors   [][]float64
}

//--------------------
//IObject interface implementation
//--------------------

var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//Generates a pseudo random id for objects that dont have one
func (object *GObject) GenerateID(n int) {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	object.ID = string(b)
}

//Sets a given GObjects position
func (object *GObject) Translate(targetP []float64) {
	object.GeometricCenter = targetP
}

//Sets objects rotation
func (object *GObject) Rotate(angle float64) {
	object.Rotation = angle
}

func (object *GObject) Fade(a float64) {
	object.Transparency = a
}
