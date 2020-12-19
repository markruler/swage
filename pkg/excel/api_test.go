package excel

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/markruler/swage/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPISheet(t *testing.T) {
	tmp := excelize.NewFile()
	var err error

	// param.Required
	err = createAPISheet(tmp, "", "", spec.Operation{
		Parameters: []spec.Parameters{
			{
				Required: true,
			},
		},
	}, nil, 1)
	assert.Error(t, err)

	err = createAPISheet(tmp, "", "", spec.Operation{
		Parameters: []spec.Parameters{
			{
				Required: false,
			},
		},
	}, nil, 1)
	assert.Error(t, err)

	// reflect.DeepEqual(spec.Response{}, response)
	err = createAPISheet(tmp, "", "", spec.Operation{}, nil, 1)
	assert.Error(t, err)

	// response.Schema.Type == "array"
	err = createAPISheet(tmp, "", "", spec.Operation{
		Responses: map[string]spec.Response{
			"200": {
				Description: "OK",
				Schema: spec.Schema{
					Type: "array",
					Items: spec.Items{
						Ref: "#/definitions/Test",
					},
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)

	// response.Schema.Type != "array"
	err = createAPISheet(tmp, "", "", spec.Operation{
		Responses: map[string]spec.Response{
			"200": {
				Description: "OK",
				Schema: spec.Schema{
					Type: "integer",
				},
			},
		},
	}, nil, 1)
	assert.NoError(t, err)
}
