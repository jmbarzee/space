package space

import "math"

// Vector is a 3D coordinate (also known as Point)
type Vector struct {
	X, Y, Z float64
}

// NewVector produces a new Vector from spherical coordinates
func NewVector(direction Orientation, radius float64) Vector {
	sinTheta, cosTheta := math.Sincos(direction.Theta)
	sinPhi, cosPhi := math.Sincos(direction.Phi)
	return Vector{
		X: radius * sinTheta * cosPhi,
		Y: radius * sinTheta * sinPhi,
		Z: radius * cosTheta,
	}
}

// Translate shifts a vector by another vector (addition)
func (v Vector) Translate(q Vector) Vector {
	return Vector{
		X: v.X + q.X,
		Y: v.Y + q.Y,
		Z: v.Z + q.Z,
	}
}

// Scale Multiplyivlies a vector by a given scale
func (v Vector) Scale(i float64) Vector {
	return Vector{
		X: v.X * i,
		Y: v.Y * i,
		Z: v.Z * i,
	}
}

// Transform Multiplyiplies a vector by a given matrix
func (v Vector) Transform(m Matrix) Vector {
	w := (v.X * m[3][0]) + (v.Y * m[3][1]) + (v.Z * m[3][2]) + (1 * m[3][3])
	return Vector{
		X: (v.X * m[0][0]) + (v.Y * m[0][1]) + (v.Z * m[0][2]) + (1 * m[0][3]),
		Y: (v.X * m[1][0]) + (v.Y * m[1][1]) + (v.Z * m[1][2]) + (1 * m[1][3]),
		Z: (v.X * m[2][0]) + (v.Y * m[2][1]) + (v.Z * m[2][2]) + (1 * m[2][3]),
	}.Scale(1 / w)
}

// Project returns the projection of u onto v
func (v Vector) Project(u Vector) Vector {
	top := (u.X * v.X) + (u.Y * v.Y) + (u.Z * v.Z)
	bot := (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
	return v.Scale(top / bot)
}

// Negative returns the vector which is directly opposite to v
func (v Vector) Negative() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

// TranslationMatrix produces a matrix which will transform by v
func (v Vector) TranslationMatrix() Matrix {
	return Matrix{
		{1, 0, 0, v.X},
		{0, 1, 0, v.Y},
		{0, 0, 1, v.Z},
		{0, 0, 0, 1},
	}
}

// Orientation return the angles of the spherical coordinates of v
func (v Vector) Orientation() Orientation {
	return Orientation{
		Theta: math.Atan(v.Y / v.X),
		Phi:   math.Atan(math.Sqrt((v.X*v.X)+(v.Y*v.Y)) / v.Z),
	}
}
