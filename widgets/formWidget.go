package widgets

import (
	"errors"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
)

type FormWidgetInterface interface {
	Widget
}

type FormWidget struct {
	baseWidget BaseWidget
}

var ErrWidgetNameNotSet = errors.New("widget name not set")

func NewFormWidget(
	width int, height int, setters ...func(
		n *FormWidget,
	),
) (FormWidgetInterface, error) {
	var widget FormWidget
	for _, setter := range setters {
		setter(&widget)
	}

	if widget.baseWidget.Name == "" {
		return nil, ErrWidgetNameNotSet
	}

	widget.setId()
	widget.setWidth(width)
	widget.setHeight(height)

	return &widget, nil
}

func (fw *FormWidget) setName(name string) {
	fw.baseWidget.Name = name
}

func (fw *FormWidget) setHeight(height int) {
	fw.baseWidget.Height = height
}

func (fw *FormWidget) setWidth(width int) {
	fw.baseWidget.Width = width
}

func (fw *FormWidget) setId() {
	uld := ulid.Make()
	id := "formWidget_" + fw.baseWidget.Name + "_" + uld.String()
	fw.baseWidget.Id = id
}

func (fw *FormWidget) setDescription(description string) {
	fw.baseWidget.Description = description
}

func (fw *FormWidget) Encode() templ.Component {
	return nil
}

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
