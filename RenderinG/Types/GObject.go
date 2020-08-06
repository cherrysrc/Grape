package Types

import "math/rand"

//Animatable objects implement these functions
type iObject interface {
	GenerateID(int)
	CalculateCenter()
	Translate([]float64)
}

//Object Configuration
type GObject struct {
	ID              string
	GeometricCenter []float64
	Vertices        [][]float64
	Colors          [][]float64
}

//--------------------
//IObject interface implementation
//--------------------

var letterBytes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//Generates a pseudo random id for objects that dont have one
func (obj *GObject) GenerateID(n int) {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	obj.ID = string(b)
}

//Calculates the geometric center of a shape
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

//Sets a given GObjects position
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
