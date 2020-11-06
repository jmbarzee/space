package space

// Vector is a point in 3D space
// Vector supports many operations that modify 3D points
type Vector interface {
	Cartesian() Cartesian
	Spherical() Spherical

	Translate(Vector) Vector
	Scale(float64) Vector
	Transform(Matrix) Vector

	Project(Vector) Vector
}
