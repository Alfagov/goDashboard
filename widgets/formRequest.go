package widgets

import "github.com/gofiber/fiber/v2"

type FormRequest interface {
	BindFormRequest(v interface{}) error
}

type formRequest struct {
	c *fiber.Ctx
}

func NewFormRequest(c *fiber.Ctx) FormRequest {
	return &formRequest{c: c}
}

func (fr *formRequest) BindFormRequest(v interface{}) error {
	return fr.c.BodyParser(v)
}
