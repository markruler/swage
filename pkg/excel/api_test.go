package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPISheet(t *testing.T) {
	xl := New("")
	var err error

	err = xl.createAPISheet("", "", &spec.Operation{}, nil, 1)
	assert.Error(t, err)

	err = xl.createAPISheet("", "", nil, nil, 1)
	assert.Error(t, err)

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
	assert.Error(t, err)

	p := parser.New("../../aio/testdata/json/dev.json")
	xl.SwaggerSpec, _ = p.Parse()
	xl.createAPISheet("", "", nil, nil, 1)
}

func TestResponseSchemaItems(t *testing.T) {
	xl := New("")
	var err error
	// spotify.json - [get] /albums
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

func TestRootSchemaFromRef(t *testing.T) {
	// zoom.us.json - [get] /accounts/{accountId}/billing
	xl := New("")
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
										// Ref: spec.MustCreateRef("#/definitions/AccountPlans"),
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

func TestRootDefinitionAllOfFromRef(t *testing.T) {
	// zoom.us.json - [get] /meetings/{meetingId}/registrants
	// zoom.us.json - [get] /accounts
	// zoom.us.json - [get] /groups
	xl := New("")
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
