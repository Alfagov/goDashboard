package dashboard

import (
	"github.com/Alfagov/goDashboard/pages"
	"github.com/gofiber/fiber/v2"
)

type dashboard struct {
	Router *fiber.App
	Pages  []pages.Page
}

type Dashboard interface {
	AddPage(page pages.Page)
	Compile()
}

func NewDashboard(app *fiber.App) Dashboard {
	return &dashboard{
		Router: app,
		Pages:  []pages.Page{},
	}
}

func (d *dashboard) AddPage(page pages.Page) {
	d.Pages = append(d.Pages, page)
}

func (d *dashboard) Compile() {
	for _, page := range d.Pages {

		d.Router.Get(
			page.GetIndexRoute(), func(c *fiber.Ctx) error {
				t := page.Encode()
				return c.Render("", t)
			},
		)

		widgets := page.GetWidgets()
		for _, widget := range widgets.NumericWidgets {
			w := widget
			d.Router.Get(
				w.GetRoute(), func(c *fiber.Ctx) error {
					update, err := w.HandleUpdate()
					if err != nil {
						return err
					}

					t := w.UpdateAction(update)

					return c.Render("", t)
				},
			)
		}
	}

}
