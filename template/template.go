package template

import (
	"github.com/go-openapi/spec"
	"github.com/markruler/swage/xlsx"
)

type Template interface {
	// GetExcel returns Excel that the template has
	GetExcel() *xlsx.Excel
	// Generate generates Excel from Open API
	Generate(*spec.Swagger) error
	// CreateIndexSheet generates index sheet
	CreateIndexSheet() error
	// CreateAPISheet generates an API sheet for each index
	CreateAPISheet(path string, method string, operation *spec.Operation, definitions spec.Definitions, sheetName int) error
}
