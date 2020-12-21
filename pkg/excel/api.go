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
	rowResponse := xl.setAPISheetResponse(rowReuqest, worksheetName)
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

		if param.Schema != nil {
			if !reflect.DeepEqual(param.Schema.Ref, spec.Ref{}) {
				// lastIndex := strings.LastIndex(param.Schema.Ref, "/")
				// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), param.Schema.Ref[lastIndex+1:])
				definitionName := strings.TrimLeft(param.Schema.Ref.GetPointer().String(), "/")
				xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), definitionName)
				definition := xl.SwaggerSpec.Definitions[definitionName]
				xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), strings.Join(definition.Type, ";"))
			}
			if !reflect.DeepEqual(param.Schema.Items, spec.Items{}) {
				// TODO: array definition
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
	row++
	return row
}

func (xl *Excel) setAPISheetResponse(row int, worksheetName string) int {
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

	// response := detail.Responses["200"]
	// if reflect.DeepEqual(spec.Response{}, response) {
	// 	return errors.New("Response is empty")
	// }

	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "body")
	xl.File.SetCellInt(worksheetName, fmt.Sprintf("%s%d", "D", row), 1)
	// TODO: Set definitions
	// if response.Schema.Type == "array" {
	// 	lastIndex := strings.LastIndex(response.Schema.Items.Ref, "/")
	// 	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), response.Schema.Items.Ref[lastIndex+1:])
	// 	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), response.Schema.Type)
	// } else {
	// 	lastIndex := strings.LastIndex(response.Schema.Ref, "/")
	// 	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), response.Schema.Ref[lastIndex+1:])
	// 	xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "object")
	// }
	// TODO: Set description
	// xl.File.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), response.Description)
	xl.File.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "E", row), xl.Style.Center)
	row++
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
