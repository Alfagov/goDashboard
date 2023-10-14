package widgets

import (
	"fmt"
	chartrender "github.com/go-echarts/go-echarts/v2/render"
	"html/template"
	"io"
)

type snippetRenderer struct {
	c      interface{}
	before []func()
}

func newSnippetRenderer(c interface{}, before ...func()) chartrender.Renderer {
	return &snippetRenderer{c: c, before: before}
}

func (r *snippetRenderer) Render(w io.Writer) error {
	const tplName = "chart"
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(
			template.New(tplName).
				Funcs(
					template.FuncMap{
						"safeJS": func(s interface{}) template.JS {
							return template.JS(fmt.Sprint(s))
						},
					},
				).
				Parse(baseTpl),
		)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}

// adapted from
// https://github.com/go-echarts/go-echarts/blob/master/templates/base.go
// https://github.com/go-echarts/go-echarts/blob/master/templates/header.go
var baseTpl = `
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>
{{- range .JSAssets.Values }}
   <script src="{{ . }}"></script>
{{- end }}
<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    let option_{{ .ChartID | safeJS }} = {{ .JSON }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
`
