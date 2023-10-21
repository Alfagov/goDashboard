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
	SetName(name string)
	setHeight(height int)
	setWidth(width int)
	SetId(id string)
	setDescription(description string)
	setRow(row int)
	setColumn(column int)
	withLayout(layout *models.WidgetLayout)
	GetRow() int
	GetLayout() *models.WidgetLayout
	GetId() string
	GetName() string
}

func NewWidget() Widget {
	var baseWidget BaseWidget
	baseWidget.Layout = &models.WidgetLayout{}
	return &baseWidget
}

func (b *BaseWidget) GetName() string {
	return b.Name
}

func (b *BaseWidget) GetLayout() *models.WidgetLayout {
	return b.Layout
}

func (b *BaseWidget) GetId() string {
	return b.Id
}

func (b *BaseWidget) SetName(name string) {
	b.Name = name
}

func (b *BaseWidget) setHeight(height int) {
	b.Layout.Height = height
}

func (b *BaseWidget) setWidth(width int) {
	b.Layout.Width = width
}

func (b *BaseWidget) SetId(id string) {
	b.Id = id
}

func (b *BaseWidget) setDescription(description string) {
	b.Description = description
}

func (b *BaseWidget) setRow(row int) {
	b.Layout.Row = row
}

func (b *BaseWidget) setColumn(column int) {
	b.Layout.Column = column
}

func (b *BaseWidget) withLayout(layout *models.WidgetLayout) {
	b.Layout = layout
}

func (b *BaseWidget) GetRow() int {
	return b.Layout.Row
}
