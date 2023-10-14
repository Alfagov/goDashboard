package dashboard

import (
	"github.com/Alfagov/goDashboard/pageContainer"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/gofiber/fiber/v2"
)

type dashboard struct {
	Router         *fiber.App
	Pages          map[string]pages.Page
	PageContainers map[string]pageContainer.PageContainer
}

type Dashboard interface {
	AddPage(page pages.Page)
	AddPageContainer(pageContainer pageContainer.PageContainer)
	Compile()
}

func NewDashboard(app *fiber.App) Dashboard {
	return &dashboard{
		Router:         app,
		Pages:          make(map[string]pages.Page),
		PageContainers: make(map[string]pageContainer.PageContainer),
	}
}
