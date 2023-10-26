package dashboard

import (
	"embed"
	"github.com/Alfagov/goDashboard/config"
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/utils"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	fLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
)

//go:embed static/*
var staticFiles embed.FS

type dashboard struct {
	Name           string
	Image          string
	Router         *fiber.App
	Pages          map[string]pages.Page
	PageContainers map[string]pages.PageContainer
	PagesSpec      []models.PageSpec
	IndexPage      func(body templ.Component) templ.Component
}

type Dashboard interface {
	AddPage(page pages.Page)
	AddPageContainer(pageContainer pages.PageContainer)
	Compile()
	Run() error
}

func NewDashboard(name string, img string) Dashboard {

	app := fiber.New(fiber.Config{Views: &utils.TemplRender{}})

	app.Use(fLogger.New())

	app.Hooks().OnRoute(
		func(r fiber.Route) error {
			logger.L.Info(
				"Added route: ", zap.String("Name", r.Path), zap.String(
					"Method",
					r.Method,
				),
			)

			return nil
		},
	)

	return &dashboard{
		Name:           name,
		Image:          img,
		Router:         app,
		Pages:          make(map[string]pages.Page),
		PageContainers: make(map[string]pages.PageContainer),
	}
}

func init() {
	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}

	logger.L.Info("Logger initialized")

	err = config.InitConfig()
	if err != nil {
		logger.L.Error("Error initializing config", zap.Error(err))
		panic(err)
	}

	logger.L.Info("Config initialized")
}
