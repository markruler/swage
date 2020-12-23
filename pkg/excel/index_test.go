package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestCreateIndexSheet(t *testing.T) {
	xl := New()
	var err error

	xl.SwaggerSpec = &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{
					"/test": {
						PathItemProps: spec.PathItemProps{
							Get: &spec.Operation{
								OperationProps: spec.OperationProps{
									Deprecated: true,
								},
							},
							Put: &spec.Operation{
								OperationProps: spec.OperationProps{
									Deprecated: true,
								},
							},
							Post: &spec.Operation{
								OperationProps: spec.OperationProps{
									Deprecated: true,
								},
							},
							Delete: &spec.Operation{
								OperationProps: spec.OperationProps{
									Deprecated: true,
								},
							},
							Options: &spec.Operation{
								OperationProps: spec.OperationProps{
									Deprecated: true,
								},
							},
							Head: &spec.Operation{
								OperationProps: spec.OperationProps{
									Deprecated: true,
								},
							},
							Patch: &spec.Operation{
								OperationProps: spec.OperationProps{
									Deprecated: true,
								},
							},
						},
					},
				},
			},
		},
	}
	err = xl.createIndexSheet()
	assert.NoError(t, err)
	prop, err := xl.File.GetDocProps()
	assert.NoError(t, err)
	assert.Equal(t, "OpenAPI", prop.Category)
	assert.Equal(t, "Swage", prop.Creator)
	assert.Equal(t, "xlsx", prop.Identifier)

	// xl.SwaggerSpec = &spec.Swagger{}
	// err = xl.createIndexSheet()
	// assert.Error(t, err)
}
