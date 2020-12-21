package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	xl := New("")
	var err error

	_, err = xl.Save(nil)
	assert.Error(t, err)

	_, err = xl.Save(&spec.Swagger{})
	assert.Error(t, err)

	xl.Verbose = true
	path, err := xl.Save(&spec.Swagger{
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
	assert.Equal(t, "swage.xlsx", path)

	xl.OutputFilePath = "excel_test.xlsx"
	// xl = New("excel_test.xlsx")
	path, err = xl.Save(&spec.Swagger{
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
	assert.Equal(t, "excel_test.xlsx", path)
}
