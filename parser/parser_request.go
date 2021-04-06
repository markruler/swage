package parser

import (
	"reflect"

	oas "github.com/go-openapi/spec"
)

func extractRequests(swagger *oas.Swagger, operation *oas.Operation) (requests []APIRequest, err error) {
	// xlsx/simple/api_request.go
	requests = []APIRequest{}
	for _, param := range operation.Parameters {
		request := APIRequest{}
		if !reflect.DeepEqual(param.Ref, oas.Ref{}) {
			// parameterFromRef(ref)
			name := DefinitionNameFromRef(param.Ref)
			param = swagger.Parameters[name]
		}
		request.Required = checkRequired(param.Required)
		request.Schema = param.Name
		request.ParameterType = param.In
		request.DataType = param.Type
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
		request.Description = param.Description
		requests = append(requests, request)
	}
	return requests, nil
}
