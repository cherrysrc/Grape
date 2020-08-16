package Structures

import (
	"math/rand"
)

//Animatable objects implement these functions
type iObject interface {
	GenerateID(int)
	Translate([]float64)
	Rotate(float64)
}

//Object Configuration
type GObject struct {
	ID              string
	Rotation        float64
	GeometricCenter []float64
	Vertices        [][]float64
	Colors          [][]float64
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

//Rotates all vertices of an object around its center point
//Subtract center position
//Rotate around origin
//Add center position
func (object *GObject) Rotate(angle float64) {
	object.Rotation = angle
}
