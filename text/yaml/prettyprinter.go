package yaml

import "gopkg.in/yaml.v2"
import "fmt"

var prettyPrinter = &PrettyPrinter{}

func init() {

}

//PrettyPrinter a yaml implementation of PrettyPrinter
type PrettyPrinter struct {
}

//Print prints in yaml format
func (p *PrettyPrinter) Print(in interface{}) string {
	b, err := yaml.Marshal(in)
	if err != nil {
		//fallback
		return fmt.Sprintf("%v", in)
	}
	return string(b)
}

//PrettyPrint uses a default implementation of PrettyPrinter to pretty print in yaml
func PrettyPrint(in interface{}) string {
	return prettyPrinter.Print(in)
}
