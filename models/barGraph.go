package models

type BarGraphData struct {
	XAxis interface{}
	YAxis []BarYAxis
}

type BarYAxis struct {
	Name string
	Data interface{}
}

type LineGraphData struct {
	XAxis interface{}
	YAxis []LineYAxis
}

type LineYAxis struct {
	Name string
	Data interface{}
}

type WaterfallGraphData struct {
	XAxis interface{}
	YAxis WaterfallYAxis
}

type WaterfallYAxis struct {
	Name    string
	PosName string
	NegName string
	Data    []int
}
