package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/markruler/swage/pkg/spec"
)

// Parse ...
func Parse(jsonPath string) (*spec.SwaggerAPI, error) {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("%s\n", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("%s\n", err)
		return nil, err
	}

	var swaggerAPI spec.SwaggerAPI
	json.Unmarshal(byteValue, &swaggerAPI)

	return &swaggerAPI, err
}
