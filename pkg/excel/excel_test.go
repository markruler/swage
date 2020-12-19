package excel

import (
	"testing"

	"github.com/markruler/swage/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	var err error
	err = Save(&spec.SwaggerAPI{}, "", false)
	assert.Error(t, err)

	err = Save(&spec.SwaggerAPI{
		Swagger: "2.0",
	}, "", false)
	assert.NoError(t, err)
}
