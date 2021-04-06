package parser

import (
	"reflect"
	"strings"

	oas "github.com/go-openapi/spec"
)

// TODO: convert "#/definitions/Def"
// @source editor.swagger.json
// @method POST
// @path /user

// TODO: AllOf
// @source docker.v1.41.json
// @method POST
// @path /containers/create
func extractRequests(swagger *oas.Swagger, operation *oas.Operation) (requests []APIRequest, err error) {
	requests = []APIRequest{}
	for _, param := range operation.Parameters {
		request := APIRequest{}
		var definition_type, definition_name string
		if param.Schema != nil && !reflect.DeepEqual(param.Schema.Ref, oas.Ref{}) {
			definition_type, definition_name = DefinitionNameFromRef(param.Schema.Ref)
		}

		if definition_type == "parameters" {
			param = swagger.Parameters[definition_name]
		}

		if definition_type == "definitions" {
			definition := swagger.Definitions[definition_name]
			param.Name = definition_name
			param.In = "body"
			param.Type = strings.Join(definition.Type, ",")
			param.Example = definition.Example
			param.Description = definition.Description
		}

		request.Required = checkRequired(param.Required)
		request.Schema = param.Name
		request.ParameterType = param.In
		request.DataType = param.Type
		request.Description = param.Description

		if param.Schema != nil && param.Schema.Type != nil {
			request.DataType = strings.Join(param.Schema.Type, ",")
			if param.Schema.Description != "" {
				request.Description = param.Schema.Description
			}
		}

		if param.Items != nil && param.Items.Enum != nil {
			request.Enum = Enum2string(param.Items.Enum...)
		}

		if param.Enum != nil {
			request.Enum = Enum2string(param.Enum...)
		}

		var example string
		if param.Example != nil {
			example, err = extractExample(param.Example)
			if err != nil {
				return nil, err
			}
		}

		request.Example = example
		requests = append(requests, request)
	}
	return requests, nil
}
