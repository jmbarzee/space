package space

// An Object is something which exists in space
type Object struct {
	// location is the location of the
	location Cartesian
	// Spherical is primary Spherical of the object
	Spherical Spherical
	// rotation is secondary Spherical of the object
	// rotation will always be orthogonal to direction
	rotation Spherical
}

// NewObject creates an object
func NewObject(location Cartesian, Spherical, rotation Spherical) *Object {
	normalizedRotation := Spherical.PortionOrtagonal(rotation)
	return &Object{
		location:  location,
		Spherical: Spherical,
		rotation:  normalizedRotation,
	}
}

// GetLocation returns the physical location of the device
func (o Object) GetLocation() Cartesian {
	return o.location
}

// SetLocation changes the physical location of the device
func (o *Object) SetLocation(newLocation Cartesian) {
	o.location = newLocation
}

// GetSpherical returns the physical Spherical of the device
func (o Object) GetSpherical() Spherical {
	return o.Spherical
}

// SetSpherical changes the physical Spherical of the device
func (o *Object) SetSpherical(newSpherical Spherical) {
	o.Spherical = newSpherical
	o.rotation = newSpherical.PortionOrtagonal(o.rotation)
}

// GetRotation returns the physical rotation of the device
func (o Object) GetRotation() Spherical {
	return o.rotation
}

// SetRotation changes the physical rotation of the device
func (o *Object) SetRotation(newSpherical Spherical) {
	o.rotation = o.Spherical.PortionOrtagonal(newSpherical)
}
