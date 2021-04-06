package simple

import (
	"errors"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/parser"
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

	swage_spec, err := parser.Convert(spec)
	if err != nil {
		return err
	}
	simple.xl.SwageSpec = swage_spec

	if err := simple.CreateIndexSheet(); err != nil {
		return err
	}

	if err := simple.CreateAPISheet(); err != nil {
		return err
	}

	return nil
}
