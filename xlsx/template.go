package xlsx

import (
	"github.com/go-openapi/spec"
)

type Template interface {
	// GetExcel returns Excel that the template has
	GetExcel() *Excel
	// Generate generates Excel from Open API
	Generate(*spec.Swagger) error
	// CreateIndexSheet generates index sheet
	CreateIndexSheet() error
	// CreateAPISheet generates an API sheet for each index
	CreateAPISheet(path string, method string, operation *spec.Operation, definitions spec.Definitions, sheetName int) error
}
