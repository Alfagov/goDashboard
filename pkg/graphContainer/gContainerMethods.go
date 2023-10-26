package graphContainer

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func (g *graphWidgetImpl) CompileRoutes(router *fiber.App) {
	router.Get(
		g.htmxOpts.GetUrl(), func(c *fiber.Ctx) error {
			return c.JSON(g.update())
		},
	)
}

func (g *graphWidgetImpl) Encode() templ.Component {
	return templates.GeneralGraph(
		g.baseWidget.GetId(),
		g.graph.Encode(g.baseWidget.GetLayout().Height),
		g.baseWidget.GetLayout(),
		g.htmxOpts.GetHtmx(),
	)
}

func (g *graphWidgetImpl) AddParentPath(path string) error {
	return g.htmxOpts.GetHtmx().AddBeforePath(path)
}

func (g *graphWidgetImpl) update() map[string]interface{} {
	return g.graph.HandleUpdate()
}

func (g *graphWidgetImpl) getHtmx() htmx.HTMX {
	return g.htmxOpts
}

func (g *graphWidgetImpl) WithSettings(settings ...func(gw GraphWidget)) GraphWidget {
	for _, setter := range settings {
		setter(g)
	}
	return g
}
