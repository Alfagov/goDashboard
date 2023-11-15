package dashboard

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
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
		if !ok {
			return c.SendStatus(404)
		}

		responseTemplate := widget.Render(models.NewReqWrapper(c))

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
		if !ok {
			return c.SendStatus(404)
		}

		responseTemplate := widget.Render(models.NewReqWrapper(c))

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
		if !ok {
			return c.SendStatus(404)
		}

		responseTemplate := widget.Render(models.NewReqWrapper(c))

		return c.Render("", responseTemplate.Component)
	})

	d.Router.Get("/:container/:page", func(c *fiber.Ctx) error {
		container, ok := d.FindChildByType(c.Params("container"), "container")
		if !ok {
			return c.SendStatus(404)
		}

		c.Locals("pageName", c.Params("page"))

		responseTemplate := IndexPage(d.name, d.image, d.GetSpec().Children,
			container.Render(models.NewReqWrapper(c)).Component)

		c.Set("HX-Push-Url", c.Path())

		return c.Render("", responseTemplate)
	})

	d.Router.Get("/:page", func(c *fiber.Ctx) error {
		page, ok := d.FindChild(c.Params("page"))
		if !ok {
			return c.SendStatus(404)
		}

		if page.Type().SuperType() == components.PageType.SuperType() {
			responseTemplate := IndexPage(d.name, d.image, d.GetSpec().Children, page.Render(nil).Component)
			c.Set("HX-Push-Url", c.Path())
			return c.Render("", responseTemplate)
		}

		return c.SendStatus(400)
	})

	d.Router.Get("/", func(c *fiber.Ctx) error {
		responseTemplate := IndexPage(d.name, d.image, d.GetSpec().Children, d.Render(nil).Component)
		c.Set("HX-Push-Url", c.Path())
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
