package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
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

// https://github.com/spf13/cobra/blob/v1.1.1/command_test.go
func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}
