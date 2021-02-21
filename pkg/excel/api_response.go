package excel

import (
	"errors"
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
		var icode int
		if icode, err = strconv.Atoi(code); err != nil {
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

		if response.Schema == nil || &response.Ref == nil {
			// TODO: write test code
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)
			xl.Context.row++
			continue
		}

		if response.Schema.Title != "" {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), response.Schema.Title)
		}

		if response.Schema.Type != nil {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(response.Schema.Type, ","))
		}

		if response.Schema.Type.Contains("array") {
			items := response.Schema.Items
			if items.Schema != nil {
				schema := items.Schema
				xl.setCellWithSchema(schema.Title, "body", strings.Join(response.Schema.Type, ","), response.Description)
				continue
			}
			for _, schema := range items.Schemas {
				if !reflect.DeepEqual(spec.Ref{}, schema.Ref) {
					definitionName, definition := xl.getDefinitionFromRef(items.Schemas[0].Ref)
					if definition == nil {
						return errors.New("not found response.Schema.Items definition")
					}
					xl.setCellWithSchema(definitionName, "body", "array", response.Description)
					continue
				}
			}
		}

		if !reflect.DeepEqual(spec.Ref{}, response.Schema.Ref) {
			schema, err := spec.ResolveRef(*xl.SwaggerSpec, &response.Schema.Ref)
			if err != nil {
				return err
			}
			if schema == nil {
				return errors.New("not found response.Schema.Ref definition")
			}
			// TODO: handle 'AllOf'
			// if len(schema.AllOf) != 0 {
			// 	for _, oneSchema := range schema.AllOf {
			// 		fmt.Println("oneSchema:", oneSchema)
			// 		schema, err := spec.ResolveResponse(xl.SwaggerSpec, oneSchema.Ref)
			// 		if err != nil {
			// 			return err
			// 		}
			// 		fmt.Println("schema:", schema)
			// 		schemaName, _ := xl.getDefinitionFromRef(oneSchema.Ref)
			// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), schemaName)
			// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
			// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "object")
			// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), schema.Description)
			// 	}
			// 	return nil
			// }
			schemaName, _ := xl.getDefinitionFromRef(response.Schema.Ref)
			xl.setCellWithSchema(schemaName, "body", "object", response.Description)
			xl.Context.row++
			continue
		}

		if response.Schema.Title != "" {
			xl.setCellWithSchema(response.Schema.Title, "body", "object", response.Description)
			xl.Context.row++
			continue
		}

		// TODO: refactoring if-hell
		// TODO: write test code
		if response.Schema.Properties != nil {
			for propertyName, propertySchema := range response.Schema.Properties {
				if !reflect.DeepEqual(spec.Ref{}, propertySchema.Ref) {
					definitionName, definition := xl.getDefinitionFromRef(propertySchema.Ref)
					if definition == nil {
						return errors.New("not found definition")
					}
					if propertySchema.Items != nil {
						definitionName, definition = xl.getDefinitionFromRef(propertySchema.Items.Schema.Ref)
						if definition == nil {
							return errors.New("not found definition")
						}
					}
					xl.setCellWithSchema(definitionName, "body", propertyName, propertySchema.Description)
					xl.Context.row++
					continue
				}
				xl.setCellWithSchema(propertyName, "body", strings.Join(response.Schema.Type, ","), response.Description)
				xl.Context.row++
			}
		}
	}
	return nil
}
