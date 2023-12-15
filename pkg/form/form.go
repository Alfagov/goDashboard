package form

import (
	"github.com/Alfagov/goDashboard/internal/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/htmx"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	ActionSelectFieldQuery  = "action_select_field"
	NameSelectFieldQuery    = "name_select_field"
	ActionSelectValue       = "select"
	ActionSelectRemoteValue = "select-remote"
	LabelStructTag          = "label"
	TypeStructTag           = "type"
)

// Form is an interface that defines a structure for form-based components.
// It provides methods to modify, display and handle form-related actions.
type Form[F any] interface {
	components.UIComponent

	// setUpdateHandler sets a custom handler used to handle the form update request.
	// It receives a components.RequestWrapper and returns an UpdateResponse.
	setUpdateHandler(handler func(c F) *UpdateResponse)

	// addFormFields allows adding multiple fields to the form.
	addFormFields(field ...*models.Field)

	// setSelectHandler sets the select handler for the field with fieldName
	setSelectHandler(fieldName string, handler func(string) []string)

	// updateAction defines the update action for the form.
	// Returns a template component for rendering.
	updateAction(data *UpdateResponse) templ.Component

	// WithSettings allows applying multiple settings to the form.
	// Returns the modified form.
	WithSettings(settings ...func(f Form[F])) Form[F]
}

type formImpl[F any] struct {
	baseWidget    widgets.Widget
	validator     *validator.Validate
	fields        []*models.Field
	popUpResponse bool
	updateHandler func(c F) *UpdateResponse
	description   string
	spec          *models.TreeSpec
	parent        components.UIComponent
	children      []components.UIComponent
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

	id := "formWidget_" + name + "_" + uuid.New().String()
	widget.baseWidget.SetId(id)
	widget.htmxOpts.AppendToPath("widget", id)

	err := widget.generate()
	if err != nil {
		logger.L.Error("error generating form", zap.Error(err))
		return nil
	}

	return widget
}
