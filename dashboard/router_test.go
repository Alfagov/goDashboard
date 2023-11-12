package dashboard

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"testing"
)

func Test_dashboard_CreateRoutes(t *testing.T) {
	type fields struct {
		id          string
		name        string
		image       string
		description string
		Router      *fiber.App
		treeSpec    *models.TreeSpec
		IndexPage   func(body templ.Component) templ.Component
		Children    map[string]components.UIComponent
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dashboard{
				id:          tt.fields.id,
				name:        tt.fields.name,
				image:       tt.fields.image,
				description: tt.fields.description,
				Router:      tt.fields.Router,
				treeSpec:    tt.fields.treeSpec,
				IndexPage:   tt.fields.IndexPage,
				Children:    tt.fields.Children,
			}
			d.CreateRoutes()
		})
	}
}
