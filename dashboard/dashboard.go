package dashboard

import (
	"embed"
	"github.com/Alfagov/goDashboard/config"
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/utils"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	fLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
)

//go:embed static/*
var staticFiles embed.FS

type dashboard struct {
	id          string
	name        string
	image       string
	description string
	Router      *fiber.App
	treeSpec    *models.TreeSpec
	IndexPage   func(body templ.Component) templ.Component
	Children    map[string]components.UIComponent
}

type Dashboard interface {
	Run() error
	WithPages(pages ...components.UIComponent) Dashboard
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
		name:     name,
		image:    img,
		Router:   app,
		Children: make(map[string]components.UIComponent),
	}
}

func InitDashboardGlobals() {
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
