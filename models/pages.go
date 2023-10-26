package models

import "github.com/a-h/templ"

type PagesDescriptor struct {
	Name     string
	Route    string
	Template templ.Component
}

type PageSpec struct {
	Name        string
	ImageRoute  string
	Description string
	Route       string
	Pages       []PageSpec
}

type ListElement struct {
	Route       string
	Name        string
	ImageRoute  string
	Description string
	Children    []ListElement
}
