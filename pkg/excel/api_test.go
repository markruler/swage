package excel

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPISheet(t *testing.T) {
	xl := New("")
	var err error

	err = xl.createAPISheet("", "", &spec.Operation{}, nil, 1)
	assert.NoError(t, err)

	err = xl.createAPISheet("", "", nil, nil, 1)
	assert.Error(t, err)

	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Required: true,
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)

	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "OK",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: spec.StringOrArray{
											"array",
										},
										Items: &spec.SchemaOrArray{
											Schemas: []spec.Schema{
												{
													SchemaProps: spec.SchemaProps{
														Ref: spec.MustCreateRef("#/definitions/Test"),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)

	err = xl.createAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: map[int]spec.Response{
						200: {
							ResponseProps: spec.ResponseProps{
								Description: "OK",
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: spec.StringOrArray{
											"integer",
											"object",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)

	p := parser.New("../../aio/testdata/json/dev.json")
	xl.SwaggerSpec, _ = p.Parse()
	xl.createAPISheet("", "", nil, nil, 1)
}
