package numeric

import (
	"errors"
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
	"strconv"
	"time"
)

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

func (n *numeric) Render(components.RequestWrapper) *components.RenderResponse {

	return &components.RenderResponse{
		Component: templates.NumericWidget(
			n.baseWidget.GetName(),
			strconv.Itoa(n.initialValue),
			n.unit,
			n.unitAfter,
			n.htmxOpts.GetHtmx(),
			n.baseWidget.GetLayout()),
	}
}

func (n *numeric) Type() components.NodeType {
	return components.NumericWidgetType
}

func (n *numeric) Name() string {
	return n.baseWidget.GetName()
}

func (n *numeric) UpdateSpec() *models.TreeSpec {
	route := components.GetRouteFromParents(n)

	n.htmxOpts.AddBeforePath(route)
	return &models.TreeSpec{
		Name:        n.Name(),
		ImageRoute:  "",
		Description: n.description,
		Route:       n.htmxOpts.GetUrl(),
		Children:    nil,
	}
}

func (n *numeric) GetSpec() *models.TreeSpec {
	return n.spec
}

func (n *numeric) GetChildren() []components.UIComponent {
	return nil
}

func (n *numeric) FindChild(string) (components.UIComponent, bool) {
	return nil, false
}

func (n *numeric) FindChildByType(string, string) (components.UIComponent, bool) {
	return nil, false
}

func (n *numeric) SetParent(parent components.UIComponent) {
	n.parent = parent
}

func (n *numeric) GetParent() components.UIComponent {
	return n.parent
}

func (n *numeric) AddChild(components.UIComponent) error {
	return errors.New("numeric widget cannot have children")
}

func (n *numeric) KillChild(components.UIComponent) error {
	return errors.New("numeric widget cannot have children")
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
}

func (n *numeric) GetName() string {
	return n.baseWidget.GetName()
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
