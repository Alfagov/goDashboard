package components

import (
	"github.com/Alfagov/goDashboard/internal/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"go.uber.org/zap"
	"os"
)

func VisualizeTree(tree UIComponent) {
	treeChart := charts.NewTree()
	treeChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: "white"}),
		charts.WithTitleOpts(opts.Title{Title: "Tree-Visualize"}),
	)

	nodes := TreeSpecToChartNodes(
		[]*models.TreeSpec{
			tree.GetSpec(),
		})

	treeChart.AddSeries("tree", nodes)

	f, _ := os.Create("tree.html")
	err := treeChart.Render(f)
	if err != nil {
		logger.L.Error("error in rendering tree", zap.Error(err))
	}
}

func TreeSpecToChartNodes(spec []*models.TreeSpec) []opts.TreeData {
	var nodes []opts.TreeData

	for _, child := range spec {
		nodes = append(nodes, opts.TreeData{
			Name:     child.Name,
			Children: treeSpecToChartNodes(child.Children),
		})
	}

	return nodes
}

func treeSpecToChartNodes(spec []*models.TreeSpec) []*opts.TreeData {
	var nodes []*opts.TreeData
	for _, child := range spec {
		nodes = append(nodes, &opts.TreeData{
			Name:     child.Name,
			Children: treeSpecToChartNodes(child.Children),
		})
	}

	return nodes
}

// GetRouteFromParents constructs a string that represents the hierarchical path of a UI component by concatenating the names
// of its parent components. It stops once it reaches a component of the DashboardType or if the component has no parent.
// The resulting route is a '/' separated string reflecting the hierarchy from the top-level parent to the given component.
func GetRouteFromParents(n UIComponent) string {
	parent := n.GetParent()
	route := ""
	for {
		if parent == nil || parent.Type().Is(DashboardType) {
			break
		}
		route = parent.Name() + "/" + route
		parent = parent.GetParent()
	}

	return route
}
