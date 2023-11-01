package dashboard

import (
	_ "embed"
	"github.com/Alfagov/goDashboard/config"
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"go.uber.org/zap"
)

// Dashboard interface implementation

func (d *dashboard) Run() error {

	d.CreateStaticRoutes()
	d.CreateRoutes()

	dashboardUrl := config.C.DashboardConfig.Host + ":" + config.C.DashboardConfig.Port

	if config.C.DashboardConfig.SSL {
		return d.Router.ListenTLS(dashboardUrl, config.C.DashboardConfig.CertPath, config.C.DashboardConfig.KeyPath)
	}

	d.treeSpec = d.UpdateSpec()

	components.VisualizeTree(d)

	return d.Router.Listen(dashboardUrl)
}

func (d *dashboard) WithPages(pages ...components.UIComponent) Dashboard {
	for _, page := range pages {
		d.AddChild(page)
	}

	return d
}

// UIComponent interface implementation

func (d *dashboard) Render(components.RequestWrapper) *components.RenderResponse {
	return &components.RenderResponse{
		Component: templates.ListGridPage(d.treeSpec.Children),
	}
}

func (d *dashboard) Type() components.NodeType {
	return components.DashboardType
}

func (d *dashboard) Name() string {
	return d.name
}

func (d *dashboard) GetChildren() []components.UIComponent {
	t := make([]components.UIComponent, 0)
	for _, child := range d.Children {
		t = append(t, child)
	}

	return t
}

func (d *dashboard) FindChild(name string) (components.UIComponent, bool) {
	child, ok := d.Children[name]
	return child, ok
}

func (d *dashboard) FindChildByType(name string, componentType string) (components.UIComponent, bool) {
	child, ok := d.Children[name]
	if !ok {
		return nil, false
	}

	if child.Type().IsType(componentType) {
		return nil, false
	}

	return child, true
}

func (d *dashboard) Id() string {
	return d.id
}

func (d *dashboard) FindChildById(id string) (components.UIComponent, bool) {
	for _, child := range d.Children {
		if child.Id() == id {
			return child, true
		}
	}

	return nil, false
}

func (d *dashboard) SetParent(components.UIComponent) {}

func (d *dashboard) GetParent() components.UIComponent {
	return nil
}

func (d *dashboard) AddChild(child components.UIComponent) error {

	if !(child.Type().SuperType() == "pages") {
		logger.L.Error("Dashboard: wrong type of child", zap.String("name", child.Name()),
			zap.String("type", child.Type().TypeName()))

		return components.ErrWrongChildType(child.Name(), components.PageType.TypeName(), child.Type().TypeName())
	}

	_, exists := d.Children[child.Name()]
	if exists {
		return components.ErrChildExists(child.Name())
	}

	child.SetParent(d)
	d.Children[child.Name()] = child

	return nil
}

func (d *dashboard) KillChild(child components.UIComponent) error {
	_, exists := d.Children[child.Name()]
	if !exists {
		return components.ErrChildNotFound(child.Name())
	}

	delete(d.Children, child.Name())

	return nil
}

func (d *dashboard) UpdateSpec() *models.TreeSpec {

	var dashboardSpec []*models.TreeSpec
	for _, child := range d.Children {
		dashboardSpec = append(dashboardSpec, child.UpdateSpec())
	}

	spec := &models.TreeSpec{
		Name:        d.name,
		ImageRoute:  d.image,
		Description: d.description,
		Route:       d.name,
		Children:    dashboardSpec,
	}

	d.treeSpec = spec

	return spec
}

func (d *dashboard) GetSpec() *models.TreeSpec {
	return d.treeSpec
}
