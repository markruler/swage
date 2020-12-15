package cmd

import (
	"reflect"
	"testing"
)

// TODO: Compare two arrays
// TODO: Compare two objects
func TestParsePetStoreAPI(t *testing.T) {
	path := "../aio/example/example.json"
	api, err := parse(path)
	t.Logf("%v\n", api)
	t.Logf("%v\n", api.Schemes)
	for path, spec := range api.Paths {
		if !reflect.DeepEqual(spec.Put, Put{}) {
			t.Logf("spec.Paths: %v\n", path)
			// t.Logf("%v\n", spec.Put.Parameters)
			for _, param := range spec.Put.Parameters {
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
