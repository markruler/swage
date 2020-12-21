package excel

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
)

func (xl *Excel) createAPISheet(path, method string, operation *spec.Operation, definitions spec.Definitions, sheetName int) error {
	if operation == nil {
		return errors.New("Operation should not be empty")
	}
	worksheetName := strconv.Itoa(sheetName)
	xl.File.NewSheet(worksheetName)

	row := 1
	rowHeader := xl.setAPISheetHeader(row, worksheetName, path, method, operation)
	if row == rowHeader {
		return errors.New("Something wrong happened")
	}
	rowReuqest := xl.setAPISheetRequest(rowHeader, worksheetName, operation)
	if rowHeader == rowReuqest {
		return errors.New("Something wrong happened")
	}
	rowResponse := xl.setAPISheetResponse(rowReuqest, worksheetName, operation)
	if rowReuqest == rowResponse {
		return errors.New("Something wrong happened")
	}
	return nil
}

func (xl *Excel) setAPISheetHeader(row int, worksheetName string, path, method string, operation *spec.Operation) int {
	xl.File.SetColWidth(worksheetName, "A", "A", 12.0)
	xl.File.SetColWidth(worksheetName, "B", "B", 13.0)
	xl.File.SetColWidth(worksheetName, "F", "F", 40.0)

	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Back to Index")
	xl.File.SetCellHyperLink(worksheetName, fmt.Sprintf("%s%d", "A", row), "INDEX!A1", "Location")
	xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "A", row), xl.Style.Button)
	row++

	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Tag")
	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	if len(operation.Tags) > 0 {
		xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), operation.Tags[0])
	}
	row++

	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Path")
	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), path)
	row++
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Method")
	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), method)
	row++
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Summary")
	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), operation.Summary)
	row++
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Description")
	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), operation.Description)
	row++
	return row
}

func (xl *Excel) setAPISheetRequest(row int, worksheetName string, operation *spec.Operation) int {
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "REQUEST")
	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.File.SetRowHeight(worksheetName, row, 15)
	xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), xl.Style.Title)
	row++

	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "required")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), "parameter")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "type")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), "level")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "data")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), "description")
	xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), xl.Style.Center)
	row++

	for _, param := range operation.Parameters {
		if param.Required {
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "O")
		} else {
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "X")
		}

		if param.Name != "" {
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), param.Name)
		}
		if param.Type != "" {
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), param.Type)
		}

		// TODO:
		if param.Schema != nil {
			if !reflect.DeepEqual(spec.Ref{}, param.Schema.Ref) {
				// lastIndex := strings.LastIndex(param.Schema.Ref, "/")
				// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), param.Schema.Ref[lastIndex+1:])
				definitionName := strings.TrimLeft(param.Schema.Ref.GetPointer().String(), "/")
				xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), definitionName)
				definition := xl.SwaggerSpec.Definitions[definitionName]
				xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), strings.Join(definition.Type, ";"))
			}
			if !reflect.DeepEqual(spec.Items{}, param.Schema.Items) {
				// definitionName := strings.TrimLeft(param.Schema.Ref.GetPointer().String(), "/")
				// xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), definitionName)
				// definition := xl.SwaggerSpec.Definitions[definitionName]
				// xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), strings.Join(definition.Type, ";"))
			}
		}

		if param.In != "" {
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), param.In)
		}

		xl.File.SetCellInt(worksheetName, fmt.Sprintf("%s%d", "D", row), 1)

		if param.Description != "" {
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), param.Description)
		}

		xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "E", row), xl.Style.Center)
		row++
	}
	return row + 1
}

func (xl *Excel) setAPISheetResponse(row int, worksheetName string, operation *spec.Operation) int {
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "RESPONSE")
	xl.File.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.File.SetRowHeight(worksheetName, row, 15)
	xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), xl.Style.Title)
	row++

	// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "required")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), "schema")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "type")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), "level")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "data")
	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), "description")
	xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), xl.Style.Center)
	row++

	// TODO:
	response := operation.Responses
	xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "E", row), xl.Style.Center)
	// fmt.Println(response == nil)
	// fmt.Println(reflect.DeepEqual(spec.Response{}, response))
	if response == nil {
		return row
	}

	// fmt.Println(response.StatusCodeResponses)
	var success spec.Response
	if _, ok := response.StatusCodeResponses[200]; !ok {
		return row
	}
	success = response.StatusCodeResponses[200]
	// fmt.Println(success.Ref)
	// fmt.Println(success.Schema)
	if success.Schema == nil || &success.Ref == nil {
		return row
	}

	// Items.Schemas
	if success.Schema.Type.Contains("array") {
		// lastIndex := strings.LastIndex(response.Schema.Items.Ref, "/")
		// xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), response.Schema.Items.Ref[lastIndex+1:])
		items := success.Schema.Items
		// fmt.Println(items)
		// fmt.Println(items.Schema) // nil
		if items.Schemas != nil {
			// fmt.Println("items.Schemas[0]:", items.Schemas[0])
			// fmt.Println("items.Schemas[0].Ref:", items.Schemas[0].Ref)
			definitionName := strings.TrimLeft(items.Schemas[0].Ref.GetPointer().String(), "/")
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), definitionName)
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "body")
			xl.File.SetCellInt(worksheetName, fmt.Sprintf("%s%d", "D", row), 1)
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "array")
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), success.Description)
		}
		return row
	}

	// Items.Schema
	// fmt.Println("success.Schema.Properties:", success.Schema.Properties)
	if success.Schema.Properties != nil {
		for propertyName, propertySchema := range success.Schema.Properties {
			// definitionName := strings.TrimLeft(propertySchema.Ref.GetPointer().String(), "/")
			lastIndex := strings.LastIndex(propertySchema.Ref.GetPointer().String(), "/")
			definitionName := propertySchema.Ref.GetPointer().String()[lastIndex+1:]
			if propertySchema.Items != nil {
				// lastIndex = strings.LastIndex(propertySchema.Items.Schema.Ref.GetPointer().String(), "/")
				lastIndex = strings.LastIndex(propertySchema.Items.Schema.Ref.GetPointer().String(), "/")
				definitionName = propertySchema.Items.Schema.Ref.GetPointer().String()[lastIndex+1:]
			}
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), definitionName)
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "body")
			xl.File.SetCellInt(worksheetName, fmt.Sprintf("%s%d", "D", row), 1)
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), propertyName)
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), propertySchema.Description)
		}
		return row
	}

	// fmt.Println("success.Schema.Items:", success.Schema.Items)
	if success.Schema.Items != nil {
		items := success.Schema.Items
		// fmt.Println("items.Schema.Type:", items.Schema.Type)
		if items.Schema.Type != nil {
			itemType := strings.Join(items.Schema.Type, ";")
			if success.Schema.Type != nil {
				schemaType := strings.Join(success.Schema.Type, ";")
				xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), schemaType)
			}
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "body")
			xl.File.SetCellInt(worksheetName, fmt.Sprintf("%s%d", "D", row), 1)
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), itemType)
			xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), items.Schema.Description)
		}
	}

	return row
}

func definitionFromRef(ref spec.Ref) string {
	url := ref.GetURL()
	if url == nil {
		return ""
	}
	fragmentParts := strings.Split(url.Fragment, "/")
	numParts := len(fragmentParts)

	return fragmentParts[numParts-1]
}
