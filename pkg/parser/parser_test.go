package parser

import (
	"reflect"
	"testing"

	"github.com/markruler/swage/pkg/spec"
)

// TODO: Compare two arrays
// TODO: Compare two objects
func TestParsePetStoreAPI(t *testing.T) {
	path := "../../aio/example/swagger.editor.json"
	api, err := Parse(path)
	t.Logf("%v\n", api)
	t.Logf("%v\n", api.Schemes)
	for path, operations := range api.Paths {
		for operation, detail := range operations {
			t.Logf("%s\n", operation)
			if !reflect.DeepEqual(detail, spec.Operation{}) {
				t.Logf("spec.Paths: %v\n", path)
				// t.Logf("%v\n", detail.Parameters)
				for _, param := range detail.Parameters {
					t.Logf("param.Put.Parameters.Name: %v\n", param.Name)
				}
			}
		}
	}

	if err != nil {
		t.FailNow()
	}
	if api == nil {
		t.FailNow()
	}
	if api.Swagger != "2.0" {
		t.FailNow()
	}
	if api.Info.Title != "Swagger Petstore" {
		t.FailNow()
	}
}

// func TestParseShortAPI(t *testing.T) {
// 	path := "../aio/example/short.json"
// }
