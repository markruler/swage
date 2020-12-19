package style

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"
)

var xl = excelize.NewFile()

func TestStyleLeft(t *testing.T) {
	styleID := Left(xl)
	assert.Equal(t, 1, styleID)
}
func TestStyleCenter(t *testing.T) {
	styleID := Center(xl)
	assert.Equal(t, 2, styleID)
}

func TestStyleTitle(t *testing.T) {
	styleID := Title(xl)
	assert.Equal(t, 3, styleID)
}

func TestStyleButton(t *testing.T) {
	styleID := Button(xl)
	assert.Equal(t, 4, styleID)
}
