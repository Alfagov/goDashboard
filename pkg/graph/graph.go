package graph

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/toolbox"
	"github.com/a-h/templ"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/oklog/ulid/v2"
)

type Graph interface {
	Encode(height int) templ.Component
	HandleUpdate() map[string]interface{}
	WithToolboxOpts(toolboxOpts toolbox.Toolbox) Graph
	GetId() string
	GetName() string
}

type lineGraphImpl struct {
	Id          string
	Name        string
	tBox        toolbox.Toolbox
	dataHandler func() *models.LineGraphData
}

type barGraphImpl struct {
	Id          string
	Name        string
	tBox        toolbox.Toolbox
	stacked     bool
	dataHandler func() *models.BarGraphData
}

func NewLineGraph(name string, dataHandler func() *models.LineGraphData) Graph {
	return &lineGraphImpl{
		Id:          "lineGraph_" + name + "_" + ulid.Make().String(),
		Name:        name,
		dataHandler: dataHandler,
	}
}

func NewBarGraph(
	name string, dataHandler func() *models.BarGraphData,
	stacked bool,
) Graph {
	return &barGraphImpl{
		Id:          "barGraph_" + name + "_" + ulid.Make().String(),
		Name:        name,
		stacked:     stacked,
		dataHandler: dataHandler,
	}
}

func (bg *barGraphImpl) Encode(h int) templ.Component {
	g := barGraphFromData(bg.dataHandler(), bg.stacked)
	g.ChartID = bg.Id
	if bg.tBox != nil {
		g.SetGlobalOptions(
			charts.WithToolboxOpts(bg.tBox.GetToolbox()),
		)
	}

	return templComponentGraph(g, h, g.Validate)
}

func (lg *lineGraphImpl) Encode(h int) templ.Component {
	g := lineGraphFromData(lg.dataHandler())
	g.ChartID = lg.Id
	if lg.tBox != nil {
		g.SetGlobalOptions(
			charts.WithToolboxOpts(lg.tBox.GetToolbox()),
		)
	}

	return templComponentGraph(g, h, g.Validate)
}

func (bg *barGraphImpl) HandleUpdate() map[string]interface{} {
	g := barGraphFromData(bg.dataHandler(), bg.stacked)

	data := g.JSON()
	data["id"] = bg.Id

	return data
}

func (lg *lineGraphImpl) HandleUpdate() map[string]interface{} {
	g := lineGraphFromData(lg.dataHandler())

	data := g.JSON()
	data["id"] = lg.Id

	return data
}

func (bg *barGraphImpl) GetId() string {
	return bg.Id
}

func (lg *lineGraphImpl) GetId() string {
	return lg.Id
}

func (bg *barGraphImpl) WithToolboxOpts(toolboxOpts toolbox.Toolbox) Graph {
	bg.tBox = toolboxOpts
	return bg
}

func (lg *lineGraphImpl) WithToolboxOpts(toolboxOpts toolbox.Toolbox) Graph {
	lg.tBox = toolboxOpts
	return lg
}

func (bg *barGraphImpl) GetName() string {
	return bg.Name
}

func (lg *lineGraphImpl) GetName() string {
	return lg.Name
}
