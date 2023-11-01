package form

import (
	"github.com/Alfagov/goDashboard/models"
)

func SetFormInitialValue[F any](value models.UpdateResponse) func(f Form[F]) {
	return func(f Form[F]) {
		f.setInitialValue(value)
	}
}

func SetFormUpdateHandler[F any](handler func(c F) *models.UpdateResponse) func(
	f Form[F],
) {
	return func(f Form[F]) {
		f.setUpdateHandler(handler)
	}
}

func SetFormFields[F any](fields ...models.Field) func(
	f Form[F],
) {
	return func(f Form[F]) {
		f.addFormFields(fields...)
	}
}
