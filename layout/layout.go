package layout

import (
	"github.com/Alfagov/goDashboard/models"
	"strconv"
)

func NewWidgetLayout(setters ...func(layout *models.WidgetLayout)) *models.WidgetLayout {
	var layout models.WidgetLayout
	for _, setter := range setters {
		setter(&layout)
	}
	return &layout
}

func SetRow(row int) func(layout *models.WidgetLayout) {
	return func(layout *models.WidgetLayout) {
		layout.Row = row
	}
}

func SetColumn(column int) func(layout *models.WidgetLayout) {
	return func(layout *models.WidgetLayout) {
		layout.Column = column
	}
}

func SetWidth(width int) func(layout *models.WidgetLayout) {
	return func(layout *models.WidgetLayout) {
		layout.Width = width
	}
}

func SetHeight(height int) func(layout *models.WidgetLayout) {
	return func(layout *models.WidgetLayout) {
		layout.Height = height
	}
}

func ToCSS(layout *models.WidgetLayout) string {

	var rowEnd string
	var rowStart string
	var colSpan string
	var colStart string

	rowStart = strconv.Itoa(layout.Row)
	rowStart = "row-start-" + rowStart + " "

	if layout.Height > 1 {
		rowEnd = strconv.Itoa(layout.Row + layout.Height)
		rowEnd = "row-end-" + rowEnd + " "
	}

	colStart = strconv.Itoa(layout.Column)
	colStart = "col-start-" + colStart + " "

	if layout.Width > 1 {
		colSpan = strconv.Itoa(layout.Width)
		colSpan = "col-span-" + colSpan + " "
	}

	outString := rowStart + rowEnd + colStart + colSpan

	return outString
}
