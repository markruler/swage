package simple

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPISheet(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	var err error

	err = simple.CreateAPISheet("", "", nil, nil, 1)
	assert.Error(t, err)

	err = simple.CreateAPISheet("", "", &spec.Operation{}, nil, 1)
	assert.Error(t, err)

	err = simple.CreateAPISheet("", "", &spec.Operation{
		OperationProps: spec.OperationProps{
			Parameters: []spec.Parameter{
				{
					ParamProps: spec.ParamProps{
						Required: true,
					},
				},
			},
			Responses: &spec.Responses{},
		},
	}, nil, 1)
	assert.NoError(t, err)

	p := parser.New("../../aio/testdata/json/dev.json")
	xl.SwaggerSpec, _ = p.Parse()
	simple.CreateAPISheet("", "", nil, nil, 1)
}
