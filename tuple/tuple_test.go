package tuple

import (
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	actual := MakePoint(4.3, -4.2, 3.1)

	if actual.X() != 4.3 ||
		actual.Y() != -4.2 ||
		actual.Z() != 3.1 ||
		actual.W() != 1.0 {
		t.Fatalf("tuple Point not instantiated correctly - %v\n", actual)
	}
}

func TestTupleIsVector(t *testing.T) {
	actual := MakeVector(4.3, -4.2, 3.1)

	if actual.X() != 4.3 ||
		actual.Y() != -4.2 ||
		actual.Z() != 3.1 ||
		actual.W() != 0.0 {
		t.Fatalf("tuple Vector not instantiated correctly - %v\n", actual)
	}
}

func TestAddingTwotuples(t *testing.T) {
	t1 := &Point{coords: [4]float64{3, -2, 5, 1.0}}
	t2 := &Vector{coords: [4]float64{-2, 3, 1, 0.0}}
	expected := &Point{coords: [4]float64{1, 1, 6, 1}}

	actual, _ := Add(t1, t2)

	if !Equals(actual, expected) {
		t.Fatalf("Sum of tuples is incorrect\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestSubtractingTwoPoints(t *testing.T) {
	p1 := MakePoint(3, 2, 1)
	p2 := MakePoint(5, 6, 7)

	expected := MakeVector(-2.0, -4.0, -6.0)
	actual, err := Subtract(p1, p2)

	if !Equals(actual, expected) || err != nil {
		t.Fatalf("Subtracting two Points is incorrect\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestSubtractingVectorFromPoint(t *testing.T) {
	p1 := MakePoint(3, 2, 1)
	p2 := MakeVector(5, 6, 7)

	expected := MakePoint(-2, -4, -6)
	actual, err := Subtract(p1, p2)

	if !Equals(actual, expected) || err != nil {
		t.Fatalf("Subtracting Vector from a Point is incorrect\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestSubtractingTwoVectors(t *testing.T) {
	v1 := MakeVector(3, 2, 1)
	v2 := MakeVector(5, 6, 7)

	expected := MakeVector(-2, -4, -6)
	actual, err := Subtract(v1, v2)

	if !Equals(actual, expected) || err != nil {
		t.Fatalf("Subtracting two Vectors is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestSubtractingVectorFromZeroVector(t *testing.T) {
	v1 := ZeroVector()
	v2 := MakeVector(1, -2, 3)

	expected := MakeVector(-1, 2, -3)
	actual, err := Subtract(v1, v2)

	if !Equals(actual, expected) || err != nil {
		t.Fatalf("Subtracting Vector from Zero Vector is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMultiplyTupleByAScalar(t *testing.T) {
	a := MakePoint(1, -2, 3)
	scalar := 3.5

	expected := MakePoint(3.5, -7, 10.5)
	actual := Multiply(a, scalar)

	if !Equals(actual, expected) {
		t.Fatalf("Multiplying a tuple by a scalar is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMultiplyTupleByAFraction(t *testing.T) {
	a := MakePoint(1, -2, 3)
	scalar := 0.5

	expected := MakePoint(0.5, -1, 1.5)
	actual := Multiply(a, scalar)

	if !Equals(actual, expected) {
		t.Fatalf("Multiplying a tuple by a fraction is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestDivideTupleByScalar(t *testing.T) {
	a := MakePoint(1, -2, 3)
	scalar := 2.0

	expected := MakePoint(0.5, -1, 1.5)
	actual, _ := Divide(a, scalar)

	if !Equals(actual, expected) {
		t.Fatalf("Dividing a tuple by a fraction is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}
