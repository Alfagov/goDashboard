package widgets

import (
	"dario.cat/mergo"
	"errors"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"github.com/oklog/ulid/v2"
)

type FormWidget interface {
	Widget
	setUpdateHandler(handler func() *models.UpdateResponse)
	setInitialValue(value models.UpdateResponse)
	UpdateAction(data *models.UpdateResponse) templ.Component
	Encode() templ.Component
	WithFormSpecs(settings ...func(f FormWidget)) FormWidget
	addFormFields(field ...*models.FormField)
	addFormButtons(button ...*models.FormButton)
	addFormCheckboxes(checkbox ...*models.FormCheckbox)
	SetPageRoute(route string)
}

type formWidget struct {
	baseWidget    BaseWidget
	fields        []*models.FormField
	buttons       []*models.FormButton
	checkboxes    []*models.FormCheckbox
	updateHandler func() *models.UpdateResponse
	initialValue  models.UpdateResponse
	popUpResponse bool
	htmx          models.HtmxPoll
}

func (n *formWidget) SetPageRoute(route string) {
	r := "/update/" + route + "/" + n.baseWidget.Id
	n.htmx.Route = r
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

var ErrWidgetNameNotSet = errors.New("widget name not set")

func NewFormWidget(
	setters ...func(
		n Widget,
	),
) FormWidget {
	var widget formWidget
	widget.baseWidget.Layout = &models.WidgetLayout{}

	for _, setter := range setters {
		setter(&widget)
	}

	widget.setId()

	return &widget
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

func (fw *formWidget) setUpdateHandler(
	handler func() *models.UpdateResponse,

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
	)

	return element
}

func SetName(name string) func(f Widget) {
	return func(f Widget) {
		f.setName(name)
	}
}

func SetDescription(description string) func(f Widget) {
	return func(f Widget) {
		f.setDescription(description)
	}
}

func SetFormInitialValue(value models.UpdateResponse) func(f FormWidget) {
	return func(f FormWidget) {
		f.setInitialValue(value)
	}
}

func SetHeight(height int) func(Widget) {
	return func(f Widget) {
		f.setHeight(height)
	}
}

func SetWidth(width int) func(Widget) {
	return func(f Widget) {
		f.setWidth(width)
	}
}

func SetLayout(layout *models.WidgetLayout) func(Widget) {
	return func(f Widget) {
		f.withLayout(layout)
	}
}

func SetRow(row int) func(Widget) {
	return func(f Widget) {
		f.setRow(row)
	}
}

func SetColumn(column int) func(Widget) {
	return func(f Widget) {
		f.setColumn(column)
	}
}

func SetFormUpdateHandler(handler func() *models.UpdateResponse) func(
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
