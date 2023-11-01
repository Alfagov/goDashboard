package pages

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
)

// PageContainer is a container for pages
type pageContainer struct {
	id          string
	name        string
	imagePath   string
	description string
	indexPage   string

	spec     *models.TreeSpec
	parent   components.UIComponent
	children map[string]components.UIComponent
}

type PageContainer interface {
	components.UIComponent
	setImagePath(path string)
	GetIndexPage() string
	SetIndexPage(indexPage string)
	GetImagePath() string
	WithPages(pages ...components.UIComponent) PageContainer
}

func NewPageContainer(
	name string, setters ...func(
		pc PageContainer,
	),
) PageContainer {
	var pc pageContainer
	pc.name = name
	pc.imagePath = "/static/img/" + name + ".png"
	pc.children = make(map[string]components.UIComponent)

	for _, setter := range setters {
		setter(&pc)
	}

	return &pc
}

// UIComponent interface implementation

func (pc *pageContainer) UpdateSpec() *models.TreeSpec {

	route := pc.parent.Name() + "/" + pc.name

	var childrenSpec []*models.TreeSpec
	for _, child := range pc.children {
		childrenSpec = append(childrenSpec, child.UpdateSpec())
	}

	spec := &models.TreeSpec{
		Name:        pc.name,
		ImageRoute:  pc.imagePath,
		Description: pc.description,
		Route:       route,
		Children:    childrenSpec,
	}

	pc.spec = spec

	return spec
}

func (pc *pageContainer) GetSpec() *models.TreeSpec {
	return pc.spec
}

func (pc *pageContainer) Render(components.RequestWrapper) *components.RenderResponse {
	return &components.RenderResponse{
		Component: templates.PageContainer(pc.children[pc.indexPage].Render(nil).Component, pc.spec.Children),
	}
}

func (pc *pageContainer) Type() components.NodeType {
	return components.PageContainerType
}

func (pc *pageContainer) Name() string {
	return pc.name
}

func (pc *pageContainer) GetChildren() []components.UIComponent {
	var t []components.UIComponent
	for _, child := range pc.children {
		t = append(t, child)
	}

	return t
}

func (pc *pageContainer) FindChild(name string) (components.UIComponent, bool) {
	child, ok := pc.children[name]
	return child, ok
}

func (p *pageContainer) Id() string {
	return p.id
}

func (p *pageContainer) FindChildById(id string) (components.UIComponent, bool) {
	for _, child := range p.children {
		if child.Id() == id {
			return child, true
		}
	}

	return nil, false
}

func (pc *pageContainer) FindChildByType(name string, componentType string) (components.UIComponent, bool) {
	child, ok := pc.children[name]
	if !ok {
		return nil, false
	}

	if child.Type().IsType(componentType) {
		return child, true
	}

	return nil, false
}

func (pc *pageContainer) SetParent(parent components.UIComponent) {
	pc.parent = parent
}

func (pc *pageContainer) GetParent() components.UIComponent {
	return pc.parent
}

func (pc *pageContainer) AddChild(child components.UIComponent) error {
	_, exists := pc.children[child.Name()]
	if exists {
		return components.ErrChildExists(child.Name())
	}

	child.SetParent(pc)
	pc.children[child.Name()] = child

	return nil
}

func (pc *pageContainer) KillChild(child components.UIComponent) error {
	_, exists := pc.children[child.Name()]
	if !exists {
		return components.ErrChildNotFound(child.Name())
	}

	delete(pc.children, child.Name())

	return nil
}

// PageContainer interface implementation

func (pc *pageContainer) setImagePath(path string) {
	pc.imagePath = path
}

func (pc *pageContainer) GetImagePath() string {
	return pc.imagePath
}

func (pc *pageContainer) SetIndexPage(indexPage string) {
	pc.indexPage = indexPage
}

func (pc *pageContainer) GetIndexPage() string {
	return pc.indexPage
}

func (pc *pageContainer) WithPages(pages ...components.UIComponent) PageContainer {
	for _, pg := range pages {
		pc.AddChild(pg)
	}

	return pc
}

// Setters

func SetIndexPage(indexPage string) func(p PageContainer) {
	return func(p PageContainer) {
		p.SetIndexPage(indexPage)
	}
}

func SetContainerImagePath(path string) func(p PageContainer) {
	return func(p PageContainer) {
		p.setImagePath(path)
	}
}
