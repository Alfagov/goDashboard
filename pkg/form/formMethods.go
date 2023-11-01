package form

import (
	"errors"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
)

// Form interface implementation

func (fw *formImpl[F]) addFormFields(field ...models.Field) {
	fw.fields = append(fw.fields, field...)
}

func (fw *formImpl[F]) setUpdateHandler(
	handler func(c F) *models.UpdateResponse,

) {
	fw.updateHandler = handler
}

func (fw *formImpl[F]) setInitialValue(value models.UpdateResponse) {
	fw.initialValue = value
}

func (fw *formImpl[F]) updateAction(data *models.UpdateResponse) templ.Component {

	if !data.Success {
		element := templates.ErrorAlert(data.Title, data.Message)
		return element
	}

	element := templates.SuccessAlert(data.Title, data.Message)
	return element
}

func (fw *formImpl[F]) WithSettings(
	settings ...func(
		f Form[F],
	),
) Form[F] {
	for _, setter := range settings {
		setter(fw)
	}

	return fw
}

func (fw *formImpl[F]) process(req components.RequestWrapper) (*F, error) {

	var data F
	if req != nil {
		err := req.BindFormRequest(&data)
		if err != nil {
			return nil, err
		}
	}

	return &data, nil
}

// UIComponent interface implementation

func (fw *formImpl[F]) Render(req components.RequestWrapper) *components.RenderResponse {

	if req != nil && req.Method() == "POST" {

		inputData, err := fw.process(req)
		if err != nil {
			return &components.RenderResponse{
				Err: err,
			}
		}

		data := fw.updateHandler(*inputData)
		return &components.RenderResponse{
			Component: fw.updateAction(data),
		}
	}

	var fieldsComponent []templ.Component
	for _, field := range fw.fields {
		fieldsComponent = append(fieldsComponent, templates.FormField(field))
	}

	return &components.RenderResponse{
		Component: templates.GenericForm(
			fw.Name(),
			fieldsComponent,
			fw.baseWidget.GetLayout(),
			fw.htmxOpts.GetHtmx(),
		),
	}
}

func (fw *formImpl[F]) Type() components.NodeType {
	return components.FormWidgetType
}

func (fw *formImpl[F]) Name() string {
	return fw.baseWidget.GetName()
}

func (fw *formImpl[F]) UpdateSpec() *models.TreeSpec {
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

func (fw *formImpl[F]) GetSpec() *models.TreeSpec {
	return fw.spec
}

func (fw *formImpl[F]) GetChildren() []components.UIComponent {
	return nil
}

func (fw *formImpl[F]) FindChild(string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl[F]) Id() string {
	return fw.baseWidget.GetId()
}

func (fw *formImpl[F]) FindChildById(string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl[F]) FindChildByType(string, string) (components.UIComponent, bool) {
	return nil, false
}

func (fw *formImpl[F]) SetParent(parent components.UIComponent) {
	fw.parent = parent
}

func (fw *formImpl[F]) GetParent() components.UIComponent {
	return fw.parent
}

func (fw *formImpl[F]) AddChild(components.UIComponent) error {
	return errors.New("not applicable")
}

func (fw *formImpl[F]) KillChild(components.UIComponent) error {
	return errors.New("not applicable")
}
