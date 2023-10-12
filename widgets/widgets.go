package widgets

import (
	"github.com/Alfagov/goDashboard/models"
)

type BaseWidget struct {
	Id          string
	Name        string
	Description string
	Route       string
	Layout      *models.WidgetLayout
}

type Widget interface {
	setName(name string)
	withLayout(layout *models.WidgetLayout)
	setHeight(height int)
	setWidth(width int)
	setId()
	setDescription(description string)
	setRow(row int)
	setColumn(column int)
	GetRow() int
}
