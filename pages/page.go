package pages

import (
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/a-h/templ"
	"go.uber.org/zap"
)

// Page represents a page in a website or an application.
// Each page has a unique ID, a name, a description and an image path.
// It also contains a specification of it and children in the form of a TreeSpec, a parent UI component
// and a set of child UI Components.
type page struct {
	// id is a unique identifier for the page.
	id string

	// name is the name of the page that might be displayed in the user interface.
	name string

	// description provides more details about what the page is intended for.
	description string

	// imagePath provides the path to the image that represents or is associated with the page.
	imagePath string

	// spec represents the specification of the page in the form of a Tree structure.
	spec *models.TreeSpec

	// parent represents the UIComponent which is the parent of this page.
	parent components.UIComponent

	// children is a slice of UIComponents that are children of this page in the UI hierarchy.
	children []components.UIComponent
}

// Page is an interface that represents a construct for a webpage screen.
// It extends the UIComponent interface.
// It includes methods for setting name, setting image path,
// retrieving image path and adding widgets.
type Page interface {
	// UIComponent is a set of common methods for User Interface components.
	components.UIComponent

	// setName sets the name of the Page.
	setName(name string)

	// setImagePath sets the path of the image that represents or is associated with the Page.
	setImagePath(path string)

	// GetImagePath retrieves the image path associated with the Page.
	GetImagePath() string

	// WithWidgets adds a set of UIComponent as widgets to the Page.
	WithWidgets(widgets ...components.UIComponent) Page
}

// NewPage creates a new page with the specified name and applies optional configuration functions.
// The page's default image path is constructed using the given name and prefixed with "/static/img/".
// Additional configurations can be applied to the page by passing setter functions that match the Page signature.
//
// Parameters:
// name: the name of the page, which is used to set the page's name and to construct the default image path.
// setters: a variadic list of functions that take a Page pointer and return nothing, used to apply additional configurations to the page.
//
// Returns a pointer to the newly created and configured page.
func NewPage(name string, setters ...func(p Page)) Page {
	var p page
	p.name = name
	p.children = []components.UIComponent{}
	p.imagePath = "/static/img/" + name + ".png"

	for _, setter := range setters {
		setter(&p)
	}

	return &p
}

// UIComponent interface implementation

func (p *page) UpdateSpec() *models.TreeSpec {

	route := components.GetRouteFromParents(p)

	var childrenSpec []*models.TreeSpec
	for _, child := range p.children {
		childrenSpec = append(childrenSpec, child.UpdateSpec())
	}

	spec := &models.TreeSpec{
		Name:        p.name,
		ImageRoute:  p.imagePath,
		Description: p.description,
		Route:       route + p.name,
		Children:    childrenSpec,
	}

	p.spec = spec

	return spec
}

func (p *page) GetSpec() *models.TreeSpec {
	return p.spec
}

func (p *page) Render(models.RequestWrapper) *components.RenderResponse {
	var componentList []templ.Component
	for _, child := range p.children {
		componentList = append(componentList, child.Render(nil).Component)
	}

	return &components.RenderResponse{
		Component: GridPage(componentList),
	}
}

func (p *page) Type() components.NodeType {
	return components.PageType
}

func (p *page) Name() string {
	return p.name
}

func (p *page) GetChildren() []components.UIComponent {
	t := make([]components.UIComponent, 0)
	for _, child := range p.children {
		t = append(t, child)
	}

	return t
}

func (p *page) FindChild(name string) (components.UIComponent, bool) {
	for _, child := range p.children {
		if child.Name() == name {
			return child, true
		}
	}

	return nil, false
}

func (p *page) FindChildByType(name string, componentType string) (components.UIComponent, bool) {
	for _, child := range p.children {
		if child.Name() == name && child.Type().IsType(componentType) {
			return child, true
		}
	}

	return nil, false
}

func (p *page) Id() string {
	return p.id
}

func (p *page) FindChildById(id string) (components.UIComponent, bool) {
	for _, child := range p.children {
		if child.Id() == id {
			return child, true
		}
	}

	return nil, false
}

func (p *page) SetParent(parent components.UIComponent) {
	p.parent = parent
}

func (p *page) GetParent() components.UIComponent {
	return p.parent
}

func (p *page) AddChild(child components.UIComponent) error {

	if !(child.Type().SuperType() == "widgets") {
		logger.L.Error("Page: wrong type of child", zap.String("name", child.Name()),
			zap.String("type", child.Type().TypeName()))

		return components.ErrWrongChildType(child.Name(), components.PageType.TypeName(), child.Type().TypeName())
	}

	_, exists := p.FindChild(child.Name())
	if exists {
		return components.ErrChildExists(child.Name())
	}

	child.SetParent(p)
	p.children = append(p.children, child)

	return nil
}

func (p *page) RemoveChild(child components.UIComponent) error {
	_, exists := p.FindChild(child.Name())
	if !exists {
		return components.ErrChildNotFound(child.Name())
	}

	for i, c := range p.children {
		if c.Name() == child.Name() {
			p.children = append(p.children[:i], p.children[i+1:]...)
		}
	}

	return nil
}

// Page interface implementation

// GetImagePath returns the image path associated with the page.
func (p *page) GetImagePath() string {
	return p.imagePath
}

// setImagePath sets the image path for the page.
// This method is unexported and can only be used within the package.
func (p *page) setImagePath(path string) {
	p.imagePath = path
}

// setName sets the name of the page.
// This method is unexported and can only be used within the package.
func (p *page) setName(name string) {
	p.name = name
}

// WithWidgets adds UIComponent widgets to the page if they are of super type "widgets".
// It panics if a component with a different super type is added.
// This is used to chain the addition of multiple widgets to a page.
func (p *page) WithWidgets(widgets ...components.UIComponent) Page {
	for _, widget := range widgets {
		if widget.Type().SuperType() != "widgets" {
			panic("cannot add non-widget component to page")
		}
		err := p.AddChild(widget)
		logger.L.Error("Page.WithWidgets: error adding widget to page", zap.Error(err))
	}

	return p
}

// SETTERS

// SetName returns a function that sets the name of the page.
// This function can be passed to other functions that accept a configuration function for a Page.
func SetName(name string) func(p Page) {
	return func(p Page) {
		p.setName(name)
	}
}

// SetImagePath returns a function that sets the image path of the page.
// This function can be passed to other functions that accept a configuration function for a Page.
func SetImagePath(path string) func(p Page) {
	return func(p Page) {
		p.setImagePath(path)
	}
}
