package mapper

import (
	"reflect"
	"testing"

	"github.com/ottogiron/ironman/template/generator/metadata/field"
)

func TestFixedListMapper(t *testing.T) {
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
			"Fixed list",
			args{
				map[string]interface{}{
					"id":    "myFixedList",
					"type":  string(field.TypeFixedList),
					"label": "My fixed list",
					"fields": []interface{}{
						map[string]interface{}{
							"id":    "myTextField1",
							"type":  string(field.TypeText),
							"label": "My text field1",
						},
						map[string]interface{}{
							"id":    "myTextField2",
							"type":  string(field.TypeText),
							"label": "My text field2",
						},
						map[string]interface{}{
							"id":    "myTextField3",
							"type":  string(field.TypeText),
							"label": "My text field3",
						},
					},
				},
			},
			field.NewFixedList(
				field.Field(
					map[string]interface{}{
						"id":    "myFixedList",
						"type":  string(field.TypeFixedList),
						"label": "My fixed list",
						"fields": []interface{}{
							map[string]interface{}{
								"id":    "myTextField1",
								"type":  string(field.TypeText),
								"label": "My text field1",
							},
							map[string]interface{}{
								"id":    "myTextField2",
								"type":  string(field.TypeText),
								"label": "My text field2",
							},
							map[string]interface{}{
								"id":    "myTextField3",
								"type":  string(field.TypeText),
								"label": "My text field3",
							},
						},
					}),
				[]interface{}{
					field.NewText(
						field.Field{
							"id":    "myTextField1",
							"type":  string(field.TypeText),
							"label": "My text field1",
						},
					),
					field.NewText(
						field.Field{
							"id":    "myTextField2",
							"type":  string(field.TypeText),
							"label": "My text field2",
						},
					),
					field.NewText(
						field.Field{
							"id":    "myTextField3",
							"type":  string(field.TypeText),
							"label": "My text field3",
						},
					),
				},
			),
			false,
		},
		{
			"Fixed list field keys mandatory fails",
			args{
				map[string]interface{}{
					"id":    "myFixedList",
					"type":  string(field.TypeFixedList),
					"label": "My fixed list",
				},
			},
			nil,
			true,
		},
		{
			"Fixed list fields should be a []interface{} fails",
			args{
				map[string]interface{}{
					"id":     "myFixedList",
					"type":   string(field.TypeFixedList),
					"label":  "My fixed list",
					"fields": "some invalid fields value",
				},
			},
			nil,
			true,
		},
		{
			"Fixed list fields should field mapping fails",
			args{
				map[string]interface{}{
					"id":    "myFixedList",
					"type":  string(field.TypeFixedList),
					"label": "My fixed list",
					"fields": []interface{}{
						map[string]interface{}{},
					},
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := FixedListMapper{}
			got, err := m.Map(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("%q. FixedListMapper() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q. FixedListMapper() = %v, want %v", tt.name, got, tt.want)
			}
		})

	}
}
