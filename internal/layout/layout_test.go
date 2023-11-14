package layout

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSettersFunc(t *testing.T) {
	type args struct {
		setter func(*WidgetLayout)
	}
	tests := []struct {
		name string
		args args
		want *WidgetLayout
	}{
		{
			name: "test set row",
			args: args{
				setter: SetRow(3),
			},
			want: &WidgetLayout{
				Row: 3,
			},
		},
		{
			name: "test set column",
			args: args{
				setter: SetColumn(3),
			},
			want: &WidgetLayout{
				Column: 3,
			},
		},
		{
			name: "test set width",
			args: args{
				setter: SetWidth(3),
			},
			want: &WidgetLayout{
				Width: 3,
			},
		},
		{
			name: "test set height",
			args: args{
				setter: SetHeight(3),
			},
			want: &WidgetLayout{
				Height: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewWidgetLayout(tt.args.setter)

			assert.Equal(t, tt.want, got)
		})
	}
}
