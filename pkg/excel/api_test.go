package excel

import (
	"testing"

	"github.com/cxsu/swage/pkg/parser"
	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPISheet(t *testing.T) {
	xl := New()
	var err error

	err = xl.createAPISheet("", "", nil, nil, 1)
	assert.Error(t, err)

	err = xl.createAPISheet("", "", &spec.Operation{}, nil, 1)
	assert.NoError(t, err)

	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Required: true,
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)

	p := parser.New("../../aio/testdata/json/dev.json")
	xl.SwaggerSpec, _ = p.Parse()
	xl.createAPISheet("", "", nil, nil, 1)
}

// @source docker.v1.41.json
// @method GET
// @path /containers/json
func TestParameterWithoutSchema(t *testing.T) {
	xl := New()
	var err error
	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Name:        "all",
						In:          "query",
						Description: "Return all containers. By default, only running containers are shown.\n",
						Required:    true,
					},
					SimpleSchema: spec.SimpleSchema{
						Type:    "boolean",
						Default: false,
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "O", row[12][0])
	assert.Equal(t, "all", row[12][1])
	assert.Equal(t, "query", row[12][2])
	assert.Equal(t, "boolean", row[12][3])
	assert.Equal(t, "Return all containers. By default, only running containers are shown.\n", row[12][6])
}

