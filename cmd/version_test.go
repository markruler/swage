package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandVersion_Snapshot(t *testing.T) {
	out, err := executeCommand(swageCmd.Root(), versionCmd.Use)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("swage %s\n", swageVersion), string(out))
}
