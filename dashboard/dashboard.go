package dashboard

import (
	"embed"
	"github.com/Alfagov/goDashboard/internal/config"
	"github.com/Alfagov/goDashboard/internal/logger"
	"github.com/Alfagov/goDashboard/internal/utils"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	fLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
)

//go:embed static/*
var staticFiles embed.FS

// dashboard represents a collection of UI components and configuration for managing a user interface dashboard.
// it is the root component of the framework.
type dashboard struct {
	// id is the unique identifier for the dashboard.
	id string
	// name is the human-readable name of the dashboard.
	name string
	// image is a link to an image associated with the dashboard.
	image string
	// description provides a brief overview of the dashboard's purpose and contents.
	description string
	// Router is the HTTP routing mechanism that the dashboard uses.
	Router *fiber.App
	// treeSpec holds the structural specification of the dashboard's component tree.
	treeSpec *models.TreeSpec
	// IndexPage is a function that takes a templ.Component and returns a templ.Component,
	// used to define the root page of the dashboard.
	IndexPage func(body templ.Component) templ.Component
	// Children is a map of UIComponent objects that make up the dashboard's elements.
	Children map[string]components.UIComponent
}

// Dashboard is an interface that defines the functionality of a dashboard instance.
type Dashboard interface {
	// Run initializes and starts the dashboard service, potentially returning an error if something goes wrong.
	Run() error
	// WithPages takes a variadic number of UIComponent instances and adds them to the dashboard, returning the updated Dashboard instance.
	WithPages(pages ...components.UIComponent) Dashboard
}

// NewDashboard creates and returns a new Dashboard instance with the specified name and image.
// It sets up a new Fiber application with a custom template renderer and attaches logger middleware.
// The function also hooks into the route registration process to log info about newly added routes.
// The returned Dashboard is initialized with a name, image, router, and an empty map for child components.
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

// InitDashboardGlobals initializes the global settings for the dashboard application.
// It sets up the logger and configuration by calling their respective initialization functions.
// If any initialization fails, it logs the error and panics, preventing the application from starting with incorrect configurations.
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
