package simple

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestGetParameterSchema(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()

	param := simple.parameterFromRef(spec.Ref{})
	assert.Nil(t, param)

	param = simple.parameterFromRef(spec.MustCreateRef(""))
	assert.Nil(t, param)

	param = simple.parameterFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Nil(t, param)

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
		},
	}
	param = simple.parameterFromRef(spec.MustCreateRef("#/asd/qwe"))
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
	param = simple.parameterFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Equal(t, xl.SwaggerSpec.Parameters["qwe"], *param)
	assert.Equal(t, xl.SwaggerSpec.Parameters["qwe"].Name, "test name")
}

func TestGetDefinitionSchema(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()

	_, def := simple.definitionFromRef(spec.Ref{})
	assert.Nil(t, def)

	_, def = simple.definitionFromRef(spec.MustCreateRef(""))
	assert.Nil(t, def)

	// FIXME: converting undefined references
	// _, def = simple.parameterFromRef(spec.MustCreateRef("#/asd/qwe"))
	// assert.Nil(t, def)

	simple.GetExcel().SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
		},
	}
	_, def = simple.definitionFromRef(spec.MustCreateRef("#/asd/qwe"))
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
	_, def = simple.definitionFromRef(spec.MustCreateRef("#/asd/qwe"))
	assert.Equal(t, xl.SwaggerSpec.Definitions["qwe"], *def)
	assert.Equal(t, xl.SwaggerSpec.Definitions["qwe"].ID, "test id")
}
