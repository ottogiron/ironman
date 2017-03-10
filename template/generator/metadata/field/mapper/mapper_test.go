package mapper

import (
	"reflect"
	"testing"

	"github.com/ironman-project/ironman/template/generator/metadata/field"
)

func TestNew(t *testing.T) {
	type args struct {
		fieldType field.Type
	}
	tests := []struct {
		name    string
		args    args
		want    Mapper
		wantErr bool
	}{
		{"New Text Mapper", args{field.TypeText}, &TextMapper{}, false},
		{"New Array Mapper", args{field.TypeArray}, &ArrayMapper{}, false},
		{"New Fixed List Mapper", args{field.TypeFixedList}, &FixedListMapper{}, false},
		{"New Mapper error", args{field.Type("")}, nil, true},
	}
	for _, tt := range tests {
		got, err := New(tt.args.fieldType)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. New() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestMapUnstructuredToField(t *testing.T) {
	type args struct {
		unstructuredField interface{}
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
					"type":  string("wutype"),
					"label": "My text field",
				},
			},
			nil,
			true,
		},
		{
			"Invalid unstructured field",
			args{
				nil,
			},
			nil,
			true,
		},
		{
			"Mising mandatory fields",
			args{
				map[string]interface{}{
					"type":  string(field.TypeText),
					"label": "My text field",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapUnstructuredToField(tt.args.unstructuredField)
			if (err != nil) != tt.wantErr {
				t.Errorf("%q. MapUnstructuredToField() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q. MapUnstructuredToField() = \n%v, \nwant \n%v", tt.name, got, tt.want)
			}
		})

	}
}
