package form

import (
	"errors"
	"github.com/Alfagov/goDashboard/internal/logger"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
	"github.com/Alfagov/goDashboard/pkg/views"
	"github.com/a-h/templ"
	"go.uber.org/zap"
	"net/http"
	"reflect"
)

// Form interface implementation

func (fw *formImpl[F]) addFormFields(field ...*models.Field) {
	fw.fields = append(fw.fields, field...)
}

func (fw *formImpl[F]) setTableLink(table components.UIComponent) {
	fw.tableLink = table
}

func (fw *formImpl[F]) setTableUpdateHandler(handler func(c F) ([][]interface{}, error)) {
	fw.tableUpdateHandler = handler
}

func (fw *formImpl[F]) setUpdateHandler(
	handler func(c F) *UpdateResponse,

) {
	fw.updateHandler = handler
}

func (fw *formImpl[F]) updateAction(data *UpdateResponse) templ.Component {

	if !data.Success {
		element := views.ErrorAlert(data.Title, data.Message)
		return element
	}

	element := views.SuccessAlert(data.Title, data.Message)
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

func (fw *formImpl[F]) process(req models.RequestWrapper) (*F, error) {

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

func (fw *formImpl[F]) Render(req models.RequestWrapper) *components.RenderResponse {

	if req != nil {
		if req.Method() == http.MethodPost {
			inputData, err := fw.process(req)
			if err != nil {
				return components.NewRenderResponse(nil, nil, err)
			}

			if fw.tableLink != nil && fw.tableUpdateHandler != nil {
				tableData, err := fw.tableUpdateHandler(*inputData)
				if err != nil {
					return components.NewRenderResponse(nil, nil, err)
				}

				req.AddAdditionalData(tableData)

				req.AddHeaders("HX-Retarget", "#"+fw.tableLink.Id())
				req.AddHeaders("HX-Reswap", "outerHTML")

				return components.NewRenderResponse(fw.tableLink.Render(req).Component, nil, nil)
			}

			data := fw.updateHandler(*inputData)

			return components.NewRenderResponse(fw.updateAction(data), nil, nil)

		}

		if req.Query(ActionSelectFieldQuery) == ActionSelectValue {
			for _, field := range fw.fields {
				if field.Name == req.Query(NameSelectFieldQuery) {
					return components.NewRenderResponse(
						SelectOptions(field.SelectHandler(req.Query(field.Label, "")), field.Name+"options"),
						nil, nil)
				}
			}
		}

		if req.Query(ActionSelectFieldQuery) == ActionSelectRemoteValue {
			for _, field := range fw.fields {
				if field.Name == req.Query(NameSelectFieldQuery) {
					return components.NewRenderResponse(
						SelectOptions(field.SelectHandler(req.Query(field.Label, "")), field.Name+"options"),
						nil, nil)
				}
			}
		}
	}

	var fieldsComponent []templ.Component
	for _, field := range fw.fields {
		fieldsComponent = append(fieldsComponent, FormField(*field, fw.spec.Route))
	}

	return components.NewRenderResponse(
		GenericForm(
			fw.Name(),
			fieldsComponent,
			fw.baseWidget.GetLayout(),
			fw.htmxOpts.GetHtmx(),
		), nil, nil)
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

func (fw *formImpl[F]) RemoveChild(components.UIComponent) error {
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

		fields = append(fields, field)
	}

	return fields
}
