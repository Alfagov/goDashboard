package main

import (
	"github.com/Alfagov/goDashboard/utils"
	"github.com/Alfagov/goDashboard/widgets"
	"github.com/gin-gonic/gin"
)

func main() {

	wd := widgets.NewNumericWidget(
		5,
		5,
		widgets.SetName("test"),
		widgets.SetDescription("test"),
	)

	router := gin.Default()
	router.HTMLRender = &utils.TemplRender{}
	router.GET(
		"/", func(c *gin.Context) {
			c.HTML(200, "", wd.Encode())
		},
	)

	router.Run(":8080")
}
