package yaml

import "gopkg.in/yaml.v2"
import "fmt"

//Print prints in yaml format
func Print(in interface{}) string {
	b, err := yaml.Marshal(in)
	if err != nil {
		//fallback
		return fmt.Sprintf("%v", in)
	}
	return string(b)
}
