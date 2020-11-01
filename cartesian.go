package space

import (
	"math"
)

// Cartesian is a 3D coordinate in cartesian form
type Cartesian struct {
	X, Y, Z float64
}

// NewCartesian produces a new Cartesian from spherical coordinates
func NewCartesian(orientation Orientation, radius float64) Cartesian {
	sinTheta, cosTheta := math.Sincos(orientation.Theta)
	sinPhi, cosPhi := math.Sincos(orientation.Phi)
	return Cartesian{
		X: radius * sinTheta * cosPhi,
		Y: radius * sinTheta * sinPhi,
		Z: radius * cosTheta,
	}
}

// Translate shifts a Cartesian by another Cartesian (addition)
func (v Cartesian) Translate(q Cartesian) Cartesian {
	return Cartesian{
		X: v.X + q.X,
		Y: v.Y + q.Y,
		Z: v.Z + q.Z,
	}
}

// Scale scales a Cartesian by i
func (v Cartesian) Scale(i float64) Cartesian {
	return Cartesian{
		X: v.X * i,
		Y: v.Y * i,
		Z: v.Z * i,
	}
}

// Transform Multiplyiplies a Cartesian by a given matrix
func (v Cartesian) Transform(m Matrix) Cartesian {
	w := (v.X * m[3][0]) + (v.Y * m[3][1]) + (v.Z * m[3][2]) + (1 * m[3][3])
	return Cartesian{
		X: (v.X * m[0][0]) + (v.Y * m[0][1]) + (v.Z * m[0][2]) + (1 * m[0][3]),
		Y: (v.X * m[1][0]) + (v.Y * m[1][1]) + (v.Z * m[1][2]) + (1 * m[1][3]),
		Z: (v.X * m[2][0]) + (v.Y * m[2][1]) + (v.Z * m[2][2]) + (1 * m[2][3]),
	}.Scale(1 / w)
}

// Project returns the projection of u onto v
func (v Cartesian) Project(u Cartesian) Cartesian {
	top := (u.X * v.X) + (u.Y * v.Y) + (u.Z * v.Z)
	bot := (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
	return v.Scale(top / bot)
}

// Negative returns the Cartesian which is directly opposite to v
func (v Cartesian) Negative() Cartesian {
	return Cartesian{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

// TranslationMatrix produces a matrix which will transform by v
func (v Cartesian) TranslationMatrix() Matrix {
	return Matrix{
		{1, 0, 0, v.X},
		{0, 1, 0, v.Y},
		{0, 0, 1, v.Z},
		{0, 0, 0, 1},
	}
}

// Orientation return the angles of the spherical coordinates of v
func (v Cartesian) Orientation() Orientation {
	return Orientation{
		Theta: math.Atan(v.Y / v.X),
		Phi:   math.Atan(math.Sqrt((v.X*v.X)+(v.Y*v.Y)) / v.Z),
	}
}
