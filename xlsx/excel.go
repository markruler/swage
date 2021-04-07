package xlsx

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	oas "github.com/go-openapi/spec"
	"github.com/markruler/swage/spec"
)

// Excel to save
type Excel struct {
	File           *excelize.File
	Style          style
	SwaggerSpec    *oas.Swagger
	SwageSpec      *spec.SwageSpec
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
