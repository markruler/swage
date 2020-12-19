package excel

import (
	"testing"

	"github.com/markruler/swage/pkg/spec"
	"github.com/stretchr/testify/assert"
)


func TestCreateIndexSheet(t *testing.T) {
	xl := createIndexSheet(&spec.SwaggerAPI{})
	prop, err := xl.GetDocProps()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "OpenAPI", prop.Category)
	assert.Equal(t, "Swage", prop.Creator)
	assert.Equal(t, "xlsx", prop.Identifier)
}