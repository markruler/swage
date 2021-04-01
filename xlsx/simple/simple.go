package simple

import (
	"errors"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/xlsx"
)

type Simple struct {
	xl *xlsx.Excel
}

func New() *Simple {
	return &Simple{
		xl: xlsx.New(),
	}
}

func (simple *Simple) GetExcel() *xlsx.Excel {
	return simple.xl
}

func (simple *Simple) Generate(spec *spec.Swagger) error {
	simple.xl = xlsx.New()

	if spec == nil {
		return errors.New("OpenAPI should not be empty")
	}

	if spec.Swagger == "" {
		return errors.New("OpenAPI version should not be empty")
	}

	if spec.Paths == nil {
		return errors.New("path sould not be empty")
	}

	simple.xl.SwaggerSpec = spec

	if err := simple.CreateIndexSheet(); err != nil {
		return err
	}
	return nil
}
