package widgets

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/a-h/templ"
)

type FormWidget interface {
	Widget

	setUpdateHandler(handler func(c FormRequest) *models.UpdateResponse)
	setInitialValue(value models.UpdateResponse)
	addFormFields(field ...*models.FormField)
	addFormButtons(button ...*models.FormButton)
	addFormCheckboxes(checkbox ...*models.FormCheckbox)

	SetPageRoute(route string)
	GetRoute() string

	HandlePost(c FormRequest) *models.UpdateResponse
	UpdateAction(data *models.UpdateResponse) templ.Component
	WithFormSpecs(settings ...func(f FormWidget)) FormWidget
	Encode() templ.Component
}

type formWidget struct {
	baseWidget    BaseWidget
	fields        []*models.FormField
	buttons       []*models.FormButton
	checkboxes    []*models.FormCheckbox
	updateHandler func(c FormRequest) *models.UpdateResponse
	initialValue  models.UpdateResponse
	popUpResponse bool
	htmx          models.HtmxPoll
}

func NewFormWidget(
	name string,
	setters ...func(
		n Widget,
	),
) FormWidget {
	var widget formWidget
	widget.baseWidget.Name = name
	widget.baseWidget.Layout = &models.WidgetLayout{}

	for _, setter := range setters {
		setter(&widget)
	}

	widget.setId()

	return &widget
}
