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
	Context        *context
}

type context struct {
	worksheetName string
	row           int
}

// New returns an Excel struct instance.
func New() *Excel {
	xl := &Excel{
		File:    excelize.NewFile(),
		Context: &context{},
	}
	xl.File.SetDefaultFont("Arial")
	xl.setStyle()
	xl.indexSheetName = "INDEX"
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
	// TODO: make templates (ex. template1, template2, ...)
	if err := xl.createIndexSheet(); err != nil {
		return err
	}
	return nil
}
