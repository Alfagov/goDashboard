// Code generated by templ@(devel) DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Alfagov/goDashboard/models"
import "github.com/Alfagov/goDashboard/htmx"
import "github.com/Alfagov/goDashboard/layout"

func NumericWidget(title string, value string, unit string, unitAfter bool, htmx *htmx.Htmx, widgetLayout *models.WidgetLayout) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var var_2 = []any{layout.ToCSS(widgetLayout) + "bg-white p-4 rounded-lg shadow-lg"}
		err = templ.RenderCSSItems(ctx, templBuffer, var_2...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_2).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-get=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(htmx.Route))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-trigger=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString("every " + htmx.Interval))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-target=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(htmx.Target))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-swap=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(htmx.Swap))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><p class=\"text-gray-600\">")
		if err != nil {
			return err
		}
		var var_3 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		if unitAfter {
			_, err = templBuffer.WriteString("<h2 class=\"text-3xl font-bold\">")
			if err != nil {
				return err
			}
			var var_4 string = value
			_, err = templBuffer.WriteString(templ.EscapeString(var_4))
			if err != nil {
				return err
			}
			var var_5 string = unit
			_, err = templBuffer.WriteString(templ.EscapeString(var_5))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</h2>")
			if err != nil {
				return err
			}
		} else {
			_, err = templBuffer.WriteString("<h2 class=\"text-3xl font-bold\">")
			if err != nil {
				return err
			}
			var var_6 string = unit
			_, err = templBuffer.WriteString(templ.EscapeString(var_6))
			if err != nil {
				return err
			}
			var var_7 string = value
			_, err = templBuffer.WriteString(templ.EscapeString(var_7))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</h2>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func GenericForm(formTitle string, fields []templ.Component, checkbox []*models.FormCheckbox, button []*models.FormButton, widgetLayout *models.WidgetLayout, htmx *htmx.Htmx) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_8 := templ.GetChildren(ctx)
		if var_8 == nil {
			var_8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var var_9 = []any{layout.ToCSS(widgetLayout) + "bg-white p-4 rounded-lg shadow-lg"}
		err = templ.RenderCSSItems(ctx, templBuffer, var_9...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_9).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><h2 class=\"text-2xl font-bold mb-4\">")
		if err != nil {
			return err
		}
		var var_10 string = formTitle
		_, err = templBuffer.WriteString(templ.EscapeString(var_10))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><div id=\"form-error\"></div><form hx-post=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(htmx.Route))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-swap=\"afterbegin\" hx-target=\"#test-id\">")
		if err != nil {
			return err
		}
		for _, field := range fields {
			err = field.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		for _, chk := range checkbox {
			err = FormCheckbox(chk).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		for _, btn := range button {
			err = FormButton(btn).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</form></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FormField(data *models.FormField) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_11 := templ.GetChildren(ctx)
		if var_11 == nil {
			var_11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"mb-4\"><label for=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(data.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"block text-sm font-medium text-gray-600\">")
		if err != nil {
			return err
		}
		var var_12 string = data.Label
		_, err = templBuffer.WriteString(templ.EscapeString(var_12))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(data.FieldType))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(data.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" name=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(data.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"mt-1 p-2 w-full border rounded-md\"></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func SelectFormField(label string, name string, options []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_13 := templ.GetChildren(ctx)
		if var_13 == nil {
			var_13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"mb-4\"><label for=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"block text-sm font-medium text-gray-600\">")
		if err != nil {
			return err
		}
		var var_14 string = label
		_, err = templBuffer.WriteString(templ.EscapeString(var_14))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><select id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" name=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"mt-1 p-2 w-full border rounded-md\">")
		if err != nil {
			return err
		}
		for _, option := range options {
			_, err = templBuffer.WriteString("<option>")
			if err != nil {
				return err
			}
			var var_15 string = option
			_, err = templBuffer.WriteString(templ.EscapeString(var_15))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</option>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</select></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FormButton(data *models.FormButton) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_16 := templ.GetChildren(ctx)
		if var_16 == nil {
			var_16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"mb-4\">")
		if err != nil {
			return err
		}
		var var_17 = []any{"bg-" + data.Color + "-500 text-white p-2 rounded-md"}
		err = templ.RenderCSSItems(ctx, templBuffer, var_17...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<button type=\"submit\" class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_17).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var var_18 string = data.Label
		_, err = templBuffer.WriteString(templ.EscapeString(var_18))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FormCheckbox(data *models.FormCheckbox) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_19 := templ.GetChildren(ctx)
		if var_19 == nil {
			var_19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"mb-4\"><label for=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(data.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"block text-sm font-medium text-gray-600\">")
		if err != nil {
			return err
		}
		var var_20 string = data.Label
		_, err = templBuffer.WriteString(templ.EscapeString(var_20))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"checkbox\" id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(data.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" name=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(data.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"mt-1\"></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
