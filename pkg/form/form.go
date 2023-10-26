package form

import (
	"github.com/Alfagov/goDashboard/htmx"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

type FormWidget interface {
	setUpdateHandler(handler func(c FormRequest) *models.UpdateResponse)
	setInitialValue(value models.UpdateResponse)
	addFormFields(field ...*models.FormField)
	addFormButtons(button ...*models.FormButton)
	addFormCheckboxes(checkbox ...*models.FormCheckbox)

	HandlePost(c FormRequest) *models.UpdateResponse
	UpdateAction(data *models.UpdateResponse) templ.Component
	WithSettings(settings ...func(f FormWidget)) FormWidget
	GetHtmx() htmx.HTMX
	Encode() templ.Component
	CompileRoutes(router *fiber.App)
	AddParentPath(path string) error
}

type formWidget struct {
	baseWidget    widgets.Widget
	fields        []*models.FormField
	buttons       []*models.FormButton
	checkboxes    []*models.FormCheckbox
	updateHandler func(c FormRequest) *models.UpdateResponse
	initialValue  models.UpdateResponse
	popUpResponse bool
	htmxOpts      htmx.HTMX
}

func (fw *formWidget) CompileRoutes(router *fiber.App) {
	router.Post(
		fw.htmxOpts.GetRoute(), func(c *fiber.Ctx) error {
			update := fw.HandlePost(NewFormRequest(c))

			return c.Render("", fw.UpdateAction(update))
		},
	)
}

func (fw *formWidget) AddParentPath(path string) error {
	return fw.htmxOpts.GetHtmx().AddBeforePath(path)
}

func newForm() *formWidget {
	var w formWidget
	w.baseWidget = widgets.NewWidget()
	w.htmxOpts = htmx.NewEmpty()
	return &w
}

func NewFormWidget(name string, setters ...func(n widgets.Widget)) FormWidget {
	widget := newForm()
	widget.baseWidget.SetName(name)

	for _, setter := range setters {
		setter(widget.baseWidget)
	}

	id := "formWidget_" + name + "_" + ulid.Make().String()
	widget.baseWidget.SetId(id)
	widget.htmxOpts.SetRoute("/update/" + id)

	return widget
}
