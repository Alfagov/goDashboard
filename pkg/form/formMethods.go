package form

import (
	"errors"
	"fmt"
	"github.com/Alfagov/goDashboard/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/templates"
	"github.com/a-h/templ"
	"go.uber.org/zap"
	"net/http"
	"reflect"
)

// Form interface implementation

func (fw *formImpl[F]) addFormFields(field ...*models.Field) {
	fw.fields = append(fw.fields, field...)
}

func (fw *formImpl[F]) setUpdateHandler(
	handler func(c F) *models.UpdateResponse,

) {
	fw.updateHandler = handler
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

func (fw *formImpl[F]) validate(data F) error {
	return fw.validator.Struct(data)
}

func (fw *formImpl[F]) generate() error {
	var data F
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return errors.New("invalid type")
	}

	t := v.Type()

	fw.fields = toFieldArray(t)

	return nil
}

func (fw *formImpl[F]) setSelectHandler(fieldName string, handler func(string) []string) {
	for _, field := range fw.fields {
		if field.Name == fieldName {
			field.SelectHandler = handler
		}
	}

	logger.L.Error("field not found")
}

// UIComponent interface implementation

func (fw *formImpl[F]) Render(req components.RequestWrapper) *components.RenderResponse {

	if req != nil {
		if req.Method() == http.MethodPost {
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

		if req.Query(ActionSelectFieldQuery) == "select" {
			for _, field := range fw.fields {
				if field.Name == req.Query(NameSelectFieldQuery) {
					return &components.RenderResponse{
						Component: templates.SelectOptions(field.SelectHandler(req.Query(field.Label, ""))),
					}
				}
			}
		}
	}

	var fieldsComponent []templ.Component
	for _, field := range fw.fields {
		fieldsComponent = append(fieldsComponent, templates.FormField(*field, fw.spec.Route))
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

	err := fw.htmxOpts.AddBeforePath(route)
	if err != nil {
		logger.L.Error("error in updating spec", zap.Error(err))
	}

	spec := &models.TreeSpec{
		Name:        fw.Name(),
		ImageRoute:  "",
		Description: fw.description,
		Route:       fw.htmxOpts.GetUrl(),
		Children:    nil,
	}

	fw.spec = spec

	return spec
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

// Utils

func toFieldArray(t reflect.Type) []*models.Field {
	var fields []*models.Field
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag

		name := t.Field(i).Name
		label := tag.Get(LabelStructTag)
		tp := tag.Get(TypeStructTag)

		field := FieldMap[tp](name, label)
		if tp == "select" {
			field.Route = fmt.Sprintf("?%s=select&%s=%s", ActionSelectFieldQuery, NameSelectFieldQuery, name)
		}

		fields = append(fields, &field)
	}

	return fields
}
