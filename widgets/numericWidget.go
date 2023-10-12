package widgets

import (
	"dario.cat/mergo"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
	"strconv"
	"time"
)

type numericWidget struct {
	baseWidget    BaseWidget
	updateHandler func() (int, error)
	initialValue  int
	htmx          models.HtmxPoll
}

type NumericWidget interface {
	Widget
	setUpdateHandler(handler func() (int, error))
	setInitialValue(value int)
	UpdateAction(value int) templ.Component
	HandleUpdate() (int, error)
	Encode() templ.Component
	SetPageRoute(route string)
	WithSpecificSettings(settings ...func(f NumericWidget)) NumericWidget
	GetRoute() string
}

func NewNumericWidget(
	updateInterval time.Duration,
	baseSetters ...func(
		n Widget,
	),
) NumericWidget {
	var widget numericWidget
	widget.baseWidget.Layout = &models.WidgetLayout{}

	for _, setter := range baseSetters {
		setter(&widget)
	}

	widget.setId()
	widget.htmx.Interval = updateInterval.String()
	widget.htmx.Route = "/update/" + widget.baseWidget.Id
	widget.htmx.Target = "this"
	widget.htmx.Swap = "outerHTML"

	return &widget
}

func (n *numericWidget) SetPageRoute(route string) {
	r := "/update/" + route + "/" + n.baseWidget.Id
	n.htmx.Route = r
}

func (n *numericWidget) HandleUpdate() (int, error) {
	return n.updateHandler()
}

func (n *numericWidget) GetRoute() string {
	return n.htmx.Route
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

func (fw *numericWidget) setRow(row int) {
	fw.baseWidget.Layout.Row = row
}

func (fw *numericWidget) setColumn(column int) {
	fw.baseWidget.Layout.Column = column
}

func (n *numericWidget) setInitialValue(value int) {
	n.initialValue = value
}

func (n *numericWidget) setUpdateHandler(handler func() (int, error)) {
	n.updateHandler = handler
}

func (n *numericWidget) UpdateAction(value int) templ.Component {
	element := templates.NumericWidget(
		n.baseWidget.Name, strconv.Itoa(value),
		&n.htmx,
		n.baseWidget.Layout,
	)
	return element
}

func (n *numericWidget) GetRow() int {
	return n.baseWidget.Layout.Row
}

func (n *numericWidget) Encode() templ.Component {
	element := templates.NumericWidget(
		n.baseWidget.Name, strconv.Itoa(n.initialValue),
		&n.htmx,
		n.baseWidget.Layout,
	)
	return element
}

func SetNumericUpdateHandler(handler func() (int, error)) func(
	f NumericWidget,
) {
	return func(f NumericWidget) {
		f.setUpdateHandler(handler)
	}
}

func SetNumericInitValue(value int) func(
	f NumericWidget,
) {
	return func(f NumericWidget) {
		f.setInitialValue(value)
	}
}
