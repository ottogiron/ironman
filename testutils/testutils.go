package testutils

import (
	"encoding/json"
	"testing"
)

//Marshal marshals the value to a string representation
func Marshal(val interface{}, t *testing.T) string {
	bytes, err := json.Marshal(val)

	if err != nil {
		t.Fatalf("Failed to marshal object %v", val)
	}
	return string(bytes)
}
