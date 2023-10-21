package graph

import (
	"context"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/utils"
	"github.com/a-h/templ"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
)

func barGraphFromData(data *models.BarGraphData, stacked bool) *charts.Bar {

	graph := charts.NewBar()

	graph.SetXAxis(data.XAxis)

	for _, series := range data.YAxis {
		items := make([]opts.BarData, 0)
		dat := series.Data.([]int)
		for _, d := range dat {
			items = append(items, opts.BarData{Value: d})
		}

		if stacked {
			graph.AddSeries(
				series.Name, items, charts.WithBarChartOpts(
					opts.BarChart{
						Stack: "x",
					},
				),
			)
			continue
		}

		graph.AddSeries(series.Name, items)
	}

	return graph
}

func lineGraphFromData(data *models.LineGraphData, stacked bool) *charts.Line {

	graph := charts.NewLine()

	graph.SetXAxis(data.XAxis)

	for _, series := range data.YAxis {
		items := make([]opts.LineData, 0)
		dat := series.Data.([]int)
		for _, d := range dat {
			items = append(items, opts.LineData{Value: d})
		}
		graph.AddSeries(series.Name, items)
	}

	return graph
}

func templComponentGraph(c interface{}, before ...func()) templ.ComponentFunc {
	renderer := utils.NewSnippetRenderer(c, before...)
	return func(ctx context.Context, w io.Writer) error {
		return renderer.Render(w)
	}
}
