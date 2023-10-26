package form

import "github.com/Alfagov/goDashboard/models"

func SetFormInitialValue(value models.UpdateResponse) func(f Form) {
	return func(f Form) {
		f.setInitialValue(value)
	}
}

func SetFormUpdateHandler(handler func(c FormRequest) *models.UpdateResponse) func(
	f Form,
) {
	return func(f Form) {
		f.setUpdateHandler(handler)
	}
}

func SetFormFields(fields ...*models.FormField) func(
	f Form,
) {
	return func(f Form) {
		f.addFormFields(fields...)
	}
}

func SetFormButtons(buttons ...*models.FormButton) func(
	f Form,
) {
	return func(f Form) {
		f.addFormButtons(buttons...)
	}
}

func SetFormCheckboxes(checkboxes ...*models.FormCheckbox) func(
	f Form,
) {
	return func(f Form) {
		f.addFormCheckboxes(checkboxes...)
	}
}
