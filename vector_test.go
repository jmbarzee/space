package space

import (
	"math"
	"testing"
)

type VectorTest struct {
	Initial   Vector
	Operation func(Vector) Vector
	Expected  Vector
}

func RunVectorTests(t *testing.T, cases []VectorTest) {
	for i, c := range cases {
		actual := c.Operation(c.Initial)
		if !VectorsEqual(c.Expected, actual) {
			t.Fatalf("Test %v failed. Vectors were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected, actual)
		}
	}
}

// VectorsEqual compares vectors
func VectorsEqual(a, b Vector) bool {
	if !FloatsEqual(a.X, b.X, MinErr) {
		return false
	}
	if !FloatsEqual(a.Y, b.Y, MinErr) {
		return false
	}
	if !FloatsEqual(a.Z, b.Z, MinErr) {
		return false
	}
	return true
}

func TestNewVector(t *testing.T) {
	sq2o2 := math.Sqrt2 / 2
	cases := []VectorTest{
		{
			Operation: func(v Vector) Vector {
				o := NewOrientation(0, 0)
				return NewVector(o, 1)
			},
			Expected: Vector{0, 0, 1},
		},
		{
			Operation: func(v Vector) Vector {
				t := 0.0 / 4.0 * math.Pi
				p := 1.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{0, 0, 1},
		},
		{
			Operation: func(v Vector) Vector {
				t := 1.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{sq2o2, 0, sq2o2},
		},
		{
			Operation: func(v Vector) Vector {
				t := 2.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{1, 0, 0},
		},
		{
			Operation: func(v Vector) Vector {
				t := 3.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{sq2o2, 0, -sq2o2},
		},
		{
			Operation: func(v Vector) Vector {
				t := 4.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{0, 0, -1},
		},
		{
			Operation: func(v Vector) Vector {
				t := 5.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{-sq2o2, 0, -sq2o2},
		},
		{
			Operation: func(v Vector) Vector {
				t := 6.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{-1, 0, 0},
		},
		{
			Operation: func(v Vector) Vector {
				t := 7.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{-sq2o2, 0, sq2o2},
		},
		{
			Operation: func(v Vector) Vector {
				t := 8.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewVector(o, 1)
			},
			Expected: Vector{0, 0, 1},
		},
	}
	RunVectorTests(t, cases)
}
func TestVectorTranslate(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Vector{2, 3, 5}
				return v.Translate(u)
			},
			Expected: Vector{2, 3, 5},
		},
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Vector{-2, -3, -5}
				return v.Translate(u)
			},
			Expected: Vector{-2, -3, -5},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 0}
				return v.Translate(u)
			},
			Expected: Vector{2, 3, 5},
		},
		{
			Initial: Vector{-2, -3, -5},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 0}
				return v.Translate(u)
			},
			Expected: Vector{-2, -3, -5},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorScale(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{-2, -3, -5},
			Operation: func(v Vector) Vector {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				i := 1.1
				return v.Scale(i)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				i := 1.1
				return v.Scale(i)
			},
			Expected: Vector{2.2, 3.3, 5.5},
		},
		{
			Initial: Vector{-2, -3, -5},
			Operation: func(v Vector) Vector {
				i := 1.1
				return v.Scale(i)
			},
			Expected: Vector{-2.2, -3.3, -5.5},
		},
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				i := -1.1
				return v.Scale(i)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				i := -1.1
				return v.Scale(i)
			},
			Expected: Vector{-2.2, -3.3, -5.5},
		},
		{
			Initial: Vector{-2, -3, -5},
			Operation: func(v Vector) Vector {
				i := -1.1
				return v.Scale(i)
			},
			Expected: Vector{2.2, 3.3, 5.5},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorTransform(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Vector{2, 3, 5},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Vector{2, 0, 0},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{0, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Vector{0, 3, 0},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Vector{0, 0, 5},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{2, 0, 0, 0},
					{0, 2, 0, 0},
					{0, 0, 2, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Vector{4, 6, 10},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Matrix{
					{1, 0, 0, 9},
					{0, 1, 0, 9},
					{0, 0, 1, 9},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Vector{11, 12, 14},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorProject(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Vector{2, 0, 0}
				return v.Project(u)
			},
			Expected: Vector{2, 0, 0},
		},
		{
			Initial: Vector{0, 1, 0},
			Operation: func(v Vector) Vector {
				u := Vector{2, 0, 0}
				return v.Project(u)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{0, 0, 1},
			Operation: func(v Vector) Vector {
				u := Vector{2, 0, 0}
				return v.Project(u)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Vector{0, 2, 0}
				return v.Project(u)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{0, 1, 0},
			Operation: func(v Vector) Vector {
				u := Vector{0, 2, 0}
				return v.Project(u)
			},
			Expected: Vector{0, 2, 0},
		},
		{
			Initial: Vector{0, 0, 1},
			Operation: func(v Vector) Vector {
				u := Vector{0, 2, 0}
				return v.Project(u)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{1, 0, 0},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 2}
				return v.Project(u)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{0, 1, 0},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 2}
				return v.Project(u)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{0, 0, 1},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 2}
				return v.Project(u)
			},
			Expected: Vector{0, 0, 2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				u := Vector{6, 0, 0}
				return v.Project(u)
			},
			Expected: Vector{2, 2, 2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				u := Vector{0, 6, 0}
				return v.Project(u)
			},
			Expected: Vector{2, 2, 2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 6}
				return v.Project(u)
			},
			Expected: Vector{2, 2, 2},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorNegative(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				return v.Negative()
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{1, 0, 0},
			Operation: func(v Vector) Vector {
				return v.Negative()
			},
			Expected: Vector{-1, 0, 0},
		},
		{
			Initial: Vector{0, 2, 3},
			Operation: func(v Vector) Vector {
				return v.Negative()
			},
			Expected: Vector{0, -2, -3},
		},
		{
			Initial: Vector{1, 2, 3},
			Operation: func(v Vector) Vector {
				return v.Negative()
			},
			Expected: Vector{-1, -2, -3},
		},
	}
	RunVectorTests(t, cases)
}

func TestVectorTranslationMatrix(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Vector{2, 3, 5}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Vector{2, 3, 5},
		},
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				u := Vector{-2, -3, -5}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Vector{-2, -3, -5},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 0}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Vector{2, 3, 5},
		},
		{
			Initial: Vector{-2, -3, -5},
			Operation: func(v Vector) Vector {
				u := Vector{0, 0, 0}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Vector{-2, -3, -5},
		},
		{
			Initial: Vector{2, 3, 5},
			Operation: func(v Vector) Vector {
				u := Vector{7, 8, 9}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Vector{9, 11, 14},
		},
	}
	RunVectorTests(t, cases)
}
