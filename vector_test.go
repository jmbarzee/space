package space

import (
	"math"
	"testing"
)

const MinErr = 0.000001

// FloatsEqual compares floats
func FloatsEqual(a, b float64, err float64) bool {
	return float64(math.Abs(a-b)) < err
}

type VectorTest struct {
	Initial           Vector
	Operation         func(Vector) Vector
	ExpectedCartesian Cartesian
	ExpectedSpherical Spherical
}

func RunVectorTests(t *testing.T, cases []VectorTest) {
	for i, c := range cases {
		v := c.Operation(c.Initial)
		actualCartesian := v.Cartesian()
		if !CartesiansEqual(c.ExpectedCartesian, actualCartesian) {
			t.Fatalf("Test %v failed. Cartesians were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.ExpectedCartesian, actualCartesian)
		}
		actualSpherical := v.Spherical()
		if !SphericalsEqual(c.ExpectedSpherical, actualSpherical) {
			t.Fatalf("Test %v failed. Sphericals were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.ExpectedSpherical, actualSpherical)
		}
	}
}
