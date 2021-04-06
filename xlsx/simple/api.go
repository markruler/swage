package simple

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/markruler/swage/parser"
)

func (simple *Simple) CreateAPISheet() (err error) {
	xl := simple.xl
	apis := xl.SwageSpec.API
	if len(apis) == 0 {
		return errors.New("api is empty")
	}

	for index, api := range apis {
		xl.WorkSheetName = strconv.Itoa(index + 1)
		xl.File.NewSheet(xl.WorkSheetName)
		xl.Context.Row = 1

		simple.setAPISheetHeader(api)
		simple.setAPISheetRequest(api.Request)
		simple.setAPISheetResponse(api.Response)
	}
	return nil
}

func (simple *Simple) setAPISheetHeader(api parser.SwageAPI) {
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
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.Tag)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "ID")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.ID)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Path")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.Path)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Method")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.Method)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Consumes")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.Consumes)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Produces")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.Produces)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Summary")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.Summary)
	xl.Context.Row++
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "Description")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), api.Header.Description)
	xl.Context.Row++

	xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", 2), fmt.Sprintf("%s%d", "A", xl.Context.Row-1), xl.Style.Column)
	xl.Context.Row++
}

func (simple *Simple) setAPISheetRequest(parameters []parser.APIRequest) {
	xl := simple.xl
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "REQUEST")
	xl.File.MergeCell(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row))
	xl.File.SetRowHeight(xl.WorkSheetName, xl.Context.Row, 20.0)
	xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row), xl.Style.Title)
	xl.Context.Row++

	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), "required")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), "schema")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "C", xl.Context.Row), "param-type")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), "data-type")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "E", xl.Context.Row), "enum")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), "example")
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), "description")
	xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row), xl.Style.Column)
	xl.Context.Row++

	for _, param := range parameters {
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "E", xl.Context.Row), xl.Style.Center)
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), fmt.Sprintf("%s%d", "F", xl.Context.Row), xl.Style.Left)
		simple.setCellWithOneRequest(param)
		xl.Context.Row++
	}

	xl.Context.Row++
}

func (simple *Simple) setAPISheetResponse(responses []parser.APIResponse) {
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

	for _, response := range responses {
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "F", xl.Context.Row), xl.Style.Center)
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), fmt.Sprintf("%s%d", "G", xl.Context.Row), xl.Style.Left)
		simple.setCellWithOneResponse(response)
	}
}

func (simple *Simple) setCellWithOneRequest(param parser.APIRequest) {
	xl := simple.xl
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), param.Required)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), param.Schema)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "C", xl.Context.Row), param.ParameterType)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), param.DataType)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "E", xl.Context.Row), param.Enum)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), param.Example)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), param.Description)
}

func (simple *Simple) setCellWithOneResponse(response parser.APIResponse) {
	xl := simple.xl
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), response.StatusCode)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), response.Schema)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "C", xl.Context.Row), response.ResponseType)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), response.DataType)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "E", xl.Context.Row), response.Enum)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), response.Example)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), response.Description)
}
