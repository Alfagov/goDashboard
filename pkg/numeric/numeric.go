package numeric

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/htmx"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/google/uuid"
	"time"
)

// numeric is a struct that encapsulates attributes and behaviors specific to a numeric widget.
// this struct is primarily responsible for displaying numeric information interactively to users.
type numeric struct {
	// baseWidget represents the underlying widget extended by numeric.
	baseWidget widgets.Widget
	// updateHandler is a function that, when called, returns the current value as an integer
	// and an error if the update operation fails.
	updateHandler func() (int, error)
	// initialValue is the starting value for the numeric widget before any user interaction.
	initialValue int
	// unit is a string that holds the measurement unit for the numeric value (e.g., "kg", "%", "cm").
	unit string
	// description provides a text explaining or describing the numeric widget.
	description string
	// unitAfter indicates whether the unit should be displayed after the numeric value (true)
	// or before it (false).
	unitAfter bool
	// htmxOpts are options specific to HTMX, an attribute extension for HTML,
	// that can enhance the widget's behavior with AJAX, WebSockets, etc.
	htmxOpts htmx.HTMX
	// spec is a pointer to a TreeSpec model that provides a structured specification
	// of the numeric widget's tree hierarchy.
	spec *models.TreeSpec
	// parent is the UIComponent that contains this numeric widget within the UI hierarchy.
	parent components.UIComponent
}

// Numeric defines the interface for components that display and interact with numeric values.
// Implementers can handle updates to values, adjust unit representations, and integrate with
// templates, routing, and HTMX functionalities.
type Numeric interface {
	// UIComponent is a set of common methods for User Interface components.
	components.UIComponent

	// setUpdateHandler defines a method to assign a function that updates the numeric value.
	// The assigned function should return an integer and an error.
	setUpdateHandler(handler func() (int, error))

	// setInitialValue assigns an initial integer value to the numeric component.
	setInitialValue(value int)

	// setUnit defines a method to set the measurement unit associated with the numeric value.
	setUnit(unit string)

	// withUnitAfter specifies the method to toggle the position of the unit to follow the numeric value.
	withUnitAfter()

	// WithSettings takes one or more setting functions and applies them to the numeric component.
	// It returns the Numeric instance with modified settings for chaining.
	WithSettings(settings ...func(f Numeric)) Numeric
}

func newNumeric() *numeric {
	var w numeric
	w.baseWidget = widgets.NewWidget()
	w.htmxOpts = htmx.NewEmpty()
	return &w
}

// NewNumeric creates a new numeric widget with a specified update interval and applies optional base setters.
// It generates a unique identifier for the widget using ULID and configures HTMX options for AJAX interactions.
//
// Parameters:
// updateInterval: the duration between each update to the numeric widget's data.
// baseSetters: a variadic list of functions that configure the base properties of the widget.
//
// Returns a newly initialized Numeric widget configured with the provided settings and ready for use.
func NewNumeric(name string, updateInterval time.Duration, baseSetters ...func(n widgets.Widget)) Numeric {
	widget := newNumeric()
	widget.baseWidget.SetName(name)

	for _, setter := range baseSetters {
		setter(widget.baseWidget)
	}

	id := "numericWidget_" + uuid.New().String()
	widget.baseWidget.SetId(id)

	widget.htmxOpts.AppendToPath("widget", id)
	widget.htmxOpts.SetInterval(updateInterval.String())
	widget.htmxOpts.SetTarget("this")
	widget.htmxOpts.SetSwap("outerHTML")

	return widget
}
