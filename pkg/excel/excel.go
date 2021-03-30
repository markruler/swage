package excel

import (
	"errors"
	"strings"

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
func (xl *Excel) Generate(swaggerAPI *spec.Swagger, template string) error {
	// fmt.Printf("template: \"%s\"\n", strings.TrimSpace(template))

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

	switch strings.TrimSpace(template) {
	case "default":
		if err := xl.createIndexSheet(); err != nil {
			return err
		}
	// TODO: make a formal template
	// case "custom":
	// if err := xl.createMyCustomSheet(); err != nil {
	// 	return err
	// }
	default:
		return errors.New("the template not found")
	}
	return nil
}
