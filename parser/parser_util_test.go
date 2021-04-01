package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMapEmpty(t *testing.T) {
	arr := SortMap("")
	assert.Nil(t, arr)
}

func TestSortMapResponses(t *testing.T) {
	testmap := map[int]interface{}{
		404: "Not Found",
		500: "Internal Server Error",
		200: "OK",
		301: "Moved Permanently",
	}
	arr := SortMap(testmap)
	assert.Equal(t, []string{"200", "301", "404", "500"}, arr)
}

func TestEnum2string(t *testing.T) {
	str := Enum2string("qwe", "asd")
	assert.Equal(t, "qwe,asd", str)
}
