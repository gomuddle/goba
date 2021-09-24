package goba

// An Image represents the current state of some Database.
type Image struct {
	// Type is the type of the Database that created the Image.
	Type DatabaseType `json:"type,omitempty"`

	// Name is the name of the Image.
	Name string `json:"name,omitempty"`

	// Content represents the Database's state.
	Content []byte `json:"content,omitempty"`
}
