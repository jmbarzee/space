package space

import (
	"testing"
)

type VectorTest struct {
	Initial   Vector
	Operation func(Vector) Vector
	Expected  vectorEquivalent
}

func RunVectorTests(t *testing.T, cases []VectorTest) {
	for i, c := range cases {
		v := c.Operation(c.Initial)
		actualCartesian := v.Cartesian()
		if !CartesiansEqual(c.Expected.Cartesian, actualCartesian) {
			t.Fatalf("Test %v failed. Cartesians were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected.Cartesian, actualCartesian)
		}
		actualSpherical := v.Spherical()
		if !SphericalsEqual(c.Expected.Spherical, actualSpherical) {
			t.Fatalf("Test %v failed. Sphericals were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected.Spherical, actualSpherical)
		}
	}
}

func TestTranslate(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Spherical{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{1, 0, 0}
				return v.Translate(u)
			},
			Expected: AxisX,
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{1, 0, 0}
				return v.Translate(u)
			},
			Expected: AxisX,
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{-2, 0, 0}
				return v.Translate(u)
			},
			Expected: AxisXN,
		},
		{
			Initial: Cartesian{-1, -1, -1},
			Operation: func(v Vector) Vector {
				u := Cartesian{1, 1, 4}
				return v.Translate(u)
			},
			Expected: AxisZ3,
		},
	}
	RunVectorTests(t, cases)
}

func TestScale(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Spherical{0, 0, 0},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{-1, 0, 0},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{0, -1, 0},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{0, 0, -1},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Vector) Vector {
				i := 1.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Vector) Vector {
				i := 1.0
				return v.Scale(i)
			},
			Expected: AxisX,
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Vector) Vector {
				i := 1.0
				return v.Scale(i)
			},
			Expected: AxisY,
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Vector) Vector {
				i := 1.0
				return v.Scale(i)
			},
			Expected: AxisZ,
		},
		{
			Initial: Cartesian{-1, 0, 0},
			Operation: func(v Vector) Vector {
				i := 1.0
				return v.Scale(i)
			},
			Expected: AxisXN,
		},
		{
			Initial: Cartesian{0, -1, 0},
			Operation: func(v Vector) Vector {
				i := 1.0
				return v.Scale(i)
			},
			Expected: AxisYN,
		},
		{
			Initial: Cartesian{0, 0, -1},
			Operation: func(v Vector) Vector {
				i := 1.0
				return v.Scale(i)
			},
			Expected: AxisZN,
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Vector) Vector {
				i := 3.0
				return v.Scale(i)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Vector) Vector {
				i := 3.0
				return v.Scale(i)
			},
			Expected: AxisX3,
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Vector) Vector {
				i := 3.0
				return v.Scale(i)
			},
			Expected: AxisY3,
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Vector) Vector {
				i := 3.0
				return v.Scale(i)
			},
			Expected: AxisZ3,
		},
		{
			Initial: Cartesian{-1, 0, 0},
			Operation: func(v Vector) Vector {
				i := 3.0
				return v.Scale(i)
			},
			Expected: AxisXN3,
		},
		{
			Initial: Cartesian{0, -1, 0},
			Operation: func(v Vector) Vector {
				i := 3.0
				return v.Scale(i)
			},
			Expected: AxisYN3,
		},
		{
			Initial: Cartesian{0, 0, -1},
			Operation: func(v Vector) Vector {
				i := 3.0
				return v.Scale(i)
			},
			Expected: AxisZN3,
		},
	}
	RunVectorTests(t, cases)
}

func TestTransform(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Spherical{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Origin,
		},
		{
			Initial: Cartesian{3, 0, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisX3,
		},
		{
			Initial: Cartesian{0, 3, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisY3,
		},
		{
			Initial: Cartesian{0, 0, 3},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisZ3,
		},
		{
			Initial: Cartesian{-3, 0, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisXN3,
		},
		{
			Initial: Cartesian{0, -3, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisYN3,
		},
		{
			Initial: Cartesian{0, 0, -3},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisZN3,
		},
		{
			Initial: Cartesian{3, 3, 3},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisX3,
		},
		{
			Initial: Cartesian{3, 3, 3},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{0, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisY3,
		},
		{
			Initial: Cartesian{3, 3, 3},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisZ3,
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{3, 0, 0, 0},
					{0, 3, 0, 0},
					{0, 0, 3, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisX3,
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{3, 0, 0, 0},
					{0, 3, 0, 0},
					{0, 0, 3, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisY3,
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{3, 0, 0, 0},
					{0, 3, 0, 0},
					{0, 0, 3, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisZ3,
		},
		{
			Initial: Cartesian{-1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{3, 0, 0, 0},
					{0, 3, 0, 0},
					{0, 0, 3, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisXN3,
		},
		{
			Initial: Cartesian{0, -1, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{3, 0, 0, 0},
					{0, 3, 0, 0},
					{0, 0, 3, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisYN3,
		},
		{
			Initial: Cartesian{0, 0, -1},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{3, 0, 0, 0},
					{0, 3, 0, 0},
					{0, 0, 3, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisZN3,
		},
		{
			Initial: Cartesian{1, 1, -1},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, -1},
					{0, 1, 0, -1},
					{0, 0, 1, 2},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: AxisZ,
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorProject(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Spherical{1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{3, 3, 3}
				return v.Project(u)
			},
			Expected: AxisZ3,
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{3, 3, 3}
				return v.Project(u)
			},
			Expected: AxisX3,
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{3, 3, 3}
				return v.Project(u)
			},
			Expected: AxisY3,
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Vector) Vector {
				u := Cartesian{3, 3, 3}
				return v.Project(u)
			},
			Expected: AxisZ3,
		},
		{
			Initial: Cartesian{-1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{-3, -3, -3}
				return v.Project(u)
			},
			Expected: AxisXN3,
		},
		{
			Initial: Cartesian{0, -1, 0},
			Operation: func(v Vector) Vector {
				u := Cartesian{-3, -3, -3}
				return v.Project(u)
			},
			Expected: AxisYN3,
		},
		{
			Initial: Cartesian{0, 0, -1},
			Operation: func(v Vector) Vector {
				u := Cartesian{-3, -3, -3}
				return v.Project(u)
			},
			Expected: AxisZN3,
		},
	}
	RunVectorTests(t, cases)
}
