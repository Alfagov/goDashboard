// Code generated by templ@v0.2.364 DO NOT EDIT.

package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Alfagov/goDashboard/models"

func GridPage(widgets []templ.Component) templ.Component {
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
		for _, widget := range widgets {
			err = widget.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Aside(descriptor []*models.TreeSpec) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<aside class=\"col-span-1 row-start-1 row-end-6 rounded bg-white p-4 shadow-lg\"><h2 class=\"mb-4 text-lg font-semibold\">")
		if err != nil {
			return err
		}
		var_3 := `Menu`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><ul>")
		if err != nil {
			return err
		}
		for _, page := range descriptor {
			_, err = templBuffer.WriteString("<li class=\"mb-2\"><a onclick=\"disposeAllCharts()\" hx-target=\"#page-content\" hx-get=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString("/" + page.Route))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\" hx-swap=\"outerHTML\" hx-select=\"#page-content\" class=\"text-gray-700\">")
			if err != nil {
				return err
			}
			var var_4 string = page.Name
			_, err = templBuffer.WriteString(templ.EscapeString(var_4))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></li>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</ul></aside>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func PageContainerView(page templ.Component, descriptor []*models.TreeSpec) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_5 := templ.GetChildren(ctx)
		if var_5 == nil {
			var_5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"drawer\"><input id=\"page-drawer\" type=\"checkbox\" class=\"drawer-toggle\"><div class=\"drawer-content\">")
		if err != nil {
			return err
		}
		err = page.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div class=\"drawer-side\"><label for=\"page-drawer\" aria-label=\"close sidebar\" class=\"drawer-overlay\"></label><ul class=\"menu p-4 w-80 min-h-full bg-base-200 text-base-content\"><p class=\"text-lg font-semibold\">")
		if err != nil {
			return err
		}
		var_6 := `Sub-Pages`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><div class=\"divider\"></div>")
		if err != nil {
			return err
		}
		for _, page := range descriptor {
			_, err = templBuffer.WriteString("<li><button class=\"btn btn-outline text-xl m-2\" onclick=\"disposeAllCharts()\" hx-target=\"#page-content\" hx-get=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString("/" + page.Route))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\" hx-swap=\"outerHTML\" hx-select=\"#page-content\">")
			if err != nil {
				return err
			}
			var var_7 string = page.Name
			_, err = templBuffer.WriteString(templ.EscapeString(var_7))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</button></li>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</ul></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}