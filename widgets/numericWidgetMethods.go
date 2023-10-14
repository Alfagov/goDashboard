package widgets

import (
	"dario.cat/mergo"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
	"strconv"
)

func (n *numericWidget) SetPageRoute(route string) {
	r := route + "/update/" + n.baseWidget.Id
	n.htmx.Route = r
}

func (n *numericWidget) HandleUpdate() (int, error) {
	return n.updateHandler()
}

func (n *numericWidget) GetRoute() string {
	return n.htmx.Route
}

func (n *numericWidget) withUnitAfter() {
	n.unitAfter = true
}

func (n *numericWidget) withLayout(layout *models.WidgetLayout) {
	err := mergo.Merge(n.baseWidget.Layout, layout, mergo.WithOverride)
	if err != nil {
		panic(err)
	}
}

func (n *numericWidget) WithSpecificSettings(
	settings ...func(
		f NumericWidget,
	),
) NumericWidget {
	for _, setter := range settings {
		setter(n)
	}

	return n
}

func (n *numericWidget) setId() {
	uld := ulid.Make()
	id := "numericWidget_" + uld.String()
	n.baseWidget.Id = id
}

func (n *numericWidget) setHeight(height int) {
	n.baseWidget.Layout.Height = height
}

func (n *numericWidget) setWidth(width int) {
	n.baseWidget.Layout.Width = width
}

func (n *numericWidget) setName(name string) {
	n.baseWidget.Name = name
}

func (n *numericWidget) setDescription(description string) {
	n.baseWidget.Description = description
}

func (n *numericWidget) setRow(row int) {
	n.baseWidget.Layout.Row = row
}

func (n *numericWidget) setColumn(column int) {
	n.baseWidget.Layout.Column = column
}

func (n *numericWidget) setInitialValue(value int) {
	n.initialValue = value
}

func (n *numericWidget) setUpdateHandler(handler func() (int, error)) {
	n.updateHandler = handler
}

func (n *numericWidget) UpdateAction(value int) templ.Component {
	element := templates.NumericWidget(
		n.baseWidget.Name, strconv.Itoa(value), n.unit, n.unitAfter,
		&n.htmx,
		n.baseWidget.Layout,
	)
	return element
}

func (n *numericWidget) setUnit(unit string) {
	n.unit = unit
}

func (n *numericWidget) GetRow() int {
	return n.baseWidget.Layout.Row
}

func (n *numericWidget) Encode() templ.Component {
	element := templates.NumericWidget(
		n.baseWidget.Name, strconv.Itoa(n.initialValue), n.unit, n.unitAfter,
		&n.htmx,
		n.baseWidget.Layout,
	)
	return element
}
