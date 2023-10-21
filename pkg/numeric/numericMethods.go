package numeric

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"strconv"
)

func (n *numeric) GetHtmx() htmx.HTMX {
	return n.htmxOpts
}

func (n *numeric) HandleUpdate() (int, error) {
	return n.updateHandler()
}

func (n *numeric) withUnitAfter() {
	n.unitAfter = true
}

func (n *numeric) WithSettings(
	settings ...func(
		f Numeric,
	),
) Numeric {
	for _, setter := range settings {
		setter(n)
	}

	return n
}

func (n *numeric) setInitialValue(value int) {
	n.initialValue = value
}

func (n *numeric) setUpdateHandler(handler func() (int, error)) {
	n.updateHandler = handler
}

func (n *numeric) UpdateAction(value int) templ.Component {
	element := templates.NumericWidget(
		n.baseWidget.GetName(), strconv.Itoa(value), n.unit, n.unitAfter,
		n.htmxOpts.GetHtmx(),
		n.baseWidget.GetLayout(),
	)
	return element
}

func (n *numeric) setUnit(unit string) {
	n.unit = unit
}

func (n *numeric) Encode() templ.Component {
	element := templates.NumericWidget(
		n.baseWidget.GetName(), strconv.Itoa(n.initialValue), n.unit, n.unitAfter,
		n.htmxOpts.GetHtmx(),
		n.baseWidget.GetLayout(),
	)
	return element
}
