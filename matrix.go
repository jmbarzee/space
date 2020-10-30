package space

import "math"

// Matrix is a transformational matrix for 3D space (3 x 3)
type Matrix [][]float64

// NewRotationMatrixX produces a matrix which will rotate about X
func NewRotationMatrixX(theta float64) Matrix {
	sin, cos := math.Sincos(theta)
	return Matrix{
		{1, 0, 0, 0},
		{0, cos, -sin, 0},
		{0, sin, cos, 0},
		{0, 0, 0, 1},
	}
}

// NewRotationMatrixY produces a matrix which will rotate about Y
func NewRotationMatrixY(theta float64) Matrix {
	sin, cos := math.Sincos(theta)
	return Matrix{
		{cos, 0, sin, 0},
		{0, 1, 0, 0},
		{-sin, 0, cos, 0},
		{0, 0, 0, 1},
	}
}

// NewRotationMatrixZ produces a matrix which will rotate about Z
func NewRotationMatrixZ(theta float64) Matrix {
	sin, cos := math.Sincos(theta)
	return Matrix{
		{cos, -sin, 0, 0},
		{sin, cos, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// Multiply will return the result of m * n
func (m Matrix) Multiply(n Matrix) Matrix {
	r := Matrix{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
	}
	for rowM := 0; rowM < 4; rowM++ {
		for colN := 0; colN < 4; colN++ {
			a := m[rowM][0] * n[0][colN]
			b := m[rowM][1] * n[1][colN]
			c := m[rowM][2] * n[2][colN]
			d := m[rowM][3] * n[3][colN]
			r[rowM][colN] = a + b + c + d
		}
	}
	return r
}
