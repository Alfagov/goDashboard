package dashboard

import (
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/test"
	"github.com/Alfagov/goDashboard/utils"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestNewDashboard(t *testing.T) {
	type args struct {
		name string
		img  string
	}
	tests := []struct {
		name string
		args args
		want Dashboard
	}{
		{
			name: "TestNewDashboard",
			args: args{
				name: "test",
				img:  "test",
			},
			want: &dashboard{
				name:      "test",
				image:     "test",
				Router:    fiber.New(fiber.Config{Views: &utils.TemplRender{}}),
				treeSpec:  nil,
				IndexPage: nil,
				Children:  make(map[string]components.UIComponent),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewDashboard(tt.args.name, tt.args.img)

			assert.NotNil(t, got)

			gotDashboard, ok := got.(*dashboard)
			assert.True(t, ok)

			assert.Equal(t, tt.args.name, gotDashboard.name)
			assert.Equal(t, tt.args.img, gotDashboard.image)

			assert.IsType(t, &fiber.App{}, gotDashboard.Router)
		})
	}
}

// TODO: finish tests
func Test_dashboard_AddChild(t *testing.T) {
	logger.L = zap.NewNop()
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
	type args struct {
		child components.UIComponent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test wrong child type",
			fields: fields{
				Children: make(map[string]components.UIComponent),
			},
			args: args{
				child: test.NewMockUIComponent("test", components.FormWidgetType),
			},
			wantErr: true,
		},
		{
			name: "test for existing child",
			fields: fields{
				Children: map[string]components.UIComponent{"test": test.NewMockUIComponent("test", components.PageType)},
			},
			args: args{
				child: test.NewMockUIComponent("test", components.PageType),
			},
			wantErr: true,
		},
		{
			name: "test success add child",
			fields: fields{
				Children: make(map[string]components.UIComponent),
			},
			args: args{
				child: test.NewMockUIComponent("test", components.PageType),
			},
			wantErr: false,
		},
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
			if err := d.AddChild(tt.args.child); (err != nil) != tt.wantErr {
				t.Errorf("AddChild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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

func Test_dashboard_CreateStaticRoutes(t *testing.T) {
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
			d.CreateStaticRoutes()
		})
	}
}

func Test_dashboard_FindChild(t *testing.T) {
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
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   components.UIComponent
		want1  bool
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
			got, got1 := d.FindChild(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindChild() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindChild() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dashboard_FindChildById(t *testing.T) {
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
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   components.UIComponent
		want1  bool
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
			got, got1 := d.FindChildById(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindChildById() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindChildById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dashboard_FindChildByType(t *testing.T) {
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
	type args struct {
		name          string
		componentType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   components.UIComponent
		want1  bool
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
			got, got1 := d.FindChildByType(tt.args.name, tt.args.componentType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindChildByType() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindChildByType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dashboard_GetChildren(t *testing.T) {
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
		want   []components.UIComponent
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
			if got := d.GetChildren(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_GetParent(t *testing.T) {
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
		want   components.UIComponent
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
			if got := d.GetParent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_GetSpec(t *testing.T) {
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
		want   *models.TreeSpec
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
			if got := d.GetSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_Id(t *testing.T) {
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
		want   string
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
			if got := d.Id(); got != tt.want {
				t.Errorf("Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_KillChild(t *testing.T) {
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
	type args struct {
		child components.UIComponent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
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
			if err := d.KillChild(tt.args.child); (err != nil) != tt.wantErr {
				t.Errorf("KillChild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dashboard_Name(t *testing.T) {
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
		want   string
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
			if got := d.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_Render(t *testing.T) {
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
	type args struct {
		in0 components.RequestWrapper
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *components.RenderResponse
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
			if got := d.Render(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_Run(t *testing.T) {
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
		name    string
		fields  fields
		wantErr bool
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
			if err := d.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dashboard_SetParent(t *testing.T) {
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
	type args struct {
		in0 components.UIComponent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			d.SetParent(tt.args.in0)
		})
	}
}

func Test_dashboard_Type(t *testing.T) {
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
		want   components.NodeType
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
			if got := d.Type(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_UpdateSpec(t *testing.T) {
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
		want   *models.TreeSpec
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
			if got := d.UpdateSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dashboard_WithPages(t *testing.T) {
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
	type args struct {
		pages []components.UIComponent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Dashboard
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
			if got := d.WithPages(tt.args.pages...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPages() = %v, want %v", got, tt.want)
			}
		})
	}
}
