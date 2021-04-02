package simple

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

// @source docker.v1.41.json
// @method head
// @path /containers/{id}/archive
func TestResponseHeaders(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error
	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			ID: "ContainerArchiveInfo",
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "no error",
								Schema:      &spec.Schema{},
								Headers: map[string]spec.Header{
									"X-Docker-Container-Path-Stat": {
										SimpleSchema: spec.SimpleSchema{
											Type: "string",
										},
										HeaderProps: spec.HeaderProps{
											Description: "A base64 - encoded JSON object with some filesystem header\ninformation about the path\n",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "200", row[15][0])
	assert.Equal(t, "X-Docker-Container-Path-Stat", row[15][1])
	assert.Equal(t, "header", row[15][2])
	assert.Equal(t, "string", row[15][3])
	assert.Equal(t, "A base64 - encoded JSON object with some filesystem header\ninformation about the path\n", row[15][6])
}

// @source docker.v1.41.json
func TestResponseWithoutSchema(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error
	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			ID: "ContainerInspect",
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						// @method post
						// @path /containers/{id}/attach
						101: {
							ResponseProps: spec.ResponseProps{
								Description: "no error, hints proxy about hijacking",
							},
						},
						// @method get
						// @path /containers/{id}/json
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "no error",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Title: "ContainerInspectResponse",
										Type:  []string{"object"},
										Properties: spec.SchemaProperties{
											"Id": {
												SchemaProps: spec.SchemaProps{
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
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "101", row[15][0])
	assert.Equal(t, "", row[15][1])
	assert.Equal(t, "", row[15][2])
	assert.Equal(t, "", row[15][3])
	assert.Equal(t, "no error, hints proxy about hijacking", row[15][6])
	assert.Equal(t, "200", row[16][0])
	assert.Equal(t, "ContainerInspectResponse", row[16][1])
	assert.Equal(t, "body", row[16][2])
	assert.Equal(t, "object", row[16][3])
	assert.Equal(t, "no error", row[16][6])
}

// @source docker.v1.41.json
// @method DELETE
// @path /containers/{id}
func TestResponseSchemaWithRef(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()

	xl.SwaggerSpec = &spec.Swagger{
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
		},
	}
	var err error
	err = simple.CreateAPISheet("", "", &spec.Operation{
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
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)

	assert.Equal(t, "204", row[15][0])
	assert.Equal(t, "", row[15][1])
	assert.Equal(t, "", row[15][2])
	assert.Equal(t, "", row[15][3])
	assert.Equal(t, "", row[15][4])
	assert.Equal(t, "", row[15][5])
	assert.Equal(t, "no error", row[15][6])

	assert.Equal(t, "400", row[16][0])
	assert.Equal(t, "ErrorResponse", row[16][1])
	assert.Equal(t, "body", row[16][2])
	assert.Equal(t, "object", row[16][3])
	assert.Equal(t, "", row[16][4])
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", row[16][5])
	assert.Equal(t, "bad parameter", row[16][6])

	assert.Equal(t, "404", row[17][0])
	assert.Equal(t, "ErrorResponse", row[17][1])
	assert.Equal(t, "body", row[17][2])
	assert.Equal(t, "object", row[17][3])
	assert.Equal(t, "", row[17][4])
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", row[17][5])
	assert.Equal(t, "bad parameter", row[17][6])

	assert.Equal(t, "409", row[18][0])
	assert.Equal(t, "Conflict", row[18][1])
	assert.Equal(t, "body", row[18][2])
	assert.Equal(t, "object", row[18][3])
	assert.Equal(t, "", row[18][4])
	assert.Equal(t, "{\n    \"application/json\": {\n        \"message\": \"You cannot remove a running container: c2ada9df5af8. Stop the\\ncontainer before attempting removal or force remove\\n\"\n    }\n}", row[18][5])
	assert.Equal(t, "indicates a request conflict with current state of the target resource.", row[18][6])

	assert.Equal(t, "500", row[19][0])
	assert.Equal(t, "ErrorResponse", row[19][1])
	assert.Equal(t, "body", row[19][2])
	assert.Equal(t, "object", row[19][3])
	assert.Equal(t, "", row[19][4])
	assert.Equal(t, "{\n    \"message\": \"Something went wrong.\"\n}", row[19][5])
	assert.Equal(t, "server error", row[19][6])
}

// @source cisco.meraki.json
// @method POST
// @path /devices/{serial}/camera/generateSnapshot
func TestResponseSchemaWithoutRef(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error

	xl.SwaggerSpec = &spec.Swagger{}
	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						202: {
							ResponseProps: spec.ResponseProps{
								Description: "Successful operation",
								Examples: map[string]interface{}{
									"application/json": map[string]string{
										"expiry": "Access to the image will expire at 2018-12-11T03:12:39Z.",
										"url":    "https://spn4.meraki.com/stream/jpeg/snapshot/b2d123asdf423qd22d2",
									},
								},
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: spec.StringOrArray{
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
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "202", row[15][0])
	assert.Equal(t, "", row[15][1])
	assert.Equal(t, "body", row[15][2])
	assert.Equal(t, "object", row[15][3])
	assert.Equal(t, "Successful operation", row[15][6])
}

// @source spotify.json
// @method get
// @path /albums
func TestResponseSchemaItemsWithRef(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: spec.Definitions{
				"Test": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Description: "Test definition",
						Type: spec.StringOrArray{
							"string",
						},
					},
				},
			},
		},
	}
	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "OK",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: spec.StringOrArray{
											"array",
										},
										Items: &spec.SchemaOrArray{
											Schemas: []spec.Schema{
												{
													SchemaProps: spec.SchemaProps{
														Ref: spec.MustCreateRef("#/definitions/Test"),
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
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "200", row[15][0])
	assert.Equal(t, "Test", row[15][1])
	assert.Equal(t, "body", row[15][2])
	assert.Equal(t, "array", row[15][3])
	assert.Equal(t, "OK", row[15][6])
}

func TestResponseSchemaItemsWithoutRef(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error

	// @source docker.v1.41.json
	// @method POST
	// @path /containers/prune
	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "No error",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:  []string{"object"},
										Title: "ContainerPruneResponse",
										Properties: spec.SchemaProperties{
											"ContainersDeleted": {
												SchemaProps: spec.SchemaProps{
													Description: "Container IDs that were deleted",
													Type:        []string{"array"},
													Items: &spec.SchemaOrArray{
														Schema: &spec.Schema{
															SchemaProps: spec.SchemaProps{
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
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "200", row[15][0])
	assert.Equal(t, "ContainerPruneResponse", row[15][1])
	assert.Equal(t, "body", row[15][2])
	assert.Equal(t, "object", row[15][3])
	assert.Equal(t, "No error", row[15][6])

	// @source docker.v1.41.json
	// @method GET
	// @path /images/search
	// @path /images/{name}/history
	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "No error",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Type:  []string{"object"},
													Title: "ImageSearchResponseItem",
													Properties: spec.SchemaProperties{
														"description": {
															SchemaProps: spec.SchemaProps{
																Type: []string{"string"},
															},
														},
														"is_official": {
															SchemaProps: spec.SchemaProps{
																Type: []string{"boolean"},
															},
														},
														"is_automated": {
															SchemaProps: spec.SchemaProps{
																Type: []string{"boolean"},
															},
														},
														"name": {
															SchemaProps: spec.SchemaProps{
																Type: []string{"string"},
															},
														},
														"star_count": {
															SchemaProps: spec.SchemaProps{
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
	}, nil, 2)
	assert.NoError(t, err)
	row, err = xl.File.GetRows("2")
	assert.NoError(t, err)
	assert.Equal(t, "200", row[15][0])
	assert.Equal(t, "ImageSearchResponseItem", row[15][1])
	assert.Equal(t, "body", row[15][2])
	assert.Equal(t, "array", row[15][3])
	assert.Equal(t, "No error", row[15][6])
}

// @source bbc.radio.json
// @method GET
// @path /broadcasts/latest
func TestResponseDefault(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: spec.Definitions{
				"BroadcastsResponse": spec.Schema{
					SchemaProps: spec.SchemaProps{},
				},
				"ErrorResponse": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type:     []string{"object"},
						Required: []string{"errors"},
					},
				},
			},
		},
	}

	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "OK",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef("#/definitions/BroadcastsResponse"),
									},
								},
							},
						},
					},
					Default: &spec.Response{
						ResponseProps: spec.ResponseProps{
							Description: "Unexpected error",
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
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "default", row[15][0])
	assert.Equal(t, "ErrorResponse", row[15][1])
	assert.Equal(t, "body", row[15][2])
	assert.Equal(t, "object", row[15][3])
	assert.Equal(t, "Unexpected error", row[15][6])
	assert.Equal(t, "200", row[16][0])
	assert.Equal(t, "BroadcastsResponse", row[16][1])
	assert.Equal(t, "body", row[16][2])
	assert.Equal(t, "object", row[16][3])
	assert.Equal(t, "OK", row[16][6])
}

func TestPropDefinitionFromSchemaRef(t *testing.T) {
	xl := New()
	err := xl.propDefinitionFromSchemaRef(spec.Response{})
	assert.EqualError(t, err, "response is empty")
}

// FIXME:
// @source zoom.us.json
// @method GET
// @path /accounts
// @path /groups
func TestAllOfDefinitionWithRef(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: spec.Definitions{
				"AccountList": spec.Schema{
					SchemaProps: spec.SchemaProps{
						AllOf: []spec.Schema{
							{
								SchemaProps: spec.SchemaProps{
									Ref: spec.MustCreateRef("#/definitions/Pagination"),
								},
							},
							{
								SchemaProps: spec.SchemaProps{
									Ref: spec.MustCreateRef("#/definitions/AccountListItem"),
								},
							},
						},
					},
				},
				"Pagination": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Description: "Pagination Object",
						Properties: spec.SchemaProperties{
							"page_count": {
								SchemaProps: spec.SchemaProps{
									Description: "The number of items returned on this page",
									Type: spec.StringOrArray{
										"integer",
									},
								},
							},
							"page_number": {
								SchemaProps: spec.SchemaProps{
									Description: "The page number of current results",
									Type: spec.StringOrArray{
										"integer",
									},
									Default: 1,
								},
							},
							"page_size": {
								SchemaProps: spec.SchemaProps{
									Description: "The number of records returned within a single API call",
									Type: spec.StringOrArray{
										"integer",
									},
									Default: 30,
									Maximum: &[]float64{300}[0],
								},
							},
							"total_records": {
								SchemaProps: spec.SchemaProps{
									Description: "The number of all records available across pages",
									Type: spec.StringOrArray{
										"integer",
									},
								},
							},
						},
						Type: spec.StringOrArray{
							"object",
						},
					},
				},
				"AccountListItem": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Description: "Account object in account list",
						Type: spec.StringOrArray{
							"object",
						},
					},
				},
			},
		},
	}
	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "Account list returned",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef("#/definitions/AccountList"),
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)

	assert.Equal(t, "200", row[15][0])
	assert.Equal(t, "AccountList", row[15][1])
	assert.Equal(t, "body", row[15][2])
	assert.Equal(t, "object", row[15][3])
	assert.Equal(t, "Account list returned", row[15][6])
}
