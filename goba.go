package goba

import "errors"

// ErrNoSuchHandler occurs whenever Goba could
// not find a handler for some DatabaseType.
var ErrNoSuchHandler = errors.New("no such handler")

// Goba manages a collection of image handlers and
// acts as the main interface of the application.
type Goba struct {
	handlers []ImageHandler
}

// New creates a new Goba managing the given handlers.
func New(handlers ...ImageHandler) *Goba {
	return &Goba{handlers: handlers}
}

// CreateImage creates a new Image of the Database with the given type.
func (g Goba) CreateImage(typ DatabaseType) (*Image, error) {
	for _, handler := range g.handlers {
		if handler.Type() == typ {
			return handler.CreateImage()
		}
	}
	return nil, ErrNoSuchHandler
}

// ApplyImage applies the Image with the given
// name to the Database with the given type.
func (g Goba) ApplyImage(typ DatabaseType, name string) error {
	for _, handler := range g.handlers {
		if handler.Type() == typ {
			return handler.ApplyImage(name)
		}
	}
	return ErrNoSuchHandler
}

// FindImage retrieves the Image with the given type and name.
func (g Goba) FindImage(typ DatabaseType, name string) (*Image, error) {
	for _, handler := range g.handlers {
		if handler.Type() == typ {
			return handler.FindImage(name)
		}
	}
	return nil, ErrNoSuchHandler
}

// DeleteImage deletes the Image with the given type and name.
func (g Goba) DeleteImage(typ DatabaseType, name string) error {
	for _, handler := range g.handlers {
		if handler.Type() == typ {
			return handler.DeleteImage(name)
		}
	}
	return ErrNoSuchHandler
}
