package numeric

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetNumericInitValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want func(
			f Numeric,
		)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SetNumericInitValue(tt.args.value), "SetNumericInitValue(%v)", tt.args.value)
		})
	}
}

func TestSetNumericUnit(t *testing.T) {
	type args struct {
		unit string
	}
	tests := []struct {
		name string
		args args
		want func(
			f Numeric,
		)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SetNumericUnit(tt.args.unit), "SetNumericUnit(%v)", tt.args.unit)
		})
	}
}

func TestSetNumericUnitAfter(t *testing.T) {
	tests := []struct {
		name string
		want func(
			f Numeric,
		)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SetNumericUnitAfter(), "SetNumericUnitAfter()")
		})
	}
}

func TestSetNumericUpdateHandler(t *testing.T) {
	type args struct {
		handler func() (int, error)
	}
	tests := []struct {
		name string
		args args
		want func(
			f Numeric,
		)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SetNumericUpdateHandler(tt.args.handler), "SetNumericUpdateHandler(%v)", tt.args.handler)
		})
	}
}
