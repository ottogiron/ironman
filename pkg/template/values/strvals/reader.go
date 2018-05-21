package strvals

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/ironman-project/ironman/pkg/template/values"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
	"k8s.io/helm/pkg/strvals"
)

var _ values.Reader = (*reader)(nil)

type reader struct {
	valueFiles []string
	values     []string
}

//New returns a new instance of a flags values reader
//flags in the form of key=value, key=value1
//
func New(valueFiles []string, values []string) values.Reader {
	return &reader{
		valueFiles: valueFiles,
		values:     values,
	}
}

func (r *reader) Read() (values.Values, error) {
	return vals(r.valueFiles, r.values)
}

// vals merges values from files specified via -f/--values and
// directly via --set, marshaling them to YAML
func vals(valueFiles []string, vals []string) (values.Values, error) {
	base := map[string]interface{}{}

	// User specified a values files via -f/--values
	for _, filePath := range valueFiles {
		currentMap := map[string]interface{}{}

		var bytes []byte
		var err error
		if strings.TrimSpace(filePath) == "-" {
			bytes, err = ioutil.ReadAll(os.Stdin)
		} else {
			bytes, err = readFile(filePath)
		}

		if err != nil {
			return nil, err
		}

		if err := yaml.Unmarshal(bytes, &currentMap); err != nil {
			return nil, errors.Errorf("failed to parse %s: %s", filePath, err)
		}
		// Merge with the previous map
		base = mergeValues(base, currentMap)
	}

	// User specified a value via --set
	for _, value := range vals {
		if err := strvals.ParseInto(value, base); err != nil {
			return nil, errors.Errorf("failed parsing --set data: %s", err)
		}
	}

	return values.Values(base), nil
}

//readFile load a file from the local directory or a remote file with a url.
func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// Merges source and destination map, preferring values from the source map
func mergeValues(dest map[string]interface{}, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		// If the key doesn't exist already, then just set the key to that value
		if _, exists := dest[k]; !exists {
			dest[k] = v
			continue
		}
		nextMap, ok := v.(map[string]interface{})
		// If it isn't another map, overwrite the value
		if !ok {
			dest[k] = v
			continue
		}
		// If the key doesn't exist already, then just set the key to that value
		if _, exists := dest[k]; !exists {
			dest[k] = nextMap
			continue
		}
		// Edge case: If the key exists in the destination, but isn't a map
		destMap, isMap := dest[k].(map[string]interface{})
		// If the source map has a map for this key, prefer it
		if !isMap {
			dest[k] = v
			continue
		}
		// If we got to this point, it is a map in both, so merge them
		dest[k] = mergeValues(destMap, nextMap)
	}
	return dest
}
