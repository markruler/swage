package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/markruler/swage/pkg/spec"
)

// Parse ...
func Parse(jsonPath string) (*spec.SwaggerAPI, error) {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var swaggerAPI spec.SwaggerAPI
	if err := json.Unmarshal(byteValue, &swaggerAPI); err != nil {
		return nil, err
	}

	return &swaggerAPI, err
}
