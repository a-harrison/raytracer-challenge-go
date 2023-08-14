package tuple

import (
	"errors"
	"math"
)

const EPSILON float64 = 0.00001

type tuple interface {
	Elems() [4]float64
	SetValue(iter int, value float64)
	IsPoint() bool
	IsVector() bool
	X() float64
	Y() float64
	Z() float64
	W() float64
}

func MakePoint(x, y, z float64) tuple {
	res := new(Point)
	res.coords = [4]float64{x, y, z, 1.0}
	return res
}

func MakeVector(x, y, z float64) tuple {
	res := new(Vector)
	res.coords = [4]float64{x, y, z, 0.0}
	return res
}

func fuzzyEquals(f1 float64, f2 float64) bool {
	if math.Abs(f1-f2) <= EPSILON {
		return true
	}
	return false
}

func Equals(t1 tuple, t2 tuple) bool {
	for iter := range t1.Elems() {
		if !fuzzyEquals(t1.Elems()[iter], t2.Elems()[iter]) {
			return false
		}
	}
	return true
}

func Add(t1 tuple, t2 tuple) (tuple, error) {
	var res tuple

	switch wSum := t1.W() + t2.W(); wSum {
	case 0:
		res = new(Vector)
	case 1:
		res = new(Point)
	default:
		return res, errors.New("Unable to add 2 Points")
	}

	for iter := range res.Elems() {
		res.SetValue(iter, t1.Elems()[iter]+t2.Elems()[iter])
	}

	return res, nil
}

func Subtract(lhs tuple, rhs tuple) (tuple, error) {
	var res tuple

	switch wSum := lhs.W() - rhs.W(); wSum {
	case 0:
		res = new(Vector)
	case 1:
		res = new(Point)
	case -1:
		return res, errors.New("Unable to subtract a Point from a Vector")
	default:
		return res, errors.New("Unidentified tuple subtraction performed")
	}

	for iter := range res.Elems() {
		res.SetValue(iter, lhs.Elems()[iter]-rhs.Elems()[iter])
	}

	return res, nil
}

func Multiply(t1 tuple, scalar float64) tuple {
	t := t1 
	t.SetValue(0, t.X() * scalar)
	t.SetValue(1, t.Y() * scalar)
	t.SetValue(2, t.Z() * scalar)

	return t
}

func Divide(t1 tuple, scalar float64) (tuple, error) {
	if scalar == 0 {
		return t1, errors.New("Unable to divide tuple by zero.")
	}

	t := t1 

	t.SetValue(0, t.X() / scalar)
	t.SetValue(1, t.Y() / scalar)
	t.SetValue(2, t.Z() / scalar)

	return t, nil
}
