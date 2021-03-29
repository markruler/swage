package excel

import (
	"fmt"
	"reflect"

	"github.com/go-openapi/spec"
)

func (xl *Excel) setAPISheetRequest(operation *spec.Operation) {
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

	for _, param := range operation.Parameters {
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "F", xl.Context.row), xl.Style.Center)

		if !reflect.DeepEqual(param.Ref, spec.Ref{}) {
			param = *xl.parameterFromRef(param.Ref)
		}

		xl.checkRequired(param.Required)

		xl.setCellWithSchema(param.Name, param.In, param.Type, param.Description)

		if param.Schema != nil {
			xl.parameterSchema(param)
			continue
		}

		if param.Items != nil && param.Items.Enum != nil {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), enum2string(param.Items.Enum...))
		}

		if param.Enum != nil {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), enum2string(param.Enum...))
		}
		xl.Context.row++
	}
	xl.Context.row++
}
