package models

import "github.com/gofiber/fiber/v2"

// NewReqWrapper creates a new RequestWrapper that wraps the Fiber context for handling HTTP requests.
func NewReqWrapper(c *fiber.Ctx) RequestWrapper {
	return &requestWrapper{c: c}
}

// BindFormRequest binds the parsed form data from the request into the provided struct pointer 'v'.
func (fr *requestWrapper) BindFormRequest(v interface{}) error {
	return fr.c.BodyParser(v)
}

// Query retrieves the value of a query string parameter by key, with an optional default value if the key is not present.
func (fr *requestWrapper) Query(key string, def ...string) string {
	return fr.c.Query(key, def...)
}

// Method returns the HTTP method of the request, allowing handlers to adapt their behavior to the method.
func (fr *requestWrapper) Method() string {
	return fr.c.Method()
}

// requestWrapper is an internal implementation that wraps a Fiber context to adhere to the RequestWrapper interface.
type requestWrapper struct {
	c *fiber.Ctx
}

// RequestWrapper defines the interface for handling HTTP request data.
type RequestWrapper interface {
	// BindFormRequest binds form data from an HTTP request to the provided struct pointer.
	BindFormRequest(v interface{}) error

	// Query returns the value of a query parameter given its key, with optional default values.
	Query(key string, def ...string) string

	// Method returns the HTTP method of the request.
	Method() string
}

// Field defines a structure for form fields, including their name, label, type, and the route where the field is to be used.
// It also contains handlers for selection and validation: SelectHandler provides options for selection-type fields based on a query,
// and ValidationHandler performs field-specific validation, returning true if the field's current value is valid.
type Field struct {
	Name              string
	Label             string
	Type              string
	Route             string
	SelectHandler     func(query string) []string
	ValidationHandler func() bool
}
