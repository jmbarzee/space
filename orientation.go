package space

import "math"

// Orientation represents the direction of spherical coordinates
type Orientation struct {
	// Phi is tilt from Z
	Phi float64
	// Theta is rotation about Z
	Theta float64
}

// NewOrientation creates a new Orientation from a rotation and tilt
func NewOrientation(theta, phi float64) Orientation {
	o := Orientation{}
	o = o.Rotate(theta)
	o = o.Tilt(phi)
	return o
}

// Rotate will adjust the rotation about Z by theta
func (o Orientation) Rotate(theta float64) Orientation {
	wrappedTheta := o.Theta + theta
	o.Theta = math.Mod(wrappedTheta, math.Pi*2)
	return o
}

// Tilt will adjust the tilt from Z by phi
func (o Orientation) Tilt(phi float64) Orientation {
	wrappedPhi := o.Phi + phi
	newPhi := math.Mod(wrappedPhi, math.Pi*2)

	// Check if tilt is negative
	if newPhi < 0 {
		newPhi = (math.Pi * 2) + newPhi
	}

	// Check if tilt is beyond range of [0, pi]
	if newPhi > math.Pi {
		o.Phi = (math.Pi * 2) - newPhi
		return o.Rotate(math.Pi)
	}

	o.Phi = newPhi
	return o
}

// Transform transforms the orientation by the given matrix
func (o Orientation) Transform(m Matrix) Orientation {
	v := o.Cartesian().Transform(m)
	return v.Orientation()
}

// RotationMatrix produces a matrix which will transform by o
func (o Orientation) RotationMatrix() Matrix {
	orientationPhi := NewRotationMatrixY(o.Phi)
	orientationTheta := NewRotationMatrixZ(o.Theta)

	return orientationTheta.Multiply(orientationPhi)
}

// Cartesian returns a Cartesian of length one in the orientation of o
func (o Orientation) Cartesian() Cartesian {
	cosT, sinT := math.Sincos(o.Theta)
	cosP, sinP := math.Sincos(o.Phi)
	return Cartesian{
		X: cosT * sinP,
		Y: sinT * sinP,
		Z: cosP,
	}
}

// PortionOrtagonal returns the portion of o2 which is orthogonal to o
func (o Orientation) PortionOrtagonal(o2 Orientation) Orientation {
	v := o.Cartesian()
	u := o2.Cartesian()
	portUOrthoV := u.Translate(v.Project(u).Negative())
	return portUOrthoV.Orientation()
}
