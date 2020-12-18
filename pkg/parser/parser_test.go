package parser

import (
	"reflect"
	"testing"

	"github.com/markruler/swage/pkg/spec"
)

func TestParseSwaggerApiV2(t *testing.T) {
	path := "../../aio/testdata/v2.0.json"
	api, err := Parse(path)

	if err != nil {
		t.FailNow()
	}
	if api == nil {
		t.FailNow()
	}
	if api.Swagger != "2.0" {
		t.FailNow()
	}
	if api.Info.Title != "Swagger Sample App" {
		t.FailNow()
	}
	if api.Info.Description != "https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md" {
		t.FailNow()
	}
	if api.Info.TermsOfService != "http://swagger.io/terms/" {
		t.FailNow()
	}
	if api.Info.Contact.Name != "API Support" {
		t.FailNow()
	}
	if api.Info.Contact.URL != "http://www.swagger.io/support" {
		t.FailNow()
	}
	if api.Info.Contact.Email != "support@swagger.io" {
		t.FailNow()
	}
	if api.Info.License.Name != "Apache 2.0" {
		t.FailNow()
	}
	if api.Info.License.URL != "http://www.apache.org/licenses/LICENSE-2.0.html" {
		t.FailNow()
	}
	if api.Info.Version != "1.0.1" {
		t.FailNow()
	}
	if api.Host != "127.0.0.1:3000" {
		t.FailNow()
	}
	if api.BasePath != "/api/v1" {
		t.FailNow()
	}
	if !reflect.DeepEqual(api.Schemes, []string{"http", "https", "ws", "wss"}) {
		t.FailNow()
	}
	if !reflect.DeepEqual(api.Tags, []spec.Tag{
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
	}) {
		t.FailNow()
	}
	assert("api.Consumes", t,
		reflect.DeepEqual(api.Consumes, []string{
			"text/plain; charset=utf-8",
			"application/json",
			"application/vnd.github+json",
			"application/vnd.github.v3+json",
		}))

	// Logging
	t.Logf("spec.Paths: %v\n", path)
}

// func TestParseOASApiV3(t *testing.T) {
// 	path := "../../aio/testdata/v3.0.json"
// }

func assert(what string, t *testing.T, expected bool) {
	t.Helper()
	if !expected {
		t.Errorf("founded an error in \"%s\"\n", what)
		t.FailNow()
	}
}
