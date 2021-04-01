package xlsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateExcel(t *testing.T) {
	xl := New()
	assert.Equal(t, "INDEX", xl.IndexSheetName)
}
