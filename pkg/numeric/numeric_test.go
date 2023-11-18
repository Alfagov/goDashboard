package numeric

import (
	"github.com/Alfagov/goDashboard/pkg/htmx"
	"github.com/Alfagov/goDashboard/pkg/layout"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestNewNumeric(t *testing.T) {
	uuid.SetRand(test.FixedRand{})
	type args struct {
		name           string
		updateInterval time.Duration
		baseSetters    []func(n widgets.Widget)
	}
	tests := []struct {
		name string
		args args
		want Numeric
	}{
		{
			name: "returns new numeric with correct name",
			args: args{
				name:           "test",
				updateInterval: time.Second,
			},
			want: &numeric{
				baseWidget: &widgets.BaseWidget{
					Name:   "test",
					Layout: &layout.WidgetLayout{},
					Id:     "numericWidget_" + uuid.New().String(),
				},
				htmxOpts: &htmx.Htmx{
					Route:    url.URL{Path: "widget/numericWidget_" + uuid.New().String()},
					Target:   "this",
					Swap:     "outerHTML",
					Interval: time.Second.String(),
					Method:   "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewNumeric(tt.args.name, tt.args.updateInterval, tt.args.baseSetters...))
		})
	}
}
