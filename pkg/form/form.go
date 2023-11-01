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
type Form[F any] interface {
	components.UIComponent

	// setUpdateHandler sets a custom handler used to handle the form update request.
	// It receives a components.RequestWrapper and returns an UpdateResponse.
	setUpdateHandler(handler func(c F) *models.UpdateResponse)

	// setInitialValue sets the initial value for the form.
	// It takes an UpdateResponse as its argument.
	setInitialValue(value models.UpdateResponse)

	// addFormFields allows adding multiple fields to the form.
	addFormFields(field ...models.Field)

	// updateAction defines the update action for the form.
	// Returns a template component for rendering.
	updateAction(data *models.UpdateResponse) templ.Component

	// WithSettings allows applying multiple settings to the form.
	// Returns the modified form.
	WithSettings(settings ...func(f Form[F])) Form[F]
}

type formImpl[F any] struct {
	baseWidget    widgets.Widget
	fields        []models.Field
	buttons       []*models.FormButton
	checkboxes    []*models.FormCheckbox
	updateHandler func(c F) *models.UpdateResponse
	initialValue  models.UpdateResponse
	description   string
	spec          *models.TreeSpec
	parent        components.UIComponent
	popUpResponse bool
	htmxOpts      htmx.HTMX
}

func newForm[F any]() *formImpl[F] {
	var w formImpl[F]
	w.baseWidget = widgets.NewWidget()
	w.htmxOpts = htmx.NewEmpty()
	return &w
}

func NewFormWidget[F any](name string, setters ...func(n widgets.Widget)) Form[F] {
	widget := newForm[F]()
	widget.baseWidget.SetName(name)

	for _, setter := range setters {
		setter(widget.baseWidget)
	}

	id := "formWidget_" + name + "_" + ulid.Make().String()
	widget.baseWidget.SetId(id)
	widget.htmxOpts.AppendToPath("widget", id)

	return widget
}
