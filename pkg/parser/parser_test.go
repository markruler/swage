package parser

import (
	"reflect"
	"testing"

	"github.com/markruler/swage/pkg/spec"
)

// TODO: Compare two arrays
// TODO: Compare two objects
func TestParsePetStoreAPI(t *testing.T) {
	path := "../../aio/example/example.json"
	api, err := Parse(path)
	t.Logf("%v\n", api)
	t.Logf("%v\n", api.Schemes)
	for path, desc := range api.Paths {
		if !reflect.DeepEqual(desc.Put, spec.Operation{}) {
			t.Logf("spec.Paths: %v\n", path)
			// t.Logf("%v\n", spec.Put.Parameters)
			for _, param := range desc.Put.Parameters {
				t.Logf("param.Put.Parameters.Name: %v\n", param.Name)
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
