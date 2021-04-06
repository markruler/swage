package parser

import (
	"testing"

	oas "github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestParseSpecV2(t *testing.T) {
	var err error

	_, err = Parse("../testdata/json/fake.js")
	assert.Error(t, err)

	_, err = Parse("../testdata/json/fake.json")
	assert.Error(t, err)

	api, err := Parse("../testdata/json/sample.pet.json")
	assert.NoError(t, err)

	assert.Equal(t, "Swagger Sample App", api.Info.InfoProps.Title)
	assert.Equal(t, "2.0", api.Swagger)
	assert.Equal(t, "Swagger Sample App", api.Info.Title)
	assert.Equal(t, "https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md", api.Info.Description)
	assert.Equal(t, "http://swagger.io/terms/", api.Info.TermsOfService)
	assert.Equal(t, "API Support", api.Info.Contact.Name)
	assert.Equal(t, "http://www.swagger.io/support", api.Info.Contact.URL)
	assert.Equal(t, "support@swagger.io", api.Info.Contact.Email)
	assert.Equal(t, "http://www.apache.org/licenses/LICENSE-2.0.html", api.Info.License.URL)
	assert.Equal(t, "1.0.1", api.Info.Version)
	assert.Equal(t, "/api/v1", api.BasePath)
	assert.Equal(t, []string{"http", "https", "ws", "wss"}, api.Schemes)

	tag := []oas.Tag{
		{
			TagProps: oas.TagProps{
				Name:        "pet",
				Description: "Everything about your Pets",
				ExternalDocs: &oas.ExternalDocumentation{
					Description: "Find out more",
					URL:         "http://swagger.io",
				},
			},
		},
		{
			TagProps: oas.TagProps{
				Name:        "store",
				Description: "Access to Petstore orders",
			},
		},
		{
			TagProps: oas.TagProps{
				Name:        "user",
				Description: "Operations about user",
				ExternalDocs: &oas.ExternalDocumentation{
					Description: "Find out more about our store",
					URL:         "http://swagger.io",
				},
			},
		},
	}
	assert.Equal(t, tag, api.Tags)

	consume := []string{
		"text/plain; charset=utf-8",
		"application/json",
		"application/vnd.github+json",
		"application/vnd.github.v3+json",
	}
	assert.Equal(t, consume, api.Consumes)

	produce := []string{
		"text/plain",
		"application/json",
		"application/vnd.github+json",
		"application/vnd.github.v3+json",
	}
	assert.Equal(t, produce, api.Produces)

	post := api.Paths.Paths["/_hello/_world/{id}"]
	assert.Equal(t, "world description!", post.Post.Description)
	assert.Equal(t, []string{"*/*"}, post.Post.Consumes)
	assert.Equal(t, []string{"application/json", "text/html"}, post.Post.Produces)
	assert.Equal(t, []string{"world"}, post.Post.Tags)
	assert.Equal(t, "world summary!", post.Post.Summary)

	postParameters := []oas.Parameter{
		{
			SimpleSchema: oas.SimpleSchema{
				Type:             "array",
				Items:            oas.NewItems().Typed("string", ""),
				CollectionFormat: "csv",
			},
			ParamProps: oas.ParamProps{
				Name:        "id",
				In:          "path",
				Description: "ID of pet to use",
				Required:    true,
			},
		},
		{
			ParamProps: oas.ParamProps{
				Name:        "pet",
				In:          "body",
				Description: "pet description!",
				Required:    false,
				Schema: &oas.Schema{
					SchemaProps: oas.SchemaProps{
						Ref: oas.MustCreateRef("definitions.json#/Pet"),
					},
				},
			},
		},
	}
	assert.Equal(t, postParameters, post.Post.Parameters)

	postResponses := &oas.Responses{
		ResponsesProps: oas.ResponsesProps{
			StatusCodeResponses: map[int]oas.Response{
				200: {
					ResponseProps: oas.ResponseProps{
						Description: "OK",
						Schema: &oas.Schema{
							SchemaProps: oas.SchemaProps{
								Ref: oas.MustCreateRef("#/definitions/Pet"),
							},
						},
					},
				},
				500: {
					ResponseProps: oas.ResponseProps{
						Description: "Internal Server Error",
						Schema: &oas.Schema{
							SchemaProps: oas.SchemaProps{
								Type: oas.StringOrArray{
									"object",
								},
								AdditionalProperties: &oas.SchemaOrBool{
									Allows: true,
									Schema: &oas.Schema{
										SchemaProps: oas.SchemaProps{
											Type: oas.StringOrArray{
												"integer",
											},
											Format: "int32",
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
	assert.Equal(t, postResponses, post.Post.Responses)

	postSecurity := []map[string][]string{
		{
			"petstore_auth": {"write:pets", "read:pets"},
		},
	}
	assert.Equal(t, postSecurity, post.Post.Security)

	get := api.Paths.Paths["/hello/swage"]
	assert.Equal(t, "swage description!", get.Get.Description)
	assert.Equal(t, []string{"application/vnd.github.v3+json"}, get.Get.Consumes)
	assert.Equal(t, []string{"application/json"}, get.Get.Produces)
	assert.Equal(t, []string{"swage"}, get.Get.Tags)
	assert.Equal(t, "swage summary!", get.Get.Summary)

	getParameters := []oas.Parameter(nil)
	assert.Equal(t, getParameters, get.Get.Parameters)

	getResponses := &oas.Responses{
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
									Schema: &oas.Schema{
										SchemaProps: oas.SchemaProps{
											Ref: oas.MustCreateRef("#/definitions/ApiResponse"),
										},
									},
								},
							},
						},
					},
				},
				500: {
					ResponseProps: oas.ResponseProps{
						Description: "Internal Server Error",
						Schema: &oas.Schema{
							SchemaProps: oas.SchemaProps{
								Type: oas.StringOrArray{
									"object",
								},
								AdditionalProperties: &oas.SchemaOrBool{
									Allows: true,
									Schema: &oas.Schema{
										SchemaProps: oas.SchemaProps{
											Type: oas.StringOrArray{
												"integer",
											},
											Format: "int32",
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
	assert.Equal(t, getResponses, get.Get.Responses)

	getSecurity := []map[string][]string{
		{
			"api_key": {},
		},
	}
	assert.Equal(t, getSecurity, get.Get.Security)

	allDefinition := oas.Definitions{
		"Category": oas.Schema{
			SchemaProps: oas.SchemaProps{
				Type: oas.StringOrArray{
					"object",
				},
				Properties: oas.SchemaProperties{
					"id": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"integer",
							},
							Format: "int64",
						},
					},
					"name": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"string",
							},
						},
					},
				},
			},
			SwaggerSchemaProps: oas.SwaggerSchemaProps{
				XML: &oas.XMLObject{
					Name: "Category",
				},
			},
		},
		"Pet": oas.Schema{
			SchemaProps: oas.SchemaProps{
				Type: oas.StringOrArray{
					"object",
				},
				Required: []string{"name", "photoUrls"},
				Properties: oas.SchemaProperties{
					"id": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"integer",
							},
							Format: "int64",
						},
					},
					"category": {
						SchemaProps: oas.SchemaProps{
							Ref: oas.MustCreateRef("#/definitions/Category"),
						},
					},
					"name": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"string",
							},
						},
						SwaggerSchemaProps: oas.SwaggerSchemaProps{
							Example: "doggie",
						},
					},
					"photoUrls": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"array",
							},
							Items: &oas.SchemaOrArray{
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Type: oas.StringOrArray{
											"string",
										},
									},
								},
							},
						},
						SwaggerSchemaProps: oas.SwaggerSchemaProps{
							XML: &oas.XMLObject{
								Name:    "photoUrls",
								Wrapped: true,
							},
						},
					},
					"age": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"integer",
							},
							Format:  "int32",
							Minimum: &[]float64{1}[0],
						},
					},
					"tags": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"array",
							},
							Items: &oas.SchemaOrArray{
								Schema: &oas.Schema{
									SchemaProps: oas.SchemaProps{
										Ref: oas.MustCreateRef("#/definitions/Tag"),
									},
								},
							},
						},
						SwaggerSchemaProps: oas.SwaggerSchemaProps{
							XML: &oas.XMLObject{
								Name:    "tags",
								Wrapped: true,
							},
						},
					},
					"status": {
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"string",
							},
							Description: "pet status in the store",
							Enum: []interface{}{
								"available",
								"pending",
								"sold",
							},
						},
					},
				},
			},
			SwaggerSchemaProps: oas.SwaggerSchemaProps{
				XML: &oas.XMLObject{
					Name: "Pet",
				},
			},
		},
		"ApiResponse": oas.Schema{
			SchemaProps: oas.SchemaProps{
				Type: oas.StringOrArray{
					"object",
				},
				Properties: oas.SchemaProperties{
					"code": oas.Schema{
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"integer",
							},
							Format: "int32",
							Enum: []interface{}{
								"00",
								"11",
								"22",
							},
						},
					},
					"type": oas.Schema{
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"string",
							},
						},
						SwaggerSchemaProps: oas.SwaggerSchemaProps{
							Example: "test type",
						},
					},
					"message": oas.Schema{
						SchemaProps: oas.SchemaProps{
							Type: oas.StringOrArray{
								"string",
							},
						},
						SwaggerSchemaProps: oas.SwaggerSchemaProps{
							Example: "test-msg",
						},
					},
				},
			},
		},
	}

	assert.Equal(t, allDefinition, api.Definitions)

	apiKey := &oas.SecurityScheme{
		SecuritySchemeProps: oas.SecuritySchemeProps{
			Type: "apiKey",
			Name: "api_key",
			In:   "header",
		},
	}
	assert.Equal(t, apiKey, api.SecurityDefinitions["api_key"])

	petstoreAuth := &oas.SecurityScheme{
		SecuritySchemeProps: oas.SecuritySchemeProps{
			Type:             "oauth2",
			AuthorizationURL: "http://petstore.swagger.io/oauth/dialog",
			Flow:             "implicit",
			Scopes: map[string]string{
				"write:pets": "modify pets in your account",
				"read:pets":  "read your pets",
			},
		},
	}
	assert.Equal(t, petstoreAuth, api.SecurityDefinitions["petstore_auth"])

	externalDocs := &oas.ExternalDocumentation{
		Description: "Find out more about Swagger",
		URL:         "http://swagger.io",
	}
	assert.Equal(t, externalDocs, api.ExternalDocs)
}

func TestExtractOperation(t *testing.T) {
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

	oas_oas := &oas.Swagger{
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

	api, err := extractOperation(oas_oas, "/", "GET", test_operation)
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
