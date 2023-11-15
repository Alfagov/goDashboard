package form

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type testFieldForm struct {
	Name string `label:"Name" type:"text"`
}

type testAllFieldsForm struct {
	ButtonField        string `label:"ButtonField" type:"button"`
	CheckboxField      string `label:"CheckboxField" type:"checkbox"`
	ColorField         string `label:"ColorField" type:"color"`
	DateField          string `label:"DateField" type:"date"`
	DateTimeLocalField string `label:"DateTimeLocalField" type:"datetime-local"`
	EmailField         string `label:"EmailField" type:"email"`
	FileField          string `label:"FileField" type:"file"`
	HiddenField        string `label:"HiddenField" type:"hidden"`
	ImageField         string `label:"ImageField" type:"image"`
	MonthField         string `label:"MonthField" type:"month"`
	NumberField        string `label:"NumberField" type:"number"`
	PasswordField      string `label:"PasswordField" type:"password"`
	RadioField         string `label:"RadioField" type:"radio"`
	RangeField         string `label:"RangeField" type:"range"`
	ResetField         string `label:"ResetField" type:"reset"`
	SearchField        string `label:"SearchField" type:"search"`
	SubmitField        string `label:"SubmitField" type:"submit"`
	TelField           string `label:"TelField" type:"tel"`
	TextField          string `label:"TextField" type:"text"`
	TimeField          string `label:"TimeField" type:"time"`
	URLField           string `label:"URLField" type:"url"`
	WeekField          string `label:"WeekField" type:"week"`
	SelectField        string `label:"SelectField" type:"select"`
	SelectRemoteField  string `label:"SelectRemoteField" type:"select-remote"`
}

func Test_formImpl_generate(t *testing.T) {
	type testCase[F any] struct {
		name    string
		fw      formImpl[F]
		want    []*models.Field
		wantErr bool
	}
	tests := []testCase[testFieldForm]{
		{
			name: "",
			fw:   formImpl[testFieldForm]{},
			want: []*models.Field{
				FieldMap["text"]("Name", "Name"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fw.generate()

			assert.NoError(t, err)

			assert.Equal(t, tt.want, tt.fw.fields)

		})
	}
}

func Test_toFieldArray(t *testing.T) {
	expectedFields := []*models.Field{
		FieldMap["button"]("ButtonField", "ButtonField"),
		FieldMap["checkbox"]("CheckboxField", "CheckboxField"),
		FieldMap["color"]("ColorField", "ColorField"),
		FieldMap["date"]("DateField", "DateField"),
		FieldMap["datetime-local"]("DateTimeLocalField", "DateTimeLocalField"),
		FieldMap["email"]("EmailField", "EmailField"),
		FieldMap["file"]("FileField", "FileField"),
		FieldMap["hidden"]("HiddenField", "HiddenField"),
		FieldMap["image"]("ImageField", "ImageField"),
		FieldMap["month"]("MonthField", "MonthField"),
		FieldMap["number"]("NumberField", "NumberField"),
		FieldMap["password"]("PasswordField", "PasswordField"),
		FieldMap["radio"]("RadioField", "RadioField"),
		FieldMap["range"]("RangeField", "RangeField"),
		FieldMap["reset"]("ResetField", "ResetField"),
		FieldMap["search"]("SearchField", "SearchField"),
		FieldMap["submit"]("SubmitField", "SubmitField"),
		FieldMap["tel"]("TelField", "TelField"),
		FieldMap["text"]("TextField", "TextField"),
		FieldMap["time"]("TimeField", "TimeField"),
		FieldMap["url"]("URLField", "URLField"),
		FieldMap["week"]("WeekField", "WeekField"),
		FieldMap["select"]("SelectField", "SelectField"),
		FieldMap["select-remote"]("SelectRemoteField", "SelectRemoteField"),
	}

	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		args args
		want []*models.Field
	}{
		{
			name: "returns correct field array",
			args: args{
				t: reflect.TypeOf(testAllFieldsForm{}),
			},
			want: expectedFields,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fields := toFieldArray(tt.args.t)

			assert.Equal(t, tt.want, fields)
			assert.Len(t, fields, len(tt.want))
		})
	}
}
