package excel

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
)

func (xl *Excel) setAPISheetResponse(operation *spec.Operation) (err error) {
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "RESPONSE")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetRowHeight(xl.Context.worksheetName, xl.Context.row, 20.0)
	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Title)
	xl.Context.row++

	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "code")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), "schema")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "param-type")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "data-type")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), "enum")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "F", xl.Context.row), "example")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), "description")
	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Column)
	xl.Context.row++

	responses := operation.Responses
	if responses == nil {
		// TODO: nil check
		// return errors.New("[spec.Responses] is empty")
		return nil
	}

	if responses.Default != nil {
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "F", xl.Context.row), xl.Style.Center)
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Left)
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "default")
		if responses.Default.Schema != nil && !reflect.DeepEqual(spec.Ref{}, responses.Default.Schema.Ref) {
			schema, err := spec.ResolveRef(xl.SwaggerSpec, &responses.Default.Schema.Ref)
			if err != nil {
				return err
			}
			schemaName, _ := xl.getDefinitionFromRef(responses.Default.Schema.Ref)
			xl.setCellWithSchema(schemaName, "body", strings.Join(schema.Type, ","), responses.Default.Description)
		} else {
			xl.setCellWithSchema("", "body", "string", responses.Default.Description)
		}
		xl.Context.row++
	}

	codes := sortMap(responses.StatusCodeResponses)
	for _, code := range codes {
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "F", xl.Context.row), xl.Style.Center)
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Left)

		icode, err := strconv.Atoi(code)
		if err != nil {
			return err
		}
		response := responses.StatusCodeResponses[icode]

		xl.File.SetCellInt(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), icode)
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)

		if reflect.DeepEqual(spec.Response{}, response) {
			continue
		}

		for headerKey, header := range response.Headers {
			xl.setCellWithSchema(headerKey, "header", header.Type, header.Description)
		}

		if response.Schema == nil /*|| &response.Ref == nil*/ {
			// TODO: write test code
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)
			xl.Context.row++
			continue
		}

		if !reflect.DeepEqual(response.Schema.Ref, spec.Ref{}) {
			xl.getResponseSchemaRef(response)
		}
		xl.getResponseSchema(response)
	}
	return nil
}
