package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/pkg/parser"
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

// @source editor.swagger.json
// @method POST
// @path /user
func TestParameterSchemaWithRef(t *testing.T) {
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
			Parameters: []spec.Parameter{
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
			},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "O", row[8][0])
	assert.Equal(t, "User", row[8][1])
	assert.Equal(t, "body", row[8][2])
	assert.Equal(t, "object", row[8][3])
	assert.Equal(t, "Created user object", row[8][6])
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
	assert.Equal(t, "O", row[8][0])
	assert.Equal(t, "users", row[8][1])
	assert.Equal(t, "body", row[8][2])
	assert.Equal(t, "array", row[8][3])
	assert.Equal(t, "List of user object", row[8][6])
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
	assert.Equal(t, "O", row[8][0])
	assert.Equal(t, "status", row[8][1])
	assert.Equal(t, "query", row[8][2])
	assert.Equal(t, "array", row[8][3])
	assert.Equal(t, "Status values that need to be considered for filter", row[8][6])
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
	assert.Equal(t, "", row[11][0])
	assert.Equal(t, "Test", row[11][1])
	assert.Equal(t, "body", row[11][2])
	assert.Equal(t, "array", row[11][3])
	assert.Equal(t, "OK", row[11][6])

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
											"integer",
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
}

// @source zoom.us.json
// @method GET
// @path /accounts/{accountId}/billing
func TestSchemaWithRef(t *testing.T) {
	xl := New()
	var err error
	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Parameters: map[string]spec.Parameter{
				"AccountId": {
					ParamProps: spec.ParamProps{
						Description: "The account ID",
						In:          "path",
						Name:        "accountId",
						Required:    true,
					},
					SimpleSchema: spec.SimpleSchema{
						Type: "string",
					},
				},
			},
			Definitions: spec.Definitions{
				"AccountPlan": spec.Schema{
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
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Ref: spec.MustCreateRef("#/parameters/AccountId"),
							},
						},
					},
				},
			},
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "OK",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef("#/definitions/AccountPlan"),
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
	assert.Equal(t, "O", row[8][0])
	assert.Equal(t, "accountId", row[8][1])
	assert.Equal(t, "path", row[8][2])
	assert.Equal(t, "string", row[8][3])
	assert.Equal(t, "The account ID", row[8][6])
}

// @source cisco.meraki.json
// @method POST
// @path /devices/{serial}/camera/generateSnapshot
func TestSchemaWithoutRef(t *testing.T) {
	xl := New()
	var err error
	xl.SwaggerSpec = &spec.Swagger{}
	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						In:       "path",
						Name:     "serial",
						Required: true,
					},
					SimpleSchema: spec.SimpleSchema{
						Type: "string",
					},
				},
			},
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
	assert.Equal(t, "O", row[8][0])
	assert.Equal(t, "serial", row[8][1])
	assert.Equal(t, "path", row[8][2])
	assert.Equal(t, "string", row[8][3])
	assert.Equal(t, "", row[8][6])
}

// @source zoom.us.json
// @method GET
// @path /meetings/{meetingId}/registrants
// @path /accounts
// @path /groups
func TestDefinitionAllOfWithRef(t *testing.T) {
	xl := New()
	var err error
	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
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
				"PageNumber": {
					ParamProps: spec.ParamProps{
						Description: "Current page number of returned records",
						In:          "query",
						Name:        "page_number",
					},
					SimpleSchema: spec.SimpleSchema{
						Type:    "integer",
						Default: 1,
					},
				},
			},
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
			},
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
	assert.Equal(t, "X", row[8][0])
	assert.Equal(t, "page_size", row[8][1])
	assert.Equal(t, "query", row[8][2])
	assert.Equal(t, "integer", row[8][3])
	assert.Equal(t, "The number of records returned within a single API call", row[8][6])
	assert.Equal(t, "", row[11][0])
	assert.Equal(t, "AccountList", row[11][1])
	assert.Equal(t, "body", row[11][2])
	// TODO:
	// assert.Equal(t, "objects", row[11][3])
	assert.Equal(t, "object", row[11][3])
	assert.Equal(t, "Account list returned", row[11][6])
}
