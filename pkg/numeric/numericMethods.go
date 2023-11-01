package numeric

import (
	"errors"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"strconv"
)

// Numeric interface implementation

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

func (n *numeric) setUnit(unit string) {
	n.unit = unit
}

// UIComponent interface implementation

func (n *numeric) Render(req components.RequestWrapper) *components.RenderResponse {

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
	return errors.New("numeric widget cannot have children")
}

func (n *numeric) KillChild(components.UIComponent) error {
	return errors.New("numeric widget cannot have children")
}
