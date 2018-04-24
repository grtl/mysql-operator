package util

import (
	"bytes"
	"github.com/grtl/mysql-operator/artifacts"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// ObjectFromFile creates a kubernetes object from a yaml file.
func ObjectFromFile(filename string, destination interface{}) error {
	assetBytes, err := artifacts.Asset(filename)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(assetBytes)
	return yaml.NewYAMLOrJSONDecoder(reader, 32).Decode(destination)
}
