package layout

import (
	"strconv"
)

// WidgetLayout defines the grid position and dimensions of a widget within a layout.
type WidgetLayout struct {
	Row    int
	Column int
	Width  int
	Height int
}

// NewWidgetLayout creates and returns a new WidgetLayout with optional layout configurations set through variadic setters.
// Each setter function is applied to the layout in the order they are provided.
func NewWidgetLayout(setters ...func(layout *WidgetLayout)) *WidgetLayout {
	var layout WidgetLayout
	for _, setter := range setters {
		setter(&layout)
	}
	return &layout
}

// SetRow returns a function that sets the row position of a WidgetLayout to the specified integer value.
func SetRow(row int) func(layout *WidgetLayout) {
	return func(layout *WidgetLayout) {
		layout.Row = row
	}
}

// SetColumn returns a function that sets the column position of a WidgetLayout to the specified integer value.
func SetColumn(column int) func(layout *WidgetLayout) {
	return func(layout *WidgetLayout) {
		layout.Column = column
	}
}

// SetWidth returns a function that sets the width dimension of a WidgetLayout to the specified integer value.
func SetWidth(width int) func(layout *WidgetLayout) {
	return func(layout *WidgetLayout) {
		layout.Width = width
	}
}

// SetHeight returns a function that sets the height dimension of a WidgetLayout to the specified integer value.
func SetHeight(height int) func(layout *WidgetLayout) {
	return func(layout *WidgetLayout) {
		layout.Height = height
	}
}

// ToCSS takes a pointer to a WidgetLayout and constructs a string representing CSS grid class names that
// denote the start and end positions for rows and columns, as well as column span if the widget spans more than one column.
// This string can be used as class names in an HTML element to position it according to the layout within a CSS Grid.
func ToCSS(layout *WidgetLayout) string {

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
