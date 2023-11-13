package widgets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWidgetSetters(t *testing.T) {
	type args struct {
		setter func(widget Widget)
	}
	tests := []struct {
		name string
		args args
		want *BaseWidget
	}{
		{
			name: "test set name",
			args: args{
				setter: SetName("test"),
			},
			want: &BaseWidget{
				Name: "test",
			},
		},
		{
			name: "test set description",
			args: args{
				setter: SetDescription("test"),
			},
			want: &BaseWidget{
				Description: "test",
			},
		},
		{
			name: "test set layout",
			args: args{
				setter: SetLayout(nil),
			},
			want: &BaseWidget{
				Layout: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.want, got)

		})
	}
}
