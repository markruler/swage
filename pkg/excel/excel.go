package excel

import (
	"errors"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/go-openapi/spec"
)

// Excel ...
type Excel struct {
	File           *excelize.File
	SwaggerSpec    *spec.Swagger
	Style          style
	OutputFilePath string
	Verbose        bool
	indexSheetName string
}

// New ...
func New(path string) *Excel {
	xl := &Excel{
		File: excelize.NewFile(),
	}
	xl.setStyle()

	if path == "" {
		xl.OutputFilePath = "swage.xlsx"
	} else {
		xl.OutputFilePath = path
	}

	xl.indexSheetName = "INDEX"

	return xl
}

// Save ...
func (xl *Excel) Save(swaggerAPI *spec.Swagger) (path string, err error) {
	if swaggerAPI == nil {
		return "", errors.New("OpenAPI should not be empty")
	}
	if swaggerAPI.Swagger == "" {
		return "", errors.New("OpenAPI version should not be empty")
	}
	if swaggerAPI.Paths == nil {
		return "", errors.New("Path sould not be empty")
	}
	xl.SwaggerSpec = swaggerAPI
	if err := xl.createIndexSheet(); err != nil {
		return "", err
	}
	if err := xl.File.SaveAs(xl.OutputFilePath); err != nil {
		log.Fatalln(err)
	}
	if xl.Verbose {
		return xl.OutputFilePath, nil
	}
	return "", nil
}
