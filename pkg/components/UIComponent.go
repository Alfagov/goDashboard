package components

import (
	"errors"
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/a-h/templ"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"go.uber.org/zap"
	"os"
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

// GetRouteFromParents constructs a string that represents the hierarchical path of a UI component by concatenating the names
// of its parent components. It stops once it reaches a component of the DashboardType or if the component has no parent.
// The resulting route is a '/' separated string reflecting the hierarchy from the top-level parent to the given component.
func GetRouteFromParents(n UIComponent) string {
	parent := n.GetParent()
	route := ""
	for {
		if parent == nil || parent.Type().Is(DashboardType) {
			break
		}
		route = parent.Name() + "/" + route
		parent = parent.GetParent()
	}

	return route
}

// RenderResponse is a struct that holds the outcome of a component's render operation.
// It encapsulates the rendered template component, any additional JSON data, and an error, if occurred.
type RenderResponse struct {
	Component templ.Component
	Json      map[string]interface{}
	Err       error
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

// nodeType defines a structure to represent the type and supertype of a UI component node.
// It facilitates type checks and hierarchical relationships within UI components.
type nodeType struct {
	superType string
	typeName  string
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

var (
	DashboardType     NodeType = &nodeType{superType: "dashboard", typeName: "dashboard"}
	PageType          NodeType = &nodeType{superType: "pages", typeName: "page"}
	PageContainerType NodeType = &nodeType{superType: "pages", typeName: "pageContainer"}
	NumericWidgetType NodeType = &nodeType{superType: "widgets", typeName: "numeric"}
	GraphWidgetType   NodeType = &nodeType{superType: "widgets", typeName: "graph"}
	FormWidgetType    NodeType = &nodeType{superType: "widgets", typeName: "form"}
)

func VisualizeTree(tree UIComponent) {
	treeChart := charts.NewTree()
	treeChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: "white"}),
		charts.WithTitleOpts(opts.Title{Title: "Tree-Visualize"}),
	)

	nodes := TreeSpecToChartNodes(
		[]*models.TreeSpec{
			tree.GetSpec(),
		})

	treeChart.AddSeries("tree", nodes)

	f, _ := os.Create("tree.html")
	err := treeChart.Render(f)
	if err != nil {
		logger.L.Error("error in rendering tree", zap.Error(err))
	}
}

func TreeSpecToChartNodes(spec []*models.TreeSpec) []opts.TreeData {
	var nodes []opts.TreeData

	for _, child := range spec {
		nodes = append(nodes, opts.TreeData{
			Name:     child.Name,
			Children: treeSpecToChartNodes(child.Children),
		})
	}

	return nodes
}

func treeSpecToChartNodes(spec []*models.TreeSpec) []*opts.TreeData {
	var nodes []*opts.TreeData
	for _, child := range spec {
		nodes = append(nodes, &opts.TreeData{
			Name:     child.Name,
			Children: treeSpecToChartNodes(child.Children),
		})
	}

	return nodes
}
