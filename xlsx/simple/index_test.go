package simple

import (
	"testing"

	"github.com/markruler/swage/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateIndexSheet_APINotExists(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()

	xl.SwageSpec = &parser.SwageSpec{
		API: []parser.SwageAPI{},
	}

	err := simple.CreateIndexSheet()
	assert.Error(t, err)
}

func TestCreateIndexSheet_APIExists(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()

	xl.SwageSpec = &parser.SwageSpec{
		API: []parser.SwageAPI{
			{
				Header: parser.APIHeader{
					Tag: "test",
				},
			},
		},
	}

	err := simple.CreateIndexSheet()
	assert.NoError(t, err)

	prop, err := xl.File.GetDocProps()
	assert.NoError(t, err)
	assert.Equal(t, "OpenAPI", prop.Category)
	assert.Equal(t, "Swage", prop.Creator)
	assert.Equal(t, "xlsx", prop.Identifier)
}
