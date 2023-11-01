package graphContainer

import (
	"errors"
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/graph"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/templates"
)

// GraphWidget represents an interface for a widget in the graph UI.
type GraphWidget interface {

	// update method returns a map representing the state of the GraphWidget.
	update() map[string]interface{}

	// WithSettings method is used to update the GraphWidget's settings.
	// It accepts a variable number of functions that mutate the GraphWidget.
	WithSettings(settings ...func(gw GraphWidget)) GraphWidget
}

type graphWidgetImpl struct {
	baseWidget  widgets.Widget
	parent      components.UIComponent
	description string
	specs       *models.TreeSpec
	htmxOpts    htmx.HTMX
	graph       graph.Graph
}

func (g *graphWidgetImpl) Render(req components.RequestWrapper) *components.RenderResponse {
	if req.Method() == "POST" {
		return &components.RenderResponse{
			Json: g.update(),
		}
	}

	return &components.RenderResponse{
		Component: templates.GeneralGraph(
			g.baseWidget.GetId(),
			g.graph.Encode(g.baseWidget.GetLayout().Height),
			g.baseWidget.GetLayout(),
			g.htmxOpts.GetHtmx(),
		),
	}
}

func (g *graphWidgetImpl) Type() components.NodeType {
	return components.GraphWidgetType
}

func (g *graphWidgetImpl) Name() string {
	return g.baseWidget.GetName()
}

func (g *graphWidgetImpl) UpdateSpec() *models.TreeSpec {
	route := components.GetRouteFromParents(g)

	g.htmxOpts.AddBeforePath(route)
	return &models.TreeSpec{
		Name:        g.Name(),
		ImageRoute:  "",
		Description: g.description,
		Route:       g.htmxOpts.GetUrl(),
		Children:    nil,
	}
}

func (g *graphWidgetImpl) GetSpec() *models.TreeSpec {
	return g.specs
}

func (g *graphWidgetImpl) GetChildren() []components.UIComponent {
	return nil
}

func (g *graphWidgetImpl) FindChild(string) (components.UIComponent, bool) {
	return nil, false
}

func (g *graphWidgetImpl) FindChildByType(string, string) (components.UIComponent, bool) {
	return nil, false
}

func (g *graphWidgetImpl) SetParent(parent components.UIComponent) {
	g.parent = parent
}

func (g *graphWidgetImpl) GetParent() components.UIComponent {
	return g.parent
}

func (g *graphWidgetImpl) AddChild(components.UIComponent) error {
	return errors.New("not applicable")
}

func (g *graphWidgetImpl) KillChild(components.UIComponent) error {
	return errors.New("not applicable")
}

func newGraphWidget() *graphWidgetImpl {
	var w graphWidgetImpl
	w.baseWidget = widgets.NewWidget()
	w.htmxOpts = htmx.NewEmpty()
	return &w
}

func NewGraphWidget(
	graphObj graph.Graph, setters ...func(
		n widgets.
			Widget,
	),
) GraphWidget {
	w := newGraphWidget()

	w.graph = graphObj

	for _, setter := range setters {
		setter(w.baseWidget)
	}

	id := "graphWidget_" + graphObj.GetId()
	name := "graphWidget_" + graphObj.GetName()

	w.baseWidget.SetId(id)
	w.baseWidget.SetName(name)
	w.htmxOpts.AppendToPath("graph", id)

	return w
}
