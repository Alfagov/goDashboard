// Code generated by templ@v0.2.364 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Alfagov/goDashboard/models"

func IndexPage(name string, dashSvg string, navElements []*models.TreeSpec, body templ.Component) templ.Component {
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
		err = header().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<body class=\"bg-gray-100 h-screen font-sans flex flex-col\"><div class=\"container mx-auto p-4 h-full\">")
		if err != nil {
			return err
		}
		err = NavbarWidget(name, dashSvg, navElements).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = body.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div id=\"alert-reference-element\"></div></body>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
