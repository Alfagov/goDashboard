package main

import (
	"fmt"
	"github.com/Alfagov/goDashboard/dashboard"
	"github.com/Alfagov/goDashboard/layout"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/utils"
	"github.com/Alfagov/goDashboard/widgets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"time"
)

func main() {

	wd := widgets.NewNumericWidget(
		time.Second*5,
		widgets.SetName("test"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(1),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSpecificSettings(
		widgets.SetNumericInitValue(1),
		widgets.SetNumericUpdateHandler(
			func() (int, error) {
				return 1, nil
			},
		),
	)

	wd1 := widgets.NewNumericWidget(
		time.Second*5,
		widgets.SetName("test"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(2),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSpecificSettings(
		widgets.SetNumericInitValue(1),
		widgets.SetNumericUpdateHandler(
			func() (int, error) {
				return 1, nil
			},
		),
	)

	wd2 := widgets.NewNumericWidget(
		time.Second*5,
		widgets.SetName("test"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(3),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSpecificSettings(
		widgets.SetNumericInitValue(1),
		widgets.SetNumericUpdateHandler(
			func() (int, error) {
				return 1, nil
			},
		),
	)

	wd3 := widgets.NewNumericWidget(
		time.Second*5,
		widgets.SetName("test"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(3),
				layout.SetRow(2),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSpecificSettings(
		widgets.SetNumericInitValue(1),
		widgets.SetNumericUpdateHandler(
			func() (int, error) {
				return 1, nil
			},
		),
	)

	formFields := []*models.FormField{
		{
			Name:      "test",
			Label:     "test",
			FieldType: "text",
		},
		{
			Name:      "test1",
			Label:     "test1",
			FieldType: "text",
		},
	}

	formButtons := []*models.FormButton{
		{
			Label: "test",
			Color: "red",
		},
	}

	f1 := widgets.NewFormWidget(
		widgets.SetName("form1"),
		widgets.SetDescription("form1"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(1),
				layout.SetRow(2),
				layout.SetHeight(2),
				layout.SetWidth(2),
			),
		),
	).WithFormSpecs(
		widgets.SetFormFields(formFields...),
		widgets.SetFormButtons(formButtons...),
	)

	f2 := widgets.NewFormWidget(
		widgets.SetName("form1"),
		widgets.SetDescription("form1"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(2),
				layout.SetRow(4),
				layout.SetHeight(2),
				layout.SetWidth(2),
			),
		),
	).WithFormSpecs(
		widgets.SetFormFields(formFields...),
		widgets.SetFormButtons(formButtons...),
	)

	log.Println(f1)

	pg := pages.NewPage(
		pages.SetName("test"),
		pages.AddNumericWidgets(wd, wd1, wd2, wd3),
		pages.AddFormWidgets(f1, f2),
	)

	app := fiber.New(
		fiber.Config{
			Views: &utils.TemplRender{},
		},
	)
	app.Use(logger.New())

	app.Hooks().OnRoute(
		func(r fiber.Route) error {
			fmt.Println("Name: " + r.Path + ", " + r.Method)

			return nil
		},
	)

	app.Static("/", "./css")

	d := dashboard.NewDashboard(app)
	d.AddPage(pg)
	d.Compile()

	/*	app.Get(
		"/", func(c *fiber.Ctx) error {
			return c.Render("", pg.Encode())
		},
	)*/

	log.Fatal(app.Listen(":8080"))

}
