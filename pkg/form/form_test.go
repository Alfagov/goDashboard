package form

import (
	"github.com/Alfagov/goDashboard/pkg/htmx"
	"github.com/Alfagov/goDashboard/pkg/layout"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

type testEmptyFormType struct {
}

func TestNewFormWidget(t *testing.T) {
	uuid.SetRand(test.FixedRand{})
	type args struct {
		name    string
		setters []func(n widgets.Widget)
	}
	type testCase[F any] struct {
		name string
		args args
		want Form[F]
	}
	tests := []testCase[testEmptyFormType]{
		{
			name: "returns new form with correct name",
			args: args{
				name: "test",
			},
			want: &formImpl[testEmptyFormType]{
				baseWidget: &widgets.BaseWidget{
					Id:     "formWidget_test_" + uuid.New().String(),
					Name:   "test",
					Layout: layout.NewWidgetLayout(),
				},
				htmxOpts: &htmx.Htmx{
					Route: url.URL{Path: "widget/formWidget_test_" + uuid.New().String()},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewFormWidget[testEmptyFormType](tt.args.name, tt.args.setters...)

			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_newForm(t *testing.T) {
	type testCase[F any] struct {
		name string
		want *formImpl[F]
	}
	tests := []testCase[testEmptyFormType]{
		{
			name: "returns empty form",
			want: &formImpl[testEmptyFormType]{
				baseWidget: widgets.NewWidget(),
				htmxOpts:   htmx.NewEmpty(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newForm[testEmptyFormType]()

			assert.Equal(t, tt.want, got)
		})
	}
}
