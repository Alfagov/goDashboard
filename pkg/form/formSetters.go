package form

import (
	"github.com/Alfagov/goDashboard/models"
)

func SetFormUpdateHandler[F any](handler func(c F) *models.UpdateResponse) func(
	f Form[F],
) {
	return func(f Form[F]) {
		f.setUpdateHandler(handler)
	}
}

func AddFormFields[F any](fields ...*models.Field) func(
	f Form[F],
) {
	return func(f Form[F]) {
		f.addFormFields(fields...)
	}
}

func WithSelectHandler[F any](fieldName string, handler func(string) []string) func(
	f Form[F],
) {
	return func(f Form[F]) {
		f.setSelectHandler(fieldName, handler)
	}
}
