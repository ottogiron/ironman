package strvals

import (
	"encoding/json"
	"testing"

	"github.com/ironman-project/ironman/template/values"
)

func Test_reader_Read(t *testing.T) {
	tests := []struct {
		name    string
		vals    string
		want    values.Values
		wantErr bool
	}{
		{"Parse value", "name1=1", values.Values{"name1": 1}, false},
		{"Parse values", "name1=1,name2=2", values.Values{"name1": 1, "name2": 2}, false},
		{"Parse invalid values", "name1=value1,,,,name2=value2,", nil, true},
		{
			"Parse inner values",
			"outer.inner1=value,outer.middle.inner=value,",
			values.Values{
				"outer": map[string]interface{}{
					"inner1": "value",
					"middle": map[string]interface{}{
						"inner": "value",
					},
				},
			},
			false,
		},
		{
			"Parse map with list value",
			"name1={value1,value2},name2={value1,value2}",
			values.Values{"name1": []string{"value1", "value2"}, "name2": []string{"value1", "value2"}},
			false,
		},
		{
			"Parse list",
			"list[0].foo=bar,list[0].hello=world",
			values.Values{
				"list": []interface{}{
					map[string]interface{}{"foo": "bar", "hello": "world"},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &reader{tt.vals}
			got, err := r.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("reader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			j1 := marshall(got, t)
			j2 := marshall(tt.want, t)
			if j1 != j2 {
				t.Errorf("reader.Read() = %v, want %v", j1, j2)
			}
		})
	}
}

func marshall(val interface{}, t *testing.T) string {
	bytes, err := json.Marshal(val)

	if err != nil {
		t.Fatalf("Failed to marshal object %v", val)
	}
	return string(bytes)
}
