package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	xl := New()
	var err error

	err = xl.Generate(nil)
	assert.Error(t, err)

	err = xl.Generate(&spec.Swagger{})
	assert.Error(t, err)

	err = xl.Generate(&spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{
					"/test": {
						PathItemProps: spec.PathItemProps{
							Get: &spec.Operation{
								OperationProps: spec.OperationProps{
									Summary: "test",
								},
							},
						},
					},
				},
			},
		},
	})
	assert.NoError(t, err)

	err = xl.Generate(&spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{
					"/test": {
						PathItemProps: spec.PathItemProps{
							Get: &spec.Operation{
								OperationProps: spec.OperationProps{
									Summary: "test",
								},
							},
						},
					},
				},
			},
		},
	})
	assert.NoError(t, err)
}
