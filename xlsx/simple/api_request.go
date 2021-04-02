package simple

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/parser"
)

func (simple *Simple) setAPISheetRequest(operation *spec.Operation) error {
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

	for _, param := range operation.Parameters {
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "A", xl.Context.Row), fmt.Sprintf("%s%d", "E", xl.Context.Row), xl.Style.Center)
		xl.File.SetCellStyle(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), fmt.Sprintf("%s%d", "F", xl.Context.Row), xl.Style.Left)

		if !reflect.DeepEqual(param.Ref, spec.Ref{}) {
			param = *simple.parameterFromRef(param.Ref)
		}

		simple.checkRequired(param.Required)

		b, err := json.MarshalIndent(param.Example, "", "    ")
		if err != nil {
			return err
		}
		simple.setCellWithSchema(param.Name, param.In, param.Type, string(b), param.Description)

		if param.Items != nil && param.Items.Enum != nil {
			xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "E", xl.Context.Row), parser.Enum2string(param.Items.Enum...))
		}

		if param.Enum != nil {
			xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "E", xl.Context.Row), parser.Enum2string(param.Enum...))
		}

		if param.Schema != nil {
			simple.parameterSchema(param)
		}

		xl.Context.Row++
	}
	xl.Context.Row++
	return nil
}
