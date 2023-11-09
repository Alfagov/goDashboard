package graphContainer

import (
	"errors"
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"go.uber.org/zap"
)

// GraphWidget implementation

func (g *graphWidgetImpl) update() map[string]interface{} {
	return g.graph.HandleUpdate()
}

func (g *graphWidgetImpl) WithSettings(settings ...func(gw GraphWidget)) GraphWidget {
	for _, setter := range settings {
		setter(g)
	}
	return g
}

// UIComponent implementation

func (g *graphWidgetImpl) Render(req models.RequestWrapper) *components.RenderResponse {
	if req != nil {
		return &components.RenderResponse{
			Json: g.update(),
		}
	}

	return &components.RenderResponse{
		Component: GeneralGraph(
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

	err := g.htmxOpts.AddBeforePath(route)
	if err != nil {
		logger.L.Error("error in updating spec", zap.Error(err))
	}

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

func (g *graphWidgetImpl) Id() string {
	return g.baseWidget.GetId()
}

func (g *graphWidgetImpl) FindChildById(string) (components.UIComponent, bool) {
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

func (g *graphWidgetImpl) RemoveChild(components.UIComponent) error {
	return errors.New("not applicable")
}
