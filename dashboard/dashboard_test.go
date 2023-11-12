package dashboard

import (
	"github.com/Alfagov/goDashboard/internal/logger"
	"github.com/Alfagov/goDashboard/internal/utils"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/test"
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
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantChild int
	}{
		{
			name: "test wrong child type",
			fields: fields{
				Children: make(map[string]components.UIComponent),
			},
			args: args{
				child: test.NewMockUIComponent("test", components.FormWidgetType, nil),
			},
			wantErr:   true,
			wantChild: 0,
		},
		{
			name: "test error for existing child",
			fields: fields{
				Children: map[string]components.UIComponent{"test": test.NewMockUIComponent("test",
					components.PageType, nil)},
			},
			args: args{
				child: test.NewMockUIComponent("test", components.PageType, nil),
			},
			wantErr:   true,
			wantChild: 1,
		},
		{
			name: "test success add child",
			fields: fields{
				Children: make(map[string]components.UIComponent),
			},
			args: args{
				child: test.NewMockUIComponent("test", components.PageType, nil),
			},
			wantErr:   false,
			wantChild: 1,
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

			assert.Len(t, d.GetChildren(), tt.wantChild)
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
		name      string
		fields    fields
		args      args
		want      components.UIComponent
		wantFound bool
	}{
		{
			name: "success finding child",
			fields: fields{
				Children: map[string]components.UIComponent{"test": test.NewMockUIComponent("test",
					components.PageType, nil)},
			},
			args: args{
				name: "test",
			},
			want:      test.NewMockUIComponent("test", components.PageType, nil),
			wantFound: true,
		},
		{
			name: "failure finding child",
			fields: fields{
				Children: map[string]components.UIComponent{},
			},
			args: args{
				name: "test",
			},
			want:      nil,
			wantFound: false,
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

			got, found := d.FindChild(tt.args.name)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantFound, found)
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
		wantLen int
	}{
		{
			name: "success killing child",
			fields: fields{
				Children: map[string]components.UIComponent{"test": test.NewMockUIComponent("test",
					components.PageType, nil)},
			},
			args: args{
				child: test.NewMockUIComponent("test", components.PageType, nil),
			},
			wantErr: false,
			wantLen: 0,
		},
		{
			name: "failure killing child",
			fields: fields{
				Children: map[string]components.UIComponent{},
			},
			args: args{
				child: test.NewMockUIComponent("test", components.PageType, nil),
			},
			wantErr: true,
			wantLen: 0,
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

			err := d.RemoveChild(tt.args.child)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Len(t, d.GetChildren(), tt.wantLen)
		})
	}
}

// TODO: FINISH ME
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
		{
			name:   "test dashboard type",
			fields: fields{},
			want:   components.DashboardType,
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
			assert.Equal(t, tt.want, d.Type())
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
		{
			name: "test update spec",
			fields: fields{
				name:        "testName",
				image:       "testImage",
				description: "testDescription",
				Children:    map[string]components.UIComponent{},
			},
			want: &models.TreeSpec{
				Name:        "testName",
				ImageRoute:  "testImage",
				Description: "testDescription",
				Route:       "testName",
				Children:    nil,
			},
		},
		{
			name: "test update spec",
			fields: fields{
				name:        "testName",
				image:       "testImage",
				description: "testDescription",
				Children: map[string]components.UIComponent{"test": test.NewMockUIComponent("test",
					components.PageType, nil)},
			},
			want: &models.TreeSpec{
				Name:        "testName",
				ImageRoute:  "testImage",
				Description: "testDescription",
				Route:       "testName",
				Children:    []*models.TreeSpec{{Name: "test"}},
			},
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
			if got := d.UpdateSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
