package dashboard

import (
	"github.com/Alfagov/goDashboard/pages"
	"github.com/gofiber/fiber/v2"
)

type Dashboard struct {
	Router *fiber.App
	Pages  []*pages.Page
}

func main() {
}
