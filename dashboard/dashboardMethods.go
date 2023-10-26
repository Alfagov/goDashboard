package dashboard

import (
	_ "embed"
	"github.com/Alfagov/goDashboard/config"
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
	"strings"
)

func (d *dashboard) AddPageContainer(pageContainer pages.PageContainer) {
	logger.L.Debug("Adding page container", zap.String("Name", pageContainer.GetName()))
	d.PageContainers[pageContainer.GetName()] = pageContainer
}

func (d *dashboard) AddPage(page pages.Page) {
	logger.L.Debug("Adding page", zap.String("Name", page.GetName()))
	d.Pages[page.GetName()] = page
}

func (d *dashboard) Compile() {
	logger.L.Debug("Compiling dashboard", zap.String("Name", ""))
	d.generateDashboardPagesSpec()
	d.generateIndexPageTemplate()
	d.createContainerPagesRoutes()
	d.createPagesRoutes()
	d.createIndexPage()
}

func (d *dashboard) generateIndexPageTemplate() {
	d.IndexPage = func(body templ.Component) templ.Component {
		name := d.Name
		img := d.Image
		return templates.IndexPage(name, img, pagesSpecToElement(d.PagesSpec), body)
	}
}

func (d *dashboard) Run() error {
	d.Router.Use("/internal_static/*", func(c *fiber.Ctx) error {
		// Get the requested file path (after /static)
		file, err := staticFiles.ReadFile("static/" + c.Params("*"))
		if err != nil {
			return c.SendStatus(404) // Not found
		}

		extension := strings.Split(c.Params("*"), ".")[1]
		if extension == "css" {
			c.Set("Content-Type", "text/css")
		}
		return c.Send(file)
	})

	d.Router.Use("/static/*", func(c *fiber.Ctx) error {

		file, err := os.ReadFile("static/" + c.Params("*"))
		if err != nil {
			return c.SendStatus(404) // Not found
		}

		return c.Send(file)
	})

	d.Compile()

	dashboardUrl := config.C.DashboardConfig.Host + ":" + config.C.DashboardConfig.Port

	if config.C.DashboardConfig.SSL {
		return d.Router.ListenTLS(dashboardUrl, config.C.DashboardConfig.CertPath, config.C.DashboardConfig.KeyPath)
	}

	return d.Router.Listen(dashboardUrl)
}

func (d *dashboard) createPagesRoutes() {
	for _, page := range d.Pages {
		page.CompilePageRoutes(d.Router, d.IndexPage)
	}
}

func (d *dashboard) createContainerPagesRoutes() {
	for _, container := range d.PageContainers {
		container.CompileRoutes(d.Router, d.IndexPage)
	}
}

func (d *dashboard) generateDashboardPagesSpec() {
	t := pagesToSpec(d.Pages)

	for _, container := range d.PageContainers {
		t = append(t, models.PageSpec{
			Name:        container.GetName(),
			ImageRoute:  container.GetImagePath(),
			Description: "",
			Route:       container.GetRoute(),
			Pages:       pagesToSpec(container.GetPages()),
		})
	}

	d.PagesSpec = t
}

func (d *dashboard) createIndexPage() {
	page := templates.ListGridPage(pagesSpecToElement(d.PagesSpec))

	d.Router.Get("/", func(c *fiber.Ctx) error {
		c.Set("HX-Push-Url", "/")
		return c.Render("", d.IndexPage(page))
	})
}

func pagesToSpec(pages map[string]pages.Page) []models.PageSpec {
	t := make([]models.PageSpec, 0)
	for _, page := range pages {
		t = append(t, models.PageSpec{
			Name:        page.GetName(),
			ImageRoute:  page.GetImagePath(),
			Description: "",
			Route:       page.GetRoute(),
			Pages:       nil,
		})
	}

	return t
}

func pagesSpecToElement(pagesSpec []models.PageSpec) []models.ListElement {
	t := make([]models.ListElement, 0)
	for _, spec := range pagesSpec {
		t = append(t, models.ListElement{
			Name:        spec.Name,
			ImageRoute:  spec.ImageRoute,
			Description: spec.Description,
			Route:       spec.Route,
			Children:    pagesSpecToElement(spec.Pages),
		})
	}

	return t
}
