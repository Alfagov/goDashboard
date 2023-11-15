package models

type TreeSpec struct {
	Name        string
	ImageRoute  string
	Description string
	Route       string
	Children    []*TreeSpec
}
