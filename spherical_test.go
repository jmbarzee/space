package space

import (
	"math"
	"testing"
)

type SphericalTest struct {
	Initial   Spherical
	Operation func(Spherical) Spherical
	Expected  Spherical
}

func RunSphericalTests(t *testing.T, cases []SphericalTest) {
	for i, c := range cases {
		actual := c.Operation(c.Initial)
		if !SphericalsEqual(c.Expected, actual) {
			t.Fatalf("Test %v failed. Sphericals were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected, actual)
		}
	}
}

// SphericalsEqual compares Sphericals
func SphericalsEqual(a, b Spherical) bool {
	if !FloatsEqual(a.R, b.R, MinErr) {
		return false
	}
	// Handle Points which are close to origin with very different directions
	if a.R < MinErr {
		return true
	}

	if !FloatsEqual(a.P, b.P, MinErr) {
		return false
	}
	// Handle Points which are close to z-axis with very different rotations
	if a.P < MinErr || (0 < a.P-math.Pi && a.P-math.Pi < MinErr) {
		return true
	}

	if !FloatsEqual(a.T, b.T, MinErr) {
		return false
	}
	return true
}
func TestNewSpherical(t *testing.T) {
	// sq2o2 := math.Sqrt2 / 2
	cases := []SphericalTest{
		{
			Operation: func(v Spherical) Spherical {
				return NewSpherical(1, 0, 0)
			},
			Expected: Spherical{
				R: 1,
				P: 0,
				T: 0,
			},
		},
		{
			Operation: func(v Spherical) Spherical {
				return NewSpherical(1, 0, 0)
			},
			Expected: Spherical{
				R: 1,
				P: 0,
				T: 0,
			},
		},
	}
	RunSphericalTests(t, cases)
}

func TestSphericalCartesian(t *testing.T) {
	cases := make([]CartesianTest, len(AllEquivalencies))
	for i, pair := range AllEquivalencies {
		s := pair.Spherical
		cases[i] = CartesianTest{
			Operation: func(Cartesian) Cartesian {
				return s.Cartesian()
			},
			Expected: pair.Cartesian,
		}
	}
	RunCartesianTests(t, cases)
}
