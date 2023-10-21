package dashboard

import (
	"github.com/Alfagov/goDashboard/pageContainer"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
)

type dashboard struct {
	Router         *fiber.App
	Logger         *zap.Logger
	Pages          map[string]pages.Page
	PageContainers map[string]pageContainer.PageContainer
}

type Dashboard interface {
	AddPage(page pages.Page)
	AddPageContainer(pageContainer pageContainer.PageContainer)
	Compile()
	Run() error
}

func NewDashboard() Dashboard {
	app := fiber.New(
		fiber.Config{
			Views: &utils.TemplRender{},
		},
	)

	app.Use(logger.New())

	l, err := initializeLogger()
	if err != nil {
		panic(err)
	}

	l.Info("Logger initialized")

	app.Hooks().OnRoute(
		func(r fiber.Route) error {
			l.Info(
				"Added route: ", zap.String("Name", r.Path), zap.String(
					"Method",
					r.Method,
				),
			)

			return nil
		},
	)

	return &dashboard{
		Router:         app,
		Logger:         l,
		Pages:          make(map[string]pages.Page),
		PageContainers: make(map[string]pageContainer.PageContainer),
	}
}
