package toolbox

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewToolbox(t *testing.T) {
	type args struct {
		options []func(t Toolbox)
	}
	tests := []struct {
		name string
		args args
		want Toolbox
	}{
		{
			name: "returns new toolbox with correct options",
			args: args{
				options: []func(t Toolbox){
					SetZoom(true),
					WithSaveImage("test"),
					WithRestore(),
				},
			},
			want: &toolbox{
				options: opts.Toolbox{
					Show: true,
					Feature: &opts.ToolBoxFeature{
						DataZoom: &opts.ToolBoxFeatureDataZoom{
							Show:       true,
							YAxisIndex: "yAxisIndex",
						},
						SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
							Show: true,
							Name: "test",
						},
						Restore: &opts.ToolBoxFeatureRestore{
							Show: true,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewToolbox(tt.args.options...)

			assert.Equal(t, tt.want, got)
		})
	}
}
