package util

import (
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"
)

// ObjectFromFile creates a kubernetes object from a yaml file.
func ObjectFromFile(filename string, destination interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	return yaml.NewYAMLOrJSONDecoder(f, 32).Decode(destination)
}
