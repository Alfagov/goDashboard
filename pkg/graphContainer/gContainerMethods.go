package graphContainer

func (g *graphWidgetImpl) update() map[string]interface{} {
	return g.graph.HandleUpdate()
}

func (g *graphWidgetImpl) WithSettings(settings ...func(gw GraphWidget)) GraphWidget {
	for _, setter := range settings {
		setter(g)
	}
	return g
}