func TestParameterSchemaWithRef(t *testing.T) {
	xl := New()
	var err error
	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			// @source zoom.us.json
			// @method PUT
			// @path /meetings/{meetingId}/recordings/{recordingId}/status
			Parameters: map[string]spec.Parameter{
				"PageSize": {
					ParamProps: spec.ParamProps{
						Description: "The number of records returned within a single API call",
						In:          "query",
						Name:        "page_size",
					},
					SimpleSchema: spec.SimpleSchema{
						Type:    "integer",
						Default: 30,
					},
					CommonValidations: spec.CommonValidations{
						Maximum: &[]float64{300}[0],
					},
				},
			},
			// @source editor.swagger.json
			// @method POST
			// @path /user
			Definitions: spec.Definitions{
				"User": {
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{
							"object",
						},
						Properties: spec.SchemaProperties{
							"id": {
								SchemaProps: spec.SchemaProps{
									Type:   []string{"integer"},
									Format: "int64",
								},
							},
							"username": {
								SchemaProps: spec.SchemaProps{
									Type: []string{"string"},
								},
							},
						},
					},
				},
			},
		},
	}
	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Ref: spec.MustCreateRef("#/parameters/PageSize"),
							},
						},
					},
				},
				{
					ParamProps: spec.ParamProps{
						Name:        "user",
						In:          "body",
						Description: "Created user object",
						Required:    true,
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Ref: spec.MustCreateRef("#/definitions/User"),
							},
						},
					},
				},
				{
					ParamProps: spec.ParamProps{
						In:       "body",
						Name:     "body",
						Required: true,
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Properties: spec.SchemaProperties{
									"action": {
										SchemaProps: spec.SchemaProps{
											Enum: []interface{}{"recover"},
											Type: []string{"string"},
										},
										// TODO: need to handle extra props? this header used in common?
										ExtraProps: map[string]interface{}{
											"x-enum-descriptions": "recover meeting recording",
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
	assert.Equal(t, "X", row[12][0])
	assert.Equal(t, "page_size", row[12][1])
	assert.Equal(t, "query", row[12][2])
	assert.Equal(t, "integer", row[12][3])
	assert.Equal(t, "The number of records returned within a single API call", row[12][6])
	assert.Equal(t, "O", row[13][0])
	assert.Equal(t, "User", row[13][1])
	assert.Equal(t, "body", row[13][2])
	assert.Equal(t, "object", row[13][3])
	assert.Equal(t, "Created user object", row[13][6])
	assert.Equal(t, "O", row[14][0])
	assert.Equal(t, "action", row[14][1])
	assert.Equal(t, "body", row[14][2])
	assert.Equal(t, "string", row[14][3])
	assert.Equal(t, "", row[14][6])
	// assert.Equal(t, "recover meeting recording", row[14][6])
}

func TestParameterSchemaWithoutRef(t *testing.T) {
	xl := New()
	var err error
	var row [][]string
	// @source docker.v1.41.json
	// @method POST
	// @path /build
	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Name:        "inputStream",
						In:          "body",
						Required:    true,
						Description: "A tar archive compressed with one of the following algorithms: identity (no compression), gzip, bzip2, xz.",
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Type: []string{
									"string",
									"strings",
								},
								Format: "binary",
							},
						},
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err = xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "O", row[12][0])
	assert.Equal(t, "inputStream", row[12][1])
	assert.Equal(t, "body", row[12][2])
	assert.Equal(t, "string;strings", row[12][3])
	assert.Equal(t, "A tar archive compressed with one of the following algorithms: identity (no compression), gzip, bzip2, xz.", row[12][6])

	// @source zoom.us.json
	// @method POST
	// @path /groups
	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Name:     "body",
						In:       "body",
						Required: true,
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Description: "Group name",
								Type:        []string{"string"},
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
	assert.Equal(t, "O", row[12][0])
	assert.Equal(t, "body", row[12][1])
	assert.Equal(t, "body", row[12][2])
	assert.Equal(t, "string", row[12][3])
	assert.Equal(t, "Group name", row[12][6])
}

// @source editor.swagger.json
// @method POST
// @path /user/createWithList
func TestParameterSchemaItemsWithRef(t *testing.T) {
	xl := New()
	var err error
	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: spec.Definitions{
				"User": {
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{
							"object",
						},
						Properties: spec.SchemaProperties{
							"id": {
								SchemaProps: spec.SchemaProps{
									Type:   []string{"integer"},
									Format: "int64",
								},
							},
							"username": {
								SchemaProps: spec.SchemaProps{
									Type: []string{"string"},
								},
							},
						},
					},
				},
			},
		},
	}
	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Tags:    []string{"user"},
			Summary: "Creates list of users with given input array",
			ID:      "createUsersWithListInput",
			Produces: []string{
				"application/xml",
				"application/json",
			},
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Name:        "users",
						In:          "body",
						Description: "List of user object",
						Required:    true,
					},
					SimpleSchema: spec.SimpleSchema{
						Type: "array",
						Items: &spec.Items{
							Refable: spec.Refable{
								Ref: spec.MustCreateRef("#/definitions/User"),
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
	assert.Equal(t, "O", row[12][0])
	assert.Equal(t, "users", row[12][1])
	assert.Equal(t, "body", row[12][2])
	assert.Equal(t, "array", row[12][3])
	assert.Equal(t, "List of user object", row[12][6])
}

// @source editor.swagger.json
// @method GET
// @path /pet/findByStatus
func TestParameterSchemaItemsWithoutRef(t *testing.T) {
	xl := New()
	var err error
	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Name:        "status",
						In:          "query",
						Description: "Status values that need to be considered for filter",
						Required:    true,
					},
					SimpleSchema: spec.SimpleSchema{
						Type: "array",
						Items: &spec.Items{
							CommonValidations: spec.CommonValidations{
								Enum: []interface{}{
									"available",
									"pending",
									"sold",
								},
							},
							SimpleSchema: spec.SimpleSchema{
								Type:    "string",
								Default: "available",
							},
						},
						CollectionFormat: "multi",
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "O", row[12][0])
	assert.Equal(t, "status", row[12][1])
	assert.Equal(t, "query", row[12][2])
	assert.Equal(t, "array", row[12][3])
	assert.Equal(t, "Status values that need to be considered for filter", row[12][6])
}

// @source docker.v1.41.json
// @method head
// @path /containers/{id}/archive
func TestResponseHeaders(t *testing.T) {
	xl := New()
	var err error
	err = xl.createAPISheet("", "", &spec.Operation{
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
	xl := New()
	var err error
	err = xl.createAPISheet("", "", &spec.Operation{
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
	xl := New()
	var err error
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
				},
			},
		},
	}
	err = xl.createAPISheet("", "", &spec.Operation{
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
								Description: "conflict",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef("#/definitions/ErrorResponse"),
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
	assert.Equal(t, "no error", row[15][6])
	assert.Equal(t, "400", row[16][0])
	assert.Equal(t, "ErrorResponse", row[16][1])
	assert.Equal(t, "body", row[16][2])
	assert.Equal(t, "object", row[16][3])
	assert.Equal(t, "bad parameter", row[16][6])
	assert.Equal(t, "404", row[17][0])
	assert.Equal(t, "ErrorResponse", row[17][1])
	assert.Equal(t, "body", row[17][2])
	assert.Equal(t, "object", row[17][3])
	assert.Equal(t, "bad parameter", row[17][6])
	assert.Equal(t, "409", row[18][0])
	assert.Equal(t, "ErrorResponse", row[18][1])
	assert.Equal(t, "body", row[18][2])
	assert.Equal(t, "object", row[18][3])
	assert.Equal(t, "conflict", row[18][6])
	assert.Equal(t, "500", row[19][0])
	assert.Equal(t, "ErrorResponse", row[19][1])
	assert.Equal(t, "body", row[19][2])
	assert.Equal(t, "object", row[19][3])
	assert.Equal(t, "server error", row[19][6])
}

// @source cisco.meraki.json
// @method POST
// @path /devices/{serial}/camera/generateSnapshot
func TestResponseSchemaWithoutRef(t *testing.T) {
	xl := New()
	var err error
	xl.SwaggerSpec = &spec.Swagger{}
	err = xl.createAPISheet("", "", &spec.Operation{
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
	xl := New()
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
	err = xl.createAPISheet("", "", &spec.Operation{
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
	xl := New()
	var err error
	// @source docker.v1.41.json
	// @method POST
	// @path /containers/prune
	err = xl.createAPISheet("", "", &spec.Operation{
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
	err = xl.createAPISheet("", "", &spec.Operation{
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
	xl := New()
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
	err = xl.createAPISheet("", "", &spec.Operation{
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

// @source zoom.us.json
// @method GET
// @path /accounts
// @path /groups
func TestAllOfDefinitionWithRef(t *testing.T) {
	xl := New()
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
	err = xl.createAPISheet("", "", &spec.Operation{
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
	// TODO:
	// assert.Equal(t, "objects", row[15][3])
	assert.Equal(t, "object", row[15][3])
	assert.Equal(t, "Account list returned", row[15][6])
}
