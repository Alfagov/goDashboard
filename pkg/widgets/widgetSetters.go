package widgets

import (
	"github.com/Alfagov/goDashboard/layout"
)

// SetName returns a function that sets the name of a Widget to the specified string.
func SetName(name string) func(f Widget) {
	return func(f Widget) {
		f.SetName(name)
	}
}

// SetDescription returns a function that sets the description of a Widget to the specified string.
func SetDescription(description string) func(f Widget) {
	return func(f Widget) {
		f.setDescription(description)
	}
}

// SetHeight returns a function that sets the height of a Widget to the specified integer value.
func SetHeight(height int) func(Widget) {
	return func(f Widget) {
		f.setHeight(height)
	}
}

// SetWidth returns a function that sets the width of a Widget to the specified integer value.
func SetWidth(width int) func(Widget) {
	return func(f Widget) {
		f.setWidth(width)
	}
}

// SetLayout returns a function that assigns a specific WidgetLayout to a Widget.
func SetLayout(layout *layout.WidgetLayout) func(Widget) {
	return func(f Widget) {
		f.withLayout(layout)
	}
}

// SetRow returns a function that sets the row position of a Widget within a grid layout to the specified integer.
func SetRow(row int) func(Widget) {
	return func(f Widget) {
		f.setRow(row)
	}
}

// SetColumn returns a function that sets the column position of a Widget within a grid layout to the specified integer.
func SetColumn(column int) func(Widget) {
	return func(f Widget) {
		f.setColumn(column)
	}
}
