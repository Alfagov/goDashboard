package form

import (
	"errors"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

// Form interface implementation

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

// UIComponent interface implementation

func (fw *formImpl) Render(req components.RequestWrapper) *components.RenderResponse {

	if req != nil && req.Method() == "POST" {
		data := fw.updateHandler(req)
		return &components.RenderResponse{
			Component: fw.updateAction(data),
		}
	}

	fields := fw.fields
	buttons := fw.buttons
	checkboxes := fw.checkboxes

	var fieldsComponent []templ.Component
	for _, field := range fields {
		fieldsComponent = append(fieldsComponent, templates.FormField(field))
	}

	return &components.RenderResponse{
		Component: templates.GenericForm(
			fw.Name(),
			fieldsComponent,
			checkboxes,
			buttons,
			fw.baseWidget.GetLayout(),
			fw.htmxOpts.GetHtmx(),
		),
	}
}

func (fw *formImpl) Type() components.NodeType {
	return components.FormWidgetType
}

func (fw *formImpl) Name() string {
	return fw.baseWidget.GetName()
}

func (fw *formImpl) UpdateSpec() *models.TreeSpec {
	route := components.GetRouteFromParents(fw)

	fw.htmxOpts.AddBeforePath(route)
	return &models.TreeSpec{
		Name:        fw.Name(),
		ImageRoute:  "",
		Description: fw.description,
		Route:       fw.htmxOpts.GetUrl(),
		Children:    nil,
	}
}

func (fw *formImpl) GetSpec() *models.TreeSpec {
	return fw.spec
}

func (fw *formImpl) GetChildren() []components.UIComponent {
	return nil
}

func (fw *formImpl) FindChild(string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl) Id() string {
	return fw.baseWidget.GetId()
}

func (fw *formImpl) FindChildById(string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl) FindChildByType(string, string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl) SetParent(parent components.UIComponent) {
	fw.parent = parent
}

func (fw *formImpl) GetParent() components.UIComponent {
	return fw.parent
}

func (fw *formImpl) AddChild(components.UIComponent) error {
	return errors.New("not applicable")
}

func (fw *formImpl) KillChild(components.UIComponent) error {
	return errors.New("not applicable")
}
