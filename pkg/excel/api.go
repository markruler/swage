package excel

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/markruler/swage/pkg/spec"
	"github.com/markruler/swage/pkg/style"
)

func createAPISheet(xl *excelize.File, path, operation string, detail spec.Operation, definitions map[string]spec.Definition, sheetName int) {
	worksheetName := strconv.Itoa(sheetName)
	xl.NewSheet(worksheetName)
	xl.SetColWidth(worksheetName, "A", "A", 12.0)
	xl.SetColWidth(worksheetName, "B", "B", 13.0)
	xl.SetColWidth(worksheetName, "F", "F", 40.0)

	row := 1
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Back to Index")
	xl.SetCellHyperLink(worksheetName, fmt.Sprintf("%s%d", "A", row), "INDEX!A1", "Location")
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "A", row), style.Button(xl))
	row++

	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Tag")
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	if len(detail.Tags) > 0 {
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), detail.Tags[0])
	}
	row++

	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Path")
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), path)
	row++
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Method")
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), operation)
	row++
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Summary")
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), detail.Summary)
	row++
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Description")
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "B", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), detail.Description)
	row++

	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "REQUEST")
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetRowHeight(worksheetName, row, 15)
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), style.Title(xl))
	row++

	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "required")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), "parameter")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "type")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), "level")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "data")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), "description")
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), style.Center(xl))
	row++

	for _, param := range detail.Parameters {
		if param.Required {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "O")
		} else {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "X")
		}
		// TODO: Set definitions
		if param.Schema.Ref != "" {
			lastIndex := strings.LastIndex(param.Schema.Ref, "/")
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), param.Schema.Ref[lastIndex+1:])
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), param.Schema.Type)
		} else if !reflect.DeepEqual(param.Schema.Items, spec.Items{}) {
			lastIndex := strings.LastIndex(param.Schema.Items.Ref, "/")
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), param.Schema.Items.Ref[lastIndex+1:])
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), param.Schema.Type)
			// fmt.Println("strings.Join(items[:], ", "):", strings.Join(param.Schema.Items.Enum[:], ", "))
		} else {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), param.Name)
		}

		if param.In != "" {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), param.In)
		}

		xl.SetCellInt(worksheetName, fmt.Sprintf("%s%d", "D", row), 1)

		if param.Type != "" {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), param.Type)
		}

		if param.Description != "" {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), param.Description)
		}

		xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "E", row), style.Center(xl))
		row++
	}
	row++

	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "RESPONSE")
	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetRowHeight(worksheetName, row, 15)
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), style.Title(xl))
	row++

	// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "required")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), "schema")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "type")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), "level")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "data")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), "description")
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), style.Center(xl))
	row++

	response := detail.Responses["200"]
	if reflect.DeepEqual(spec.Response{}, response) {
		return
	}

	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "body")
	xl.SetCellInt(worksheetName, fmt.Sprintf("%s%d", "D", row), 1)
	if response.Schema.Type == "array" {
		lastIndex := strings.LastIndex(response.Schema.Items.Ref, "/")
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), response.Schema.Items.Ref[lastIndex+1:])
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), response.Schema.Type)
	} else {
		// TODO: Set definitions
		lastIndex := strings.LastIndex(response.Schema.Ref, "/")
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), response.Schema.Ref[lastIndex+1:])
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "object")
	}
	// TODO: Set description
	// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), response.Description)
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "E", row), style.Center(xl))
	row++
}
