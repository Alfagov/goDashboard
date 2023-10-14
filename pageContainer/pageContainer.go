package pageContainer

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

type pageContainer struct {
	Name           string
	ContainerRoute string
	Pages          map[string]pages.Page
	IndexPage      string
}

type PageContainer interface {
	AddPage(page pages.Page)
	GetPages() map[string]pages.Page
	GetName() string
	GetPage(name string) pages.Page
	GetIndexPage() string
	SetIndexPage(indexPage string)
	Encode(page string) templ.Component
	GetRoute() string
}

func (p *pageContainer) SetIndexPage(indexPage string) {
	p.IndexPage = indexPage
}

func (p *pageContainer) GetPage(name string) pages.Page {
	return p.Pages[name]
}

func (p *pageContainer) GetIndexPage() string {
	return p.IndexPage
}

func AddPage(page pages.Page) func(p PageContainer) {
	return func(p PageContainer) {
		p.AddPage(page)
	}
}

func SetIndexPage(indexPage string) func(p PageContainer) {
	return func(p PageContainer) {
		p.SetIndexPage(indexPage)
	}
}

func NewPageContainer(
	name string, setters ...func(
		p PageContainer,
	),
) PageContainer {
	var p pageContainer

	p.Pages = make(map[string]pages.Page)
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

func (p *pageContainer) AddPage(page pages.Page) {
	page.SetRoute(p.ContainerRoute + page.GetRoute())

	widgets := page.GetWidgets()
	for _, widget := range widgets.NumericWidgets {
		widget.SetPageRoute(page.GetRoute())
	}

	for _, widget := range widgets.FormWidgets {
		widget.SetPageRoute(page.GetRoute())
	}

	p.Pages[page.GetName()] = page
}

func (p *pageContainer) GetPages() map[string]pages.Page {
	return p.Pages
}

func (p *pageContainer) Encode(page string) templ.Component {

	descriptor := p.generatePagesDescriptor()
	pageTemplate := p.Pages[page].Encode()

	pContainer := templates.PageContainer(pageTemplate, descriptor)

	return pContainer
}
