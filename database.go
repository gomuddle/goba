package goba

// A Database represents a collection of data.
type Database interface {
	// Type returns the Database's type.
	Type() DatabaseType

	// CreateImage returns an Image containing the Database's current state.
	CreateImage() (*Image, error)

	// ApplyImage applies the given Image to the Database.
	ApplyImage(image Image) error
}
