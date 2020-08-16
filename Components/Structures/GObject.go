package Structures

import (
	"math"
	"math/rand"
)

//Animatable objects implement these functions
type iObject interface {
	GenerateID(int)
	CalculateCenter()
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

//Calculates the geometric center of a shape
func (object *GObject) CalculateCenter() {
	avgX := 0.0
	avgY := 0.0

	for i := range object.Vertices {
		avgX += object.Vertices[i][0]
		avgY += object.Vertices[i][1]
	}

	vertexCount := float64(len(object.Vertices))
	avgX /= vertexCount
	avgY /= vertexCount

	object.GeometricCenter = []float64{
		avgX,
		avgY,
	}
}

//Sets a given GObjects position
func (object *GObject) Translate(targetP []float64) {
	for i := range object.Vertices {
		offset := []float64{
			object.Vertices[i][0] - object.GeometricCenter[0],
			object.Vertices[i][1] - object.GeometricCenter[1],
		}

		object.Vertices[i][0] = targetP[0] + offset[0]
		object.Vertices[i][1] = targetP[1] + offset[1]
	}
	object.GeometricCenter = targetP
}

//Rotates all vertices of an object around its center point
//Subtract center position
//Rotate around origin
//Add center position
func (object *GObject) Rotate(angle float64) {
	for i := range object.Vertices {
		originX := object.Vertices[i][0] - object.GeometricCenter[0]
		originY := object.Vertices[i][1] - object.GeometricCenter[1]

		rotatedX := originX*math.Cos(angle) - originY*math.Sin(angle)
		rotatedY := originX*math.Sin(angle) + originY*math.Cos(angle)

		object.Vertices[i][0] = rotatedX + object.GeometricCenter[0]
		object.Vertices[i][1] = rotatedY + object.GeometricCenter[1]
	}
	object.Rotation = angle
}
