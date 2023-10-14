package main

import (
	"fmt"
	"github.com/Alfagov/goDashboard/dashboard"
	"github.com/Alfagov/goDashboard/layout"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pageContainer"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/utils"
	"github.com/Alfagov/goDashboard/widgets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
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
				return rand.Intn(500), nil
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
				return rand.Intn(500), nil
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
				return rand.Intn(500), nil
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
				return rand.Intn(500), nil
			},
		),
	)

	wd11 := widgets.NewNumericWidget(
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
				return rand.Intn(500), nil
			},
		),
	)

	wd21 := widgets.NewNumericWidget(
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
				return rand.Intn(500), nil
			},
		),
	)

	wd31 := widgets.NewNumericWidget(
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
				return rand.Intn(500), nil
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
		"form1",
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
		widgets.SetFormUpdateHandler(
			func(c widgets.FormRequest) *models.UpdateResponse {
				return &models.UpdateResponse{
					Success: true,
					Message: "success",
					Title:   "Working",
				}
			},
		),
	)

	bw := widgets.NewBarGraphWidget()

	pg := pages.NewPage(
		pages.SetName("test"),
		pages.AddNumericWidgets(wd, wd1, wd2, wd3),
		pages.AddFormWidgets(f1),
		pages.AddBarGraphWidgets(bw),
	)

	pg1 := pages.NewPage(
		pages.SetName("test1"),
		pages.AddNumericWidgets(wd11, wd21, wd31),
	)

	pgContainer := pageContainer.NewPageContainer(
		"testContainer",
		pageContainer.AddPage(pg),
		pageContainer.AddPage(pg1),
		pageContainer.SetIndexPage("test"),
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
	d.AddPageContainer(pgContainer)
	d.Compile()

	log.Fatal(app.Listen(":8080"))

}
