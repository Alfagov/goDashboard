package dashboard

import (
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

func (d *dashboard) CreateRoutes() {
	d.Router.Get("/:container/:page/widget/:widget", func(c *fiber.Ctx) error {

		container, ok := d.FindChildByType(c.Params("container"), "container")
		if !ok {
			return c.SendStatus(404)
		}

		page, ok := container.FindChildByType(c.Params("page"), "page")
		if !ok {
			return c.SendStatus(404)
		}

		widget, ok := page.FindChildById(c.Params("widget"))

		responseTemplate := widget.Render(components.NewReqWrapper(c))

		if responseTemplate.Json != nil {
			return c.JSON(responseTemplate.Json)
		}

		return c.Render("", responseTemplate.Component)
	})

	d.Router.Post("/:container/:page/widget/:widget", func(c *fiber.Ctx) error {

		container, ok := d.FindChildByType(c.Params("container"), "container")
		if !ok {
			return c.SendStatus(404)
		}

		page, ok := container.FindChildByType(c.Params("page"), "page")
		if !ok {
			return c.SendStatus(404)
		}

		widget, ok := page.FindChildById(c.Params("widget"))

		responseTemplate := widget.Render(components.NewReqWrapper(c))

		if responseTemplate.Json != nil {
			return c.JSON(responseTemplate.Json)
		}

		return c.Render("", responseTemplate.Component)
	})

	d.Router.Get("/:page/widget/:widget", func(c *fiber.Ctx) error {
		page, ok := d.FindChildByType(c.Params("page"), "page")
		if !ok {
			return c.SendStatus(404)
		}

		widget, ok := page.FindChildById(c.Params("widget"))

		responseTemplate := widget.Render(components.NewReqWrapper(c))

		return c.Render("", responseTemplate.Component)
	})

	d.Router.Get("/:container/:page", func(c *fiber.Ctx) error {
		logger.L.Info("Page container page called")
		container, ok := d.FindChildByType(c.Params("container"), "container")
		if !ok {
			return c.SendStatus(404)
		}

		page, ok := container.FindChildByType(c.Params("page"), "page")
		if !ok {
			return c.SendStatus(404)
		}

		responseTemplate := templates.PageContainer(page.Render(components.NewReqWrapper(c)).Component, page.GetSpec().Children)

		responseTemplate = templates.IndexPage(d.name, d.image, container.GetSpec().Children, responseTemplate)

		return c.Render("", responseTemplate)
	})
}

func (d *dashboard) CreateStaticRoutes() {
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
}
