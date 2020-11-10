package space

import (
	"testing"
)

func TestNewObject(t *testing.T) {
	cases := []struct {
		Location    Cartesian
		Orientation Spherical
		Rotation    Spherical
		Expected    Object
	}{
		{
			Location:    Origin.Cartesian,
			Orientation: AxisX.Spherical,
			Rotation:    AxisY.Spherical,
			Expected: Object{
				location:    Origin.Cartesian,
				orientation: AxisX.Spherical,
				rotation:    AxisY.Spherical,
			},
		},
		{
			Location:    Origin.Cartesian,
			Orientation: AxisX.Spherical,
			Rotation:    AxisZ.Spherical,
			Expected: Object{
				location:    Origin.Cartesian,
				orientation: AxisX.Spherical,
				rotation:    AxisZ.Spherical,
			},
		},
	}
	for i, c := range cases {
		actual := NewObject(c.Location, c.Orientation, c.Rotation)
		if !ObjectsEqual(&c.Expected, actual) {
			t.Fatalf("NewObject %v failed. Objects were not equal:\n\tExpected: %v,\n\tActual: %v", i, &c.Expected, actual)
		}
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
