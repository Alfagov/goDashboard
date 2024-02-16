package form

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
)

func SetFormUpdateHandler[F any](handler func(c F) *UpdateResponse) func(
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

func SetTableLink[F any](table components.UIComponent) func(
	f Form[F]) {
	return func(f Form[F]) {
		f.setTableLink(table)
	}
}

func SetTableUpdateHandler[F any](handler func(c F) ([][]interface{}, error)) func(
	f Form[F]) {
	return func(f Form[F]) {
		f.setTableUpdateHandler(handler)
	}
}
