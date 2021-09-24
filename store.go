package goba

// A Store manages database images.
type Store interface {
	// SaveImage saves the given Image. If an Image with the
	// given Image's name already exists, SaveImage should
	// return a non-nil error.
	SaveImage(image Image) error

	// FindImage retrieves an Image with the given name.
	// If no such Image exists, FindImage should return
	// a non-nil error.
	FindImage(name string) (*Image, error)

	// DeleteImage deletes the Image with the given name.
	// If no such Image exists, DeleteImage should be a no-op.
	DeleteImage(name string) error
}
