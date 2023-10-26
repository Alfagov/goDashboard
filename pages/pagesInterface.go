package pages

import (
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/pkg/form"
	"github.com/Alfagov/goDashboard/pkg/graphContainer"
	"github.com/Alfagov/goDashboard/pkg/numeric"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type page struct {
	Name      string
	PageRoute string
	ImagePath string
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
	setImagePath(path string)
	Encode() templ.Component
	GetWidgets() *WidgetsContainer
	GetRoute() string
	SetRoute(route string)
	GetName() string
	CompileWidgetsRoutes(router *fiber.App)
	CompilePageRoutes(router *fiber.App, indexRenderer func(component templ.Component) templ.Component)
	GetImagePath() string
}

func (p *page) CompileWidgetsRoutes(router *fiber.App) {
	for _, widget := range p.Widgets.NumericWidgets {
		widget.CompileRoutes(router)
	}

	for _, widget := range p.Widgets.FormWidgets {
		widget.CompileRoutes(router)
	}

	for _, widget := range p.Widgets.GraphWidgets {
		widget.CompileRoutes(router)
	}
}

func (p *page) GetImagePath() string {
	return p.ImagePath
}

func (p *page) CompilePageRoutes(router *fiber.App, indexRenderer func(component templ.Component) templ.Component) {

	router.Get(p.GetRoute(), func(c *fiber.Ctx) error {
		c.Set("HX-Push-Url", p.GetRoute())
		return c.Render("", indexRenderer(p.Encode()))
	})

	p.CompileWidgetsRoutes(router)

}

func NewPage(name string, setters ...func(p Page)) Page {
	var p page
	p.Name = name
	p.ImagePath = "/static/img/" + name + ".png"
	p.Widgets = &WidgetsContainer{
		NumericWidgets: []numeric.Numeric{},
		FormWidgets:    []form.FormWidget{},
		GraphWidgets:   []graphContainer.GraphWidget{},
	}

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

func (p *page) setImagePath(path string) {
	p.ImagePath = path
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

	err := widget.AddParentPath(p.Name)
	if err != nil {
		logger.L.Error("Error adding parent path", zap.Error(err))
	}

	p.Widgets.NumericWidgets = append(p.Widgets.NumericWidgets, widget)
}

func (p *page) addFormWidgets(widget form.FormWidget) {
	err := widget.AddParentPath(p.Name)
	if err != nil {
		logger.L.Error("Error adding parent path", zap.Error(err))
	}

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

func SetImagePath(path string) func(p Page) {
	return func(p Page) {
		p.setImagePath(path)
	}
}
