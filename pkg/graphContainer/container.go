package graphContainer

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/pkg/graph"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

// GraphWidget represents an interface for a widget in the graph UI.
type GraphWidget interface {

	// update method returns a map representing the state of the GraphWidget.
	update() map[string]interface{}

	// getHtmx method returns an HTMX object associated with the GraphWidget.
	getHtmx() htmx.HTMX

	// WithSettings method is used to update the GraphWidget's settings.
	// It accepts a variable number of functions that mutate the GraphWidget.
	WithSettings(settings ...func(gw GraphWidget)) GraphWidget

	// CompileRoutes method registers the necessary routes with the provided *fiber.App.
	CompileRoutes(router *fiber.App)

	// AddParentPath method appends a new path to the GraphWidget's parent
	// paths list.
	//It returns an error in case of failure.
	AddParentPath(path string) error

	// Encode method returns a templ.Component which represents the GraphWidget's HTML component.
	Encode() templ.Component
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
	w.htmxOpts.AppendToPath("graph", id)

	return w
}
