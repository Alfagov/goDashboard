package widgets

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/a-h/templ"
	"time"
)

type numericWidget struct {
	baseWidget    BaseWidget
	updateHandler func() (int, error)
	initialValue  int
	unit          string
	unitAfter     bool
	htmx          models.HtmxPoll
}

type NumericWidget interface {
	Widget

	setUpdateHandler(handler func() (int, error))
	setInitialValue(value int)
	setUnit(unit string)
	withUnitAfter()

	UpdateAction(value int) templ.Component
	HandleUpdate() (int, error)

	SetPageRoute(route string)
	GetRoute() string

	WithSpecificSettings(settings ...func(f NumericWidget)) NumericWidget
	Encode() templ.Component
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
