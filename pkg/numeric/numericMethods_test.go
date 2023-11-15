package numeric

import (
	"bytes"
	"context"
	"errors"
	"github.com/Alfagov/goDashboard/internal/htmx"
	"github.com/Alfagov/goDashboard/internal/layout"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/test"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func Test_numeric_AddChild(t *testing.T) {
	type fields struct {
		baseWidget    widgets.Widget
		updateHandler func() (int, error)
		initialValue  int
		unit          string
		description   string
		unitAfter     bool
		htmxOpts      htmx.HTMX
		spec          *models.TreeSpec
		parent        components.UIComponent
	}
	type args struct {
		in0 components.UIComponent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "returns correct error",
			fields: fields{},
			args: args{
				in0: test.NewMockUIComponent("test", components.NumericWidgetType, nil),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNumeric()

			err := n.AddChild(tt.args.in0)
			if tt.wantErr {
				assert.EqualError(t, err, components.ErrCannotHaveChildren(n.Type().TypeName()).Error())
				return
			}

			assert.NoError(t, err)
		})
	}
}

func Test_numeric_FindChild(t *testing.T) {
	type args struct {
		in0 string
	}
	tests := []struct {
		name  string
		args  args
		want  components.UIComponent
		want1 bool
	}{
		{
			name: "returns correct error",
			args: args{
				in0: "test",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNumeric()

			got, got1 := n.FindChild(tt.args.in0)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_numeric_FindChildById(t *testing.T) {
	type args struct {
		in0 string
	}
	tests := []struct {
		name  string
		args  args
		want  components.UIComponent
		want1 bool
	}{
		{
			name: "returns correct error",
			args: args{
				in0: "test",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNumeric()

			got, got1 := n.FindChildById(tt.args.in0)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_numeric_FindChildByType(t *testing.T) {
	type args struct {
		in0 string
		in1 string
	}
	tests := []struct {
		name  string
		args  args
		want  components.UIComponent
		want1 bool
	}{
		{
			name: "returns correct error",
			args: args{
				in0: "test",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNumeric()
			got, got1 := n.FindChildByType(tt.args.in0, tt.args.in1)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_numeric_GetChildren(t *testing.T) {
	tests := []struct {
		name string
		want []components.UIComponent
	}{
		{
			name: "returns correct error",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNumeric()

			got := n.GetChildren()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_numeric_GetParent(t *testing.T) {
	type fields struct {
		baseWidget    widgets.Widget
		updateHandler func() (int, error)
		initialValue  int
		unit          string
		description   string
		unitAfter     bool
		htmxOpts      htmx.HTMX
		spec          *models.TreeSpec
		parent        components.UIComponent
	}
	tests := []struct {
		name   string
		fields fields
		want   components.UIComponent
	}{
		{
			name: "returns correct parent",
			fields: fields{
				parent: test.NewMockUIComponent("test", components.NumericWidgetType, nil),
			},
			want: test.NewMockUIComponent("test", components.NumericWidgetType, nil),
		},
		{
			name: "returns nil",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &numeric{
				baseWidget:    tt.fields.baseWidget,
				updateHandler: tt.fields.updateHandler,
				initialValue:  tt.fields.initialValue,
				unit:          tt.fields.unit,
				description:   tt.fields.description,
				unitAfter:     tt.fields.unitAfter,
				htmxOpts:      tt.fields.htmxOpts,
				spec:          tt.fields.spec,
				parent:        tt.fields.parent,
			}
			got := n.GetParent()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_numeric_RemoveChild(t *testing.T) {
	type args struct {
		in0 components.UIComponent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "returns correct error",
			args: args{
				in0: test.NewMockUIComponent("test", components.NumericWidgetType, nil),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNumeric()
			err := n.RemoveChild(tt.args.in0)
			if tt.wantErr {
				assert.EqualError(t, err, components.ErrCannotHaveChildren(n.Type().TypeName()).Error())
				return
			}

			assert.NoError(t, err)
		})
	}
}

func mockUpdateHandler(newValue int, err error) func() (int, error) {
	return func() (int, error) {
		return newValue, err
	}
}

func Test_numeric_Render(t *testing.T) {

	type fields struct {
		baseWidget    widgets.Widget
		updateHandler func() (int, error)
		initialValue  int
		unit          string
		description   string
		unitAfter     bool
		htmxOpts      htmx.HTMX
		spec          *models.TreeSpec
		parent        components.UIComponent
	}
	type args struct {
		req models.RequestWrapper
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          *components.RenderResponse
		wantComponent bool
	}{
		{
			name: "returns component, no update",
			fields: fields{
				baseWidget: &widgets.BaseWidget{
					Id:          "test",
					Name:        "test",
					Description: "",
					Route:       "",
					Layout: &layout.WidgetLayout{
						Row:    1,
						Column: 1,
						Width:  1,
						Height: 1,
					},
				},
				updateHandler: nil,
				initialValue:  0,
				unit:          "s",
				description:   "",
				unitAfter:     false,
				htmxOpts: &htmx.Htmx{
					Route:    url.URL{Path: "test/test"},
					Method:   "GET",
					Target:   "this",
					Interval: time.Duration(5).String(),
					Swap:     "outerHTML",
				},
				spec:   nil,
				parent: nil,
			},
			args: args{
				req: nil,
			},
			want: &components.RenderResponse{
				Component: NumericWidget("test", "0", "s", false,
					&htmx.Htmx{
						Route:    url.URL{Path: "test/test"},
						Method:   "GET",
						Target:   "this",
						Interval: time.Duration(5).String(),
						Swap:     "outerHTML",
					}, &layout.WidgetLayout{
						Row:    1,
						Column: 1,
						Width:  1,
						Height: 1,
					}),
			},
			wantComponent: true,
		},
		{
			name: "returns component, update",
			fields: fields{
				baseWidget: &widgets.BaseWidget{
					Id:          "test",
					Name:        "test",
					Description: "",
					Route:       "",
					Layout: &layout.WidgetLayout{
						Row:    1,
						Column: 1,
						Width:  1,
						Height: 1,
					},
				},
				updateHandler: mockUpdateHandler(10, nil),
				initialValue:  0,
				unit:          "s",
				description:   "",
				unitAfter:     false,
				htmxOpts: &htmx.Htmx{
					Route:    url.URL{Path: "test/test"},
					Method:   "GET",
					Target:   "this",
					Interval: time.Duration(5).String(),
					Swap:     "outerHTML",
				},
				spec:   nil,
				parent: nil,
			},
			args: args{
				req: models.NewReqWrapper(&fiber.Ctx{}),
			},
			want: &components.RenderResponse{
				Component: NumericWidget("test", "10", "s", false,
					&htmx.Htmx{
						Route:    url.URL{Path: "test/test"},
						Method:   "GET",
						Target:   "this",
						Interval: time.Duration(5).String(),
						Swap:     "outerHTML",
					}, &layout.WidgetLayout{
						Row:    1,
						Column: 1,
						Width:  1,
						Height: 1,
					}),
			},
			wantComponent: true,
		},
		{
			name: "returns error on update error",
			fields: fields{
				baseWidget: &widgets.BaseWidget{
					Id:          "test",
					Name:        "test",
					Description: "",
					Route:       "",
					Layout: &layout.WidgetLayout{
						Row:    1,
						Column: 1,
						Width:  1,
						Height: 1,
					},
				},
				updateHandler: mockUpdateHandler(0, errors.New("testError")),
				initialValue:  0,
				unit:          "s",
				description:   "",
				unitAfter:     false,
				htmxOpts: &htmx.Htmx{
					Route:    url.URL{Path: "test/test"},
					Method:   "GET",
					Target:   "this",
					Interval: time.Duration(5).String(),
					Swap:     "outerHTML",
				},
				spec:   nil,
				parent: nil,
			},
			args: args{
				req: models.NewReqWrapper(&fiber.Ctx{}),
			},
			want: &components.RenderResponse{
				Err: errors.New("testError"),
			},
			wantComponent: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &numeric{
				baseWidget:    tt.fields.baseWidget,
				updateHandler: tt.fields.updateHandler,
				initialValue:  tt.fields.initialValue,
				unit:          tt.fields.unit,
				description:   tt.fields.description,
				unitAfter:     tt.fields.unitAfter,
				htmxOpts:      tt.fields.htmxOpts,
				spec:          tt.fields.spec,
				parent:        tt.fields.parent,
			}

			got := n.Render(tt.args.req)

			if tt.wantComponent {
				var gotW bytes.Buffer
				var excW bytes.Buffer

				assert.NoError(t, tt.want.Component.Render(context.Background(), &gotW))
				assert.NoError(t, got.Component.Render(context.Background(), &excW))

				assert.Equal(t, excW, gotW)
				return
			}

			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_numeric_Type(t *testing.T) {
	type fields struct {
		baseWidget    widgets.Widget
		updateHandler func() (int, error)
		initialValue  int
		unit          string
		description   string
		unitAfter     bool
		htmxOpts      htmx.HTMX
		spec          *models.TreeSpec
		parent        components.UIComponent
	}
	tests := []struct {
		name   string
		fields fields
		want   components.NodeType
	}{
		{
			name: "returns numeric type",
			want: components.NumericWidgetType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNumeric()
			got := n.Type()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_numeric_UpdateSpec(t *testing.T) {
	type fields struct {
		baseWidget    widgets.Widget
		updateHandler func() (int, error)
		initialValue  int
		unit          string
		description   string
		unitAfter     bool
		htmxOpts      htmx.HTMX
		spec          *models.TreeSpec
		parent        components.UIComponent
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.TreeSpec
	}{
		{
			name: "update spec",
			fields: fields{
				baseWidget: &widgets.BaseWidget{
					Name: "test",
				},
				description: "testDescription",
				htmxOpts: &htmx.Htmx{
					Route: url.URL{Path: "test"},
				},
				spec:   nil,
				parent: test.NewMockUIComponent("testParent", components.DashboardType, nil),
			},
			want: &models.TreeSpec{
				Name:        "test",
				ImageRoute:  "",
				Description: "testDescription",
				Route:       "test",
				Children:    nil,
			},
		},
		{
			name: "update spec, parent is page",
			fields: fields{
				baseWidget: &widgets.BaseWidget{
					Name: "test",
				},
				description: "testDescription",
				htmxOpts: &htmx.Htmx{
					Route: url.URL{Path: "test"},
				},
				spec: nil,
				parent: test.NewMockUIComponent("testParent", components.PageType,
					test.NewMockUIComponent("testParent", components.DashboardType, nil)),
			},
			want: &models.TreeSpec{
				Name:        "test",
				ImageRoute:  "",
				Description: "testDescription",
				Route:       "testParent/test",
				Children:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &numeric{
				baseWidget:    tt.fields.baseWidget,
				updateHandler: tt.fields.updateHandler,
				initialValue:  tt.fields.initialValue,
				unit:          tt.fields.unit,
				description:   tt.fields.description,
				unitAfter:     tt.fields.unitAfter,
				htmxOpts:      tt.fields.htmxOpts,
				spec:          tt.fields.spec,
				parent:        tt.fields.parent,
			}
			if got := n.UpdateSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
