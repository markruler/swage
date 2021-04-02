package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandGen_NoFlag(t *testing.T) {
	output, err := executeCommand(genCmd, "../testdata/json/sample.pet.json")
	assert.NoError(t, err)
	assert.Empty(t, output)
}

func TestCommandGen_FlagOutput(t *testing.T) {
	output, err := executeCommand(genCmd, "../testdata/json/sample.pet.json", "--output", "swage.xlsx")
	assert.NoError(t, err)
	assert.Empty(t, output)
}

func TestCommandGen_Help(t *testing.T) {
	output, err := executeCommand(genCmd, "--help")
	assert.NoError(t, err)
	assert.Empty(t, output)
}

func TestCommandGen_EmptyPath(t *testing.T) {
	err := genRun(genCmd, []string{})
	assert.Error(t, err, "PATH is required")
}

func TestCommandGen_FileNotFound(t *testing.T) {
	err := genRun(genCmd, []string{"../testdata/json/dev.js"})
	assert.Error(t, err)
}

func TestCommandGen_NormalSpec(t *testing.T) {
	err := genRun(genCmd, []string{"../testdata/json/sample.pet.json"})
	if err := genCmd.Flags().Set("verbose", "true"); err != nil {
		t.Error(err)
	}
	assert.NoError(t, err)
}
