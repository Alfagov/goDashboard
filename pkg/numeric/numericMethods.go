package numeric

import (
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"go.uber.org/zap"
	"strconv"
)

// Numeric interface implementation

// withUnitAfter sets the display flag for the unit to appear after the numeric value.
// This is an unexported method meant for internal use to toggle the position of the unit indicator.
func (n *numeric) withUnitAfter() {
	n.unitAfter = true
}

// WithSettings applies a series of configuration functions to the numeric widget.
// Each function in the 'settings' slice is called with the numeric instance to apply specific settings.
// Returns the numeric instance with the applied settings.
func (n *numeric) WithSettings(settings ...func(f Numeric)) Numeric {
	for _, setter := range settings {
		setter(n)
	}

	return n
}

// setInitialValue assigns an initial integer value to the numeric widget.
// This is an unexported method intended for internal use to set the widget's starting value.
func (n *numeric) setInitialValue(value int) {
	n.initialValue = value
}

// setUpdateHandler defines a handler function that updates the value of the numeric widget.
// The handler is expected to return an integer value and an error. This method is unexported
// and should be used internally within the package.
func (n *numeric) setUpdateHandler(handler func() (int, error)) {
	n.updateHandler = handler
}

// setUnit sets the measurement unit associated with the numeric value.
// This is an unexported method for internal use to define the unit displayed with the numeric value.
func (n *numeric) setUnit(unit string) {
	n.unit = unit
}

// UIComponent interface implementation

// Render generates the HTML for the numeric widget based on the current state and a request context.
// If a request is provided, it calls the updateHandler to get the latest value to be rendered.
// If the updateHandler encounters an error, the error is returned in the RenderResponse.
// Otherwise, it populates the RenderResponse with the appropriate HTML component using the NumericWidget template,
// including the name, current value, unit, unit position, and HTMX options.
// When no request context is present, it renders the widget with the initial value.
// Returns a RenderResponse containing the component's HTML or an error if one occurred during rendering.
func (n *numeric) Render(req models.RequestWrapper) *components.RenderResponse {

	if req != nil {
		value, err := n.updateHandler()
		if err != nil {
			return &components.RenderResponse{
				Err: err,
			}
		}

		return &components.RenderResponse{
			Component: templates.NumericWidget(
				n.baseWidget.GetName(),
				strconv.Itoa(value),
				n.unit,
				n.unitAfter,
				n.htmxOpts.GetHtmx(),
				n.baseWidget.GetLayout()),
		}
	}

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

	err := n.htmxOpts.AddBeforePath(route)
	if err != nil {
		logger.L.Error("error adding route to htmxOpts", zap.Error(err))
	}

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

func (n *numeric) Id() string {
	return n.baseWidget.GetId()
}

func (n *numeric) FindChildById(string) (components.UIComponent, bool) {
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
	return components.ErrCannotHaveChildren(n.Type().TypeName())
}

func (n *numeric) KillChild(components.UIComponent) error {
	return components.ErrCannotHaveChildren(n.Type().TypeName())
}
