package widgets

import (
	"github.com/Alfagov/goDashboard/pkg/layout"
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

// SetLayout returns a function that assigns a specific WidgetLayout to a Widget.
func SetLayout(layout *layout.WidgetLayout) func(Widget) {
	return func(f Widget) {
		f.withLayout(layout)
	}
}
