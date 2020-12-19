package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandVersion(t *testing.T) {
	_, err := executeCommand(versionCmd)
	assert.Nil(t, err)
}
