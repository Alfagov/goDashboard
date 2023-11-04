package form

import (
	"github.com/Alfagov/goDashboard/models"
)

var FieldMap = map[string]func(name string, label string) models.Field{
	"button":         ButtonField,
	"checkbox":       CheckboxField,
	"color":          ColorField,
	"date":           DateField,
	"datetime-local": DateTimeLocalField,
	"email":          EmailField,
	"file":           FileField,
	"hidden":         HiddenField,
	"image":          ImageField,
	"month":          MonthField,
	"number":         NumberField,
	"password":       PasswordField,
	"radio":          RadioField,
	"range":          RangeField,
	"reset":          ResetField,
	"search":         SearchField,
	"submit":         SubmitField,
	"tel":            TelField,
	"text":           TextField,
	"time":           TimeField,
	"url":            URLField,
	"week":           WeekField,
	"select":         SelectField,
}

var (
	ButtonField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "button",
		}
	}

	CheckboxField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "checkbox",
		}
	}

	ColorField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "color",
		}
	}

	DateField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "date",
		}
	}

	DateTimeLocalField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "datetime-local",
		}
	}

	EmailField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "email",
		}
	}

	FileField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "file",
		}
	}

	HiddenField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "hidden",
		}
	}

	ImageField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "image",
		}
	}

	MonthField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "month",
		}
	}

	NumberField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "number",
		}
	}

	PasswordField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "password",
		}
	}

	RadioField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "radio",
		}
	}

	RangeField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "range",
		}
	}

	ResetField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "reset",
		}
	}

	SearchField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "search",
		}
	}

	SubmitField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "submit",
		}
	}

	TelField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "tel",
		}
	}

	TextField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "text",
		}
	}

	TimeField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "time",
		}
	}

	URLField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "url",
		}
	}

	WeekField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "week",
		}
	}

	SelectField = func(name string, label string) models.Field {
		return models.Field{
			Name:  name,
			Label: label,
			Type:  "select",
			Route: "",
		}
	}
)
