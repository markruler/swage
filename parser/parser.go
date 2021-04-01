package parser

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

// Parser ...
type Parser struct {
	sourcePath string
}

// New returns a Parser instance with a resource path
func New(path string) *Parser {
	return &Parser{
		sourcePath: path,
	}
}

// Parse ...
func (p *Parser) Parse() (*spec.Swagger, error) {
	doc, err := loads.Spec(p.sourcePath)
	if err != nil {
		return nil, err
	}
	return doc.Spec(), nil
}
