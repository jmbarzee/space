package space

import (
	"math"
)

// Cartesian represents a point in space in cartesian coordinates
type Cartesian struct {
	X, Y, Z float64
}

var _ Vector = (*Cartesian)(nil)

// NewCartesian produces a new Cartesian from spherical coordinates
func NewCartesian(x, y, z float64) Cartesian {
	return Cartesian{
		X: x,
		Y: y,
		Z: z,
	}
}

// Cartesian returns c
func (c Cartesian) Cartesian() Cartesian {
	return c
}

// Spherical returns the Spherical version of c
func (c Cartesian) Spherical() Spherical {
	x2 := c.X * c.X
	y2 := c.Y * c.Y
	z2 := c.Z * c.Z
	return NewSpherical(
		math.Sqrt(x2+y2+z2),
		math.Atan2(c.Y, c.X),
		math.Atan2(math.Sqrt(x2+y2), c.Z),
	)
}

// Translate shifts a Cartesian by a Vector (addition)
func (c Cartesian) Translate(v Vector) Vector {
	d := v.Cartesian()
	return Cartesian{
		X: c.X + d.X,
		Y: c.Y + d.Y,
		Z: c.Z + d.Z,
	}
}

// Scale scales a Cartesian by i
func (c Cartesian) Scale(i float64) Vector {
	return Cartesian{
		X: c.X * i,
		Y: c.Y * i,
		Z: c.Z * i,
	}
}

// Transform Multiplyiplies a Cartesian by a given matrix
func (c Cartesian) Transform(m Matrix) Vector {
	w := (c.X * m[3][0]) + (c.Y * m[3][1]) + (c.Z * m[3][2]) + (1 * m[3][3])
	return Cartesian{
		X: (c.X * m[0][0]) + (c.Y * m[0][1]) + (c.Z * m[0][2]) + (1 * m[0][3]),
		Y: (c.X * m[1][0]) + (c.Y * m[1][1]) + (c.Z * m[1][2]) + (1 * m[1][3]),
		Z: (c.X * m[2][0]) + (c.Y * m[2][1]) + (c.Z * m[2][2]) + (1 * m[2][3]),
	}.Scale(1 / w)
}

// Project returns the projection of v onto c
func (c Cartesian) Project(v Vector) Vector {
	d := v.Cartesian()
	top := (d.X * c.X) + (d.Y * c.Y) + (d.Z * c.Z)
	bot := (c.X * c.X) + (c.Y * c.Y) + (c.Z * c.Z)
	return c.Scale(top / bot)
}

// TranslationMatrix produces a matrix which will transform by v
func (c Cartesian) TranslationMatrix() Matrix {
	return Matrix{
		{1, 0, 0, c.X},
		{0, 1, 0, c.Y},
		{0, 0, 1, c.Z},
		{0, 0, 0, 1},
	}
}
