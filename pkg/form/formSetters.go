package form

import "github.com/Alfagov/goDashboard/models"

func SetFormInitialValue(value models.UpdateResponse) func(f FormWidget) {
	return func(f FormWidget) {
		f.setInitialValue(value)
	}
}

func SetFormUpdateHandler(handler func(c FormRequest) *models.UpdateResponse) func(
	f FormWidget,
) {
	return func(f FormWidget) {
		f.setUpdateHandler(handler)
	}
}

func SetFormFields(fields ...*models.FormField) func(
	f FormWidget,
) {
	return func(f FormWidget) {
		f.addFormFields(fields...)
	}
}

func SetFormButtons(buttons ...*models.FormButton) func(
	f FormWidget,
) {
	return func(f FormWidget) {
		f.addFormButtons(buttons...)
	}
}

func SetFormCheckboxes(checkboxes ...*models.FormCheckbox) func(
	f FormWidget,
) {
	return func(f FormWidget) {
		f.addFormCheckboxes(checkboxes...)
	}
}
