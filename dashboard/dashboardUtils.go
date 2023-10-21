package dashboard

import (
	"github.com/Alfagov/goDashboard/pageContainer"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/pkg/form"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func createContainerPagesRoutes(
	router *fiber.App,
	containers map[string]pageContainer.PageContainer,
) {
	for _, container := range containers {
		router.Get(
			container.GetRoute(), func(c *fiber.Ctx) error {
				t := container.Encode(container.GetIndexPage())
				c.Set("HX-Push-Url", container.GetRoute())
				return c.Render("", templates.IndexPage(t))
			},
		)

		createPagesRoutes(router, container.GetPages(), container)
	}
}

func createPagesRoutes(
	router *fiber.App,
	pages map[string]pages.Page,
	pageContainer pageContainer.PageContainer,
) {
	for _, page := range pages {

		var tIndex templ.Component
		if pageContainer != nil {
			tIndex = pageContainer.Encode(page.GetName())
		} else {
			tIndex = page.Encode()
		}

		route := page.GetRoute()

		router.Get(
			page.GetRoute(), func(c *fiber.Ctx) error {
				c.Set("HX-Push-Url", route)
				return c.Render("", templates.IndexPage(tIndex))
			},
		)

		widgetList := page.GetWidgets()
		for _, widget := range widgetList.NumericWidgets {
			w := widget
			router.Get(
				w.GetHtmx().GetRoute(), func(c *fiber.Ctx) error {
					update, err := w.HandleUpdate()
					if err != nil {
						return err
					}

					t := w.UpdateAction(update)

					return c.Render("", t)
				},
			)
		}

		for _, widget := range widgetList.FormWidgets {
			w := widget
			router.Post(
				w.GetHtmx().GetRoute(), func(c *fiber.Ctx) error {
					update := w.HandlePost(form.NewFormRequest(c))

					t := w.UpdateAction(update)

					return c.Render("", t)
				},
			)
		}

		for _, widget := range widgetList.GraphWidgets {
			w := widget
			router.Get(
				w.GetHtmx().GetRoute(), func(c *fiber.Ctx) error {
					t := w.Update()

					return c.JSON(t)
				},
			)
		}
	}
}
