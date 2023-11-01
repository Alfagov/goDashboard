package form

import (
	"errors"
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
)

// Form is an interface that defines a structure for form-based components.
// It provides methods to modify, display and handle form-related actions.
type Form interface {

	// setUpdateHandler sets a custom handler used to handle the form update request.
	// It receives a components.RequestWrapper and returns an UpdateResponse.
	setUpdateHandler(handler func(c components.RequestWrapper) *models.UpdateResponse)

	// setInitialValue sets the initial value for the form.
	// It takes an UpdateResponse as its argument.
	setInitialValue(value models.UpdateResponse)

	// addFormFields allows adding multiple fields to the form.
	addFormFields(field ...*models.FormField)

	// addFormButtons allows adding multiple buttons to the form.
	addFormButtons(button ...*models.FormButton)

	// addFormCheckboxes allows adding multiple checkboxes to the form.
	addFormCheckboxes(checkbox ...*models.FormCheckbox)

	// updateAction defines the update action for the form.
	// Returns a template component for rendering.
	updateAction(data *models.UpdateResponse) templ.Component

	// WithSettings allows applying multiple settings to the form.
	// Returns the modified form.
	WithSettings(settings ...func(f Form)) Form
}

type formImpl struct {
	baseWidget    widgets.Widget
	fields        []*models.FormField
	buttons       []*models.FormButton
	checkboxes    []*models.FormCheckbox
	updateHandler func(c components.RequestWrapper) *models.UpdateResponse
	initialValue  models.UpdateResponse
	description   string
	spec          *models.TreeSpec
	parent        components.UIComponent
	popUpResponse bool
	htmxOpts      htmx.HTMX
}

func (fw *formImpl) Render(req components.RequestWrapper) *components.RenderResponse {

	if req != nil && req.Method() == "POST" {
		data := fw.updateHandler(req)
		return &components.RenderResponse{
			Component: fw.updateAction(data),
		}
	}

	fields := fw.fields
	buttons := fw.buttons
	checkboxes := fw.checkboxes

	var fieldsComponent []templ.Component
	for _, field := range fields {
		fieldsComponent = append(fieldsComponent, templates.FormField(field))
	}

	return &components.RenderResponse{
		Component: templates.GenericForm(
			fw.Name(),
			fieldsComponent,
			checkboxes,
			buttons,
			fw.baseWidget.GetLayout(),
			fw.htmxOpts.GetHtmx(),
		),
	}
}

func (fw *formImpl) Type() components.NodeType {
	return components.FormWidgetType
}

func (fw *formImpl) Name() string {
	return fw.baseWidget.GetName()
}

func (fw *formImpl) UpdateSpec() *models.TreeSpec {
	route := components.GetRouteFromParents(fw)

	fw.htmxOpts.AddBeforePath(route)
	return &models.TreeSpec{
		Name:        fw.Name(),
		ImageRoute:  "",
		Description: fw.description,
		Route:       fw.htmxOpts.GetUrl(),
		Children:    nil,
	}
}

func (fw *formImpl) GetSpec() *models.TreeSpec {
	return fw.spec
}

func (fw *formImpl) GetChildren() []components.UIComponent {
	return nil
}

func (fw *formImpl) FindChild(string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl) FindChildByType(string, string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl) SetParent(parent components.UIComponent) {
	fw.parent = parent
}

func (fw *formImpl) GetParent() components.UIComponent {
	return fw.parent
}

func (fw *formImpl) AddChild(components.UIComponent) error {
	return errors.New("not applicable")
}

func (fw *formImpl) KillChild(components.UIComponent) error {
	return errors.New("not applicable")
}

func newForm() *formImpl {
	var w formImpl
	w.baseWidget = widgets.NewWidget()
	w.htmxOpts = htmx.NewEmpty()
	return &w
}

func NewFormWidget(name string, setters ...func(n widgets.Widget)) Form {
	widget := newForm()
	widget.baseWidget.SetName(name)

	for _, setter := range setters {
		setter(widget.baseWidget)
	}

	id := "formWidget_" + name + "_" + ulid.Make().String()
	widget.baseWidget.SetId(id)
	widget.htmxOpts.AppendToPath("update", id)

	return widget
}
