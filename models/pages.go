package models

import "github.com/a-h/templ"

type PagesDescriptor struct {
	Name     string
	Route    string
	Template templ.Component
}
