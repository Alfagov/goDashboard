package utils

import (
	"context"
	"github.com/a-h/templ"
	"io"
	"net/http"
)

type TemplRender struct {
	Code int
	Data templ.Component
}

func (t TemplRender) Render(
	w io.Writer, n string, data interface{},
	l ...string,
) error {

	d, ok := data.(templ.Component)
	if !ok {
		return nil
	}

	if d == nil {
		return nil
	}

	return d.Render(context.Background(), w)

}

func (t TemplRender) Load() error {
	return nil
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
