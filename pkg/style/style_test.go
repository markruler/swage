package style

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"
)

func TestStyleLeft(t *testing.T) {
	xl := excelize.NewFile()
	var styleID int
	styleID = Left(xl)
	assert.Equal(t, 1, styleID)
	styleID = Center(xl)
	assert.Equal(t, 2, styleID)
	styleID = Title(xl)
	assert.Equal(t, 3, styleID)
	styleID = Button(xl)
	assert.Equal(t, 4, styleID)
}
