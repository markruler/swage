package spec

import (
	"github.com/go-openapi/loads"
	oas "github.com/go-openapi/spec"
)

// Parse `JSON`, `YAML` to Go struct
func Parse(sourcePath string) (*oas.Swagger, error) {
	doc, err := loads.Spec(sourcePath)
	if err != nil {
		return nil, err
	}
	return doc.Spec(), nil
}
