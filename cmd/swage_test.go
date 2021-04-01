package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandSwage(t *testing.T) {
	_, err := executeCommand(swageCmd, "--help")
	assert.NoError(t, err)
}

func TestExecuteCommand(t *testing.T) {
	err := Execute()
	assert.NoError(t, err)
}
