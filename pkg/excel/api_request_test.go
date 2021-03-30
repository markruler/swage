package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

// @source docker.v1.41.json
// @method GET
// @path /containers/json
func TestParameterWithoutSchema(t *testing.T) {
	xl := New()
	err := xl.createAPISheet("", "", &spec.Operation{
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
			Responses: &spec.Responses{},
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
	err := xl.createAPISheet("", "", &spec.Operation{
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
						Name:        "limit",
						In:          "query",
						Description: "Return this number of most recently created containers, including\nnon-running ones.\n",
					},
					SimpleSchema: spec.SimpleSchema{
						Type: "integer",
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
			Responses: &spec.Responses{},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)

	// t.Log(row)
	assert.Equal(t, "X", row[12][0])
	assert.Equal(t, "page_size", row[12][1])
	assert.Equal(t, "query", row[12][2])
	assert.Equal(t, "integer", row[12][3])
	assert.Equal(t, "The number of records returned within a single API call", row[12][6])

	assert.Equal(t, "X", row[13][0])
	assert.Equal(t, "limit", row[13][1])
	assert.Equal(t, "query", row[13][2])
	assert.Equal(t, "integer", row[13][3])
	assert.Equal(t, "Return this number of most recently created containers, including\nnon-running ones.\n", row[13][6])

	assert.Equal(t, "O", row[14][0])
	assert.Equal(t, "User", row[14][1])
	assert.Equal(t, "body", row[14][2])
	assert.Equal(t, "object", row[14][3])
	assert.Equal(t, "Created user object", row[14][6])

	assert.Equal(t, "O", row[15][0])
	assert.Equal(t, "action", row[15][1])
	assert.Equal(t, "body", row[15][2])
	assert.Equal(t, "string", row[15][3])
	assert.Equal(t, "", row[15][6])

	// @source docker.v1.41.json
	// @method POST
	// @path /containers/create
	// TODO: AllOf

	// TODO: ExtraProps
	// assert.Equal(t, "recover meeting recording", row[15][6])
}

func TestParameterSchemaWithoutRef(t *testing.T) {
	xl := New()
	var row [][]string
	// @source docker.v1.41.json
	// @method POST
	// @path /build
	err := xl.createAPISheet("", "", &spec.Operation{
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
			Responses: &spec.Responses{},
		},
	}, nil, 1)
	assert.NoError(t, err)
	row, err = xl.File.GetRows("1")
	assert.NoError(t, err)
	assert.Equal(t, "O", row[12][0])
	assert.Equal(t, "inputStream", row[12][1])
	assert.Equal(t, "body", row[12][2])
	assert.Equal(t, "string,strings", row[12][3])
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
			Responses: &spec.Responses{},
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
	err := xl.createAPISheet("", "", &spec.Operation{
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
			Responses: &spec.Responses{},
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
	err := xl.createAPISheet("", "", &spec.Operation{
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
			Responses: &spec.Responses{},
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
