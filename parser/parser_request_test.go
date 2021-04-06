package parser

import (
	"testing"

	oas "github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestExtractRequests(t *testing.T) {
	test_operation := &oas.Operation{
		OperationProps: oas.OperationProps{
			Parameters: []oas.Parameter{
				{
					ParamProps: oas.ParamProps{
						Name:        "all",
						In:          "query",
						Description: "Return all containers. By default, only running containers are shown.\n",
						Required:    true,
					},
					SimpleSchema: oas.SimpleSchema{
						Type:    "boolean",
						Default: false,
					},
				},
			},
			Responses: &oas.Responses{},
		},
	}

	oas_spec := &oas.Swagger{
		SwaggerProps: oas.SwaggerProps{
			Paths: &oas.Paths{
				Paths: map[string]oas.PathItem{
					"/": {
						PathItemProps: oas.PathItemProps{
							Get: test_operation,
						},
					},
				},
			},
		},
	}

	api_request, err := extractRequests(oas_spec, test_operation)
	assert.NoError(t, err)

	assert.Equal(t, "O", api_request[0].Required)
	assert.Equal(t, "all", api_request[0].Schema)
	assert.Equal(t, "query", api_request[0].ParameterType)
	assert.Equal(t, "boolean", api_request[0].DataType)
	assert.Equal(t, "", api_request[0].Enum)
	assert.Equal(t, "", api_request[0].Example)
	assert.Equal(t, "Return all containers. By default, only running containers are shown.\n", api_request[0].Description)

}
