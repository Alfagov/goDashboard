package numeric

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/oklog/ulid/v2"
	"time"
)

// numeric is a struct that encapsulates the numeric widget's behaviors and attributes.
// it is mainly used for displaying interactive numeric information to the users.
type numeric struct {
	baseWidget    widgets.Widget
	updateHandler func() (int, error)
	initialValue  int
	unit          string
	description   string
	unitAfter     bool
	htmxOpts      htmx.HTMX
	spec          *models.TreeSpec
	parent        components.UIComponent
}

// Numeric is an interface that defines methods for numeric-based components.
// The type implementing this interface can set and update values, handle unit formatting,
// and additionally has methods to interface with templates, routes, and HTMX.
type Numeric interface {
	components.UIComponent

	// setUpdateHandler sets a handler function that returns an integer and an error.
	// This handler function runs when an update operation is called for the implementing type.
	setUpdateHandler(handler func() (int, error))

	// setInitialValue sets initial integer value for the implementing type.
	setInitialValue(value int)

	// setUnit sets a string representing the unit of value for the implementing type.
	setUnit(unit string)

	// withUnitAfter toggles unit positioning to place after the value in the implementing type.
	withUnitAfter()

	// WithSettings accepts a series of settings functions and applies them to the implementing type,
	// returning the modified implementing type.
	WithSettings(settings ...func(f Numeric)) Numeric
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

	widget.htmxOpts.AppendToPath("widget", id)
	widget.htmxOpts.SetInterval(updateInterval.String())
	widget.htmxOpts.SetTarget("this")
	widget.htmxOpts.SetSwap("outerHTML")

	return widget
}
