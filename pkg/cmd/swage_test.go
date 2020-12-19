package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandSwage(t *testing.T) {
	_, err := executeCommand(swageCmd, "--help")
	assert.Nil(t, err)
}
