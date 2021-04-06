package parser

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestExtractOperation(t *testing.T) {
	test_operation := &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						204: {
							ResponseProps: spec.ResponseProps{
								Description: "no error",
							},
						},
						400: {
							ResponseProps: spec.ResponseProps{
								Description: "bad parameter",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef("#/definitions/ErrorResponse"),
									},
								},
							},
						},
						404: {
							ResponseProps: spec.ResponseProps{
								Description: "bad parameter",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Description: "no such container",
										Ref:         spec.MustCreateRef("#/definitions/ErrorResponse"),
									},
									SwaggerSchemaProps: spec.SwaggerSchemaProps{
										Example: map[string]map[string]interface{}{
											"application/json": {
												"message": "No such container: c2ada9df5af8",
											},
										},
									},
								},
							},
						},
						409: {
							ResponseProps: spec.ResponseProps{
								Description: "indicates a request conflict with current state of the target resource.",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Title: "Conflict",
										Type:  spec.StringOrArray{"body"},
									},
									SwaggerSchemaProps: spec.SwaggerSchemaProps{
										Example: map[string]map[string]interface{}{
											"application/json": {
												"message": "You cannot remove a running container: c2ada9df5af8. Stop the\ncontainer before attempting removal or force remove\n",
											},
										},
									},
								},
							},
						},
						500: {
							ResponseProps: spec.ResponseProps{
								Description: "server error",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef("#/definitions/ErrorResponse"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	oas_spec := &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: spec.Definitions{
				"ErrorResponse": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Description: "Account plan object",
						Type: spec.StringOrArray{
							"object",
						},
						Properties: spec.SchemaProperties{
							"hosts": spec.Schema{
								SchemaProps: spec.SchemaProps{
									Description: "Account plan number of hosts",
									Type: spec.StringOrArray{
										"integer",
									},
								},
							},
						},
					},
					SwaggerSchemaProps: spec.SwaggerSchemaProps{
						Example: map[string]interface{}{
							"message": "Something went wrong.",
						},
					},
				},
			},
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{
					"/": {
						PathItemProps: spec.PathItemProps{
							Get: test_operation,
						},
					},
				},
			},
		},
	}

	api, err := extractOperation(oas_spec, "/", "GET", test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "204", api.Response[0].StatusCode)
	assert.Equal(t, "", api.Response[0].Schema)
	assert.Equal(t, "", api.Response[0].ResponseType)
	assert.Equal(t, "", api.Response[0].DataType)
	assert.Equal(t, "", api.Response[0].Enum)
	assert.Equal(t, "", api.Response[0].Example)
	assert.Equal(t, "no error", api.Response[0].Description)

	assert.Equal(t, "400", api.Response[1].StatusCode)
	assert.Equal(t, "ErrorResponse", api.Response[1].Schema)
	assert.Equal(t, "body", api.Response[1].ResponseType)
	assert.Equal(t, "object", api.Response[1].DataType)
	assert.Equal(t, "", api.Response[1].Enum)
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", api.Response[1].Example)
	assert.Equal(t, "Account plan object", api.Response[1].Description)

	assert.Equal(t, "404", api.Response[2].StatusCode)
	assert.Equal(t, "ErrorResponse", api.Response[2].Schema)
	assert.Equal(t, "body", api.Response[2].ResponseType)
	assert.Equal(t, "object", api.Response[2].DataType)
	assert.Equal(t, "", api.Response[2].Enum)
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", api.Response[2].Example)
	assert.Equal(t, "Account plan object", api.Response[2].Description)

	assert.Equal(t, "409", api.Response[3].StatusCode)
	assert.Equal(t, "Conflict", api.Response[3].Schema)
	assert.Equal(t, "body", api.Response[3].ResponseType)
	assert.Equal(t, "object", api.Response[3].DataType)
	assert.Equal(t, "", api.Response[3].Enum)
	assert.Equal(t, "{\n    \"application/json\": {\n        \"message\": \"You cannot remove a running container: c2ada9df5af8. Stop the\\ncontainer before attempting removal or force remove\\n\"\n    }\n}", api.Response[3].Example)
	assert.Equal(t, "indicates a request conflict with current state of the target resource.", api.Response[3].Description)

	assert.Equal(t, "500", api.Response[4].StatusCode)
	assert.Equal(t, "ErrorResponse", api.Response[4].Schema)
	assert.Equal(t, "body", api.Response[4].ResponseType)
	assert.Equal(t, "object", api.Response[4].DataType)
	assert.Equal(t, "", api.Response[4].Enum)
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", api.Response[4].Example)
	assert.Equal(t, "Account plan object", api.Response[4].Description)
}
