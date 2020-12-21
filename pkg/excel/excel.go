package excel

import (
	"errors"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/go-openapi/spec"
)

// Excel is needed to generate a spreadsheet.
type Excel struct {
	File           *excelize.File
	Style          style
	SwaggerSpec    *spec.Swagger
	indexSheetName string
	OutputFilePath string
}

// New returns an Excel struct instance.
func New(path string) *Excel {
	xl := &Excel{
		File: excelize.NewFile(),
	}

	xl.indexSheetName = "INDEX"

	xl.setStyle()

	if path == "" {
		xl.OutputFilePath = "swage.xlsx"
	} else {
		xl.OutputFilePath = path
	}

	return xl
}

// Generate function generates a spreadsheet and then saves Excel file
// with the specified path.
func (xl *Excel) Generate(swaggerAPI *spec.Swagger) error {
	if swaggerAPI == nil {
		return errors.New("OpenAPI should not be empty")
	}
	if swaggerAPI.Swagger == "" {
		return errors.New("OpenAPI version should not be empty")
	}
	if swaggerAPI.Paths == nil {
		return errors.New("Path sould not be empty")
	}
	xl.SwaggerSpec = swaggerAPI
	if err := xl.createIndexSheet(); err != nil {
		return err
	}
	return nil
}
