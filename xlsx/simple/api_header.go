package simple

import (
	"fmt"
	"strings"

	"github.com/go-openapi/spec"
)

func (simple *Simple) setAPISheetHeader(path, method string, operation *spec.Operation) {
	xl := simple.xl
	xl.File.SetColWidth(xl.WorkSheetName, "A", "A", 12.0)
	xl.File.SetColWidth(xl.WorkSheetName, "B", "B", 33.0)
	xl.File.SetColWidth(xl.WorkSheetName, "C", "C", 12.0)
	xl.File.SetColWidth(xl.WorkSheetName, "D", "D", 12.0)
	xl.File.SetColWidth(xl.WorkSheetName, "E", "E", 20.0)
	xl.File.SetColWidth(xl.WorkSheetName, "F", "F", 40.0)
	xl.File.SetColWidth(xl.WorkSheetName, "G", "G", 90.0)
	// xl.File.SetColStyle(xl.WorkSheetName, "A:G", xl.Style.Line)

	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "A", xl.Context.Row), xl.Style.Button)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Back to Index")
	xl.File.SetCellHyperLink(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "INDEX!A1", "Location")
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Tag")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	if len(operation.Tags) > 0 {
		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), operation.Tags[0])
	}
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "ID")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), operation.ID)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Path")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), path)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Method")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), method)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Consumes")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), strings.Join(operation.Consumes, ", "))
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Produces")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), strings.Join(operation.Produces, ", "))
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Summary")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), operation.Summary)
	xl.Context.Row++
	// https://github.com/360EntSecGroup-Skylar/excelize/issues/573
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Description")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), operation.Description)
	xl.Context.Row++

	xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", 2), fmt.Sprintf("%s%d", "A", xl.Context.Row-1), xl.Style.Column)
	xl.Context.Row++
}
