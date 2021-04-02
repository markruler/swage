package simple

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPISheet_EmptyResponse(t *testing.T) {
	err := New().CreateAPISheet("", "", &spec.Operation{}, nil, 1)
	assert.Error(t, err, "response should not be empty")
}

func TestCreateAPISheet_EmptyOperation(t *testing.T) {
	err := New().CreateAPISheet("", "", nil, nil, 1)
	assert.Error(t, err, "operation should not be empty")
}

func TestCreateAPISheet_NormalSpec(t *testing.T) {
	err := New().CreateAPISheet("", "", &spec.Operation{
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
}
