package widgets

import (
	"github.com/Alfagov/goDashboard/pkg/layout"
)

// BaseWidget represents the fundamental structure of a UI widget.
type BaseWidget struct {
	// Id is the unique identifier of the widget.
	Id string
	// Name is the display name of the widget.
	Name string
	// Description provides details about the widget.
	Description string
	// Route is the URL path to which the widget is associated.
	Route string
	// Layout defines the visual arrangement of the widget on the screen.
	Layout *layout.WidgetLayout
}

// Widget is an interface that specifies the methods to manipulate the properties of a widget.
type Widget interface {
	// SetName assigns a new display name to the widget.
	SetName(name string)
	// setHeight sets the height property of the widget.
	setHeight(height int)
	// setWidth sets the width property of the widget.
	setWidth(width int)
	// SetId defines the unique identifier for the widget.
	SetId(id string)
	// setDescription provides a detailed description of the widget.
	setDescription(description string)
	// setRow assigns the widget's row position in a grid layout.
	setRow(row int)
	// setColumn assigns the widget's column position in a grid layout.
	setColumn(column int)
	// withLayout applies a WidgetLayout to the widget, defining its arrangement.
	withLayout(layout *layout.WidgetLayout)
	// GetRow retrieves the row position of the widget in a grid layout.
	GetRow() int
	// GetLayout fetches the WidgetLayout associated with the widget.
	GetLayout() *layout.WidgetLayout
	// GetId retrieves the unique identifier of the widget.
	GetId() string
	// GetName returns the display name of the widget.
	GetName() string
	// GetDescription returns the description of the widget
	GetDescription() string
}

// NewWidget creates and returns a new instance of a Widget with a default layout configuration.
// It initializes a BaseWidget and sets its layout to an empty WidgetLayout model before returning it as a Widget interface.
func NewWidget() Widget {
	var baseWidget BaseWidget
	baseWidget.Layout = &layout.WidgetLayout{}
	return &baseWidget
}

func (b *BaseWidget) GetName() string {
	return b.Name
}

func (b *BaseWidget) GetLayout() *layout.WidgetLayout {
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

func (b *BaseWidget) withLayout(layout *layout.WidgetLayout) {
	b.Layout = layout
}

func (b *BaseWidget) GetRow() int {
	return b.Layout.Row
}

func (b *BaseWidget) GetDescription() string {
	return b.Description
}
