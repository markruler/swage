package xlsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleLeft(t *testing.T) {
	xl := New()
	assert.Equal(t, 1, xl.Style.Left)
}
func TestStyleCenter(t *testing.T) {
	xl := New()
	assert.Equal(t, 2, xl.Style.Center)
}

func TestStyleButton(t *testing.T) {
	xl := New()
	assert.Equal(t, 3, xl.Style.Button)
}

func TestStyleTitle(t *testing.T) {
	xl := New()
	assert.Equal(t, 4, xl.Style.Title)
}

func TestStyleColumn(t *testing.T) {
	xl := New()
	assert.Equal(t, 5, xl.Style.Column)
}
