package widgets

import (
	"dario.cat/mergo"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
)

func (fw *formWidget) HandlePost(c FormRequest) *models.UpdateResponse {
	return fw.updateHandler(c)
}

func (fw *formWidget) SetPageRoute(route string) {
	r := route + "/update/" + fw.baseWidget.Id
	fw.htmx.Route = r
}

func (fw *formWidget) GetRoute() string {
	return fw.htmx.Route
}

func (fw *formWidget) withLayout(layout *models.WidgetLayout) {

	err := mergo.Merge(fw.baseWidget.Layout, layout, mergo.WithOverride)
	if err != nil {
		panic(err)
	}
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

func (fw *formWidget) setName(name string) {
	fw.baseWidget.Name = name
}

func (fw *formWidget) setHeight(height int) {
	fw.baseWidget.Layout.Height = height
}

func (fw *formWidget) setWidth(width int) {
	fw.baseWidget.Layout.Width = width
}

func (fw *formWidget) setRow(row int) {
	fw.baseWidget.Layout.Row = row
}

func (fw *formWidget) setColumn(column int) {
	fw.baseWidget.Layout.Column = column
}

func (fw *formWidget) GetRow() int {
	return fw.baseWidget.Layout.Row
}

func (fw *formWidget) setId() {
	uld := ulid.Make()
	id := "formWidget_" + fw.baseWidget.Name + "_" + uld.String()
	fw.baseWidget.Id = id
}

func (fw *formWidget) setDescription(description string) {
	fw.baseWidget.Description = description
}

func (fw *formWidget) setInitialValue(value models.UpdateResponse) {
	fw.initialValue = value
}

func (fw *formWidget) UpdateAction(data *models.UpdateResponse) templ.Component {

	if !data.Success {
		element := templates.ErrorAlert(data.Title, data.Message)
		return element
	}

	element := templates.SuccessAlert(data.Title, data.Message)
	return element
}

func (fw *formWidget) WithFormSpecs(
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
		fw.baseWidget.Name,
		fieldsComponent,
		checkboxes,
		buttons,
		fw.baseWidget.Layout,
		&fw.htmx,
	)

	return element
}
