package converter

import (
	"testing"

	oas "github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestExtractRequests_SimpleSchema(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						Name:        "all",
						In:          "query",
						Description: "Return all containers. By default, only running containers are shown.\n",
						Required:    true,
					},
					SimpleSchema: oas.SimpleSchema{
						Type:    "boolean",
						Default: false,
					},
				},
			},
			Responses: &oas.Responses{},
		},
	}

	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Paths: &oas.Paths{
				Paths: map[string]oas.PathItem{
					"/": {
						PathItemProps: oas.PathItemProps{
							Get: test_operation,
						},
					},
				},
			},
		},
	}

	api_request, err := extractRequests(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "O", api_request[0].Required)
	assert.Equal(t, "all", api_request[0].Schema)
	assert.Equal(t, "query", api_request[0].ParameterType)
	assert.Equal(t, "boolean", api_request[0].DataType)
	assert.Equal(t, "", api_request[0].Enum)
	assert.Equal(t, "", api_request[0].Example)
	assert.Equal(t, "Return all containers. By default, only running containers are shown.\n", api_request[0].Description)
}

// @source zoom.us.json
// @method PUT
// @path /meetings/{meetingId}/recordings/{recordingId}/status
func TestExtractRequests_SchemaRef(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						Schema: &oas.Schema{
							SchemaProps: oas.SchemaProps{
								Ref: oas.MustCreateRef("#/parameters/PageSize"),
							},
						},
					},
				},
			},
			Responses: &oas.Responses{},
		},
	}
	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Parameters: map[string]oas.Parameter{
				"PageSize": {
					ParamProps: oas.ParamProps{
						Description: "The number of records returned within a single API call",
						In:          "query",
						Name:        "page_size",
					},
					SimpleSchema: oas.SimpleSchema{
						Type:    "integer",
						Default: 30,
					},
					CommonValidations: oas.CommonValidations{
						Maximum: &[]float64{300}[0],
					},
				},
			},
			Paths: &oas.Paths{
				Paths: map[string]oas.PathItem{
					"/": {
						PathItemProps: oas.PathItemProps{
							Get: test_operation,
						},
					},
				},
			},
		},
	}
	api_request, err := extractRequests(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "X", api_request[0].Required)
	assert.Equal(t, "page_size", api_request[0].Schema)
	assert.Equal(t, "query", api_request[0].ParameterType)
	assert.Equal(t, "integer", api_request[0].DataType)
	assert.Equal(t, "", api_request[0].Enum)
	assert.Equal(t, "", api_request[0].Example)
	assert.Equal(t, "The number of records returned within a single API call", api_request[0].Description)
}

// @source docker.v1.41.json
// @method GET
// @path /containers/json
func TestExtractRequests_ExtraProps(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						In:       "body",
						Name:     "body",
						Required: true,
						Schema: &oas.Schema{
							SchemaProps: oas.SchemaProps{
								Properties: oas.SchemaProperties{
									"action": {
										SchemaProps: oas.SchemaProps{
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
			Responses: &oas.Responses{},
		},
	}
	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Paths: &oas.Paths{
				Paths: map[string]oas.PathItem{
					"/": {
						PathItemProps: oas.PathItemProps{
							Get: test_operation,
						},
					},
				},
			},
		},
	}
	api_request, err := extractRequests(oas_spec, test_operation)
	assert.NoError(t, err)

	// TODO: ExtraProps
	assert.Equal(t, "O", api_request[0].Required)
	// assert.Equal(t, "action", api_request[0].Schema)
	assert.Equal(t, "body", api_request[0].ParameterType)
	// assert.Equal(t, "string", api_request[0].DataType)
	// assert.Equal(t, "recover", api_request[0].Enum)
	assert.Equal(t, "", api_request[0].Example)
	// assert.Equal(t, "recover meeting recording", api_request[0].Description)
}

// @source docker.v1.41.json
// @method POST
// @path /build
func TestExtractRequests_MultipleSchemaType(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						Name:        "inputStream",
						In:          "body",
						Required:    true,
						Description: "A tar archive compressed with one of the following algorithms: identity (no compression), gzip, bzip2, xz.",
						Schema: &oas.Schema{
							SchemaProps: oas.SchemaProps{
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
			Responses: &oas.Responses{},
		},
	}

	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Paths: &oas.Paths{
				Paths: map[string]oas.PathItem{
					"/": {
						PathItemProps: oas.PathItemProps{
							Get: test_operation,
						},
					},
				},
			},
		},
	}
	api_request, err := extractRequests(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "O", api_request[0].Required)
	assert.Equal(t, "inputStream", api_request[0].Schema)
	assert.Equal(t, "body", api_request[0].ParameterType)
	assert.Equal(t, "string,strings", api_request[0].DataType)
	assert.Equal(t, "", api_request[0].Enum)
	assert.Equal(t, "", api_request[0].Example)
	assert.Equal(t, "A tar archive compressed with one of the following algorithms: identity (no compression), gzip, bzip2, xz.", api_request[0].Description)
}

