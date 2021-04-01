package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandGen(t *testing.T) {
	var output string
	var err error

	output, err = executeCommand(genCmd, "../testdata/json/dev.json", "--output", "swage.xlsx")
	assert.NoError(t, err)
	assert.Empty(t, output)

	output, err = executeCommand(genCmd, "../testdata/json/dev.json")
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

	err = genRun(genCmd, []string{"../testdata/json/dev.js"})
	assert.Error(t, err)

	err = genRun(genCmd, []string{"../testdata/json/dev.json"})
	if err := genCmd.Flags().Set("verbose", "true"); err != nil {
		t.Error(err)
	}

	assert.NoError(t, err)
}
