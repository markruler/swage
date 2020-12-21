package style

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleLeft(t *testing.T) {
	assert.Equal(t, 1, IDLeft)
}
func TestStyleCenter(t *testing.T) {
	assert.Equal(t, 2, IDCenter)
}

func TestStyleTitle(t *testing.T) {
	assert.Equal(t, 3, IDTitle)
}

func TestStyleButton(t *testing.T) {
	assert.Equal(t, 4, IDButton)
}
