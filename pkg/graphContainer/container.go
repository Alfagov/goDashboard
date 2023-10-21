package graphContainer

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/pkg/graph"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

type GraphWidget interface {
	WithSettings(settings ...func(gw GraphWidget)) GraphWidget
	Update() map[string]interface{}
	Encode() templ.Component
	GetHtmx() htmx.HTMX
}

func (g *graphWidgetImpl) Encode() templ.Component {
	return templates.GeneralGraph(
		g.baseWidget.GetId(),
		g.graph.Encode(),
		g.baseWidget.GetLayout(),
		g.htmxOpts.GetHtmx(),
	)
}

func (g *graphWidgetImpl) Update() map[string]interface{} {
	return g.graph.HandleUpdate()
}

func (g *graphWidgetImpl) GetHtmx() htmx.HTMX {
	return g.htmxOpts
}

func (g *graphWidgetImpl) WithSettings(settings ...func(gw GraphWidget)) GraphWidget {
	for _, setter := range settings {
		setter(g)
	}
	return g
}

type graphWidgetImpl struct {
	baseWidget widgets.Widget
	htmxOpts   htmx.HTMX
	graph      graph.Graph
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
	w.htmxOpts.SetRoute("/graph/" + id)

	return w
}
