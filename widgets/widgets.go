package widgets

import (
	"context"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"os"
	"testing"
)

type BaseWidget struct {
	Id          string
	Name        string
	Description string
	Route       string
	Width       int
	Height      int
}

type Widget interface {
	setName(name string)
	setHeight(height int)
	setWidth(width int)
	setId()
	setDescription(description string)
	Encode() templ.Component
}

func Test(test *testing.T) {
	simpleFields := templates.FormField("test", "test", "test")
	selectFields := templates.SelectFormField(
		"test", "t", []string{"test",
			"test"},
	)
	formButton := &models.FormButton{
		Label: "test1",
		Color: "gray",
	}

	t := templates.GenericForm(
		"test",
		[]templ.Component{simpleFields, selectFields},
		nil,
		formButton,
	)

	t.Render(context.Background(), os.Stdout)
}
