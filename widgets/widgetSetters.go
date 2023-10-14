package widgets

import "github.com/Alfagov/goDashboard/models"

func SetName(name string) func(f Widget) {
	return func(f Widget) {
		f.setName(name)
	}
}

func SetDescription(description string) func(f Widget) {
	return func(f Widget) {
		f.setDescription(description)
	}
}
func SetHeight(height int) func(Widget) {
	return func(f Widget) {
		f.setHeight(height)
	}
}

func SetWidth(width int) func(Widget) {
	return func(f Widget) {
		f.setWidth(width)
	}
}

func SetLayout(layout *models.WidgetLayout) func(Widget) {
	return func(f Widget) {
		f.withLayout(layout)
	}
}

func SetRow(row int) func(Widget) {
	return func(f Widget) {
		f.setRow(row)
	}
}

func SetColumn(column int) func(Widget) {
	return func(f Widget) {
		f.setColumn(column)
	}
}
