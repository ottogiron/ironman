package field

import (
	"reflect"
	"testing"

	"github.com/ottogiron/ironman/text/yaml"
)

func TestNewArray(t *testing.T) {
	type args struct {
		f               Field
		size            int
		fieldDefinition interface{}
	}
	tests := []struct {
		name string
		args args
		want *Array
	}{
		{
			"New Array",
			args{
				Field{"id": "someid"},
				3,
				"some field definition",
			},
			&Array{
				Field:           Field{"id": "someid"},
				Size:            3,
				FieldDefinition: "some field definition",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArray(tt.args.f, tt.args.size, tt.args.fieldDefinition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q. NewArray() = %v, want %v", tt.name, got, tt.want)
			}
		})

	}
}

func TestArray_String(t *testing.T) {
	type fields struct {
		Field           Field
		size            int
		FieldDefinition interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Create 'pretty' object output",
			fields{
				Field{"id": "some_id"},
				3,
				"some field definiton",
			},
			yaml.PrettyPrint(map[string]interface{}{
				"field":            map[string]string{"id": "some_id"},
				"field_definition": "some field definition",
			}),
		},
	}
	for _, tt := range tests {
		a := &Array{
			Field:           tt.fields.Field,
			Size:            tt.fields.size,
			FieldDefinition: tt.fields.FieldDefinition,
		}
		if got := a.String(); got != tt.want {
			t.Errorf("%q. Array.String() = \n%v want \n%v", tt.name, got, tt.want)
		}
	}
}
