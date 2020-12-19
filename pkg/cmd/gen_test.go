package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCommandGen(t *testing.T) {
	var output string
	var err error
	output, err = executeCommand(genCmd, "../../aio/testdata/v2.0.json", "--output", "swage.xlsx")
	assert.NoError(t, err)
	assert.Empty(t, output)
	output, err = executeCommand(genCmd, "../../aio/testdata/v2.0.json")
	assert.NoError(t, err)
	assert.Empty(t, output)
	output, err = executeCommand(genCmd, "--help")
	assert.NoError(t, err)
	assert.Empty(t, output)
}

func TestGenRun(t *testing.T) {
	var err error
	err = genRun(genCmd, []string{})
	assert.Error(t, err)
	err = genRun(genCmd, []string{"../../aio/testdata/v2.0.js"})
	assert.Error(t, err)
	err = genRun(genCmd, []string{"../../aio/testdata/v2.0.json"})
	if err := genCmd.Flags().Set("verbose", "true"); err != nil {
		t.Error(err)
	}
	err = genRun(genCmd, []string{"../../aio/testdata/v2.0.json"})
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
