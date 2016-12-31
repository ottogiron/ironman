package mapper

import (
	"reflect"
	"testing"

	"github.com/ottogiron/ironman/template/generator/metadata/field"
)

func TestTextMapper(t *testing.T) {
	type args struct {
		f field.Field
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"Text field",
			args{
				map[string]interface{}{
					"id":    "myTextField",
					"type":  string(field.TypeText),
					"label": "My text field",
				},
			},
			field.NewText(
				field.Field{
					"id":    "myTextField",
					"type":  string(field.TypeText),
					"label": "My text field",
				},
			),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := TextMapper{}
			got, err := m.Map(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("%q. TextMapper() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q. TextMapper() = %v, want %v", tt.name, got, tt.want)
			}
		})

	}
}
