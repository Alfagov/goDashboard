package models

type FormButton struct {
	Label string
	Color string
}

type FormCheckbox struct {
	Label string
	Name  string
}

type FormField struct {
	Label     string
	Name      string
	FieldType string
}

type FormSelectField struct {
	Label   string
	Name    string
	Options []string
}

type UpdateResponse struct {
	Success bool
	Message string
	Title   string
}

type Field struct {
	Name  string
	Label string
	Type  string
}
