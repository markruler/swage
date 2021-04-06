package parser

import (
	"fmt"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

// Parse `JSON`, `YAML` to Go struct
func Parse(sourcePath string) (*spec.Swagger, error) {
	doc, err := loads.Spec(sourcePath)
	if err != nil {
		return nil, err
	}
	return doc.Spec(), nil
}

// Convert Open API Spec to Swage Spec
func Convert(swagger *spec.Swagger) (*SwageSpec, error) {
	spec := &SwageSpec{}

	paths := SortMap(swagger.Paths.Paths)
	for _, path := range paths {
		// xlsx/simple/index.go
		operations := swagger.Paths.Paths[path]
		if operations.PathItemProps.Get != nil {
			spec.API = append(spec.API, SwageAPI{
				// xlsx/simple/api_header.go
				Header: APIHeader{
					Tag:         strings.Join(operations.Get.Tags, ","),
					ID:          operations.Get.ID,
					Path:        path,
					Method:      "GET",
					Consumes:    strings.Join(operations.Get.Consumes, ", "),
					Produces:    strings.Join(operations.Get.Produces, ", "),
					Summary:     operations.Get.Summary,
					Description: operations.Get.Description,
				},
				// xlsx/simple/api_request.go
				Request: APIRequest{
					// Required: operations.Get.Parameters[0].Required,
				},
				// xlsx/simple/api_response.go
				Response: APIResponse{
					// StatusCode: operations.Get.Responses,
				},
			})
		}
	}
	for _, api := range spec.API {
		fmt.Printf("%v\n", api)
	}
	return spec, nil
}
