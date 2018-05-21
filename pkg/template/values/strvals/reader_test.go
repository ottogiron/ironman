package strvals

import (
	"path/filepath"
	"testing"

	"github.com/ironman-project/ironman/pkg/template/values"
	"github.com/ironman-project/ironman/pkg/testutils"
)

func Test_reader_Read(t *testing.T) {
	tests := []struct {
		name       string
		valueFiles []string
		vals       []string
		want       values.Values
		wantErr    bool
	}{
		{"Parse value", []string{}, []string{"name1=1"}, values.Values{"name1": 1}, false},
		{"Parse values", []string{}, []string{"name1=1,name2=2"}, values.Values{"name1": 1, "name2": 2}, false},
		{"Parse invalid values", []string{}, []string{"name1=value1,,,,name2=value2,"}, nil, true},
		{
			"Parse inner values",
			[]string{},
			[]string{"outer.inner1=value,outer.middle.inner=value,"},
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
			[]string{},
			[]string{"name1={value1,value2},name2={value1,value2}"},
			values.Values{"name1": []string{"value1", "value2"}, "name2": []string{"value1", "value2"}},
			false,
		},
		{
			"Parse list",
			[]string{},
			[]string{"list[0].foo=bar,list[0].hello=world"},
			values.Values{
				"list": []interface{}{
					map[string]interface{}{"foo": "bar", "hello": "world"},
				},
			},
			false,
		},
		{
			"Parse value files with override",
			[]string{filepath.Join("testing", "values", "values.yaml"), filepath.Join("testing", "values", "values.override.yaml")},
			[]string{"foo=foo_inline"},
			values.Values{"foo": "foo_inline", "bar": "bar_override"},
			false,
		},
		{
			"Parse value files ",
			[]string{filepath.Join("testing", "values", "values.yaml")},
			[]string{},
			values.Values{"foo": "foo", "bar": "bar"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &reader{valueFiles: tt.valueFiles, values: tt.vals}
			got, err := r.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("reader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			j1 := testutils.Marshal(got, t)
			j2 := testutils.Marshal(tt.want, t)
			if j1 != j2 {
				t.Errorf("reader.Read() = %v, want %v", j1, j2)
			}
		})
	}
}
