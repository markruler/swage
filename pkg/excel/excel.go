package excel

import (
	"errors"
	"log"

	"github.com/markruler/swage/pkg/spec"
)

const (
	defaultFilePath = "swage.xlsx"
)

// Save ...
func Save(swaggerAPI *spec.SwaggerAPI, outputFilePath string, verbose bool) (string, error) {
	if swaggerAPI == nil {
		return "", errors.New("OpenAPI should not be nil")
	}
	if swaggerAPI.Swagger == "" {
		return "", errors.New("OpenAPI version should not be nil")
	}
	xl := createIndexSheet(swaggerAPI)
	var path string
	if outputFilePath != "" {
		path = outputFilePath
	} else {
		path = defaultFilePath
	}

	if err := xl.SaveAs(path); err != nil {
		log.Fatalln(err)
	}
	if verbose {
		return path, nil
	}
	return "", nil
}
