package space

import (
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
	if !near(a.R, b.R) {
		return false
	}
	// Handle Points which are close to origin with very different directions
	if near(a.R, 0) {
		return true
	}

	if !near(a.P, b.P) {
		return false
	}
	// Handle Points which are close to z-axis with very different rotations
	if near(a.P, 0) {
		return true
	}
	if near(a.P, rad(1, 1)) {
		return true
	}

	if !near(a.T, b.T) {
		// Handle Points which are close to 0 and 2pi theta
		if near(a.T, b.T-rad(2, 1)) {
			return true
		}
		if near(a.T-rad(2, 1), b.T) {
			return true
		}
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

func TestSphericalPortionOrthagonal(t *testing.T) {
	cases := []SphericalTest{
		{
			Initial: AxisX.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisX3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: Origin.Spherical,
		},
		{
			Initial: AxisX.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisXN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: Origin.Spherical,
		},
		{
			Initial: AxisX.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisY3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisY3.Spherical,
		},
		{
			Initial: AxisX.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisYN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisYN3.Spherical,
		},
		{
			Initial: AxisX.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisZ3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisZ3.Spherical,
		},
		{
			Initial: AxisX.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisZN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisZN3.Spherical,
		},
		{
			Initial: AxisY.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisX3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisX3.Spherical,
		},
		{
			Initial: AxisY.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisXN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisXN3.Spherical,
		},
		{
			Initial: AxisY.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisY.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: Origin.Spherical,
		},
		{
			Initial: AxisY.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisYN.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: Origin.Spherical,
		},
		{
			Initial: AxisY.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisZ3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisZ3.Spherical,
		},
		{
			Initial: AxisY.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisZN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisZN3.Spherical,
		},
		{
			Initial: AxisZ.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisX3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisX3.Spherical,
		},
		{
			Initial: AxisZ.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisXN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisXN3.Spherical,
		},
		{
			Initial: AxisZ.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisY3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisY3.Spherical,
		},
		{
			Initial: AxisZ.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisYN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: AxisYN3.Spherical,
		},
		{
			Initial: AxisZ.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisZ3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: Origin.Spherical,
		},
		{
			Initial: AxisZ.Spherical,
			Operation: func(v Spherical) Spherical {
				u := AxisZN3.Spherical
				return v.PortionOrtagonal(u)
			},
			Expected: Origin.Spherical,
		},
	}
	RunSphericalTests(t, cases)
}

func TestSphericalRotationMatrix(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					0,
					0,
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZ.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(1, 2),
					0,
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZ.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(2, 2),
					0,
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZ.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(3, 2),
					0,
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZ.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(4, 2),
					0,
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZ.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					0,
					rad(1, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisX.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(1, 2),
					rad(1, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisY.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(2, 2),
					rad(1, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisXN.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(3, 2),
					rad(1, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisYN.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(4, 2),
					rad(1, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisX.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					0,
					rad(2, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZN.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(1, 2),
					rad(2, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZN.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(2, 2),
					rad(2, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZN.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(3, 2),
					rad(2, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZN.Cartesian,
		},
		{
			Initial: AxisZ.Cartesian,
			Operation: func(v Cartesian) Cartesian {
				s := NewSpherical(
					0,
					rad(4, 2),
					rad(2, 2),
				)
				m := s.RotationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: AxisZN.Cartesian,
		},
	}
	RunCartesianTests(t, cases)
}
