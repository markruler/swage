package simple

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/parser"
)

func (simple *Simple) setAPISheetResponse(operation *spec.Operation) (err error) {
	xl := simple.xl
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "RESPONSE")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetRowHeight(xl.WorkSheetName, xl.Context.Row, 20.0)
	xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row), xl.Style.Title)
	xl.Context.Row++

	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "code")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), "schema")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "C", xl.Context.Row), "param-type")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), "data-type")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "E", xl.Context.Row), "enum")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), "example")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), "description")
	xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row), xl.Style.Column)
	xl.Context.Row++

	responses := operation.Responses
	if responses == nil {
		return errors.New("response is empty")
	}

	if responses.Default != nil {
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "F", xl.Context.Row), xl.Style.Center)
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row), xl.Style.Left)
		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "default")
		if responses.Default.Schema != nil && !reflect.DeepEqual(spec.Ref{}, responses.Default.Schema.Ref) {
			schema, err := spec.ResolveRef(xl.SwaggerSpec, &responses.Default.Schema.Ref)
			if err != nil {
				return err
			}
			schemaName, _ := simple.definitionFromRef(responses.Default.Schema.Ref)
			simple.setCellWithSchema(schemaName, "body", strings.Join(schema.Type, ","), responses.Default.Description)
		} else {
			simple.setCellWithSchema("", "body", "string", responses.Default.Description)
		}
		xl.Context.Row++
	}

	codes := parser.SortMap(responses.StatusCodeResponses)
	for _, code := range codes {
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "F", xl.Context.Row), xl.Style.Center)
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row), xl.Style.Left)

		icode, err := strconv.Atoi(code)
		if err != nil {
			return err
		}
		response := responses.StatusCodeResponses[icode]

		xl.File.SetCellInt(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), icode)
		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), response.Description)

		if reflect.DeepEqual(spec.Response{}, response) {
			continue
		}

		for headerKey, header := range response.Headers {
			simple.setCellWithSchema(headerKey, "header", header.Type, header.Description)
		}

		if response.Schema == nil {
			xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), response.Description)
			xl.Context.Row++
			continue
		}

		if !reflect.DeepEqual(response.Schema.Ref, spec.Ref{}) {
			simple.responseSchemaRef(response)
		}
		simple.responseSchema(response)
	}
	return nil
}
