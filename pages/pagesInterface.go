package pages

import (
	"github.com/Alfagov/goDashboard/pkg/form"
	"github.com/Alfagov/goDashboard/pkg/graphContainer"
	"github.com/Alfagov/goDashboard/pkg/numeric"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

type page struct {
	Name      string
	PageRoute string
	Widgets   *WidgetsContainer
}

type WidgetsContainer struct {
	NumericWidgets []numeric.Numeric
	FormWidgets    []form.FormWidget
	GraphWidgets   []graphContainer.GraphWidget
}

type Page interface {
	addNumericWidgets(widget numeric.Numeric)
	addFormWidgets(widget form.FormWidget)
	addGraphWidget(widget graphContainer.GraphWidget)
	setName(name string)
	Encode() templ.Component
	GetWidgets() *WidgetsContainer
	GetRoute() string
	SetRoute(route string)
	GetName() string
}

func NewPage(setters ...func(p Page)) Page {
	var p page

	for _, setter := range setters {
		setter(&p)
	}

	p.PageRoute = "/" + p.Name

	return &p
}

func (p *page) addGraphWidget(widget graphContainer.GraphWidget) {
	htm := widget.GetHtmx()
	htm.SetRoute(p.Name + htm.GetRoute())

	p.Widgets.GraphWidgets = append(p.Widgets.GraphWidgets, widget)
}

func (p *page) SetRoute(route string) {
	p.PageRoute = route
}

func (p *page) GetName() string {
	return p.Name
}

func (p *page) GetRoute() string {
	return p.PageRoute
}

func (p *page) GetWidgets() *WidgetsContainer {
	return p.Widgets
}

func (p *page) setName(name string) {
	p.Name = name
}

func (p *page) addNumericWidgets(widget numeric.Numeric) {
	htm := widget.GetHtmx()
	htm.SetRoute(p.Name + htm.GetRoute())

	p.Widgets.NumericWidgets = append(p.Widgets.NumericWidgets, widget)
}

func (p *page) addFormWidgets(widget form.FormWidget) {
	htm := widget.GetHtmx()
	htm.SetRoute(p.Name + htm.GetRoute())

	p.Widgets.FormWidgets = append(p.Widgets.FormWidgets, widget)
}

func (p *page) Encode() templ.Component {
	var components []templ.Component

	for _, widget := range p.Widgets.NumericWidgets {
		components = append(components, widget.Encode())
	}

	for _, widget := range p.Widgets.FormWidgets {
		components = append(components, widget.Encode())
	}

	for _, widget := range p.Widgets.GraphWidgets {
		components = append(components, widget.Encode())
	}

	pageOut := templates.GridPage(components)

	return pageOut
}

func AddNumericWidgets(widgets ...numeric.Numeric) func(p Page) {
	return func(p Page) {
		for _, widget := range widgets {
			p.addNumericWidgets(widget)
		}
	}
}

func AddFormWidgets(widgets ...form.FormWidget) func(p Page) {
	return func(p Page) {
		for _, widget := range widgets {
			p.addFormWidgets(widget)
		}
	}
}

func AddGraphWidgets(widgets ...graphContainer.GraphWidget) func(p Page) {
	return func(p Page) {
		for _, widget := range widgets {
			p.addGraphWidget(widget)
		}
	}
}

func SetName(name string) func(p Page) {
	return func(p Page) {
		p.setName(name)
	}
}
