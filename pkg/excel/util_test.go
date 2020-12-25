package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestSortMap(t *testing.T) {
	arr := sortMap("")
	assert.Nil(t, arr)
}

func TestGetParameterSchema(t *testing.T) {
	xl := New()
	// composite literal uses unkeyed fields
	// type Ref struct {
	// 	jsonreference.Ref
	// }
	// param := xl.getParameterFromRef(spec.Ref{
	// 	jsonreference.MustCreateRef("#/re/re"),
	// })

	param := xl.getParameterFromRef(spec.Ref{})
	assert.Nil(t, param)

	param = xl.getParameterFromRef(spec.MustCreateRef(""))
	assert.Nil(t, param)

	param = xl.getParameterFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Nil(t, param)

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
		},
	}
	param = xl.getParameterFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Nil(t, param)

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Parameters: map[string]spec.Parameter{
				"qwe": {
					ParamProps: spec.ParamProps{
						Name: "test name",
					},
				},
			},
		},
	}
	param = xl.getParameterFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Equal(t, xl.SwaggerSpec.Parameters["qwe"], *param)
	assert.Equal(t, xl.SwaggerSpec.Parameters["qwe"].Name, "test name")
}

func TestGetDefinitionSchema(t *testing.T) {
	xl := New()

	_, def := xl.getDefinitionFromRef(spec.Ref{})
	assert.Nil(t, def)

	_, def = xl.getDefinitionFromRef(spec.MustCreateRef(""))
	assert.Nil(t, def)

	// _, def = xl.getParameterFromRef(spec.MustCreateRef("#/asd/qwe"))
	// assert.Nil(t, def)

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
		},
	}
	_, def = xl.getDefinitionFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Nil(t, def)

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: spec.Definitions{
				"qwe": spec.Schema{
					SchemaProps: spec.SchemaProps{
						ID: "test id",
					},
				},
			},
		},
	}
	_, def = xl.getDefinitionFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Equal(t, xl.SwaggerSpec.Definitions["qwe"], *def)
	assert.Equal(t, xl.SwaggerSpec.Definitions["qwe"].ID, "test id")
}
