package tuple

import (
	"fmt"
)

type Point struct {
	coords [4]float64
}

func (p *Point) Elems() [4]float64 {
	return p.coords
}

func (p *Point) SetValue(i int, value float64) {
	p.coords[i] = value
}


func (p *Point) X() float64 { return p.coords[0] }
func (p *Point) Y() float64 { return p.coords[1] }
func (p *Point) Z() float64 { return p.coords[2] }
func (p *Point) W() float64 { return p.coords[3] }

func (p *Point) IsPoint() bool { return true }
func (p *Point) IsVector() bool { return false }

func (p Point) String() string {
	return fmt.Sprintf("[%v, %v, %v]", p.X(), p.Y(), p.Z())
}

func NewPoint(x, y, z float64) Point {
	return Point{ coords: [4]float64 { x, y, z, 1}}
}