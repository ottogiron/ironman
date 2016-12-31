package mapper

import (
	"reflect"
	"testing"

	"github.com/ottogiron/ironman/template/generator/metadata/field"
)

func TestArrayMapper(t *testing.T) {
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
			"Array created",
			args{
				map[string]interface{}{
					"id":    "myFixedArray",
					"type":  string(field.TypeArray),
					"label": "My fixed array",
					"size":  3,
					"field_definition": map[string]interface{}{
						"type":  "text",
						"label": "enter a value",
					},
				}},
			field.NewArray(
				field.Field(map[string]interface{}{
					"id":    "myFixedArray",
					"type":  string(field.TypeArray),
					"label": "My fixed array",
					"size":  3,
					"field_definition": map[string]interface{}{
						"id":    "placeholder",
						"type":  "text",
						"label": "enter a value",
					},
				}),
				3,
				field.NewText(
					field.Field{
						"id":    "placeholder",
						"type":  string(field.TypeText),
						"label": "enter a value",
					},
				),
			),
			false,
		},
		{
			"Array size missing",
			args{
				map[string]interface{}{
					"id":    "myFixedArray",
					"type":  string(field.TypeArray),
					"label": "My fixed array",
					"field_definition": map[string]interface{}{
						"type":  "text",
						"label": "enter a value",
					},
				}},
			nil,
			true,
		},
		{
			"Array field_definition missing",
			args{
				map[string]interface{}{
					"id":    "myFixedArray",
					"type":  string(field.TypeArray),
					"label": "My fixed array",
					"size":  3,
				}},
			nil,
			true,
		},
		{
			"Array size should be int",
			args{
				map[string]interface{}{
					"id":    "myFixedArray",
					"type":  string(field.TypeArray),
					"label": "My fixed array",
					"size":  "3",
					"field_definition": map[string]interface{}{
						"type":  "text",
						"label": "enter a value",
					},
				}},
			nil,
			true,
		},
		{
			"Array field_definition should be a map[string]interface{}",
			args{
				map[string]interface{}{
					"id":               "myFixedArray",
					"type":             string(field.TypeArray),
					"label":            "My fixed array",
					"size":             3,
					"field_definition": []interface{}{},
				}},
			nil,
			true,
		},
		{
			"Array field_definition mapping should fail",
			args{
				map[string]interface{}{
					"id":               "myFixedArray",
					"type":             string(field.TypeArray),
					"label":            "My fixed array",
					"size":             3,
					"field_definition": map[string]interface{}{},
				}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ArrayMapper{}
			got, err := m.Map(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("%q. ArrayMapper() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q. ArrayMapper() = %v, want %v", tt.name, got, tt.want)
			}
		})

	}
}
