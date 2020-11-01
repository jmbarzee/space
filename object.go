package space

// An Object is something which exists in space
type Object struct {
	// location is the location of the
	location Vector
	// orientation is primary orientation of the object
	orientation Orientation
	// rotation is secondary orientation of the object
	// rotation will always be orthogonal to direction
	rotation Orientation
}

// NewObject creates an object
func NewObject(location Vector, orientation, rotation Orientation) *Object {
	normalizedRotation := orientation.PortionOrtagonal(rotation)
	return &Object{
		location:    location,
		orientation: orientation,
		rotation:    normalizedRotation,
	}
}

// GetLocation returns the physical location of the device
func (o Object) GetLocation() Vector {
	return o.location
}

// SetLocation changes the physical location of the device
func (o *Object) SetLocation(newLocation Vector) {
	o.location = newLocation
}

// GetOrientation returns the physical orientation of the device
func (o Object) GetOrientation() Orientation {
	return o.orientation
}

// SetOrientation changes the physical orientation of the device
func (o *Object) SetOrientation(newOrientation Orientation) {
	o.orientation = newOrientation
	o.rotation = newOrientation.PortionOrtagonal(o.rotation)
}

// GetRotation returns the physical rotation of the device
func (o Object) GetRotation() Orientation {
	return o.rotation
}

// SetRotation changes the physical rotation of the device
func (o *Object) SetRotation(newOrientation Orientation) {
	o.rotation = o.orientation.PortionOrtagonal(newOrientation)
}
