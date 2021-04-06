package parser

import (
	"testing"

	oas "github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

// @source docker.v1.41.json
// @method head
// @path /containers/{id}/archive
func TestExtractResponses_WithoutSchema(t *testing.T) {
	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Definitions: oas.Definitions{},
		},
	}

	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			ID: "ContainerInspect",
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						// @method post
						// @path /containers/{id}/attach
						101: {
							ResponseProps: oas.ResponseProps{
								Description: "no error, hints proxy about hijacking",
							},
						},
						// @method get
						// @path /containers/{id}/json
						200: {
							ResponseProps: oas.ResponseProps{
								Description: "no error",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Title: "ContainerInspectResponse",
										Type:  []string{"object"},
										Properties: oas.SchemaProperties{
											"Id": {
												SchemaProps: oas.SchemaProps{
													Description: "The ID of the container",
													Type:        []string{"string"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	api_response, err := extractResponses(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "101", api_response[0].StatusCode)
	assert.Equal(t, "", api_response[0].Schema)
	assert.Equal(t, "", api_response[0].ResponseType)
	assert.Equal(t, "", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	assert.Equal(t, "no error, hints proxy about hijacking", api_response[0].Description)

	assert.Equal(t, "200", api_response[1].StatusCode)
	assert.Equal(t, "ContainerInspectResponse", api_response[1].Schema)
	assert.Equal(t, "body", api_response[1].ResponseType)
	assert.Equal(t, "object", api_response[1].DataType)
	assert.Equal(t, "", api_response[1].Enum)
	assert.Equal(t, "", api_response[1].Example)
	assert.Equal(t, "no error", api_response[1].Description)
}

// @source cisco.meraki.json
// @method POST
// @path /devices/{serial}/camera/generateSnapshot
func TestExtractResponses_Schema(t *testing.T) {
	oas_spec := &oas.Swagger{}
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						202: {
							ResponseProps: oas.ResponseProps{
								Description: "Successful operation",
								Examples: map[string]interface{}{
									"application/json": map[string]string{
										"expiry": "Access to the image will expire at 2018-12-11T03:12:39Z.",
										"url":    "https://spn4.meraki.com/stream/jpeg/snapshot/b2d123asdf423qd22d2",
									},
								},
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Type: oas.StringOrArray{
											"object",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	api_response, err := extractResponses(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "202", api_response[0].StatusCode)
	assert.Equal(t, "", api_response[0].Schema)
	assert.Equal(t, "body", api_response[0].ResponseType)
	assert.Equal(t, "object", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	assert.Equal(t, "Successful operation", api_response[0].Description)
}

// @source docker.v1.41.json
// @method DELETE
// @path /containers/{id}
func TestExtractResponses_SchemaRef(t *testing.T) {
	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Definitions: oas.Definitions{
				"ErrorResponse": oas.Schema{
					SchemaProps: oas.SchemaProps{
						Description: "Account plan object",
						Type: oas.StringOrArray{
							"object",
						},
						Properties: oas.SchemaProperties{
							"hosts": oas.Schema{
								SchemaProps: oas.SchemaProps{
									Description: "Account plan number of hosts",
									Type: oas.StringOrArray{
										"integer",
									},
								},
							},
						},
					},
					SwaggerSchemaProps: oas.SwaggerSchemaProps{
						Example: map[string]interface{}{
							"message": "Something went wrong.",
						},
					},
				},
			},
		},
	}

	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						204: {
							ResponseProps: oas.ResponseProps{
								Description: "no error",
							},
						},
						400: {
							ResponseProps: oas.ResponseProps{
								Description: "bad parameter",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Ref: oas.MustCreateRef("#/definitions/ErrorResponse"),
									},
								},
							},
						},
						404: {
							ResponseProps: oas.ResponseProps{
								Description: "bad parameter",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Description: "no such container",
										Ref:         oas.MustCreateRef("#/definitions/ErrorResponse"),
									},
									SwaggerSchemaProps: oas.SwaggerSchemaProps{
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
							ResponseProps: oas.ResponseProps{
								Description: "indicates a request conflict with current state of the target resource.",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Title: "Conflict",
										Type:  oas.StringOrArray{"body"},
									},
									SwaggerSchemaProps: oas.SwaggerSchemaProps{
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
							ResponseProps: oas.ResponseProps{
								Description: "server error",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Ref: oas.MustCreateRef("#/definitions/ErrorResponse"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	api_response, err := extractResponses(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "204", api_response[0].StatusCode)
	assert.Equal(t, "", api_response[0].Schema)
	assert.Equal(t, "", api_response[0].ResponseType)
	assert.Equal(t, "", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	assert.Equal(t, "no error", api_response[0].Description)

	assert.Equal(t, "400", api_response[1].StatusCode)
	assert.Equal(t, "ErrorResponse", api_response[1].Schema)
	assert.Equal(t, "body", api_response[1].ResponseType)
	assert.Equal(t, "object", api_response[1].DataType)
	assert.Equal(t, "", api_response[1].Enum)
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", api_response[1].Example)
	assert.Equal(t, "Account plan object", api_response[1].Description)

	assert.Equal(t, "404", api_response[2].StatusCode)
	assert.Equal(t, "ErrorResponse", api_response[2].Schema)
	assert.Equal(t, "body", api_response[2].ResponseType)
	assert.Equal(t, "object", api_response[2].DataType)
	assert.Equal(t, "", api_response[2].Enum)
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", api_response[2].Example)
	assert.Equal(t, "Account plan object", api_response[2].Description)

	assert.Equal(t, "409", api_response[3].StatusCode)
	assert.Equal(t, "Conflict", api_response[3].Schema)
	assert.Equal(t, "body", api_response[3].ResponseType)
	assert.Equal(t, "object", api_response[3].DataType)
	assert.Equal(t, "", api_response[3].Enum)
	assert.Equal(t, "{\n    \"application/json\": {\n        \"message\": \"You cannot remove a running container: c2ada9df5af8. Stop the\\ncontainer before attempting removal or force remove\\n\"\n    }\n}", api_response[3].Example)
	assert.Equal(t, "indicates a request conflict with current state of the target resource.", api_response[3].Description)

	assert.Equal(t, "500", api_response[4].StatusCode)
	assert.Equal(t, "ErrorResponse", api_response[4].Schema)
	assert.Equal(t, "body", api_response[4].ResponseType)
	assert.Equal(t, "object", api_response[4].DataType)
	assert.Equal(t, "", api_response[4].Enum)
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", api_response[4].Example)
	assert.Equal(t, "Account plan object", api_response[4].Description)
}

// @source spotify.json
// @method get
// @path /albums
func TestExtractResponses_ItemsWithRef(t *testing.T) {
	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Definitions: oas.Definitions{
				"Test": oas.Schema{
					SchemaProps: oas.SchemaProps{
						Description: "Test definition",
						Type: oas.StringOrArray{
							"string",
						},
					},
				},
			},
		},
	}

	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						200: {
							ResponseProps: oas.ResponseProps{
								Description: "OK",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Type: oas.StringOrArray{
											"array",
										},
										Items: &oas.SchemaOrArray{
											Schemas: []oas.Schema{
												{
													SchemaProps: oas.SchemaProps{
														Ref: oas.MustCreateRef("#/definitions/Test"),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	api_response, err := extractResponses(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "200", api_response[0].StatusCode)
	assert.Equal(t, "Test", api_response[0].Schema)
	assert.Equal(t, "body", api_response[0].ResponseType)
	assert.Equal(t, "array", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	assert.Equal(t, "OK", api_response[0].Description)
}

// @source docker.v1.41.json
// @method POST
// @path /containers/prune
func TestExtractResponses_ItemsWithoutRef(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						200: {
							ResponseProps: oas.ResponseProps{
								Description: "No error",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Type:  []string{"object"},
										Title: "ContainerPruneResponse",
										Properties: oas.SchemaProperties{
											"ContainersDeleted": {
												SchemaProps: oas.SchemaProps{
													Description: "Container IDs that were deleted",
													Type:        []string{"array"},
													Items: &oas.SchemaOrArray{
														Schema: &oas.Schema{
															SchemaProps: oas.SchemaProps{
																Type: []string{"string"},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	api_response, err := extractResponses(&oas.Swagger{}, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "200", api_response[0].StatusCode)
	assert.Equal(t, "ContainerPruneResponse", api_response[0].Schema)
	assert.Equal(t, "body", api_response[0].ResponseType)
	assert.Equal(t, "object", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	assert.Equal(t, "No error", api_response[0].Description)
}

// @source docker.v1.41.json
// @method GET
// @path /images/search
// @path /images/{name}/history
func TestExtractResponses_NestedItems(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						200: {
							ResponseProps: oas.ResponseProps{
								Description: "No error",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Type: []string{"array"},
										Items: &oas.SchemaOrArray{
											Schema: &oas.Schema{
												SchemaProps: oas.SchemaProps{
													Type:  []string{"object"},
													Title: "ImageSearchResponseItem",
													Properties: oas.SchemaProperties{
														"description": {
															SchemaProps: oas.SchemaProps{
																Type: []string{"string"},
															},
														},
														"is_official": {
															SchemaProps: oas.SchemaProps{
																Type: []string{"boolean"},
															},
														},
														"is_automated": {
															SchemaProps: oas.SchemaProps{
																Type: []string{"boolean"},
															},
														},
														"name": {
															SchemaProps: oas.SchemaProps{
																Type: []string{"string"},
															},
														},
														"star_count": {
															SchemaProps: oas.SchemaProps{
																Type: []string{"integer"},
															},
														},
													},
												},
											},
										},
									},
								},
								Examples: map[string]interface{}{
									"application/json": map[string]interface{}{
										"description":  "",
										"is_official":  false,
										"is_automated": false,
										"name":         "wma55/u1210sshd",
										"star_cound":   0,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	api_response, err := extractResponses(&oas.Swagger{}, test_operation)
	assert.NoError(t, err)

	// FIXME:
	assert.Equal(t, "200", api_response[0].StatusCode)
	assert.Equal(t, "ImageSearchResponseItem", api_response[0].Schema)
	assert.Equal(t, "body", api_response[0].ResponseType)
	// assert.Equal(t, "array", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	// assert.Equal(t, "No error", api_response[0].Description)
}

// @source bbc.radio.json
// @method GET
// @path /broadcasts/latest
func TestExtractResponses_Default(t *testing.T) {
	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Definitions: oas.Definitions{
				"BroadcastsResponse": oas.Schema{
					SchemaProps: oas.SchemaProps{},
				},
				"ErrorResponse": oas.Schema{
					SchemaProps: oas.SchemaProps{
						Type:     []string{"object"},
						Required: []string{"errors"},
					},
				},
			},
		},
	}

	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						200: {
							ResponseProps: oas.ResponseProps{
								Description: "OK",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Ref: oas.MustCreateRef("#/definitions/BroadcastsResponse"),
									},
								},
							},
						},
					},
					Default: &oas.Response{
						ResponseProps: oas.ResponseProps{
							Description: "Unexpected error",
							Schema: &oas.Schema{
								SchemaProps: oas.SchemaProps{
									Ref: oas.MustCreateRef("#/definitions/ErrorResponse"),
								},
							},
						},
					},
				},
			},
		},
	}

	api_response, err := extractResponses(oas_spec, test_operation)
	assert.NoError(t, err)

	// FIXME:
	// assert.Equal(t, "default", api_response[0].StatusCode)
	assert.Equal(t, "ErrorResponse", api_response[0].Schema)
	assert.Equal(t, "body", api_response[0].ResponseType)
	assert.Equal(t, "object", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	assert.Equal(t, "Unexpected error", api_response[0].Description)

	assert.Equal(t, "200", api_response[1].StatusCode)
	assert.Equal(t, "BroadcastsResponse", api_response[1].Schema)
	assert.Equal(t, "body", api_response[1].ResponseType)
	assert.Equal(t, "object", api_response[1].DataType)
	assert.Equal(t, "", api_response[1].Enum)
	assert.Equal(t, "", api_response[1].Example)
	// assert.Equal(t, "OK", api_response[1].Description)
}

// @source zoom.us.json
// @method GET
// @path /accounts
// @path /groups
func TestExtractResponses_AllOfDefinitionWithRef(t *testing.T) {
	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Definitions: oas.Definitions{
				"AccountList": oas.Schema{
					SchemaProps: oas.SchemaProps{
						// FIXME:
						AllOf: []oas.Schema{
							{
								SchemaProps: oas.SchemaProps{
									Ref: oas.MustCreateRef("#/definitions/Pagination"),
								},
							},
							{
								SchemaProps: oas.SchemaProps{
									Ref: oas.MustCreateRef("#/definitions/AccountListItem"),
								},
							},
						},
					},
				},
				"Pagination": oas.Schema{
					SchemaProps: oas.SchemaProps{
						Description: "Pagination Object",
						Properties: oas.SchemaProperties{
							"page_count": {
								SchemaProps: oas.SchemaProps{
									Description: "The number of items returned on this page",
									Type: oas.StringOrArray{
										"integer",
									},
								},
							},
							"page_number": {
								SchemaProps: oas.SchemaProps{
									Description: "The page number of current results",
									Type: oas.StringOrArray{
										"integer",
									},
									Default: 1,
								},
							},
							"page_size": {
								SchemaProps: oas.SchemaProps{
									Description: "The number of records returned within a single API call",
									Type: oas.StringOrArray{
										"integer",
									},
									Default: 30,
									Maximum: &[]float64{300}[0],
								},
							},
							"total_records": {
								SchemaProps: oas.SchemaProps{
									Description: "The number of all records available across pages",
									Type: oas.StringOrArray{
										"integer",
									},
								},
							},
						},
						Type: oas.StringOrArray{
							"object",
						},
					},
				},
				"AccountListItem": oas.Schema{
					SchemaProps: oas.SchemaProps{
						Description: "Account object in account list",
						Type: oas.StringOrArray{
							"object",
						},
					},
				},
			},
		},
	}

	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Responses: &oas.Responses{
				ResponsesProps: oas.ResponsesProps{
					StatusCodeResponses: map[int]oas.Response{
						200: {
							ResponseProps: oas.ResponseProps{
								Description: "Account list returned",
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Ref: oas.MustCreateRef("#/definitions/AccountList"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	api_response, err := extractResponses(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "200", api_response[0].StatusCode)
	assert.Equal(t, "AccountList", api_response[0].Schema)
	assert.Equal(t, "body", api_response[0].ResponseType)
	assert.Equal(t, "object", api_response[0].DataType)
	assert.Equal(t, "", api_response[0].Enum)
	assert.Equal(t, "", api_response[0].Example)
	// assert.Equal(t, "Account list returned", api_response[0].Description)
}
