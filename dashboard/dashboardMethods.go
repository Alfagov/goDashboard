package dashboard

import (
	"github.com/Alfagov/goDashboard/pageContainer"
	"github.com/Alfagov/goDashboard/pages"
)

func (d *dashboard) AddPageContainer(pageContainer pageContainer.PageContainer) {
	d.PageContainers[pageContainer.GetName()] = pageContainer
}

func (d *dashboard) AddPage(page pages.Page) {
	d.Pages[page.GetName()] = page
}

func (d *dashboard) Compile() {
	createContainerPagesRoutes(d.Router, d.PageContainers)
	createPagesRoutes(d.Router, d.Pages, nil)
}

func (d *dashboard) Run() error {
	d.Router.Static("/", "./css")
	d.Router.Static("/", "./js")

	d.Compile()

	return d.Router.Listen(":8080")
}
