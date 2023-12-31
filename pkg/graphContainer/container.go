package graphContainer

import (
	"github.com/Alfagov/goDashboard/internal/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/graph"
	"github.com/Alfagov/goDashboard/pkg/widgets"
)

// GraphWidget represents an interface for a widget in the graph UI.
type GraphWidget interface {
	components.UIComponent

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
	w.htmxOpts.AppendToPath("widget", id)

	return w
}
