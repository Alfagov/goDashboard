package form

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddFormFields(t *testing.T) {
	type args struct {
		fields []*models.Field
	}
	type testCase[F any] struct {
		name string
		args args
		want func(
			f Form[F],
		)
	}
	tests := []testCase[testFieldForm]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AddFormFields[testFieldForm](tt.args.fields...))
		})
	}
}

func TestSetFormUpdateHandler(t *testing.T) {
	type args[F any] struct {
		handler func(c F) *UpdateResponse
	}
	type testCase[F any] struct {
		name string
		args args[F]
		want func(
			f Form[F],
		)
	}
	tests := []testCase[testFieldForm]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SetFormUpdateHandler(tt.args.handler), "SetFormUpdateHandler(%v)", tt.args.handler)
		})
	}
}

func TestWithSelectHandler(t *testing.T) {
	type args struct {
		fieldName string
		handler   func(string) []string
	}
	type testCase[F any] struct {
		name string
		args args
		want func(
			f Form[F],
		)
	}
	tests := []testCase[testFieldForm]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, WithSelectHandler[testFieldForm](tt.args.fieldName, tt.args.handler),
				"WithSelectHandler(%v, "+
					"%v)", tt.args.fieldName, tt.args.handler)
		})
	}
}
