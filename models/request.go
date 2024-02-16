package models

import (
	"github.com/gofiber/fiber/v2"
)

// NewReqWrapper creates a new RequestWrapper that wraps the Fiber context for handling HTTP requests.
func NewReqWrapper(c *fiber.Ctx) RequestWrapper {
	return &requestWrapper{C: c}
}

// BindFormRequest binds the parsed form additionalData from the request into the provided struct pointer 'v'.
func (fr *requestWrapper) BindFormRequest(v interface{}) error {
	return fr.C.BodyParser(v)
}

func (fr *requestWrapper) AddAdditionalData(v [][]interface{}) {
	fr.additionalData = v
}

func (fr *requestWrapper) GetAdditionalData() [][]interface{} {
	if fr.additionalData == nil {
		return nil
	}
	return fr.additionalData.([][]interface{})
}

func (fr *requestWrapper) AddHeaders(key, value string) {
	fr.C.Set(key, value)
}

// Query retrieves the value of a query string parameter by key, with an optional default value if the key is not present.
func (fr *requestWrapper) Query(key string, def ...string) string {
	return fr.C.Query(key, def...)
}

// Method returns the HTTP method of the request, allowing handlers to adapt their behavior to the method.
func (fr *requestWrapper) Method() string {
	return fr.C.Method()
}

// Locals return the value of a local variable stored in the request context.
func (fr *requestWrapper) Locals(key string) interface{} {
	return fr.C.Locals(key)
}

// NewRequestWrapper creates a new RequestWrapper that wraps the Fiber context for handling HTTP requests.
func NewRequestWrapper(c *fiber.Ctx) RequestWrapper {
	return &requestWrapper{C: c}
}

// requestWrapper is an internal implementation that wraps a Fiber context to adhere to the RequestWrapper interface.
type requestWrapper struct {
	C              *fiber.Ctx
	additionalData interface{}
}

// RequestWrapper defines the interface for handling HTTP request additionalData.
type RequestWrapper interface {
	// BindFormRequest binds form additionalData from an HTTP request to the provided struct pointer.
	BindFormRequest(v interface{}) error

	// AddData adds additionalData to the request.
	AddAdditionalData(v [][]interface{})
	GetAdditionalData() [][]interface{}

	// AddHeaders adds headers to the request
	AddHeaders(key, value string)

	// Query returns the value of a query parameter given its key, with optional default values.
	Query(key string, def ...string) string

	// Method returns the HTTP method of the request.
	Method() string

	// Locals return the value of a local variable stored in the request context.
	Locals(key string) interface{}
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
