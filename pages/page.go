package pages

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

// page is a container of widgets
type page struct {
	name        string
	description string
	imagePath   string

	spec     *models.TreeSpec
	parent   components.UIComponent
	children []components.UIComponent
}

type Page interface {
	setName(name string)
	setImagePath(path string)
	GetImagePath() string
}

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

	return spec
}

func (p *page) GetSpec() *models.TreeSpec {
	return p.spec
}

func (p *page) Render(components.RequestWrapper) *components.RenderResponse {
	var componentList []templ.Component
	for _, child := range p.children {
		componentList = append(componentList, child.Render(nil).Component)
	}

	return &components.RenderResponse{
		Component: templates.GridPage(componentList),
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

func (p *page) SetParent(parent components.UIComponent) {
	p.parent = parent
}

func (p *page) GetParent() components.UIComponent {
	return p.parent
}

func (p *page) AddChild(child components.UIComponent) error {
	_, exists := p.FindChild(child.Name())
	if exists {
		return components.ErrChildExists(child.Name())
	}

	child.SetParent(p)
	p.children = append(p.children, child)

	return nil
}

func (p *page) KillChild(child components.UIComponent) error {
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

func (p *page) GetImagePath() string {
	return p.imagePath
}

func (p *page) setImagePath(path string) {
	p.imagePath = path
}

func (p *page) setName(name string) {
	p.name = name
}

// SETTERS

func SetName(name string) func(p Page) {
	return func(p Page) {
		p.setName(name)
	}
}

func SetImagePath(path string) func(p Page) {
	return func(p Page) {
		p.setImagePath(path)
	}
}
