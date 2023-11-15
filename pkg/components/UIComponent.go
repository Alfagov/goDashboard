package components

import (
	"github.com/Alfagov/goDashboard/models"
)

// UIComponent defines an interface for UI components within an application.
// It provides methods to manage component rendering, hierarchy, identification,
// and specification within a UI tree structure.
type UIComponent interface {
	// Render generates a RenderResponse based on the current state of the component
	// and the provided request context. This method is responsible for creating
	// a representation of the component that can be sent to the client for rendering.
	Render(req models.RequestWrapper) *RenderResponse

	// Type returns the NodeType of the component, indicating its specific type within
	// the UI framework or application logic.
	Type() NodeType

	// Name returns the name of the component, often used as a human-readable identifier.
	Name() string

	// UpdateSpec is used to update the component's specification in place and returns
	// the updated TreeSpec.
	UpdateSpec() *models.TreeSpec

	// GetSpec returns the current TreeSpec of the component, representing its
	// configuration and hierarchical relationships within the UI.
	GetSpec() *models.TreeSpec

	// GetChildren retrieves a slice of UIComponent instances that are children
	// of this component, representing the hierarchical structure of the UI.
	GetChildren() []UIComponent

	// FindChild searches for a child with the specified name and returns it along
	// with a boolean indicating if the child was found.
	FindChild(name string) (UIComponent, bool)

	// FindChildByType searches for a child with the specified name and componentType,
	// returning the child and a boolean indicating if such a child was found.
	FindChildByType(name string, componentType string) (UIComponent, bool)

	// GetParent returns the parent component of this component, allowing for
	// navigation up the UI component tree.
	GetParent() UIComponent

	// Id returns the unique identifier of the component, used to distinguish it
	// within the component tree.
	Id() string

	// FindChildById looks for a child with the given id within the component's
	// descendants and returns the child and a boolean indicating success.
	FindChildById(id string) (UIComponent, bool)

	// SetParent associates a new parent UIComponent with this component, effectively
	// repositioning it within the component tree.
	SetParent(parent UIComponent)

	// AddChild incorporates a new child UIComponent to this component's children,
	// potentially returning an error if the addition is not possible.
	AddChild(child UIComponent) error

	// RemoveChild removes a child UIComponent from this component's children.
	// This method is named whimsically and the action should be taken seriously
	// as it alters the component tree.
	RemoveChild(child UIComponent) error
}

// NodeType is an interface that defines the behavior for node type classification.
// It allows for type identification and comparison.
//
//goland:noinspection Annotator
type NodeType interface {
	// SuperType returns the name of the super type of the current node type
	SuperType() string

	// TypeName returns the unique name of the current node type
	TypeName() string

	// Is checks if the provided NodeType is equivalent to the current node type
	Is(cmp NodeType) bool

	// IsType determines if the current node type matches a specified type name.
	IsType(typeName string) bool
}

// SuperType returns the supertype of the node, allowing for categorization of node types.
func (nt *nodeType) SuperType() string {
	return nt.superType
}

// TypeName returns the specific type name of the node, which is used to identify the node type more precisely.
func (nt *nodeType) TypeName() string {
	return nt.typeName
}

// Is checks if the node's type and supertype match those of another NodeType, enabling comparison between node types.
func (nt *nodeType) Is(cmp NodeType) bool {
	return nt.superType == cmp.SuperType() && nt.typeName == cmp.TypeName()
}

// IsSuperType checks if the node's supertype matches a specified supertype, useful for asserting the node's general category.
func (nt *nodeType) IsSuperType(superType string) bool {
	return nt.superType == superType
}

// IsType checks if the node's type name matches a specified type name, allowing for specific node type assertions.
func (nt *nodeType) IsType(typeName string) bool {
	return nt.typeName == typeName
}
