package toolbox

import "github.com/go-echarts/go-echarts/v2/opts"

type toolbox struct {
	options opts.Toolbox
}

type Toolbox interface {
	setZoom(y string)
	withSaveImage(name string)
	withRestore()
	GetToolbox() opts.Toolbox
}

func NewToolbox(options ...func(t Toolbox)) Toolbox {
	var t toolbox
	if options == nil || len(options) == 0 {
		return &t
	}
	t.options.Show = true
	t.options.Feature = &opts.ToolBoxFeature{}

	for _, option := range options {
		option(&t)
	}

	return &t
}

func (t *toolbox) GetToolbox() opts.Toolbox {
	return t.options
}

func SetZoom(y bool) func(t Toolbox) {
	return func(t Toolbox) {
		if y {
			t.setZoom("yAxisIndex")
			return
		}

		t.setZoom("none")
	}
}

func WithSaveImage(name string) func(t Toolbox) {
	return func(t Toolbox) {
		t.withSaveImage(name)
	}
}

func WithRestore() func(t Toolbox) {
	return func(t Toolbox) {
		t.withRestore()
	}
}

func (t *toolbox) setZoom(y string) {
	t.options.Feature.DataZoom = &opts.ToolBoxFeatureDataZoom{
		Show:       true,
		YAxisIndex: y,
	}
}

func (t *toolbox) withSaveImage(name string) {
	t.options.Feature.SaveAsImage = &opts.ToolBoxFeatureSaveAsImage{
		Show: true,
		Name: name,
	}
}

func (t *toolbox) withRestore() {
	t.options.Feature.Restore = &opts.ToolBoxFeatureRestore{
		Show: true,
	}
}
