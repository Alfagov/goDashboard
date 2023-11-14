package utils

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/go-echarts/go-echarts/v2/render"
	"html/template"
	"io"
	"net/http"
)

type TemplRender struct {
	Code int
	Data templ.Component
}

func (t TemplRender) Render(
	w io.Writer, _ string, data interface{},
	_ ...string,
) error {

	d, ok := data.(templ.Component)
	if !ok {
		return nil
	}

	if d == nil {
		return nil
	}

	return d.Render(context.Background(), w)

}

func (t TemplRender) Load() error {
	return nil
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

type snippetRenderer struct {
	c      interface{}
	before []func()
	height int
}

func NewSnippetRenderer(c interface{}, height int, before ...func()) render.Renderer {
	return &snippetRenderer{c: c, before: before, height: height * 150}
}

func (r *snippetRenderer) Render(w io.Writer) error {
	const tplName = "chart"
	for _, fn := range r.before {
		fn()
	}

	tpl := template.Must(
		template.New(tplName).
			Funcs(
				template.FuncMap{
					"safeJS": func(s interface{}) template.JS {
						return template.JS(fmt.Sprint(s))
					},
				},
			).
			Parse(getSizedTemplate(r.height)),
	)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}

// adapted from
// https://github.com/go-echarts/go-echarts/blob/master/templates/base.go
// https://github.com/go-echarts/go-echarts/blob/master/templates/header.go
var baseTpl = `
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:100%;height:400px;"></div>
</div>
<script type="text/javascript">
    "use strict";
    var chart{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    var option_{{ .ChartID | safeJS }} = {{ .JSON }};
    chart{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
	addChart(chart{{ .ChartID | safeJS }}, {{ .ChartID }}, document.getElementById('{{ .ChartID | safeJS }}'));
</script>
`

func getSizedTemplate(h int) string {
	return fmt.Sprintf(`
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:%s;height:%dpx;"></div>
</div>
<script type="text/javascript">
    "use strict";
    var chart{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    var option_{{ .ChartID | safeJS }} = {{ .JSON }};
    chart{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
	addChart(chart{{ .ChartID | safeJS }}, {{ .ChartID }}, document.getElementById('{{ .ChartID | safeJS }}'));
</script>
`, "100%", h)
}

//{{- range .JSAssets.Values }}
//   <script src="{{ . }}"></script>
//{{- end }}