// @source editor.swagger.json
// @method POST
// @path /user/createWithList
func TestParameterSchema_RefItems_WithoutDefinitions(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Tags:    []string{"user"},
			Summary: "Creates list of users with given input array",
			ID:      "createUsersWithListInput",
			Produces: []string{
				"application/xml",
				"application/json",
			},
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						Name:        "users",
						In:          "body",
						Description: "List of user object",
						Required:    true,
					},
					SimpleSchema: oas.SimpleSchema{
						Type: "array",
						Items: &oas.Items{
							Refable: oas.Refable{
								Ref: oas.MustCreateRef("#/definitions/User"),
							},
						},
					},
				},
			},
			Responses: &oas.Responses{},
		},
	}
	api_request, err := extractRequests(&oas.Swagger{}, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "O", api_request[0].Required)
	assert.Equal(t, "users", api_request[0].Schema)
	assert.Equal(t, "body", api_request[0].ParameterType)
	assert.Equal(t, "array", api_request[0].DataType)
	assert.Equal(t, "", api_request[0].Enum)
	assert.Equal(t, "", api_request[0].Example)
	assert.Equal(t, "List of user object", api_request[0].Description)
}

// @source editor.swagger.json
// @method GET
// @path /pet/findByStatus
func TestParameterSchema_ItemsWithEnums(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						Name:        "status",
						In:          "query",
						Description: "Status values that need to be considered for filter",
						Required:    true,
					},
					SimpleSchema: oas.SimpleSchema{
						Type: "array",
						Items: &oas.Items{
							CommonValidations: oas.CommonValidations{
								Enum: []interface{}{
									"available",
									"pending",
									"sold",
								},
							},
							SimpleSchema: oas.SimpleSchema{
								Type:    "string",
								Default: "available",
							},
						},
						CollectionFormat: "multi",
					},
				},
			},
			Responses: &oas.Responses{},
		},
	}

	api_request, err := extractRequests(&oas.Swagger{}, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "O", api_request[0].Required)
	assert.Equal(t, "status", api_request[0].Schema)
	assert.Equal(t, "query", api_request[0].ParameterType)
	assert.Equal(t, "array", api_request[0].DataType)
	assert.Equal(t, "available,pending,sold", api_request[0].Enum)
	assert.Equal(t, "", api_request[0].Example)
	assert.Equal(t, "Status values that need to be considered for filter", api_request[0].Description)
}

// docker.v1.41.json
func TestRequestDefinition_Example(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Tags:        []string{"Container"},
			Summary:     "List containers",
			Description: "Returns a list of containers. For details on the format, see the\n[inspect endpoint](#operation/ContainerInspect).\n\nNote that it uses a different, smaller representation of a container\nthan inspecting a single container. For example, the list of linked\ncontainers is not propagated .\n",
			ID:          "ContainerList",
			Produces:    []string{"application/json"},
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						Name:        "port",
						In:          "body",
						Description: "Container port",
						Required:    true,
						Schema: &oas.Schema{
							SchemaProps: oas.SchemaProps{
								Ref: oas.MustCreateRef("#/definitions/Port"),
							},
						},
					},
				},
			},
			Responses: &oas.Responses{},
		},
	}

	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Definitions: oas.Definitions{
				"Port": {
					SchemaProps: oas.SchemaProps{
						Type:        oas.StringOrArray{"object"},
						Description: "An open port on a container",
						Required:    []string{"PrivatePort", "Type"},
						Properties: oas.SchemaProperties{
							"IP": {
								SchemaProps: oas.SchemaProps{
									Type:        []string{"string"},
									Format:      "ip-address",
									Description: "Host IP address that the container's port is mapped to",
								},
							},
							"PrivatePort": {
								SchemaProps: oas.SchemaProps{
									Type:        []string{"integer"},
									Format:      "uint16",
									Description: "Port on the container",
								},
							},
							"PublicPort": {
								SchemaProps: oas.SchemaProps{
									Type:        oas.StringOrArray{"integer"},
									Format:      "uint16",
									Description: "Port exposed on the host",
								},
							},
							"Type": {
								SchemaProps: oas.SchemaProps{
									Type: oas.StringOrArray{"string"},
									Enum: []interface{}{"tcp", "udp", "sctp"},
								},
							},
						},
					},
					SwaggerSchemaProps: oas.SwaggerSchemaProps{
						Example: map[string]interface{}{
							"PrivatePort": 8080,
							"PublicPort":  80,
							"Type":        "tcp",
						},
					},
				},
			},
		},
	}

	api_request, err := extractRequests(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "O", api_request[0].Required)
	assert.Equal(t, "Port", api_request[0].Schema)
	assert.Equal(t, "body", api_request[0].ParameterType)
	assert.Equal(t, "object", api_request[0].DataType)
	assert.Equal(t, "", api_request[0].Enum)
	assert.Equal(t, "{\n    \"PrivatePort\": 8080,\n    \"PublicPort\": 80,\n    \"Type\": \"tcp\"\n}", api_request[0].Example)
	assert.Equal(t, "An open port on a container", api_request[0].Description)
}
