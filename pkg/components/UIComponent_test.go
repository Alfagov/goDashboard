package components_test

import (
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRouteFromParents(t *testing.T) {
	type args struct {
		n components.UIComponent
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test correct routes with pageContainer parent",
			args: args{
				n: test.NewMockUIComponent("test", components.PageType,
					test.NewMockUIComponent("parent", components.PageContainerType, nil)),
			},
			want: "parent/",
		},
		{
			name: "test correct routes with Dashboard parent",
			args: args{
				n: test.NewMockUIComponent("test", components.PageType,
					test.NewMockUIComponent("parent", components.DashboardType, nil)),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := components.GetRouteFromParents(tt.args.n)

			assert.Equal(t, tt.want, got)
		})
	}
}
