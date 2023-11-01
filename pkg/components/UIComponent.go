package components

import (
	"errors"
	"github.com/Alfagov/goDashboard/models"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrChildExists = func(name string) error {
		return errors.New("child with name " + name + " already exists")
	}

	ErrChildNotFound = func(name string) error {
		return errors.New("child with name " + name + " not found")
	}
)

type UIComponent interface {
	// Render templ.Component returns a component that will be rendered by the client
	Render(req RequestWrapper) *RenderResponse

	// Type returns the type of the component
	Type() NodeType
	// Name returns the name of the component
	Name() string
	// UpdateSpec returns the spec of the component
	UpdateSpec() *models.TreeSpec
	// GetSpec returns the spec of the component
	GetSpec() *models.TreeSpec

	// GetChildren returns the children of the component
	GetChildren() []UIComponent
	// FindChild returns the child with the given name
	FindChild(name string) (UIComponent, bool)
	// FindChildByType returns the child with the given name and type
	FindChildByType(name string, componentType string) (UIComponent, bool)
	// GetParent returns a pointer to the parent of the component
	GetParent() UIComponent

	// SetParent sets the parent of the component
	SetParent(parent UIComponent)
	// AddChild adds a child to the component
	AddChild(child UIComponent) error
	// KillChild removes a child from the component :-(
	KillChild(child UIComponent) error
}

type RequestWrapper interface {
	BindFormRequest(v interface{}) error
	Method() string
}

type requestWrapper struct {
	c *fiber.Ctx
}

func NewReqWrapper(c *fiber.Ctx) RequestWrapper {
	return &requestWrapper{c: c}
}

func (fr *requestWrapper) BindFormRequest(v interface{}) error {
	return fr.c.BodyParser(v)
}

func (fr *requestWrapper) Method() string {
	return fr.c.Method()
}

func GetRouteFromParents(n UIComponent) string {
	parent := n.GetParent()
	route := ""
	for {
		if parent == nil || parent.Type().Is(DashboardType) {
			break
		}
		route = route + parent.Name() + "/"
		parent = parent.GetParent()
	}

	return route
}

type RenderResponse struct {
	Component templ.Component
	Json      map[string]interface{}
	Err       error
}

type DashboardTypeString string
type PageTypeString string
type WidgetTypeString string

type nodeType struct {
	superType string
	typeName  string
}

func (nt *nodeType) SuperType() string {
	return nt.superType
}

func (nt *nodeType) TypeName() string {
	return nt.typeName
}

func (nt *nodeType) Is(cmp NodeType) bool {
	return nt.superType == cmp.SuperType() && nt.typeName == cmp.TypeName()
}

func (nt *nodeType) IsSuperType(superType string) bool {
	return nt.superType == superType
}

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

type NodeType interface {
	SuperType() string
	TypeName() string
	Is(cmp NodeType) bool
	IsType(typeName string) bool
}
