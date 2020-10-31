package space

import (
	"math"
	"testing"
)

func TestNewRotationMatrixX(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				t := 0.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{1, 0, 0},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 0, 0},
		},
		{
			Initial: Vector{0, 1, 0},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 0, 1},
		},
		{
			Initial: Vector{0, 0, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{0, -1, 0},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{1, -1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, 1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{2, -2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, 2, -2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 1, -1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, -1, 1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{2, 2, -2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, -2, 2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{1, -1, -1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, 1, 1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{2, -2, -2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, 2, 2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, -1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{2, 2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixX(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, -2, -2},
		},
	}
	RunVectorTests(t, cases)
}
func TestNewRotationMatrixY(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				t := 0.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{1, 0, 0},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 0, -1},
		},
		{
			Initial: Vector{0, 1, 0},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 1, 0},
		},
		{
			Initial: Vector{0, 0, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 0, 0},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 1, -1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, -1, 1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{2, 2, -2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, -2, 2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, 1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{1, -1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, 2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{2, -2, -2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, 1, -1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{1, -1, 1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, 2, -2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{2, -2, 2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, -1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{2, 2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixY(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, -2, -2},
		},
	}
	RunVectorTests(t, cases)
}
func TestNewRotationMatrixZ(t *testing.T) {
	cases := []VectorTest{
		{
			Initial: Vector{0, 0, 0},
			Operation: func(v Vector) Vector {
				t := 0.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 0, 0},
		},
		{
			Initial: Vector{1, 0, 0},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 1, 0},
		},
		{
			Initial: Vector{0, 1, 0},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, 0, 0},
		},
		{
			Initial: Vector{0, 0, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{0, 0, 1},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, 1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{1, -1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, 2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{2, -2, -2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{1, -1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, 1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{2, -2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := -1.0 / 2.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, 2, -2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, -1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, -2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 1.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{2, 2, -2},
		},
		{
			Initial: Vector{1, 1, 1},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{1, 1, 1},
		},
		{
			Initial: Vector{-1, -1, -1},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-1, -1, -1},
		},
		{
			Initial: Vector{2, 2, 2},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{2, 2, 2},
		},
		{
			Initial: Vector{-2, -2, -2},
			Operation: func(v Vector) Vector {
				t := 2.0 / 1.0 * math.Pi
				m := NewRotationMatrixZ(t)
				return v.Transform(m)
			},
			Expected: Vector{-2, -2, -2},
		},
	}
	RunVectorTests(t, cases)
}

func TestMatrixMultiply(t *testing.T) {

	cases := []struct {
		A        Matrix
		B        Matrix
		Expected Matrix
	}{
		{
			A: Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
			B: Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
			Expected: Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
		},
		{
			A: Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{2, 3, 5, 1},
			},
			B: Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{7, 8, 9, 1},
			},
			Expected: Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{9, 11, 14, 1},
			},
		},
		{
			A: Matrix{
				{5, 7, 9, 10},
				{2, 3, 3, 8},
				{8, 10, 2, 3},
				{3, 3, 4, 8},
			},
			B: Matrix{
				{3, 10, 12, 18},
				{12, 1, 4, 9},
				{9, 10, 12, 2},
				{3, 12, 4, 10},
			},
			Expected: Matrix{
				{210, 267, 236, 271},
				{93, 149, 104, 149},
				{171, 146, 172, 268},
				{105, 169, 128, 169},
			},
		},
	}
	for i, c := range cases {
		// fmt.Printf("A:%v\nB:%v\n", c.A, c.B)
		actual := c.A.Multiply(c.B)
		if !MatriciesEqual(c.Expected, actual) {
			t.Fatalf("Multiply %v failed. Matricies were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected, actual)
		}
	}
}

const (
	MinErr = 0.000001
)

// MatriciesEqual compares matricies
func MatriciesEqual(a, b Matrix) bool {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if !FloatsEqual(a[row][col], b[row][col], MinErr) {
				return false
			}
		}
	}
	return true
}

// FloatsEqual compares floats
func FloatsEqual(a, b float64, err float64) bool {
	return float64(math.Abs(float64(a-b))) < err
}
