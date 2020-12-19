package excel

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/markruler/swage/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestCreateIndexSheet(t *testing.T) {
	var xl *excelize.File

	xl = createIndexSheet(&spec.SwaggerAPI{})
	assert.Nil(t, xl)

	xl = createIndexSheet(&spec.SwaggerAPI{
		Paths: map[string]map[string]spec.Operation{
			"/test": {
				"get": spec.Operation{
					Deprecated: "true",
				},
			},
		},
	})
	prop, err := xl.GetDocProps()
	assert.NoError(t, err)
	assert.Equal(t, "OpenAPI", prop.Category)
	assert.Equal(t, "Swage", prop.Creator)
	assert.Equal(t, "xlsx", prop.Identifier)
}
