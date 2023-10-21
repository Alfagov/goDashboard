package form

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

func (fw *formWidget) HandlePost(c FormRequest) *models.UpdateResponse {
	return fw.updateHandler(c)
}

func (fw *formWidget) addFormFields(field ...*models.FormField) {
	fw.fields = append(fw.fields, field...)
}

func (fw *formWidget) addFormButtons(button ...*models.FormButton) {
	fw.buttons = append(fw.buttons, button...)
}

func (fw *formWidget) addFormCheckboxes(checkbox ...*models.FormCheckbox) {
	fw.checkboxes = append(fw.checkboxes, checkbox...)
}

func (fw *formWidget) setUpdateHandler(
	handler func(c FormRequest) *models.UpdateResponse,

) {
	fw.updateHandler = handler
}

func (fw *formWidget) setInitialValue(value models.UpdateResponse) {
	fw.initialValue = value
}

func (fw *formWidget) GetHtmx() htmx.HTMX {
	return fw.htmxOpts
}

func (fw *formWidget) UpdateAction(data *models.UpdateResponse) templ.Component {

	if !data.Success {
		element := templates.ErrorAlert(data.Title, data.Message)
		return element
	}

	element := templates.SuccessAlert(data.Title, data.Message)
	return element
}

func (fw *formWidget) WithSettings(
	settings ...func(
		f FormWidget,
	),
) FormWidget {
	for _, setter := range settings {
		setter(fw)
	}

	return fw
}

func (fw *formWidget) Encode() templ.Component {
	fields := fw.fields
	buttons := fw.buttons
	checkboxes := fw.checkboxes

	var fieldsComponent []templ.Component
	for _, field := range fields {
		fieldsComponent = append(fieldsComponent, templates.FormField(field))
	}

	element := templates.GenericForm(
		fw.baseWidget.GetName(),
		fieldsComponent,
		checkboxes,
		buttons,
		fw.baseWidget.GetLayout(),
		fw.htmxOpts.GetHtmx(),
	)

	return element
}
