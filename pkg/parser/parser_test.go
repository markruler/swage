package parser

import (
	"testing"

	"github.com/markruler/swage/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestParseSpecV2(t *testing.T) {
	path := "../../aio/testdata/v2.0.json"
	t.Logf("Swage parse... %s\n", path)
	api, err := Parse(path)

	assert.Equal(t, nil, err, "Error should be nil")
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
			Name:        "pet",
			Description: "Everything about your Pets",
			ExternalDocs: spec.ExternalDocs{
				Description: "Find out more",
				URL:         "http://swagger.io",
			},
		},
		{
			Name:        "store",
			Description: "Access to Petstore orders",
		},
		{
			Name:        "user",
			Description: "Operations about user",
			ExternalDocs: spec.ExternalDocs{
				Description: "Find out more about our store",
				URL:         "http://swagger.io",
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

	post := api.Paths["/_hello/_world/{id}"]
	assert.Equal(t, "world description!", post["post"].Description)
	assert.Equal(t, []string{"*/*"}, post["post"].Consumes)
	assert.Equal(t, []string{"application/json", "text/html"}, post["post"].Produces)
	assert.Equal(t, []string{"world"}, post["post"].Tags)
	assert.Equal(t, "world summary!", post["post"].Summary)

	postParameters := []spec.Parameters{
		{
			Name:        "id",
			In:          "path",
			Description: "ID of pet to use",
			Required:    true,
			Type:        "array",
			Items: spec.Items{
				Type: "string",
			},
			CollectionFormat: "csv",
		},
		{
			Name:        "pet",
			In:          "body",
			Description: "pet description!",
			Required:    false,
			Schema: spec.Schema{
				Ref: "definitions.json#/Pet",
			},
		},
	}
	assert.Equal(t, postParameters, post["post"].Parameters)

	postResponses := map[string]spec.Response{
		"200": {
			Description: "OK",
			Schema: spec.Schema{
				Ref: "#/definitions/Pet",
			},
		},
		"500": {
			Description: "Internal Server Error",
			Schema: spec.Schema{
				Type: "object",
				AdditionalProperties: spec.Property{
					Type:   "integer",
					Format: "int32",
				},
			},
		},
	}
	assert.Equal(t, postResponses, post["post"].Responses)

	postSecurity := []map[string][]string{
		{
			"petstore_auth": {"write:pets", "read:pets"},
		},
	}
	assert.Equal(t, postSecurity, post["post"].Security)

	get := api.Paths["/hello/swage"]
	assert.Equal(t, "swage description!", get["get"].Description)
	assert.Equal(t, []string{"application/vnd.github.v3+json"}, get["get"].Consumes)
	assert.Equal(t, []string{"application/json"}, get["get"].Produces)
	assert.Equal(t, []string{"swage"}, get["get"].Tags)
	assert.Equal(t, "swage summary!", get["get"].Summary)

	getParameters := []spec.Parameters(nil)
	assert.Equal(t, getParameters, get["get"].Parameters)

	getResponses := map[string]spec.Response{
		"200": {
			Description: "OK",
			Schema: spec.Schema{
				Type: "array",
				Items: spec.Items{
					Ref: "#/definitions/ApiResponse",
				},
			},
		},
		"500": {
			Description: "Internal Server Error",
			Schema: spec.Schema{
				Type: "object",
				AdditionalProperties: spec.Property{
					Type:   "integer",
					Format: "int32",
				},
			},
		},
	}
	assert.Equal(t, getResponses, get["get"].Responses)

	getSecurity := []map[string][]string{
		{
			"api_key": {},
		},
	}
	assert.Equal(t, getSecurity, get["get"].Security)

	categoryDefinition := spec.Definition{
		Type: "object",
		Properties: map[string]spec.Property{
			"id": {
				Type:   "integer",
				Format: "int64",
			},
			"name": {
				Type: "string",
			},
		},
		XML: spec.XML{
			Name: "Category",
		},
	}
	assert.Equal(t, categoryDefinition, api.Definitions["Category"])

	petDefinition := spec.Definition{
		Type: "object",
		Required: []string{"name", "photoUrls"},
		Properties: map[string]spec.Property{
			"id": {
				Type: "integer",
				Format: "int64",
			},
			"category": {
				Ref: "#/definitions/Category",
			},
			"name": {
				Type: "string",
				Example: "doggie",
			},
			"photoUrls": {
				Type: "array",
				XML: spec.XML{
					Name: "photoUrl",
					Wrapped: true,
				},
				Items: spec.Items{
					Type: "string",
				},
			},
			"age": {
				Type: "integer",
				Format: "int32",
				Minimum: 1,
			},
			"tags": {
				Type: "array",
				XML: spec.XML{
					Name: "tag",
					Wrapped: true,
				},
				Items: spec.Items{
					Ref: "#/definitions/Tag",
				},
			},
			"status": {
				Type: "string",
				Description: "pet status in the store",
				Enum: []string{"available", "pending", "sold"},
			},
		},
		XML: spec.XML{
			Name: "Pet",
		},
	}
	assert.Equal(t, petDefinition, api.Definitions["Pet"])
	
	apiResponseDefinition := spec.Definition{
		Type: "object",
		Properties: map[string]spec.Property{
			"code": {
				Type: "integer",
				Format: "int32",
				Enum: []string{"00", "11", "22"},
			},
			"type": {
				Type: "string",
				Example: "test type",
			},
			"message": {
				Type: "string",
				Example: "test-msg",
			},
		},
	}
	assert.Equal(t, apiResponseDefinition, api.Definitions["ApiResponse"])

	apiKey := spec.SecurityDefinition{
		Type: "apiKey",
		Name: "api_key",
		In: "header",
	}
	assert.Equal(t, apiKey, api.SecurityDefinitions["api_key"])

	petstoreAuth := spec.SecurityDefinition{
		Type: "oauth2",
		AuthorizationURL: "http://petstore.swagger.io/oauth/dialog",
		Flow: "implicit",
		Scopes: map[string]string{
			"write:pets": "modify pets in your account",
			"read:pets": "read your pets",
		},
	}
	assert.Equal(t, petstoreAuth, api.SecurityDefinitions["petstore_auth"])

	externalDocs := spec.ExternalDocs{
		Description: "Find out more about Swagger",
		URL: "http://swagger.io",
	}
	assert.Equal(t, externalDocs, api.ExternalDocs)
}

// func TestParseSpecV3(t *testing.T) {
// 	path := "../../aio/testdata/v3.0.json"
// }
