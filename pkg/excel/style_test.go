package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleLeft(t *testing.T) {
	xl := New("")
	assert.Equal(t, 1, xl.Style.Left)
}
func TestStyleCenter(t *testing.T) {
	xl := New("")
	assert.Equal(t, 2, xl.Style.Center)
}

func TestStyleTitle(t *testing.T) {
	xl := New("")
	assert.Equal(t, 3, xl.Style.Title)
}

func TestStyleButton(t *testing.T) {
	xl := New("")
	assert.Equal(t, 4, xl.Style.Button)
}
