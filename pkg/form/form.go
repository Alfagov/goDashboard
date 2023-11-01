package form

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
)

// Form is an interface that defines a structure for form-based components.
// It provides methods to modify, display and handle form-related actions.
type Form interface {
	components.UIComponent

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
	widget.htmxOpts.AppendToPath("widget", id)

	return widget
}
