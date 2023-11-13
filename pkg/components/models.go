package components

import (
	"errors"
	"github.com/a-h/templ"
)

var (
	// ErrChildExists returns an error indicating that a child with the specified name already exists within a parent component.
	// It is used to enforce unique child names within a component.
	ErrChildExists = func(name string) error {
		return errors.New("child with name " + name + " already exists")
	}

	// ErrChildNotFound returns an error when an attempt is made to access a child component by name that does not exist.
	ErrChildNotFound = func(name string) error {
		return errors.New("child with name " + name + " not found")
	}

	// ErrWrongChildType returns an error indicating that a child component's type does not match the expected type.
	// It includes both the expected type and the actual type of the component for clear debugging information.
	ErrWrongChildType = func(name string, expectedType, componentType string) error {
		return errors.New("child with name " + name + " is not of type " + expectedType + " but of type " + componentType)
	}

	// ErrCannotHaveChildren returns an error stating that a component of a specified type cannot have children.
	// This is used to enforce component composition rules.
	ErrCannotHaveChildren = func(typeName string) error {
		return errors.New("component of type " + typeName + " cannot have children")
	}

	/* TYPE DEFINITIONS FOR UI COMPONENTS */

	DashboardType     NodeType = &nodeType{superType: "dashboard", typeName: "dashboard"}
	PageType          NodeType = &nodeType{superType: "pages", typeName: "page"}
	PageContainerType NodeType = &nodeType{superType: "pages", typeName: "pageContainer"}
	NumericWidgetType NodeType = &nodeType{superType: "widgets", typeName: "numeric"}
	GraphWidgetType   NodeType = &nodeType{superType: "widgets", typeName: "graph"}
	FormWidgetType    NodeType = &nodeType{superType: "widgets", typeName: "form"}
)

// RenderResponse is a struct that holds the outcome of a component's render operation.
// It encapsulates the rendered template component, any additional JSON data, and an error, if occurred.
type RenderResponse struct {
	Component templ.Component
	Json      map[string]interface{}
	Err       error
}

func NewRenderResponse(component templ.Component, json map[string]interface{}, err error) *RenderResponse {
	return &RenderResponse{
		Component: component,
		Json:      json,
		Err:       err,
	}
}

// nodeType defines a structure to represent the type and supertype of a UI component node.
// It facilitates type checks and hierarchical relationships within UI components.
type nodeType struct {
	superType string
	typeName  string
}
