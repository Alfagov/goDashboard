package numeric

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
	"time"
)

type numeric struct {
	baseWidget    widgets.Widget
	updateHandler func() (int, error)
	initialValue  int
	unit          string
	unitAfter     bool
	htmxOpts      htmx.HTMX
}

type Numeric interface {
	setUpdateHandler(handler func() (int, error))
	setInitialValue(value int)
	setUnit(unit string)
	withUnitAfter()

	UpdateAction(value int) templ.Component
	HandleUpdate() (int, error)

	WithSettings(settings ...func(f Numeric)) Numeric
	Encode() templ.Component

	GetHtmx() htmx.HTMX
}

func newNumeric() *numeric {
	var w numeric
	w.baseWidget = widgets.NewWidget()
	w.htmxOpts = htmx.NewEmpty()
	return &w
}

func NewNumeric(
	updateInterval time.Duration,
	baseSetters ...func(
		n widgets.Widget,
	),
) Numeric {
	widget := newNumeric()

	for _, setter := range baseSetters {
		setter(widget.baseWidget)
	}

	id := "numericWidget_" + ulid.Make().String()
	widget.baseWidget.SetId(id)

	widget.htmxOpts.SetRoute("/update/" + id)
	widget.htmxOpts.SetInterval(updateInterval.String())
	widget.htmxOpts.SetTarget("this")
	widget.htmxOpts.SetSwap("outerHTML")

	return widget
}
