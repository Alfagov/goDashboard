package form

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func (fw *formImpl) handlePost(c FormRequest) *models.UpdateResponse {
	return fw.updateHandler(c)
}

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
	handler func(c FormRequest) *models.UpdateResponse,

) {
	fw.updateHandler = handler
}

func (fw *formImpl) setInitialValue(value models.UpdateResponse) {
	fw.initialValue = value
}

func (fw *formImpl) getHtmx() htmx.HTMX {
	return fw.htmxOpts
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

func (fw *formImpl) Encode() templ.Component {
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

func (fw *formImpl) CompileRoutes(router *fiber.App) {
	router.Post(
		fw.htmxOpts.GetUrl(), func(c *fiber.Ctx) error {
			update := fw.handlePost(NewFormRequest(c))

			return c.Render("", fw.updateAction(update))
		},
	)
}

func (fw *formImpl) AddParentPath(path string) error {
	return fw.htmxOpts.GetHtmx().AddBeforePath(path)
}
