package pages

import (
	"github.com/Alfagov/goDashboard/templates"
	"github.com/Alfagov/goDashboard/widgets"
	"github.com/a-h/templ"
)

type page struct {
	Name           string
	PageRoute      string
	NumericWidgets []widgets.NumericWidget
	FormWidgets    []widgets.FormWidget
	BarGraphWidget []widgets.BarGraphWidget
}

type WidgetsContainer struct {
	NumericWidgets []widgets.NumericWidget
	FormWidgets    []widgets.FormWidget
	BarGraphWidget []widgets.BarGraphWidget
}

type Page interface {
	addNumericWidgets(widget widgets.NumericWidget)
	addFormWidgets(widget widgets.FormWidget)
	addGraphBarWidget(widget widgets.BarGraphWidget)
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
	return &WidgetsContainer{
		NumericWidgets: p.NumericWidgets,
		FormWidgets:    p.FormWidgets,
		BarGraphWidget: p.BarGraphWidget,
	}
}

func (p *page) setName(name string) {
	p.Name = name
}

func (p *page) addNumericWidgets(widget widgets.NumericWidget) {
	widget.SetPageRoute(p.Name)
	p.NumericWidgets = append(p.NumericWidgets, widget)
}

func (p *page) addFormWidgets(widget widgets.FormWidget) {
	widget.SetPageRoute(p.Name)
	p.FormWidgets = append(p.FormWidgets, widget)
}

func (p *page) addGraphBarWidget(widget widgets.BarGraphWidget) {
	// TODO: SET PAGE ROUTE
	p.BarGraphWidget = append(p.BarGraphWidget, widget)
}

func (p *page) Encode() templ.Component {
	var components []templ.Component

	for _, widget := range p.NumericWidgets {
		components = append(components, widget.Encode())
	}

	for _, widget := range p.FormWidgets {
		components = append(components, widget.Encode())
	}

	for _, widget := range p.BarGraphWidget {
		components = append(components, widget.Encode())
	}

	pageOut := templates.GridPage(components)

	return pageOut
}

func AddNumericWidgets(widgets ...widgets.NumericWidget) func(p Page) {
	return func(p Page) {
		for _, widget := range widgets {
			p.addNumericWidgets(widget)
		}
	}
}

func AddFormWidgets(widgets ...widgets.FormWidget) func(p Page) {
	return func(p Page) {
		for _, widget := range widgets {
			p.addFormWidgets(widget)
		}
	}
}

func AddBarGraphWidgets(widgets ...widgets.BarGraphWidget) func(p Page) {
	return func(p Page) {
		for _, widget := range widgets {
			p.addGraphBarWidget(widget)
		}
	}
}

func SetName(name string) func(p Page) {
	return func(p Page) {
		p.setName(name)
	}
}
