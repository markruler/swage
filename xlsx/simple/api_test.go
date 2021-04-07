package simple

import (
	"testing"

	"github.com/markruler/swage/spec"
	"github.com/stretchr/testify/assert"
)

func TestCreateAPISheet_NormalSpec(t *testing.T) {
	simple := New()
	xl := simple.xl
	xl.SwageSpec = &spec.SwageSpec{
		API: []spec.SwageAPI{
			{
				Header: spec.APIHeader{
					Tag: "test",
				},
			},
		},
	}
	err := simple.CreateAPISheet()
	assert.NoError(t, err)
}

func TestCreateSimpleAPISheet_Request(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	xl.SwageSpec = &spec.SwageSpec{
		API: []spec.SwageAPI{
			{
				Header: spec.APIHeader{
					ID: "ContainerArchiveInfo",
				},
				Request: []spec.APIRequest{
					{
						Required:      "O",
						Schema:        "all",
						ParameterType: "query",
						DataType:      "boolean",
						Description:   "Return all containers. By default, only running containers are shown.\n",
					},
				},
			},
		},
	}
	err := simple.CreateAPISheet()
	assert.NoError(t, err)

	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)

	assert.Equal(t, "O", row[12][0])
	assert.Equal(t, "all", row[12][1])
	assert.Equal(t, "query", row[12][2])
	assert.Equal(t, "boolean", row[12][3])
	assert.Equal(t, "Return all containers. By default, only running containers are shown.\n", row[12][6])
}

func TestCreateSimpleAPISheet_HeaderTypeResponse(t *testing.T) {
	simple := New()
	xl := simple.GetExcel()
	xl.SwageSpec = &spec.SwageSpec{
		API: []spec.SwageAPI{
			{
				Header: spec.APIHeader{
					ID: "ContainerArchiveInfo",
				},
				Response: []spec.APIResponse{
					{
						StatusCode:   "200",
						Schema:       "X-Docker-Container-Path-Stat",
						ResponseType: "header",
						DataType:     "string",
						Description:  "A base64 - encoded JSON object with some filesystem header\ninformation about the path\n",
					},
					{
						StatusCode:   "400",
						Schema:       "ErrorResponse",
						ResponseType: "body",
						DataType:     "object",
						Description:  "error",
					},
				},
			},
		},
	}

	err := simple.CreateAPISheet()
	assert.NoError(t, err)

	row, err := xl.File.GetRows("1")
	assert.NoError(t, err)

	assert.Equal(t, "200", row[15][0])
	assert.Equal(t, "X-Docker-Container-Path-Stat", row[15][1])
	assert.Equal(t, "header", row[15][2])
	assert.Equal(t, "string", row[15][3])
	assert.Equal(t, "", row[15][4])
	assert.Equal(t, "", row[15][5])
	assert.Equal(t, "A base64 - encoded JSON object with some filesystem header\ninformation about the path\n", row[15][6])

	assert.Equal(t, "400", row[16][0])
	assert.Equal(t, "ErrorResponse", row[16][1])
	assert.Equal(t, "body", row[16][2])
	assert.Equal(t, "object", row[16][3])
	assert.Equal(t, "", row[16][4])
	assert.Equal(t, "", row[16][5])
	assert.Equal(t, "error", row[16][6])
}
