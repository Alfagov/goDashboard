package pages

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type pageContainer struct {
	Name           string
	ImagePath      string
	ContainerRoute string
	Pages          map[string]Page
	IndexPage      string
}

type PageContainer interface {
	AddPage(page Page)
	GetPages() map[string]Page
	setImagePath(path string)
	GetName() string
	GetPage(name string) Page
	GetIndexPage() string
	SetIndexPage(indexPage string)
	Encode(page string) templ.Component
	GetRoute() string
	CompileRoutes(router *fiber.App)
	GetImagePath() string
}

func (p *pageContainer) CompileRoutes(router *fiber.App) {
	router.Get(
		p.GetRoute(), func(c *fiber.Ctx) error {
			t := p.Encode(p.GetIndexPage())
			c.Set("HX-Push-Url", p.GetRoute())
			return c.Render("", templates.IndexPage(t))
		},
	)

	for _, pg := range p.Pages {

		tmpPage := pg
		router.Get(
			pg.GetRoute(), func(c *fiber.Ctx) error {
				c.Set("HX-Push-Url", tmpPage.GetRoute())
				return c.Render("", templates.IndexPage(p.Encode(tmpPage.GetName())))
			},
		)
		tmpPage.CompileWidgetsRoutes(router)
	}
}

func (p *pageContainer) setImagePath(path string) {
	p.ImagePath = path
}

func (p *pageContainer) GetImagePath() string {
	return p.ImagePath
}

func (p *pageContainer) SetIndexPage(indexPage string) {
	p.IndexPage = indexPage
}

func (p *pageContainer) GetPage(name string) Page {
	return p.Pages[name]
}

func (p *pageContainer) GetIndexPage() string {
	return p.IndexPage
}

func AddPage(page Page) func(p PageContainer) {
	return func(p PageContainer) {
		p.AddPage(page)
	}
}

func SetIndexPage(indexPage string) func(p PageContainer) {
	return func(p PageContainer) {
		p.SetIndexPage(indexPage)
	}
}

func SetContainerImagePath(path string) func(p PageContainer) {
	return func(p PageContainer) {
		p.setImagePath(path)
	}
}

func NewPageContainer(
	name string, setters ...func(
		p PageContainer,
	),
) PageContainer {
	var p pageContainer

	p.Pages = make(map[string]Page)
	p.Name = name
	p.ContainerRoute = "/" + p.Name

	for _, setter := range setters {
		setter(&p)
	}

	return &p
}

func (p *pageContainer) GetName() string {
	return p.Name
}

func (p *pageContainer) GetRoute() string {
	return p.ContainerRoute
}

func (p *pageContainer) generatePagesDescriptor() []models.PagesDescriptor {
	var pagesDescriptor []models.PagesDescriptor
	for _, page := range p.Pages {
		pagesDescriptor = append(
			pagesDescriptor, models.PagesDescriptor{
				Name:     page.GetName(),
				Route:    page.GetRoute(),
				Template: page.Encode(),
			},
		)
	}

	return pagesDescriptor
}

func (p *pageContainer) AddPage(page Page) {
	page.SetRoute(p.ContainerRoute + page.GetRoute())

	widgets := page.GetWidgets()
	for _, widget := range widgets.NumericWidgets {
		htm := widget.GetHtmx()
		htm.SetRoute(p.ContainerRoute + htm.GetRoute())
	}

	for _, widget := range widgets.FormWidgets {
		htm := widget.GetHtmx()
		htm.SetRoute(p.ContainerRoute + htm.GetRoute())
	}

	for _, widget := range widgets.GraphWidgets {
		htm := widget.GetHtmx()
		htm.SetRoute(p.ContainerRoute + htm.GetRoute())
	}

	p.Pages[page.GetName()] = page
}

func (p *pageContainer) GetPages() map[string]Page {
	return p.Pages
}

func (p *pageContainer) Encode(page string) templ.Component {

	descriptor := p.generatePagesDescriptor()
	pageTemplate := p.Pages[page].Encode()

	pContainer := templates.PageContainer(pageTemplate, descriptor)

	return pContainer
}
