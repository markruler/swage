package converter

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"

	oas "github.com/go-openapi/spec"
	"github.com/markruler/swage/spec"
)

// FIXME: refactor
func extractResponses(swagger *oas.Swagger, operation *oas.Operation) (swageResponses []spec.APIResponse, err error) {
	swageResponses = []spec.APIResponse{}
	oas_responses := operation.Responses
	if oas_responses == nil {
		return nil, errors.New("response is empty")
	}
	if oas_responses.Default != nil {
		if oas_responses.Default.Schema != nil && !reflect.DeepEqual(oas.Ref{}, oas_responses.Default.Schema.Ref) {
			schema, err := oas.ResolveRef(swagger, &oas_responses.Default.Schema.Ref)
			if err != nil {
				return nil, err
			}
			_, definition_name := DefinitionNameFromRef(oas_responses.Default.Schema.Ref)
			swageResponses = append(swageResponses, spec.APIResponse{
				Schema:       definition_name,
				ResponseType: "body",
				DataType:     strings.Join(schema.Type, ","),
				Enum:         "",
				Example:      "",
				Description:  oas_responses.Default.Description,
			})
		} else {
			swageResponses = append(swageResponses, spec.APIResponse{
				Schema:       "",
				ResponseType: "body",
				DataType:     "string",
				Enum:         "",
				Example:      "",
				Description:  oas_responses.Default.Description,
			})
		}
	}

	codes := SortMap(oas_responses.StatusCodeResponses)
	for _, code := range codes {
		swageResponse := &spec.APIResponse{
			StatusCode: code,
		}
		icode, err := strconv.Atoi(code)
		if err != nil {
			return nil, err
		}
		oasResponse := oas_responses.StatusCodeResponses[icode]
		if reflect.DeepEqual(oas.Response{}, oasResponse) {
			continue
		}

		for key, val := range oasResponse.Headers {
			b, err := json.MarshalIndent(val.Example, "", "    ")
			if err != nil {
				return nil, err
			}
			swageResponse.Schema = key
			swageResponse.ResponseType = "header"
			swageResponse.DataType = val.Type
			swageResponse.Enum = Enum2string(val.Enum...)
			swageResponse.Example = string(b)
			swageResponse.Description = val.Description
			swageResponses = append(swageResponses, *swageResponse)
		}

		if oasResponse.Schema != nil && !reflect.DeepEqual(oasResponse.Schema.Ref, oas.Ref{}) {
			swageResponse, err = responseSchemaRef(swageResponse, oasResponse.Schema.Ref, swagger)
			if err != nil {
				return nil, err
			}
			if swageResponse != nil {
				swageResponses = append(swageResponses, *swageResponse)
			}
			continue
		}

		if oasResponse.Schema != nil {
			swageResponse.ResponseType = strings.Join(oasResponse.Schema.Type, ",")
			swageResponse.Description = oasResponse.Description
			swageResponse, err = responseSchema(swageResponse, *oasResponse.Schema, swagger)
			if err != nil {
				return nil, err
			}
			if swageResponse != nil {
				swageResponses = append(swageResponses, *swageResponse)
			}
			continue
		}

		swageResponse.Schema = ""
		swageResponse.ResponseType = ""
		swageResponse.DataType = ""
		swageResponse.Enum = ""
		swageResponse.Example = ""
		swageResponse.Description = oasResponse.Description
		swageResponses = append(swageResponses, *swageResponse)
	}
	return swageResponses, nil
}

func responseSchema(swageResponse *spec.APIResponse, schema oas.Schema, swagger *oas.Swagger) (*spec.APIResponse, error) {
	if schema.Type != nil {
		swageResponse.ResponseType = "body"
		swageResponse.DataType = strings.Join(schema.Type, ",")
	}

	if schema.Type.Contains("array") {
		return arrayDefinitionFromSchemaRef(swageResponse, schema, swagger)
	}

	example, _ := extractExample(schema.Example)

	swageResponse.Schema = schema.Title
	swageResponse.ResponseType = "body"
	swageResponse.DataType = "object"
	swageResponse.Example = example
	if schema.Description != "" {
		swageResponse.Description = schema.Description
	}
	return swageResponse, nil
}

func responseSchemaRef(swageResponse *spec.APIResponse, ref oas.Ref, swagger *oas.Swagger) (*spec.APIResponse, error) {
	schema, err := oas.ResolveRef(swagger, &ref)
	if err != nil {
		return nil, err
	}

	if schema == nil {
		return nil, errors.New("not found response.Schema.Ref definition")
	}

	example, _ := extractExample(schema.Example)

	_, swageResponse.Schema = DefinitionNameFromRef(ref)
	swageResponse.ResponseType = "body"
	swageResponse.DataType = "object"
	swageResponse.Example = example
	swageResponse.Description = schema.Description
	return swageResponse, nil
}

func arrayDefinitionFromSchemaRef(swageResponse *spec.APIResponse, schema oas.Schema, swagger *oas.Swagger) (*spec.APIResponse, error) {
	items := schema.Items
	if items.Schema != nil {
		schema := items.Schema
		swageResponse.Schema = schema.Title
		swageResponse.ResponseType = "body"
		swageResponse.DataType = strings.Join(schema.Type, ",")
		swageResponse.Description = schema.Description
		return swageResponse, nil
	}
	for _, schema := range items.Schemas {
		if !reflect.DeepEqual(oas.Ref{}, schema.Ref) {
			_, definition_name := DefinitionNameFromRef(items.Schemas[0].Ref)
			definition := swagger.Definitions[definition_name]
			if reflect.DeepEqual(oas.Schema{}, definition) {
				return nil, errors.New("not found definition")
			}
			if schema.Example != nil {
				b, err := json.MarshalIndent(schema.Example, "", "    ")
				if err != nil {
					return nil, err
				}
				swageResponse.Example = string(b)
			}
			swageResponse.Schema = definition_name
			swageResponse.ResponseType = "body"
			swageResponse.DataType = "array"
			if schema.Description != "" {
				swageResponse.Description = schema.Description
			}
		}
		// append
	}
	return swageResponse, nil
}
