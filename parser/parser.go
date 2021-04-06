package parser

import (
	"strings"

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

// Convert Open API Spec to Swage Spec
func Convert(swagger *oas.Swagger) (*SwageSpec, error) {
	swageSpec := &SwageSpec{}
	paths := SortMap(swagger.Paths.Paths)
	for _, path := range paths {
		// xlsx/simple/index.go
		// TODO: support all methods
		operations := swagger.Paths.Paths[path]
		if operations.PathItemProps.Get != nil {
			swageAPI, err := extractOperation(swagger, path, "GET", operations.PathItemProps.Get)
			if err != nil {
				return nil, err
			}
			swageSpec.API = append(swageSpec.API, *swageAPI)
		}
	}
	return swageSpec, nil
}

func extractOperation(swagger *oas.Swagger, path, method string, operation *oas.Operation) (api *SwageAPI, err error) {
	var requests []APIRequest
	if requests, err = extractRequests(swagger, operation); err != nil {
		return nil, err
	}

	var responses []APIResponse
	if responses, err = extractResponses(swagger, operation); err != nil {
		return nil, err
	}

	// xlsx/simple/api_header.go
	return &SwageAPI{
		Header: APIHeader{
			Tag:         strings.Join(operation.Tags, ","),
			ID:          operation.ID,
			Path:        path,
			Method:      method,
			Consumes:    strings.Join(operation.Consumes, ", "),
			Produces:    strings.Join(operation.Produces, ", "),
			Summary:     operation.Summary,
			Description: operation.Description,
		},
		Request:  requests,
		Response: responses,
	}, nil
}
