package excel

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
)

func (xl *Excel) createAPISheet(path, method string, operation *spec.Operation, definitions spec.Definitions, sheetName int) (err error) {
	if operation == nil {
		return errors.New("Operation should not be empty")
	}
	xl.Context.worksheetName = strconv.Itoa(sheetName)
	xl.File.NewSheet(xl.Context.worksheetName)

	xl.Context.row = 1
	xl.setAPISheetHeader(path, method, operation)
	if err = xl.setAPISheetRequest(operation); err != nil {
		return err
	}
	if err = xl.setAPISheetResponse(operation); err != nil {
		return err
	}
	return nil
}

func (xl *Excel) setAPISheetHeader(path, method string, operation *spec.Operation) error {
	xl.File.SetColWidth(xl.Context.worksheetName, "A", "A", 12.0)
	xl.File.SetColWidth(xl.Context.worksheetName, "B", "B", 33.0)
	xl.File.SetColWidth(xl.Context.worksheetName, "C", "C", 12.0)
	xl.File.SetColWidth(xl.Context.worksheetName, "D", "D", 12.0)
	xl.File.SetColWidth(xl.Context.worksheetName, "E", "E", 20.0)
	xl.File.SetColWidth(xl.Context.worksheetName, "F", "F", 40.0)
	xl.File.SetColWidth(xl.Context.worksheetName, "G", "G", 90.0)
	// xl.File.SetColStyle(xl.Context.worksheetName, "A:G", xl.Style.Line)

	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "A", xl.Context.row), xl.Style.Button)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Back to Index")
	xl.File.SetCellHyperLink(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "INDEX!A1", "Location")
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Tag")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	if len(operation.Tags) > 0 {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), operation.Tags[0])
	}
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "ID")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), operation.ID)
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Path")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), path)
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Method")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), method)
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Consumes")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), strings.Join(operation.Consumes, ", "))
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Produces")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), strings.Join(operation.Produces, ", "))
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Summary")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), operation.Summary)
	xl.Context.row++
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Description")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), operation.Description)
	xl.Context.row++

	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", 2), fmt.Sprintf("%s%d", "A", xl.Context.row-1), xl.Style.Column)
	xl.Context.row++
	return nil
}

func (xl *Excel) setAPISheetRequest(operation *spec.Operation) (err error) {
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "REQUEST")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetRowHeight(xl.Context.worksheetName, xl.Context.row, 20.0)
	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Title)
	xl.Context.row++

	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "required")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), "schema")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "param-type")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "data-type")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), "enum")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "F", xl.Context.row), "example")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), "description")
	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Column)
	xl.Context.row++

	// TODO: refactoring if-hell
	for _, param := range operation.Parameters {
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "F", xl.Context.row), xl.Style.Center)

		if !reflect.DeepEqual(param.Ref, spec.Ref{}) {
			param = *xl.getParameterFromRef(param.Ref)
		}

		if param.Required {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "O")
		} else {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "X")
		}
		xl.setCellWithSchema(param.Name, param.In, param.Type, param.Description)

		if param.Schema != nil {
			if param.Schema.Items != nil {
				if param.Schema.Items.Schema != nil {
					xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ";"))
				}
				// TODO: Schema's'
				if param.Schema.Items.Schemas != nil {
					xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ";"))
				}
			}
			if !reflect.DeepEqual(spec.Ref{}, param.Schema.Ref) {
				if strings.Contains(param.Schema.Ref.GetPointer().String(), "definitions") {
					schema, err := spec.ResolveRef(xl.SwaggerSpec, &param.Schema.Ref)
					if err != nil {
						return err
					}
					if param.Required {
						xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "O")
					} else {
						xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "X")
					}
					schemaName, _ := xl.getDefinitionFromRef(param.Schema.Ref)
					xl.setCellWithSchema(schemaName, param.In, strings.Join(schema.Type, ";"), param.Description)
					xl.Context.row++
				}
				if strings.Contains(param.Schema.Ref.GetPointer().String(), "parameters") {
					schema, err := spec.ResolveParameter(xl.SwaggerSpec, param.Schema.Ref)
					if err != nil {
						return err
					}
					if schema.Required {
						xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "O")
					} else {
						xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "X")
					}
					xl.setCellWithSchema(schema.Name, schema.In, schema.Type, schema.Description)
					xl.Context.row++
				}
				continue
			}
			if param.Schema.Properties != nil {
				for k, v := range param.Schema.Properties {
					xl.setCellWithSchema(k, param.In, strings.Join(v.Type, ";"), "")
				}
			}
			if param.Schema.Type != nil {
				xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ";"))
			}
			if param.Schema.Description != "" {
				xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), param.Schema.Description)
			}
		}
		xl.Context.row++
	}
	xl.Context.row++
	return nil
}

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

	// TODO: refactoring if-hell
	responses := operation.Responses
	if responses == nil {
		// return errors.New("[spec.Responses] is empty")
		return nil
	}

	if responses.Default != nil {
		schema, err := spec.ResolveRef(xl.SwaggerSpec, &responses.Default.Schema.Ref)
		if err != nil {
			return err
		}
		schemaName, _ := xl.getDefinitionFromRef(responses.Default.Schema.Ref)
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "F", xl.Context.row), xl.Style.Center)
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Left)
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "default")
		xl.setCellWithSchema(schemaName, "body", strings.Join(schema.Type, ";"), responses.Default.Description)
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
		xl.File.SetCellInt(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), icode)
		response := responses.StatusCodeResponses[icode]
		// fmt.Println(icode, response)
		switch icode {
		// case 101:
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)
		// case 200, 201, 204:
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "string")
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)
		// case 400, 403, 404:
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "string")
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)
		// case 500:
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "string")
		// 	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)
		default:
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), response.Description)
		}

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
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(response.Schema.Type, ";"))
		}

		if response.Schema.Type.Contains("array") {
			items := response.Schema.Items
			if items.Schema != nil {
				// fmt.Println("items.Schema:", items.Schema)
				schema := items.Schema
				xl.setCellWithSchema(schema.Title, "body", strings.Join(response.Schema.Type, ";"), response.Description)
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
				xl.setCellWithSchema(propertyName, "body", strings.Join(response.Schema.Type, ";"), response.Description)
				xl.Context.row++
			}
		}
	}
	return nil
}

func (xl *Excel) getParameterFromRef(ref spec.Ref) *spec.Parameter {
	url := ref.GetURL()
	if url == nil {
		return nil
	}
	lastIndex := strings.LastIndex(url.Fragment, "/")
	parameterName := url.Fragment[lastIndex+1:]
	param := xl.SwaggerSpec.Parameters[parameterName]
	return &param
}

func (xl *Excel) getDefinitionFromRef(ref spec.Ref) (definitionName string, definition *spec.Schema) {
	url := ref.GetURL()
	if url == nil {
		return "", nil
	}
	lastIndex := strings.LastIndex(url.Fragment, "/")
	defName := url.Fragment[lastIndex+1:]
	def := xl.SwaggerSpec.Definitions[defName]
	return defName, &def
}

func (xl *Excel) setCellWithSchema(schemaName, paramType, dataType, description string) {
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), schemaName)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), paramType)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), dataType)
	// xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), enum)
	// xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "F", xl.Context.row), example)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), description)
}
