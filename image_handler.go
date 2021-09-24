package goba

// An ImageHandler combines a Database with a Store to ensure
// that any images created by that Database are submitted to
// the store and that an Image can be loaded into the Database
// by name.
type ImageHandler struct {
	DB    Database
	Store Store
}

// Type returns the type of h's Database.
func (h ImageHandler) Type() DatabaseType {
	return h.DB.Type()
}

// CreateImage creates a new Image of h's Database,
// submits it to h's Store and returns it.
func (h ImageHandler) CreateImage() (*Image, error) {
	image, err := h.DB.CreateImage()
	if err != nil {
		return nil, err
	}
	return image, h.Store.SaveImage(*image)
}

// ApplyImage retrieves the Image with the given name
// from h's Store and applies it to h's Database.
func (h ImageHandler) ApplyImage(name string) error {
	image, err := h.Store.FindImage(name)
	if err != nil {
		return err
	}
	return h.DB.ApplyImage(*image)
}

// FindImage retrieves the Image with given name from h's Store.
func (h ImageHandler) FindImage(name string) (*Image, error) {
	return h.Store.FindImage(name)
}

// DeleteImage deletes the Image with the given name from h's Store.
func (h ImageHandler) DeleteImage(name string) error {
	return h.Store.DeleteImage(name)
}
