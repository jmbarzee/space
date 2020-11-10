package space

import (
	"testing"
)

func TestNewObject(t *testing.T) {
	cases := []struct {
		Location  Cartesian
		Spherical Spherical
		Rotation  Spherical
		Expected  Object
	}{}
	for i, c := range cases {
		actual := NewObject(c.Location, c.Spherical, c.Rotation)
		if !ObjectsEqual(&c.Expected, actual) {
			t.Fatalf("NewObject %v failed. Objects were not equal:\n\tExpected: %v,\n\tActual: %v", i, &c.Expected, actual)
		}
		t.Fatalf("NewObject %v failed. Objects were not equal:\n\tExpected: %v,\n\tActual: %v", i, &c.Expected, actual)
	}
}

// ObjectsEqual compares objects
func ObjectsEqual(a, b *Object) bool {
	if !CartesiansEqual(a.GetLocation(), b.GetLocation()) {
		return false
	}
	if !SphericalsEqual(a.GetOrientation(), b.GetOrientation()) {
		return false
	}
	if !SphericalsEqual(a.GetRotation(), b.GetRotation()) {
		return false
	}
	return true
}
