package form

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

// Form is an interface that defines a structure for form-based components.
// It provides methods to modify, display and handle form-related actions.
type Form interface {

	// setUpdateHandler sets a custom handler that is used to handle the form update request.
	// It receives a FormRequest and returns an UpdateResponse.
	setUpdateHandler(handler func(c FormRequest) *models.UpdateResponse)

	// setInitialValue sets the initial value for the form.
	// It takes an UpdateResponse as its argument.
	setInitialValue(value models.UpdateResponse)

	// addFormFields allows adding multiple fields to the form.
	addFormFields(field ...*models.FormField)

	// addFormButtons allows adding multiple buttons to the form.
	addFormButtons(button ...*models.FormButton)

	// addFormCheckboxes allows adding multiple checkboxes to the form.
	addFormCheckboxes(checkbox ...*models.FormCheckbox)

	// handlePost handles the POST request of the form.
	// Returns an UpdateResponse after the form data has been processed.
	handlePost(c FormRequest) *models.UpdateResponse

	// updateAction defines the update action for the form.
	// Returns a template component for rendering.
	updateAction(data *models.UpdateResponse) templ.Component

	// getHtmx returns the HTMX instance for the form.
	getHtmx() htmx.HTMX

	// WithSettings allows applying multiple settings to the form.
	// Returns the modified form.
	WithSettings(settings ...func(f Form)) Form

	// Encode transforms form into a template component for rendering.
	Encode() templ.Component

	// CompileRoutes adds the form-related routes to the provided router.
	CompileRoutes(router *fiber.App)

	// AddParentPath adds a parent path to the current form path.
	// Returns error if any.
	AddParentPath(path string) error
}

type formImpl struct {
	baseWidget    widgets.Widget
	fields        []*models.FormField
	buttons       []*models.FormButton
	checkboxes    []*models.FormCheckbox
	updateHandler func(c FormRequest) *models.UpdateResponse
	initialValue  models.UpdateResponse
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
	widget.htmxOpts.AppendToPath("update", id)

	return widget
}
