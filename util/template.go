package util

import (
	"bytes"
	"text/template"

	"k8s.io/apimachinery/pkg/util/yaml"
)

// ObjectFromTemplate executes Go template with given source object and
// parses the result into the destination object structure.
func ObjectFromTemplate(source interface{}, destination interface{}, templateFile string) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	var stringBuffer string
	buffer := bytes.NewBufferString(stringBuffer)
	err = tmpl.Execute(buffer, source)
	if err != nil {
		return err
	}

	return yaml.NewYAMLOrJSONDecoder(buffer, 64).Decode(destination)
}
