package numeric

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
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

// Numeric is an interface that defines methods for numeric-based components.
// The type implementing this interface can set and update values, handle unit formatting,
// and additionally has methods to interface with templates, routes, and HTMX.
type Numeric interface {

	// setUpdateHandler sets a handler function that returns an integer and an error.
	// This handler function runs when an update operation is called for the implementing type.
	setUpdateHandler(handler func() (int, error))

	// setInitialValue sets initial integer value for the implementing type.
	setInitialValue(value int)

	// setUnit sets a string representing the unit of value for the implementing type.
	setUnit(unit string)

	// withUnitAfter toggles unit positioning to place after the value in the implementing type.
	withUnitAfter()

	// updateAction performs an update operation with an integer value and returns a templ.Component.
	updateAction(value int) templ.Component

	// handleUpdate handles an update operation and returns an updated integer and an error if it occurs.
	handleUpdate() (int, error)

	// WithSettings accepts a series of settings functions and applies them to the implementing type,
	// returning the modified implementing type.
	WithSettings(settings ...func(f Numeric)) Numeric

	// Encode translates the implementing type into a templ.Component for usage in templates.
	Encode() templ.Component

	// getHtmx returns an htmx.HTMX related to the implementing type.
	getHtmx() htmx.HTMX

	// CompileRoutes adds routes to the router for the implementing type.
	CompileRoutes(router *fiber.App)

	// AddParentPath adds a parent path string to the implementing type and returns an error if this operation fails.
	AddParentPath(path string) error
}

func (n *numeric) AddParentPath(path string) error {
	return n.htmxOpts.GetHtmx().AddBeforePath(path)
}

func (n *numeric) CompileRoutes(router *fiber.App) {
	router.Get(
		n.htmxOpts.GetUrl(), func(c *fiber.Ctx) error {
			update, err := n.handleUpdate()
			if err != nil {
				return err
			}

			return c.Render("", n.updateAction(update))
		},
	)
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

	widget.htmxOpts.AppendToPath("update", id)
	widget.htmxOpts.SetInterval(updateInterval.String())
	widget.htmxOpts.SetTarget("this")
	widget.htmxOpts.SetSwap("outerHTML")

	return widget
}
