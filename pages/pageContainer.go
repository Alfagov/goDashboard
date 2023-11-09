package pages

import (
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"go.uber.org/zap"
)

// pageContainer is a private struct that serves as a container for UI components
// that represent a container of pages.
type pageContainer struct {
	// id is the unique identifier for the page container.
	id string
	// name is the display name of the page container.
	name string
	// imagePath is the file path to the image associated with the page container.
	imagePath string
	// description provides a brief summary of the page container's content.
	description string
	// indexPage specifies the entry page's identifier within the container.
	indexPage string

	// spec holds the tree specification for the structure of the pages.
	spec *models.TreeSpec
	// parent is the UIComponent that acts as a parent to this container.
	parent components.UIComponent
	// children is a map of UIComponent children keyed by their string identifiers.
	children map[string]components.UIComponent
}

// PageContainer defines the interface for a container that groups together
// UI components representing pages, providing mechanisms to manage and retrieve
// these pages within a user interface.
type PageContainer interface {
	// UIComponent is a set of common methods for User Interface components.
	components.UIComponent

	// setImagePath sets the file path to the image associated with the page container.
	setImagePath(path string)

	// GetIndexPage retrieves the identifier of the entry page in the page container.
	GetIndexPage() string

	// SetIndexPage sets the identifier of the entry page in the page container.
	SetIndexPage(indexPage string)

	// GetImagePath retrieves the file path to the image associated with the page container.
	GetImagePath() string

	// WithPages adds the given UIComponent pages to the page container and returns
	// the PageContainer instance for method chaining.
	WithPages(pages ...components.UIComponent) PageContainer
}

// NewPageContainer creates and initializes a new PageContainer with the specified name.
// It applies a list of setter functions to configure the PageContainer.
// By default, the imagePath is set to a default location under "/static/img/"
// with the image name derived from the specified name parameter.
//
// Parameters:
// name: The name is used to identify the page container and to construct the default imagePath.
// setters: A variadic list of setter functions that apply additional configurations to the page container.
//
// Returns a pointer to the newly created page container with the applied settings.
func NewPageContainer(name string, setters ...func(pc PageContainer)) PageContainer {
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

func (pc *pageContainer) Render(req models.RequestWrapper) *components.RenderResponse {

	pageName := req.Locals("pageName")
	if pageName != nil {
		p, ok := pc.FindChildByType(pageName.(string), "page")
		if ok {
			return &components.RenderResponse{
				Component: PageContainerView(p.Render(nil).Component, pc.spec.Children),
			}
		}
	}

	return &components.RenderResponse{
		Component: PageContainerView(pc.children[pc.indexPage].Render(nil).Component, pc.spec.Children),
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

func (pc *pageContainer) Id() string {
	return pc.id
}

func (pc *pageContainer) FindChildById(id string) (components.UIComponent, bool) {
	for _, child := range pc.children {
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
	if !child.Type().Is(components.PageType) {
		logger.L.Error("PageContainer.WithPages: wrong type of page", zap.String("page", child.Name()),
			zap.String("type", child.Type().TypeName()))

		return components.ErrWrongChildType(child.Name(), components.PageType.TypeName(), child.Type().TypeName())
	}

	_, exists := pc.children[child.Name()]
	if exists {
		return components.ErrChildExists(child.Name())
	}

	child.SetParent(pc)
	pc.children[child.Name()] = child

	return nil
}

func (pc *pageContainer) RemoveChild(child components.UIComponent) error {
	_, exists := pc.children[child.Name()]
	if !exists {
		return components.ErrChildNotFound(child.Name())
	}

	delete(pc.children, child.Name())

	return nil
}

// setImagePath sets the image path for the page container.
// This method is unexported and intended for internal use within the package.
func (pc *pageContainer) setImagePath(path string) {
	pc.imagePath = path
}

// GetImagePath retrieves the image path of the page container.
func (pc *pageContainer) GetImagePath() string {
	return pc.imagePath
}

// SetIndexPage defines the index page for the page container.
func (pc *pageContainer) SetIndexPage(indexPage string) {
	pc.indexPage = indexPage
}

// GetIndexPage retrieves the identifier of the index page of the page container.
func (pc *pageContainer) GetIndexPage() string {
	return pc.indexPage
}

// WithPages adds multiple UIComponent pages to the page container.
// This method can be used to chain the addition of pages to the container.
func (pc *pageContainer) WithPages(pages ...components.UIComponent) PageContainer {
	for _, pg := range pages {
		err := pc.AddChild(pg)
		logger.L.Error("error adding page to page container", zap.Error(err))
	}

	return pc
}

// Setters

// SetIndexPage returns a function that sets the index page of the page container.
// This can be used as an argument to functions that accept a configuration function for PageContainer.
func SetIndexPage(indexPage string) func(p PageContainer) {
	return func(p PageContainer) {
		p.SetIndexPage(indexPage)
	}
}

// SetContainerImagePath returns a function that sets the image path for the page container.
// This can be used as an argument to functions that accept a configuration function for PageContainer.
func SetContainerImagePath(path string) func(p PageContainer) {
	return func(p PageContainer) {
		p.setImagePath(path) // Note: This assumes the setImagePath method is exported or adjusted accordingly.
	}
}
