package parser

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

type Parser struct {
	JsonPath string
}

// Parse ...
func (p *Parser) Parse() (*spec.Swagger, error) {
	doc, err := loads.Spec(p.JsonPath)
	if err != nil {
		return nil, err
	}
	// an := analysis.New(doc.Spec())
	// opt := analysis.FlattenOpts{
	// 	Spec: an, BasePath: p.JsonPath,
	// 	Expand: true,
	// }
	// erf := analysis.Flatten(opt)
	// if erf != nil {
	// 	return nil, erf
	// }
	return doc.Spec(), nil
}
