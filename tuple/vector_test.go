package tuple

import (
	"math"
	"testing"
)

func TestNegatingAVector(t *testing.T) {
	v := NewVector(1, -2, 3)

	expected := MakeVector(-1, 2, -3)
	actual := Negate(&v)

	if !Equals(actual, expected) {
		t.Fatalf("Negating a Vector is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMagnitude100(t *testing.T) {
	v := NewVector(1, 0, 0)

	expected := 1.0
	actual := Magnitude(&v)

	if expected != actual {
		t.Fatalf("Magnitude of Vector (100) is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMagnitude010(t *testing.T) {
	v := NewVector(0, 1, 0)

	expected := 1.0
	actual := Magnitude(&v)

	if expected != actual {
		t.Fatalf("Magnitude of Vector (010) is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMagnitude001(t *testing.T) {
	v := NewVector(0, 0, 1)

	expected := 1.0
	actual := Magnitude(&v)

	if expected != actual {
		t.Fatalf("Magnitude of Vector (001) is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMagnitude123(t *testing.T) {
	v := NewVector(1, 2, 3)

	expected := math.Sqrt(14)
	actual := Magnitude(&v)

	if expected != actual {
		t.Fatalf("Magnitude of Vector (123) is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMagnitudeNegative123(t *testing.T) {
	v := NewVector(-1, -2, -3)

	expected := math.Sqrt(14)
	actual := Magnitude(&v)

	if expected != actual {
		t.Fatalf("Magnitude of Vector (-1-2-3) is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestNormalizingVector400(t *testing.T) {
	v := NewVector(4, 0, 0)

	actual := Normalize(&v)
	expected := NewVector(1, 0, 0)

	if !Equals(&actual, &expected) {
		t.Fatalf("Normalizing vector 400 is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestNormalizingVector123(t *testing.T) {
	v := NewVector(1, 2, 3)

	actual := Normalize(&v)
	expected := NewVector(
		1/math.Sqrt(14),
		2/math.Sqrt(14),
		3/math.Sqrt(14),
	)

	if !Equals(&actual, &expected) {
		t.Fatalf("Normalizing vector 123 is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestMagnitudeOfNormalVector(t *testing.T) {
	v := NewVector(1, 2, 3)
	norm := Normalize(&v)

	actual := Magnitude(&norm)
	expected := 1.0

	if actual != expected {
		t.Fatalf("Magnitude of a Normal vector is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestDotProductOfTwoVectors(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	actual := Dot(&v1, &v2)
	expected := 20.0

	if actual != expected {
		t.Fatalf("Dot product of 2 vectors is incorrect\nExpected: %v\n Actual: %v\n", expected, actual)
	}
}

func TestCrossProductOfTwoVectors(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	actual1 := Cross(&v1, &v2)
	actual2 := Cross(&v2, &v1)

	expected1 := NewVector(-1, 2, -1)
	expected2 := NewVector(1, -2, 1)

	if actual1 != expected1 {
		t.Fatalf("Cross product of 2 vectors is incorrect\nExpected: %v\n Actual: %v\n", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Fatalf("Cross product of 2 vectors is incorrect\nExpected: %v\n Actual: %v\n", expected2, actual2)
	}
}
