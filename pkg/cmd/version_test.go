package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandVersion(t *testing.T) {
	_, err := executeCommand(versionCmd)
	assert.Nil(t, err)
}

func TestVersionRun(t *testing.T) {
	versionRun(nil, nil)
}
