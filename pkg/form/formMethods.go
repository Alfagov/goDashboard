package form

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

func (fw *formImpl) addFormFields(field ...*models.FormField) {
	fw.fields = append(fw.fields, field...)
}

func (fw *formImpl) addFormButtons(button ...*models.FormButton) {
	fw.buttons = append(fw.buttons, button...)
}

func (fw *formImpl) addFormCheckboxes(checkbox ...*models.FormCheckbox) {
	fw.checkboxes = append(fw.checkboxes, checkbox...)
}

func (fw *formImpl) setUpdateHandler(
	handler func(c components.RequestWrapper) *models.UpdateResponse,

) {
	fw.updateHandler = handler
}

func (fw *formImpl) setInitialValue(value models.UpdateResponse) {
	fw.initialValue = value
}

func (fw *formImpl) updateAction(data *models.UpdateResponse) templ.Component {

	if !data.Success {
		element := templates.ErrorAlert(data.Title, data.Message)
		return element
	}

	element := templates.SuccessAlert(data.Title, data.Message)
	return element
}

func (fw *formImpl) WithSettings(
	settings ...func(
		f Form,
	),
) Form {
	for _, setter := range settings {
		setter(fw)
	}

	return fw
}
