package excel

import (
	"fmt"
	"strings"

	"github.com/go-openapi/spec"
)

func (xl *Excel) setAPISheetHeader(path, method string, operation *spec.Operation) {
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
	// https://github.com/360EntSecGroup-Skylar/excelize/issues/573
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "Description")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), operation.Description)
	xl.Context.row++

	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", 2), fmt.Sprintf("%s%d", "A", xl.Context.row-1), xl.Style.Column)
	xl.Context.row++
}
