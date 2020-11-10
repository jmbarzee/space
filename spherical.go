package space

import (
	"fmt"
	"math"
)

// Spherical represents a point in space in spherical coordinates
type Spherical struct {
	// R is distance from the origin
	R float64
	// T is rotation about Z
	T float64
	// P is tilt from Z
	P float64
}

var _ Vector = (*Spherical)(nil)

// NewSpherical creates a new Spherical from a rotation and tilt
func NewSpherical(radius, theta, phi float64) Spherical {
	s := Spherical{
		R: radius,
	}
	s = s.Rotate(theta)
	s = s.Tilt(phi)
	return s
}

// Cartesian returns the Cartesian version of s
func (s Spherical) Cartesian() Cartesian {
	sinT, cosT := math.Sincos(s.T)
	sinP, cosP := math.Sincos(s.P)
	radius := s.R
	return Cartesian{
		X: radius * sinP * cosT,
		Y: radius * sinP * sinT,
		Z: radius * cosP,
	}
}

// Spherical return s
func (s Spherical) Spherical() Spherical {
	return s
}

// Translate shifts a Spherical by a Vector (addition in cartesian space)
func (s Spherical) Translate(v Vector) Vector {
	c := s.Cartesian()
	return c.Translate(v)
}

// Scale scales a Spherical by i
func (s Spherical) Scale(i float64) Vector {
	s.R *= i
	return s
}

// Transform Multiplyiplies a Spherical by a given matrix
func (s Spherical) Transform(m Matrix) Vector {
	c := s.Cartesian()
	return c.Transform(m)
}

// Project returns the projection of v onto s
func (s Spherical) Project(v Vector) Vector {
	c := s.Cartesian()
	return c.Project(v)
}

// Rotate will adjust the rotation about Z by theta
func (s Spherical) Rotate(theta float64) Spherical {
	wrappedT := s.T + theta
	unwrappedT := math.Mod(wrappedT, math.Pi*2)
	if unwrappedT >= 0 {
		s.T = unwrappedT
	} else {
		s.T = 2*math.Pi + unwrappedT
	}
	return s
}

// Tilt will adjust the tilt from Z by phi
func (s Spherical) Tilt(phi float64) Spherical {
	wrappedP := s.P + phi
	newP := math.Mod(wrappedP, math.Pi*2)

	// Check if tilt is negative
	if newP < 0 {
		newP = (math.Pi * 2) + newP
	}

	// Check if tilt is beyond range of [0, pi]
	if newP > math.Pi {
		s.P = (math.Pi * 2) - newP
		return s.Rotate(math.Pi)
	}

	s.P = newP
	return s
}

// PortionOrtagonal returns the portion of o2 which is orthogonal to o
func (s Spherical) PortionOrtagonal(o2 Spherical) Spherical {
	v := s.Cartesian()
	u := o2.Cartesian()
	fmt.Println("v", v)
	fmt.Println("u", u)
	vProju := v.Project(u)
	fmt.Println("vProju", vProju)
	vProjuN := vProju.Scale(-1.0)
	fmt.Println("vProjuN", vProjuN)
	portUOrthoV := u.Translate(vProjuN)
	fmt.Println("portUOrthoV", portUOrthoV)
	sp := portUOrthoV.Spherical()
	fmt.Println("sp", sp)
	return sp
}

// RotationMatrix produces a matrix which will rotate based on the spherical coordinates
func (s Spherical) RotationMatrix() Matrix {
	rotateZbyT := NewRotationMatrixZ(s.T)
	rotateYbyP := NewRotationMatrixY(s.P)
	rotateZbyTinverse := NewRotationMatrixZ(-s.T)

	return rotateZbyT.Multiply(rotateYbyP).Multiply(rotateZbyTinverse)
}

func (s Spherical) String() string {
	return fmt.Sprintf("{R:%4.2f, T:%4.2f, P:%4.2f}", s.R, s.T, s.P)
}
