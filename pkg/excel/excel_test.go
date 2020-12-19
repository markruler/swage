package excel

import (
	"testing"

	"github.com/markruler/swage/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	var err error
	_, err = Save(nil, "", false)
	assert.Error(t, err)

	_, err = Save(&spec.SwaggerAPI{}, "", false)
	assert.Error(t, err)

	path, err := Save(&spec.SwaggerAPI{
		Swagger: "2.0",
		Paths: map[string]map[string]spec.Operation{
			"/test": {
				"get": {
					Summary: "test",
				},
			},
		},
	}, "", true)
	assert.NoError(t, err)
	assert.Equal(t, "swage.xlsx", path)

	path, err = Save(&spec.SwaggerAPI{
		Swagger: "2.0",
		Paths: map[string]map[string]spec.Operation{
			"/test": {
				"get": {
					Summary: "test",
				},
			},
		},
	}, "excel_test.xlsx", true)
	assert.NoError(t, err)
	assert.Equal(t, "excel_test.xlsx", path)

	path, err = Save(&spec.SwaggerAPI{
		Swagger: "2.0",
		Paths: map[string]map[string]spec.Operation{
			"/test": {
				"get": {
					Summary: "test",
				},
			},
		},
	}, "excel_test.xlsx", false)
	assert.NoError(t, err)
	assert.Equal(t, "", path)
}
