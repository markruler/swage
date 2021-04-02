package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMap_Empty(t *testing.T) {
	arr := SortMap("")
	assert.Nil(t, arr)
}

func TestSortMap_Responses(t *testing.T) {
	testmap := map[int]interface{}{
		404: "Not Found",
		500: "Internal Server Error",
		200: "OK",
		301: "Moved Permanently",
	}
	arr := SortMap(testmap)
	assert.Equal(t, []string{"200", "301", "404", "500"}, arr)
}

func TestEnum2string_string(t *testing.T) {
	str := Enum2string("qwe", "asd")
	assert.Equal(t, "qwe,asd", str)
}

// @source cisco.meraki.yaml
// @method get
// @path /networks/{networkId}/wireless/channelUtilizationHistory
func TestEnum2string_float64(t *testing.T) {
	str := Enum2string(2.4, 5)
	assert.Equal(t, "2.4,5", str)
}
