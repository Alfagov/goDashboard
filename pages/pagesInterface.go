package pages

import (
	"github.com/Alfagov/goDashboard/templates"
	"github.com/Alfagov/goDashboard/widgets"
	"github.com/a-h/templ"
)

type page struct {
	Name           string
	Widgets        []string
	IndexRoute     string
	NumericWidgets []widgets.NumericWidget
	FormWidgets    []widgets.FormWidget
}

type WidgetsContainer struct {
	NumericWidgets []widgets.NumericWidget
	FormWidgets    []widgets.FormWidget
}

type Page interface {
	addNumericWidgets(widget widgets.NumericWidget)
	addFormWidgets(widget widgets.FormWidget)
	setName(name string)
	Encode() templ.Component
	GetWidgets() *WidgetsContainer
	GetIndexRoute() string
}

func NewPage(setters ...func(p Page)) Page {
	var p page
	for _, setter := range setters {
		setter(&p)
	}

	p.IndexRoute = "/" + p.Name

	return &p
}

func (p *page) GetIndexRoute() string {
	return p.IndexRoute
}

func (p *page) GetWidgets() *WidgetsContainer {
	return &WidgetsContainer{
		NumericWidgets: p.NumericWidgets,
		FormWidgets:    p.FormWidgets,
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

func (p *page) Encode() templ.Component {
	var components []templ.Component

	for _, widget := range p.NumericWidgets {
		components = append(components, widget.Encode())
	}

	for _, widget := range p.FormWidgets {
		components = append(components, widget.Encode())
	}

	pageOut := templates.GridPage(components, true)
	idxPage := templates.IndexPage(pageOut)

	return idxPage
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

func SetName(name string) func(p Page) {
	return func(p Page) {
		p.setName(name)
	}
}
