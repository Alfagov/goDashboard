package widgets

import (
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
	"strconv"
)

type NumericWidget struct {
	baseWidget BaseWidget
}

func NewNumericWidget(
	width int, height int, setters ...func(
		n Widget,
	),
) Widget {
	var widget NumericWidget
	for _, setter := range setters {
		setter(&widget)
	}

	widget.setId()
	widget.setWidth(width)
	widget.setHeight(height)

	return &widget
}

func (n *NumericWidget) setId() {
	uld := ulid.Make()
	id := "numericWidget_" + uld.String()
	n.baseWidget.Id = id
}

func (n *NumericWidget) setHeight(height int) {
	n.baseWidget.Height = height
}

func (n *NumericWidget) setWidth(width int) {
	n.baseWidget.Width = width
}

func (n *NumericWidget) setName(name string) {
	n.baseWidget.Name = name
}

func (n *NumericWidget) setDescription(description string) {
	n.baseWidget.Description = description
}

func (n *NumericWidget) Encode() templ.Component {
	element := templates.NumericWidget(n.baseWidget.Name, strconv.Itoa(10))
	return element
}
