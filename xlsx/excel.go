package xlsx

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/go-openapi/spec"
)

// Excel to save
type Excel struct {
	File           *excelize.File
	Style          style
	SwaggerSpec    *spec.Swagger
	IndexSheetName string
	WorkSheetName  string
	Context        *context
}

type context struct {
	Row int
}

// New returns an Excel struct instance
func New() *Excel {
	xl := &Excel{
		File:    excelize.NewFile(),
		Context: &context{},
	}
	xl.File.SetDefaultFont("Arial")
	xl.setStyle()
	xl.IndexSheetName = "INDEX"
	return xl
}
