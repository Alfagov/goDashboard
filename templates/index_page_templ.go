// Code generated by templ@v0.2.364 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Alfagov/goDashboard/models"

func SimpleGridPage(pages []models.PageSpec) templ.Component {
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
		_, err = templBuffer.WriteString("<main class=\"grid h-full grid-cols-3 gap-6\"><section class=\"col-span-3 rounded p-4\"><div class=\"grid grid-cols-3 gap-6\">")
		if err != nil {
			return err
		}
		for _, item := range pages {
			err = PageGridComponent(item).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div></section></main>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func PageGridComponent(page models.PageSpec) templ.Component {
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
		_, err = templBuffer.WriteString("<a href=\"")
		if err != nil {
			return err
		}
		var var_3 templ.SafeURL = templ.URL(page.Route)
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_3)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"block focus:outline-none\"><div class=\"bg-white p-4 rounded shadow transform transition-all duration-300 hover:scale-105 hover:bg-gray-200 hover:shadow-lg active:scale-95\"><div class=\"text-xl mb-2\">")
		if err != nil {
			return err
		}
		var var_4 string = page.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_4))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><img src=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(page.ImageRoute))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" alt=\"Placeholder image\" class=\"mb-4 rounded\"><p>")
		if err != nil {
			return err
		}
		var var_5 string = page.Description
		_, err = templBuffer.WriteString(templ.EscapeString(var_5))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div></a>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
