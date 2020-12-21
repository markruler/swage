package parser

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestParseSpecV2(t *testing.T) {
	var err error
	var parser Parser

	fakePath := "../../aio/testdata/json/fake.js"
	parser = *New(fakePath)
	api, err := parser.Parse()
	assert.Error(t, err)

	fakeJSON := "../../aio/testdata/json/fake.json"
	parser = *New(fakeJSON)
	api, err = parser.Parse()
	assert.Error(t, err)

	realJSONPath := "../../aio/testdata/json/dev.json"
	parser = *New(realJSONPath)
	api, err = parser.Parse()
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

	tag := []spec.Tag{
		{
			TagProps: spec.TagProps{
				Name:        "pet",
				Description: "Everything about your Pets",
				ExternalDocs: &spec.ExternalDocumentation{
					Description: "Find out more",
					URL:         "http://swagger.io",
				},
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "store",
				Description: "Access to Petstore orders",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "user",
				Description: "Operations about user",
				ExternalDocs: &spec.ExternalDocumentation{
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

	postParameters := []spec.Parameter{
		{
			SimpleSchema: spec.SimpleSchema{
				Type:             "array",
				Items:            spec.NewItems().Typed("string", ""),
				CollectionFormat: "csv",
			},
			ParamProps: spec.ParamProps{
				Name:        "id",
				In:          "path",
				Description: "ID of pet to use",
				Required:    true,
			},
		},
		{
			ParamProps: spec.ParamProps{
				Name:        "pet",
				In:          "body",
				Description: "pet description!",
				Required:    false,
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Ref: spec.MustCreateRef("definitions.json#/Pet"),
					},
				},
			},
		},
	}
	assert.Equal(t, postParameters, post.Post.Parameters)

	postResponses := &spec.Responses{
		ResponsesProps: spec.ResponsesProps{
			StatusCodeResponses: map[int]spec.Response{
				200: {
					ResponseProps: spec.ResponseProps{
						Description: "OK",
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Ref: spec.MustCreateRef("#/definitions/Pet"),
							},
						},
					},
				},
				500: {
					ResponseProps: spec.ResponseProps{
						Description: "Internal Server Error",
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Type: spec.StringOrArray{
									"object",
								},
								AdditionalProperties: &spec.SchemaOrBool{
									Allows: true,
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type: spec.StringOrArray{
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

	getParameters := []spec.Parameter(nil)
	assert.Equal(t, getParameters, get.Get.Parameters)

	getResponses := &spec.Responses{
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
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: spec.MustCreateRef("#/definitions/ApiResponse"),
										},
									},
								},
							},
						},
					},
				},
				500: {
					ResponseProps: spec.ResponseProps{
						Description: "Internal Server Error",
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Type: spec.StringOrArray{
									"object",
								},
								AdditionalProperties: &spec.SchemaOrBool{
									Allows: true,
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type: spec.StringOrArray{
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

	allDefinition := spec.Definitions{
		"Category": spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: spec.StringOrArray{
					"object",
				},
				Properties: spec.SchemaProperties{
					"id": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"integer",
							},
							Format: "int64",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"string",
							},
						},
					},
				},
			},
			SwaggerSchemaProps: spec.SwaggerSchemaProps{
				XML: &spec.XMLObject{
					Name: "Category",
				},
			},
		},
		"Pet": spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: spec.StringOrArray{
					"object",
				},
				Required: []string{"name", "photoUrls"},
				Properties: spec.SchemaProperties{
					"id": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"integer",
							},
							Format: "int64",
						},
					},
					"category": {
						SchemaProps: spec.SchemaProps{
							Ref: spec.MustCreateRef("#/definitions/Category"),
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"string",
							},
						},
						SwaggerSchemaProps: spec.SwaggerSchemaProps{
							Example: "doggie",
						},
					},
					"photoUrls": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"array",
							},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: spec.StringOrArray{
											"string",
										},
									},
								},
							},
						},
						SwaggerSchemaProps: spec.SwaggerSchemaProps{
							XML: &spec.XMLObject{
								Name:    "photoUrls",
								Wrapped: true,
							},
						},
					},
					"age": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"integer",
							},
							Format:  "int32",
							Minimum: &[]float64{1}[0],
						},
					},
					"tags": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"array",
							},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef("#/definitions/Tag"),
									},
								},
							},
						},
						SwaggerSchemaProps: spec.SwaggerSchemaProps{
							XML: &spec.XMLObject{
								Name:    "tags",
								Wrapped: true,
							},
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
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
			SwaggerSchemaProps: spec.SwaggerSchemaProps{
				XML: &spec.XMLObject{
					Name: "Pet",
				},
			},
		},
		"ApiResponse": spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: spec.StringOrArray{
					"object",
				},
				Properties: spec.SchemaProperties{
					"code": spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
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
					"type": spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"string",
							},
						},
						SwaggerSchemaProps: spec.SwaggerSchemaProps{
							Example: "test type",
						},
					},
					"message": spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"string",
							},
						},
						SwaggerSchemaProps: spec.SwaggerSchemaProps{
							Example: "test-msg",
						},
					},
				},
			},
		},
	}

	assert.Equal(t, allDefinition, api.Definitions)

	apiKey := &spec.SecurityScheme{
		SecuritySchemeProps: spec.SecuritySchemeProps{
			Type: "apiKey",
			Name: "api_key",
			In:   "header",
		},
	}
	assert.Equal(t, apiKey, api.SecurityDefinitions["api_key"])

	petstoreAuth := &spec.SecurityScheme{
		SecuritySchemeProps: spec.SecuritySchemeProps{
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

	externalDocs := &spec.ExternalDocumentation{
		Description: "Find out more about Swagger",
		URL:         "http://swagger.io",
	}
	assert.Equal(t, externalDocs, api.ExternalDocs)
}
