package parser

import (
	"net/http"
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
		operations := swagger.Paths.Paths[path]
		if operations.PathItemProps.Get != nil {
			swageAPI, err := extractOperation(swagger, path, http.MethodGet, operations.PathItemProps.Get)
			if err != nil {
				return nil, err
			}
			swageSpec.API = append(swageSpec.API, *swageAPI)
		}
		if operations.PathItemProps.Put != nil {
			swageAPI, err := extractOperation(swagger, path, http.MethodPut, operations.PathItemProps.Put)
			if err != nil {
				return nil, err
			}
			swageSpec.API = append(swageSpec.API, *swageAPI)
		}
		if operations.PathItemProps.Post != nil {
			swageAPI, err := extractOperation(swagger, path, http.MethodPost, operations.PathItemProps.Post)
			if err != nil {
				return nil, err
			}
			swageSpec.API = append(swageSpec.API, *swageAPI)
		}
		if operations.PathItemProps.Delete != nil {
			swageAPI, err := extractOperation(swagger, path, http.MethodDelete, operations.PathItemProps.Delete)
			if err != nil {
				return nil, err
			}
			swageSpec.API = append(swageSpec.API, *swageAPI)
		}
		if operations.PathItemProps.Options != nil {
			swageAPI, err := extractOperation(swagger, path, http.MethodOptions, operations.PathItemProps.Options)
			if err != nil {
				return nil, err
			}
			swageSpec.API = append(swageSpec.API, *swageAPI)
		}
		if operations.PathItemProps.Head != nil {
			swageAPI, err := extractOperation(swagger, path, http.MethodHead, operations.PathItemProps.Head)
			if err != nil {
				return nil, err
			}
			swageSpec.API = append(swageSpec.API, *swageAPI)
		}
		if operations.PathItemProps.Patch != nil {
			swageAPI, err := extractOperation(swagger, path, http.MethodPatch, operations.PathItemProps.Patch)
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
