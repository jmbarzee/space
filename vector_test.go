package space

import (
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
		if !VecotrsEqual(c.Expected, actual) {
			t.Fatalf("Test %v failed. Vectors were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected, actual)
		}
	}
}

// VecotrsEqual compares vectors
func VecotrsEqual(a, b Vector) bool {
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
