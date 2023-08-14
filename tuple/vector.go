package tuple

import (
	"math"
)

type Vector struct {
	coords [4]float64
}

func (v *Vector) Elems() [4]float64 {
	return v.coords
}

func (v *Vector) SetValue(i int, value float64) {
	v.coords[i] = value
}

func (v *Vector) X() float64 { return v.coords[0] }
func (v *Vector) Y() float64 { return v.coords[1] }
func (v *Vector) Z() float64 { return v.coords[2] }
func (v *Vector) W() float64 { return v.coords[3] }

func (v *Vector) IsPoint() bool  { return false }
func (v *Vector) IsVector() bool { return true }

func NewVector(x, y, z float64) Vector {
	return Vector{coords: [4]float64{x, y, z, 0}}
}

func ZeroVector() *Vector {
	return &Vector{coords: [4]float64{0, 0, 0, 0}}
}

func Negate(v *Vector) tuple {
	n, _ := Subtract(ZeroVector(), v)
	// negation := n.(Vector)
	return n
}

func Magnitude(v *Vector) float64 {
	var sum float64 

	for _, value := range v.Elems() {
		sum += ( value * value )
	}

	return math.Sqrt(sum)
}

func Normalize(v *Vector) Vector {
	var normalized Vector
	magnitude := Magnitude(v)

	for iter, value := range v.Elems() {
		normalized.SetValue(iter, value / magnitude)
	}

	return normalized
}

func Dot(v1, v2 *Vector) float64 {
	sum := 0.0

	for iter := range v1.Elems() {
		sum += ( v1.Elems()[iter] * v2.Elems()[iter] )
	}

	return sum
}

func Cross(v1, v2 *Vector) Vector {
	var tmp Vector

	tmp.SetValue(0, ( ( v1.Y() * v2.Z() ) - ( v1.Z() * v2.Y() ) ) )
	tmp.SetValue(1, ( ( v1.Z() * v2.X() ) - ( v1.X() * v2.Z() ) ) )
	tmp.SetValue(2, ( ( v1.X() * v2.Y() ) - ( v1.Y() * v2.X() ) ) )

	return tmp
}