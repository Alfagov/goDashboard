package models

import "github.com/a-h/templ"

type PagesDescriptor struct {
	Name     string
	Route    string
	Template templ.Component
}

type TreeSpec struct {
	Name        string
	ImageRoute  string
	Description string
	Route       string
	Children    []*TreeSpec
}

type ListElement struct {
	Route       string
	Name        string
	ImageRoute  string
	Description string
	Children    []ListElement
}
